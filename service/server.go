package service

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"
	"os"
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
	server               *grpc.Server
	actionBoolMap        map[string]bool
	actions              map[string]StrictAction
	engines              map[string]types.WorkflowEngine
	actionInfoMap        map[string]*models.ActionInfo
	actionNodeOptionsMap map[string]ActionOptions
	actionContextMap     map[string]*ActionContext
	projectInfo          *models.ProjectInfo
	Logger               *logger.Logger
	apiClient            *apiClient.APIClient
}

func (a *actionServer) PushMessage(server services.ActionsService_PushMessageServer) error {
	for {
		if actionMessage, err := server.Recv(); err == nil {
			a.Logger.Println(actionMessage)
			if actionMessage.Type == services.ActionMessageType_PushData {
				actionName := actionMessage.Header["actionName"]
				action := a.actions[actionName]
				if action == nil {
					a.Logger.Error(errors.New("action " + actionName + " not found"))
				} else {
					pushMessageContext := &PushMessageContext{
						ActionContext: a.actionContextMap[actionName],
						MessageHeader: actionMessage.Header,
						MessageData:   actionMessage.Data,
						NodeId: actionMessage.NodeId,
					}
					action.OnMessage(pushMessageContext)
				}
			}
		} else {
			a.Logger.Error(err)
			return err
		}
	}
}

func (a *actionServer) GetRequiredResources(ctx context.Context, info *models.WorkflowInfo) (*models.WorkflowResources, error) {
	panic("implement me")
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
	//a.actions[name] = action

	if action.(StrictAction) != nil {
		a.actions[name] = action.(StrictAction)
	} else {
		a.actions[name] = newStrictAction(action)
	}

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

	requiredResources := engine.RequiredResources(_workflow)

	requiredResources = workflow.ResourcesPreProcess(requiredResources, _workflow)

	_context.Info(requiredResources)

	if requiredResources != nil {
		_, err := a.apiClient.CreateOrUpdateWorkflowExecution(
			&models.WorkflowExecution{
				XId:              tiopsConfigs.ExecutionID,
				WorkflowResource: requiredResources,
			})
		_context.Error(err)
	}

	engine.WaitForResources(_workflow)
	//engine := engines.NewBasicChanEngine(context)
	engine.Exec(_workflow)
}

func (a *actionServer) getCurrentEngine() types.WorkflowEngine {
	if a.engines[tiopsConfigs.EngineName] != nil {
		return a.engines[tiopsConfigs.EngineName]
	}
	for _, engine := range a.engines {
		return engine
	}
	return engines.NewBasicChanEngine()
}

func (a *actionServer) Start() {
	if tiopsConfigs.InMainEngine() {
		//_workflow, err := workflow.Current()
		//context := workflow.Context()
		a.runMainEngine()
	} else {
		a.init()
		sock, err := net.Listen("tcp", fmt.Sprintf(":%d", tiopsConfigs.ActionServerPort))
		if err != nil {
			log.Fatal(err)
			return
		}
		if err := a.server.Serve(sock); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
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
		actionBoolMap:        map[string]bool{},
		apiClient:            tiopsApiClient,
		Logger:               remoteLogger,
		actionNodeOptionsMap: map[string]ActionOptions{},
		actionContextMap:     map[string]*ActionContext{},
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
