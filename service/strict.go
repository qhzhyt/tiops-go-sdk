package service

import (
	"context"
	"google.golang.org/protobuf/proto"
	"tiops/common/action-client"
	actionClient "tiops/common/action-client"
	"tiops/common/logger"
	"tiops/common/services"
)

type nodeData struct {
	nodeId             string
	nextActions        []*services.NextActions
	messageCache       map[int32]map[string]*services.ActionData
	actionOptions      ActionOptions
	actionDataMapQueue chan ServiceActionDataMap
	serviceClients     map[string]*actionClient.RemoteActionClient
	pushMessageClient  map[string]services.ActionsService_PushMessageClient
}

type defaultStrictAction struct {
	action      Action
	nodeDataMap map[string]*nodeData
	logger      *logger.Logger
}

func (a *defaultStrictAction) sendOutputs(nodeCache *nodeData) {
	for data := range nodeCache.actionDataMapQueue {

		for _, nextAction := range nodeCache.nextActions {
			output := data[nextAction.OutputName]
			outputData, _ := proto.Marshal(output)

			for _, action := range nextAction.Actions {


				creatClient := func() {
					nodeCache.serviceClients[action.Service] = actionClient.NewRemoteActionClient(action.Service, 5555)

					pushClient, err := nodeCache.serviceClients[action.Service].PushMessage(context.Background())

					if err != nil {
						a.logger.Error(err)
					}

					nodeCache.pushMessageClient[action.Service] = pushClient
				}

				if nodeCache.serviceClients[action.Service] == nil {
					creatClient()
				}

				send := func() error {
					return nodeCache.pushMessageClient[action.Service].Send(&services.ActionMessage{
						NodeId:  action.NodeId,
						Type:    0,
						Message: action.Action,
						Header: map[string]string{
							"inputName": action.InputName,
						},
						Data: outputData,
					})
				}

				if nodeCache.pushMessageClient[action.Service] != nil {
					if err := send(); err != nil {
						a.logger.Error(err)
						creatClient()
						if nodeCache.pushMessageClient[action.Service] != nil {
							if err := send(); err != nil {
								a.logger.Error(err)
							}
						}
					}
				}

			}

		}
	}
}

func (a *defaultStrictAction) Init(ctx *InitContext) {
	a.action.Init(ctx)
	a.nodeDataMap = map[string]*nodeData{}
}

func (a *defaultStrictAction) RegisterNode(ctx *NodeRegisterContext) error {
	nd := &nodeData{
		nodeId:             ctx.NodeId,
		nextActions:        ctx.NextActions,
		messageCache:       map[int32]map[string]*services.ActionData{},
		actionOptions:      ctx.ActionOptions,
		actionDataMapQueue: make(chan ServiceActionDataMap),
		serviceClients:     map[string]*action_client.RemoteActionClient{},
		pushMessageClient:  map[string]services.ActionsService_PushMessageClient{},
	}
	a.nodeDataMap[ctx.NodeId] = nd
	go a.sendOutputs(nd)
	return a.action.RegisterNode(ctx)
}

func (a *defaultStrictAction) Call(ctx *RequestContext) ActionDataItem {
	return a.action.Call(ctx)
}

func (a *defaultStrictAction) CallBatch(ctx *BatchRequestContext) ActionDataBatch {
	return a.action.CallBatch(ctx)
}

func (a *defaultStrictAction) OnMessage(ctx *PushMessageContext) error {

	actionData := &services.ActionData{}

	proto.Unmarshal(ctx.MessageData, actionData)

	traceId := actionData.TraceId

	nodeData0 := a.nodeDataMap[ctx.NodeId]

	inputName := ctx.MessageHeader["inputName"]

	if nodeData0.messageCache[traceId] == nil {
		nodeData0.messageCache[traceId] = map[string]*services.ActionData{}
	}

	dataCache := nodeData0.messageCache[traceId]

	if inputName != "" {
		dataCache[inputName] = actionData
	}

	if len(dataCache) >= len(ctx.InputNames) {
		delete(nodeData0.messageCache, traceId)
		inputDataMap := TransActionDataMap(dataCache, ctx.ActionContext.Info.Inputs)

		result := a.CallBatch(
			&BatchRequestContext{
				ActionContext: ctx.ActionContext,
				NodeId:        ctx.NodeId,
				Inputs:        inputDataMap,
				ActionOptions: nodeData0.actionOptions,
			})
		outputs := ToServiceActionDataMap(actionData.Id, traceId, result, ctx.ActionContext.Info.Outputs)
		nodeData0.actionDataMapQueue <- outputs
	}

	return nil
}

func newStrictAction(action Action) StrictAction {

	return &defaultStrictAction{action: action, logger: logger.GetDefaultLogger()}
}
