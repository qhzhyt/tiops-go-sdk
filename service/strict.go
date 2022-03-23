package service

import (
	"context"
	"google.golang.org/protobuf/proto"
	"tiops/common/action-client"
	actionClient "tiops/common/action-client"
	"tiops/common/logger"
	"tiops/common/services"
)

const (
	BatchSizeName = "batchSize"
	BatchesName   = "batches"
	BatchName     = "batch"
)

type nodeData struct {
	nodeId             string
	nextActions        []*services.NextActions
	messageCache       map[int64]map[string]*services.ActionData
	actionOptions      ActionOptions
	actionDataMapQueue chan ServiceActionDataMap
	serviceClients     map[string]*actionClient.RemoteActionClient
	pushMessageClient  map[string]services.ActionsService_PushMessageClient
}

type defaultStrictAction struct {
	action         Action
	nodeDataMap    map[string]*nodeData
	logger         *logger.Logger
	done           bool
	processedCount int64
}

func (a *defaultStrictAction) Status() *services.ActionStatus {
	if sp, ok := a.action.(StatusProvider); ok {
		return sp.Status()
	}

	count := 0

	for _, nodeData0 := range a.nodeDataMap {
		count += len(nodeData0.actionDataMapQueue) + len(nodeData0.messageCache)
	}

	res := &services.ActionStatus{
		RestCount:      int64(count),
		ProcessedCount: a.processedCount,
		Done:           a.done,
	}

	return res
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

	if i, ok := a.action.(ActionInit); ok {
		i.Init(ctx)
	}

	//a.action.Init(ctx)
	a.nodeDataMap = map[string]*nodeData{}
}

func (a *defaultStrictAction) RegisterNode(ctx *NodeRegisterContext) error {
	nd := &nodeData{
		nodeId:             ctx.NodeId,
		nextActions:        ctx.NextActions,
		messageCache:       map[int64]map[string]*services.ActionData{},
		actionOptions:      ctx.ActionOptions,
		actionDataMapQueue: make(chan ServiceActionDataMap),
		serviceClients:     map[string]*action_client.RemoteActionClient{},
		pushMessageClient:  map[string]services.ActionsService_PushMessageClient{},
	}
	a.nodeDataMap[ctx.NodeId] = nd
	go a.sendOutputs(nd)

	if rn, ok := a.action.(RegisterNode); ok {
		return rn.RegisterNode(ctx)
	}

	return nil
}

func (a *defaultStrictAction) Call(ctx *RequestContext) ActionDataItem {
	a.processedCount++
	return a.action.Call(ctx)
}

func (a *defaultStrictAction) CallBatch(ctx *BatchRequestContext) ActionDataBatch {
	if ba, ok := a.action.(BatchAction); ok {
		a.processedCount += int64(ctx.Inputs.Count())
		return ba.CallBatch(ctx)
	}
	requestContext := &RequestContext{ActionContext: ctx.ActionContext, NodeId: ctx.NodeId, ActionOptions: ctx.ActionOptions, Store: ctx.Store}

	//a.logger.Warning(len(ctx.Inputs))

	if len(ctx.Inputs) < 1 {
		//	数据源
		batchSize := ctx.ActionOptions.GetIntOrDefault(BatchSizeName, 1)
		batches := ctx.ActionOptions.GetIntOrDefault(BatchesName, 1)

		//a.logger.Warning(batchSize)
		//a.logger.Warning(batches)

		batch := ctx.Store.GetIntOrDefault(BatchName, 0)
		result := make(ActionDataBatch, batchSize)
		for i := 0; i < batchSize; i++ {
			result[i] = a.Call(requestContext)
		}
		batch++
		if batch >= batches {
			ctx.Done()
		}
		ctx.Store.PutValue(BatchName, batch)

		return result
	} else {
		return ctx.Inputs.Map(func(item ActionDataItem) ActionDataItem {
			requestContext.Input = item
			return a.Call(requestContext)
		})
	}
}

func (a *defaultStrictAction) OnMessage(ctx *PushMessageContext) error {

	if pmp, ok := a.action.(PushMessageProcess); ok {
		return pmp.OnMessage(ctx)
	}

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

		batchCtx := &BatchRequestContext{
			ActionContext: ctx.ActionContext,
			NodeId:        ctx.NodeId,
			Inputs:        inputDataMap,
			ActionOptions: nodeData0.actionOptions,
		}

		result := a.CallBatch(batchCtx)

		outputs := ToServiceActionDataMap(actionData.Id, traceId, result, ctx.ActionContext.Info.Outputs)
		nodeData0.actionDataMapQueue <- outputs

		if batchCtx.HasDone() {
			a.done = true
		}
	}

	return nil
}

func newStrictAction(action Action) StrictAction {

	return &defaultStrictAction{
		action: action,
		logger: logger.GetDefaultLogger(),
	}
}
