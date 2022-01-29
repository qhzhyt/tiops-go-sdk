package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"time"
	"tiops/buildin/engines"
	apiClient "tiops/common/api-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/logger"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/utils"
	"tiops/engine/types"
	"tiops/engine/workflow"
)

const (
	ActionMessageTypePushData = iota
)

var _actionServer = newActionServer()

type actionServer struct {
	server                  *grpc.Server
	actions                 map[string]StrictAction
	engines                 map[string]types.WorkflowEngine
	actionInfoMap           map[string]*models.ActionInfo
	actionNodeOptionsMap    map[string]ActionOptions
	actionContextMap        map[string]*ActionContext
	projectInfo             *models.ProjectInfo
	Logger                  *logger.Logger
	apiClient               *apiClient.APIClient
	updatingExecutionRecord bool
	requiredResourcesMap    map[int]*models.WorkflowResources
}

func (a *actionServer) RunEngine(ctx context.Context, request *services.RunEngineRequest) (*services.StatusResponse, error) {
	engine := a.engines[request.EngineName]
	if engine == nil {
		return &services.StatusResponse{Status: 404, Message: "Engine " + request.EngineName + " Not Found"}, nil
	}
	_workflow, err := workflow.New(request.WorkflowId)
	if err != nil {
		return &services.StatusResponse{Status: 404, Message: err.Error()}, nil
	}
	go func() {
		_context := workflow.Context()
		engine.Init(_context)
		engine.WaitForResources(_workflow)
		engine.Exec(_workflow)
	}()
	return &services.StatusResponse{Status: 200, Message: "ok"}, nil
}

func (a *actionServer) GetEngineStatus(ctx context.Context, request *services.EngineStatusRequest) (*services.EngineStatusResponse, error) {
	engine := a.engines[request.EngineName]
	if engine == nil {
		return &services.EngineStatusResponse{Code: -1, Message: "Engine " + request.EngineName + " Not Found"}, nil
	}
	status, msg := engine.Status()
	return &services.EngineStatusResponse{Code: int32(status), Message: msg}, nil
}

func (a *actionServer) updateExecutionRecord(ctx context.Context, record *models.ExecutionRecord) {
	if !a.updatingExecutionRecord {
		a.updatingExecutionRecord = true
		_, err := a.apiClient.CreateOrUpdateExecutionRecord(ctx, record)
		if err != nil {
			a.Logger.Error(err.Error())
		}
		a.updatingExecutionRecord = false
	}
}

func (a *actionServer) GetExecutionRecord(ctx context.Context, request *services.EmptyRequest) (*models.ExecutionRecord, error) {
	record := a.getCurrentEngine().ExecutionRecord()
	go a.updateExecutionRecord(ctx, record)
	return record, nil
}

func (a *actionServer) GetServiceStatus(ctx context.Context, request *services.EmptyRequest) (*services.ServiceStatus, error) {
	res := &services.ServiceStatus{ActionsStatus: map[string]*services.ActionStatus{}}
	for name, action := range a.actions {
		res.ActionsStatus[name] = action.Status()
	}
	return res, nil
}

func (a *actionServer) CallHttpAction(ctx context.Context, request *services.HttpRequest) (*services.HttpResponse, error) {
	panic("implement me")
}

var NoMoreDataError = errors.New("no more data")

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
				action := a.actions[actionName]
				if action == nil {
					a.Logger.Error(errors.New("action " + actionName + " not found"))
				} else {
					pushMessageContext := &PushMessageContext{
						ActionContext: a.actionContextMap[actionName],
						MessageHeader: actionMessage.Header,
						MessageData:   actionMessage.Data,
						NodeId:        actionMessage.NodeId,
					}
					action.OnMessage(pushMessageContext)
				}
			case services.ActionMessageType_StreamCmd:
				actionName := actionMessage.Header["actionName"]
				action := a.actions[actionName]
				if actionMessage.Message == "start" {

					if action == nil {
						a.Logger.Error(errors.New("action " + actionName + " not found"))
					} else {

						actionData := &services.ActionData{}
						go func() {
							for {
								id := utils.SnowflakeID()
								idString := fmt.Sprintf("%x", id)
								actionData.TraceId = id
								actionData.Id = idString
								messageData, _ := proto.Marshal(actionData)
								pushMessageContext := &PushMessageContext{
									ActionContext: a.actionContextMap[actionName],
									MessageHeader: actionMessage.Header,
									MessageData:   messageData,
									NodeId:        actionMessage.NodeId,
								}
								err = action.OnMessage(pushMessageContext)
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

func (a *actionServer) GetRequiredResources(ctx context.Context, query *services.QueryRequest) (*models.WorkflowResources, error) {
	stage, _ := strconv.Atoi(query.Extra["stage"])
	if res, ok := a.requiredResourcesMap[stage]; ok {
		return res, nil
	}
	engine := a.getCurrentEngine()
	wf, err := workflow.Current()
	if err != nil {
		return nil, err
	}
	requiredResources := engine.RequiredResources(wf, stage)
	requiredResources = workflow.ResourcesPreProcess(requiredResources, wf)
	a.requiredResourcesMap[stage] = requiredResources
	return requiredResources, nil
}

func (a *actionServer) getActionOptions(nodeId string) ActionOptions {
	return a.actionNodeOptionsMap[nodeId]
}

func (a *actionServer) CallAction(ctx context.Context, request *services.ActionRequest) (*services.ActionResponse, error) {
	actionName := request.ActionName
	a.Logger.Info(inputLog(actionName, request.Inputs))
	inputDataMap := TransActionDataMap(request.Inputs, a.actionInfoMap[actionName].Inputs)
	actionContext := a.actionContextMap[actionName]

	var result ActionDataBatch
	if a.actions[actionName] != nil {
		result = a.actions[actionName].CallBatch(
			&BatchRequestContext{
				ActionContext: actionContext,
				NodeId:        request.NodeId,
				Inputs:        inputDataMap,
				ActionOptions: a.getActionOptions(request.NodeId),
			})
	} else {
		return nil, errors.New("action " + actionName + " not found")
	}

	outputs := ToServiceActionDataMap(request.Id, request.TraceId, result, a.actionInfoMap[actionName].Outputs)
	a.Logger.Info(outputLog(actionName, outputs))

	return &services.ActionResponse{Id: request.Id, Outputs: outputs, Done: actionContext.HasDone(), TraceId: request.TraceId}, nil
}

func (a *actionServer) RegisterActionNode(ctx context.Context, request *services.RegisterActionNodeRequest) (*services.StatusResponse, error) {

	actionName := request.ActionName

	a.actionNodeOptionsMap[request.NodeId] = request.ActionOptions
	a.Logger.Info(fmt.Sprint("register node ", request.NodeId, " with options", request.ActionOptions))

	if action := a.actions[actionName]; action != nil {
		err := action.RegisterNode(&NodeRegisterContext{
			ActionContext: a.actionContextMap[actionName],
			NodeId:        request.NodeId,
			ActionOptions: request.ActionOptions,
			NextActions:   request.NextActions,
		},
		)
		if err != nil {
			return &services.StatusResponse{Status: services.Status_Failed}, err
		}
	} else {
		return nil, errors.New("action " + actionName + " not found")
	}

	return &services.StatusResponse{Status: services.Status_Ok}, nil
}

type actionServerCtl struct {
	actionServer *actionServer
}

func (a *actionServer) Register(name string, action Action) *actionServer {
	a.actions[name] = newStrictAction(action)
	return a
}

func (a *actionServer) RegisterFunction(name string, function ActionFunction) *actionServer {
	return a.Register(name, function)
}

func (a *actionServer) RegisterEngine(name string, engine types.WorkflowEngine) *actionServer {
	a.engines[name] = engine
	return a
}

func (a *actionServer) init() {
	for name, app := range a.actions {
		app.Init(&InitContext{a.actionContextMap[name]})
	}
}

func (a *actionServer) runMainEngine() {
	engine := a.getCurrentEngine()
	_workflow, err := workflow.Current()
	_context := workflow.Context()
	if err != nil {
		_context.Error(err)
		utils.SleepAndExit(time.Second*6, 3)
	}

	engine.Init(_context)

	//if requiredResources != nil {
	//	_, err := a.apiClient.CreateOrUpdateWorkflowExecution(
	//		&models.WorkflowExecution{
	//			XId:              tiopsConfigs.ExecutionID,
	//			WorkflowResource: requiredResources,
	//		})
	//	_context.Error(err)
	//}

	engine.WaitForResources(_workflow)
	//engine := engines.NewBasicChanEngine(context)
	engine.Exec(_workflow)
}

func (a *actionServer) getCurrentEngine() types.WorkflowEngine {
	a.Logger.Info("current engine: " + tiopsConfigs.EngineName)
	if a.engines[tiopsConfigs.EngineName] != nil {
		return a.engines[tiopsConfigs.EngineName]
	}
	for _, engine := range a.engines {
		return engine
	}
	return engines.NewBasicChanEngine()
}

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

func (a *actionServer) afterEngineExec() {
	record := a.getCurrentEngine().ExecutionRecord()
	a.Logger.Info(record)
	a.updateExecutionRecord(context.TODO(), record)
}

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

func newActionServer() *actionServer {

	tiopsApiClient := apiClient.NewAPIClient(tiopsConfigs.ApiServerHost, tiopsConfigs.ApiServerGrpcPort)
	remoteLogger := logger.NewRemoteLogger(
		"action-server",
		tiopsConfigs.LogId,
		os.Getenv(tiopsConfigs.EnvNameProjectId),
		logger.StringToLogLevel(tiopsConfigs.LogLevel),
		tiopsApiClient,
	)

	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(math.MaxInt32),
		grpc.MaxSendMsgSize(1073741824),
	}

	s := grpc.NewServer(options...)

	myServer := &actionServer{
		server:               s,
		actions:              map[string]StrictAction{},
		engines:              map[string]types.WorkflowEngine{},
		apiClient:            tiopsApiClient,
		Logger:               remoteLogger,
		actionNodeOptionsMap: map[string]ActionOptions{},
		actionContextMap:     map[string]*ActionContext{},
		requiredResourcesMap: map[int]*models.WorkflowResources{},
	}

	myServer.projectInfo = &models.ProjectInfo{}
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
		myServer.actionContextMap[actionInfo.Name] = &ActionContext{Logger: remoteLogger, Info: actionInfo, InputNames: inputs, OutputNames: outputs}
	}
	services.RegisterActionsServiceServer(s, myServer)
	return myServer
}

func ActionServer() *actionServerCtl {
	return &actionServerCtl{actionServer: _actionServer}
}

func (a *actionServerCtl) Register(name string, action Action) *actionServerCtl {
	a.actionServer.Register(name, action)
	return a
}

func (a *actionServerCtl) RegisterFunction(name string, function ActionFunction) *actionServerCtl {
	return a.Register(name, function)
}

func (a *actionServerCtl) Start() {
	a.actionServer.Start()
}

func (a *actionServerCtl) RegisterEngine(name string, engine types.WorkflowEngine) *actionServerCtl {
	a.actionServer.RegisterEngine(name, engine)
	return a
}
