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
	server               *grpc.Server
	actionBoolMap        map[string]bool
	actionFuncMap        map[string]ActionFunction
	actionAppMap         map[string]ActionApplication
	actionInfoMap        map[string]*models.ActionInfo
	actionOutputsMap     map[string][]string
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

	actionContext := &ActionContext{Logger: a.Logger}

	var result BatchResult

	if a.actionAppMap[actionName] != nil {
		result = a.actionAppMap[actionName].CallBatch(
			&BatchRequestContext{
				ActionContext: actionContext,
				NodeId:        request.NodeId,
				Inputs:        inputDataMap,
				ActionOptions: a.getActionOptions(request.NodeId),
			})
	} else {
		action := a.actionFuncMap[actionName]
		if action == nil {
			return nil, errors.New("action " + actionName + " not fount")
		}
		result = inputDataMap.MapTrans(
			func(m map[string]interface{}) map[string]interface{} {
				return action(&RequestContext{
					ActionContext: actionContext,
					NodeId:        request.NodeId,
					Inputs:        m,
					ActionOptions: a.actionNodeOptionsMap[request.NodeId],
				})
			}, a.actionOutputsMap[actionName])
	}

	outputs := ToServiceActionDataMap(request.Id, result, a.actionInfoMap[actionName].Outputs)
	a.Logger.Info(outputLog(actionName, outputs))

	return &services.ActionResponse{Id: request.Id, Outputs: outputs, Done: actionContext.HasDone()}, nil
}

func (a *actionServer) RegisterActionNode(ctx context.Context, request *services.RegisterActionNodeRequest) (*services.StatusResponse, error) {

	actionName := request.ActionName

	if !a.actionBoolMap[actionName] {
		return nil, errors.New("action " + actionName + " not found")
	}

	actionContext := &ActionContext{Logger: a.Logger}

	a.actionNodeOptionsMap[request.NodeId] = request.ActionOptions
	a.Logger.Info(fmt.Sprint("register node ", request.NodeId, " with options", request.ActionOptions))

	if app := a.actionAppMap[actionName]; app != nil {
		err := app.RegisterNode(&NodeRegisterContext{
			ActionContext: actionContext,
			NodeId:        request.NodeId,
			ActionOptions: request.ActionOptions,
		},
		)
		if err != nil {
			return &services.StatusResponse{Status: services.Status_Failed}, err
		}
	}

	return &services.StatusResponse{Status: services.Status_Ok}, nil
}

func ActionServer() *actionServer {
	return _actionServer
}

func (a *actionServer) Register(name string, obj interface{}) *actionServer {

	switch obj.(type) {
	case ActionApplication:
		return a.RegisterApplication(name, obj.(ActionApplication))
	case ActionFunction:
		return a.RegisterFunction(name, obj.(ActionFunction))
	default:
		log.Println(obj, "is not a action")
	}
	return a
}

func (a *actionServer) RegisterApplication(name string, application ActionApplication) *actionServer {
	//a.actionFuncMap[name] = function
	a.actionAppMap[name] = application
	a.actionBoolMap[name] = true
	return a
}

func (a *actionServer) RegisterFunction(name string, function ActionFunction) *actionServer {
	a.actionFuncMap[name] = function
	a.actionBoolMap[name] = true
	return a
}

func (a *actionServer) init() {
	actionContext := &ActionContext{Logger: a.Logger}
	for _, app := range a.actionAppMap {
		app.Init(&InitContext{actionContext})
	}
}

func (a *actionServer) Serve() {
	a.init()
	sock, err := net.Listen("tcp", fmt.Sprintf(":%d", 5555))
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
		os.Getenv(tiopsConfigs.ProjectIdEnvName),
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
		actionFuncMap:        map[string]ActionFunction{},
		actionAppMap:         map[string]ActionApplication{},
		actionBoolMap:        map[string]bool{},
		apiClient:            tiopsApiClient,
		Logger:               remoteLogger,
		actionNodeOptionsMap: map[string]ActionOptions{},
	}

	myServer.projectInfo = &models.ProjectInfo{}
	_ = utils.UnmarshalYAMLFile(myServer.projectInfo, "manifest.yaml")

	myServer.actionInfoMap = map[string]*models.ActionInfo{}
	myServer.actionOutputsMap = map[string][]string{}
	for _, actionInfo := range myServer.projectInfo.Actions {
		myServer.actionInfoMap[actionInfo.Name] = actionInfo
		outputs := make([]string, len(actionInfo.Outputs))
		for i, output := range actionInfo.Outputs {
			outputs[i] = output.Name
		}
		myServer.actionOutputsMap[actionInfo.Name] = outputs
	}

	services.RegisterActionsServiceServer(s, myServer)
	return myServer
}
