// @Title  server.go
// @Description  ActionService实现
// @Create  heyitong  2022/6/17 17:33
// @Update  heyitong  2022/6/17 17:33

package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"math"
	"net"
	"os"
	"time"
	"tiops/action"
	actionTypes "tiops/action/types"
	apiClient "tiops/common/api-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/logger"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/utils"
	engineTypes "tiops/engine/types"
)

const (
	ActionMessageTypePushData = iota
)

var _actionServer = newActionServer()

type actionServer struct {
	server        *grpc.Server
	actions       map[string]actionTypes.StrictAction
	engines       map[string]engineTypes.WorkflowEngine
	actionInfoMap map[string]*models.ActionInfo
	//actionNodeOptionsMap    map[string]ActionOptions
	actionNodeContextMap    map[string]*actionTypes.ActionNodeContext
	actionContextMap        map[string]*actionTypes.ActionContext
	projectInfo             *models.ProjectInfo
	Logger                  *logger.Logger
	apiClient               *apiClient.APIClient
	updatingExecutionRecord bool
	requiredResourcesMap    map[string]*models.WorkflowResources
	//nodeStores              map[string]stores.DataStore
	//workspaceDataStore      stores.DataStore
	//workflowDataStore       DataStore
	//jobDataStore            DataStore
}

// GetServiceStatus 处理服务状态请求
func (a *actionServer) GetServiceStatus(ctx context.Context, request *services.EmptyRequest) (*services.ServiceStatus, error) {
	res := &services.ServiceStatus{ActionsStatus: map[string]*services.ActionStatus{}}
	for id, actionNodeCtx := range a.actionNodeContextMap {
		action0 := a.actions[actionNodeCtx.Info.Name]
		status := action0.Status(actionNodeCtx)
		res.ActionsStatus[id] = &services.ActionStatus{
			ProcessedCount: status.ProcessedCount,
			RestCount:      status.RestCount,
			Done:           status.Done,
			Message:        status.Message,
			Extra:          status.Extra,
		}
	}
	return res, nil
}

// PushMessage 处理消息请求
func (a *actionServer) PushMessage(server services.ActionsService_PushMessageServer) error {
	for {
		//fmt.Println(server)
		if actionMessage, err := server.Recv(); err == nil {
			//fmt.Println(actionMessage)
			//a.Logger.Info(actionMessage)
			switch actionMessage.Type {

			case services.ActionMessageType_PushData:
				//a.Logger.Info(actionMessage.Header)
				actionName := actionMessage.Message //actionMessage.Header["actionName"]
				action0 := a.actions[actionName]
				if action0 == nil {
					a.Logger.Error(errors.New("Action " + actionName + " not found"))
				} else {
					pushMessageContext := &actionTypes.PushMessageContext{
						ActionNodeContext: a.actionNodeContextMap[actionName],
						MessageHeader:     actionMessage.Header,
						MessageData:       actionMessage.Data,
						NodeId:            actionMessage.NodeId,
					}
					action0.OnMessage(pushMessageContext)
				}
			case services.ActionMessageType_StreamCmd:
				actionName := actionMessage.Header["actionName"]
				action0 := a.actions[actionName]
				if actionMessage.Message == "start" {

					if action0 == nil {
						a.Logger.Error(errors.New("Action " + actionName + " not found"))
					} else {

						actionData := &services.ActionData{}
						go func() {
							for {
								id := utils.SnowflakeID()
								idString := fmt.Sprintf("%x", id)
								actionData.TraceId = id
								actionData.Id = idString
								messageData, _ := proto.Marshal(actionData)
								pushMessageContext := &actionTypes.PushMessageContext{
									ActionNodeContext: a.actionNodeContextMap[actionName],
									MessageHeader:     actionMessage.Header,
									MessageData:       messageData,
									NodeId:            actionMessage.NodeId,
								}
								err = action0.OnMessage(pushMessageContext)
								if err != nil {
									a.Logger.Error(err)
								}
								if errors.Is(err, NoMoreDataError) {
									break
								}
							}

						}()
					}
				}

				return server.SendAndClose(&services.StatusResponse{Status: 1})
			}
		} else {
			a.Logger.Error(err)
			return err
		}
	}
}

// RegisterAction 注册普通处理模块
func (a *actionServer) RegisterAction(name string, action0 actionTypes.Action) *actionServer {
	a.actions[name] = action.NewStrictAction(action0)
	return a
}

//func (a *actionServer) RegisterFunction(name string, function actionTypes.ActionFunction) *actionServer {
//	return a.Register(name, function)
//}

// RegisterEngine 注册流程引擎
func (a *actionServer) RegisterEngine(name string, engine engineTypes.WorkflowEngine) *actionServer {
	a.engines[name] = engine
	return a
}

// init 初始化
func (a *actionServer) init() {
	/*for name, app := range a.actions {
		app.Init(&actionTypes.InitContext{ActionNodeContext: a.actionNodeContextMap[name]})
	}*/
}

// startServer 启动 ActionService server
func (a *actionServer) startServer() {
	sock, err := net.Listen("tcp", fmt.Sprintf(":%d", tiopsConfigs.ActionServerPort))
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := a.server.Serve(sock); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Start 启动 ActionService server 或运行流程引擎
func (a *actionServer) Start() {
	if tiopsConfigs.InMainEngine() || tiopsConfigs.RunEngineAtStartup() {
		//_workflow, err := workflow.Current()
		//context := workflow.Context()
		go a.startServer()
		a.runMainEngine()
		a.afterEngineExec()
	} else {
		a.init()
		a.startServer()
	}
}

// newActionServer 创建ActionService server
func newActionServer() *actionServer {

	tiopsApiClient := apiClient.NewAPIClient(tiopsConfigs.ApiServerHost, tiopsConfigs.ApiServerGrpcPort)
	remoteLogger := logger.NewRemoteLogger(
		"action-server",
		tiopsConfigs.LogId,
		tiopsConfigs.AppName,
		logger.StringToLogLevel(tiopsConfigs.LogLevel),
		tiopsApiClient,
	)

	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(math.MaxInt32),
		grpc.MaxSendMsgSize(1073741824),
	}

	s := grpc.NewServer(options...)

	myServer := &actionServer{
		server:  s,
		actions: map[string]actionTypes.StrictAction{},
		engines: map[string]engineTypes.WorkflowEngine{},
		//nodeStores:           map[string]stores.DataStore{},
		apiClient:        tiopsApiClient,
		Logger:           remoteLogger,
		actionContextMap: map[string]*actionTypes.ActionContext{},
		actionInfoMap:    map[string]*models.ActionInfo{},
		//actionNodeOptionsMap: map[string]ActionOptions{},
		actionNodeContextMap: map[string]*actionTypes.ActionNodeContext{},
		requiredResourcesMap: map[string]*models.WorkflowResources{},
	}

	/*myServer.projectInfo = &models.ProjectInfo{}
	_ = utils.UnmarshalYAMLFile(myServer.projectInfo, tiopsConfigs.ManifestPath)
	myServer.actionInfoMap = map[string]*models.ActionInfo{}
	for _, actionInfo := range myServer.projectInfo.Actions {
		myServer.actionInfoMap[actionInfo.Name] = actionInfo
		outputs := make([]string, len(actionInfo.Outputs))
		for i, output := range actionInfo.Outputs {
			outputs[i] = output.Name
		}
		inputs := make([]string, len(actionInfo.Inputs))
		for i, input := range actionInfo.Inputs {
			inputs[i] = input.Name
		}
		myServer.actionContextMap[actionInfo.Name] = &actionTypes.ActionContext{Logger: remoteLogger, Info: actionInfo, InputNames: inputs, OutputNames: outputs}
	}*/
	services.RegisterActionsServiceServer(s, myServer)

	if tiopsConfigs.RedirectStdOutErr {
		rErr, wErr, _ := os.Pipe()
		rOut, wOut, _ := os.Pipe()
		os.Stderr = wErr
		os.Stdout = wOut
		go func() {
			for true {
				var buf bytes.Buffer
				io.Copy(&buf, rErr)
				if buf.Len() > 0 {
					remoteLogger.Error(buf.String())
				} else {
					time.Sleep(time.Second)
				}
			}
		}()

		go func() {
			for true {
				var buf bytes.Buffer
				io.Copy(&buf, rOut)
				if buf.Len() > 0 {
					remoteLogger.Info(buf.String())
				} else {
					time.Sleep(time.Second)
				}
			}
		}()
	}

	return myServer
}
