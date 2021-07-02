package action_server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"
	"os"
	apiClient "tiops/common/api-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/logger"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/utils"
)

var _actionServer = newActionServer()

type actionServer struct {
	server        *grpc.Server
	actionBoolMap map[string]bool
	//actionFuncMap        map[string]ActionFunction
	//actionAppMap         map[string]ActionApplication

	actions map[string]Action

	actionInfoMap map[string]*models.ActionInfo
	//actionOutputsMap     map[string][]string
	actionNodeOptionsMap map[string]ActionOptions
	actionContextMap     map[string]*ActionContext
	projectInfo          *models.ProjectInfo
	Logger               *logger.Logger
	apiClient            *apiClient.APIClient
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

	outputs := ToServiceActionDataMap(request.Id, result, a.actionInfoMap[actionName].Outputs)
	a.Logger.Info(outputLog(actionName, outputs))

	return &services.ActionResponse{Id: request.Id, Outputs: outputs, Done: actionContext.HasDone()}, nil
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

func ActionServer() *actionServer {
	return _actionServer
}

func (a *actionServer) Register(name string, action Action) *actionServer {
	a.actions[name] = action
	return a
}

//func (a *actionServer) RegisterApplication(name string, application ActionApplication) *actionServer {
//	//a.actionFuncMap[name] = function
//	a.actionAppMap[name] = application
//	a.actionBoolMap[name] = true
//	return a
//}
//

func (a *actionServer) RegisterFunction(name string, function ActionFunction) *actionServer {
	return a.Register(name, function)
}

func (a *actionServer) init() {
	for name, app := range a.actions {
		app.Init(&InitContext{a.actionContextMap[name]})
	}
}

func (a *actionServer) Serve() {
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
		actions:              map[string]Action{},
		actionBoolMap:        map[string]bool{},
		apiClient:            tiopsApiClient,
		Logger:               remoteLogger,
		actionNodeOptionsMap: map[string]ActionOptions{},
		actionContextMap: map[string]*ActionContext{},
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
