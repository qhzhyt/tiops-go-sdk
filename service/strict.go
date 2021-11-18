package service

import (
	"google.golang.org/protobuf/proto"
	"tiops/common/services"
)

type nodeData struct {
	nodeId      string
	nextActions []*services.NextActions
	messageCache map[int32]map[string]*services.ActionData
	actionOptions ActionOptions
}

type defaultStrictAction struct {
	action      Action
	nodeDataMap map[string]*nodeData
}

func (a *defaultStrictAction) sendOutputs(nodeCache *nodeData, outputs ServiceActionDataMap)  {

}

func (a *defaultStrictAction) Init(ctx *InitContext) {
	a.action.Init(ctx)
	a.nodeDataMap = map[string]*nodeData{}
}

func (a *defaultStrictAction) RegisterNode(ctx *NodeRegisterContext) error {
	a.nodeDataMap[ctx.NodeId] = &nodeData{
		nodeId:      ctx.NodeId,
		nextActions: ctx.NextActions,
		messageCache: map[int32]map[string]*services.ActionData{},
		actionOptions: ctx.ActionOptions,
	}
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

	if inputName != ""{
		dataCache[inputName] = actionData
	}

	if len(dataCache) >= len(ctx.InputNames){
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
		a.sendOutputs(nodeData0, outputs)
	}

	return nil
}

func newStrictAction(action Action) StrictAction {

	return &defaultStrictAction{action: action}
}
