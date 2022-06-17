// @Title  strict.go
// @Description  完整处理组件封装，调用NewStrictAction可将部分实现的组件封装为完整的处理组件
// @Create  heyitong  2022/6/17 17:39
// @Update  heyitong  2022/6/17 17:39

package action

import (
	"context"
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"tiops/action/types"
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
	actionOptions      types.ActionOptions
	actionDataMapQueue chan types.ServiceActionDataMap
	serviceClients     map[string]*actionClient.RemoteActionClient
	pushMessageClient  map[string]services.ActionsService_PushMessageClient
}

type defaultStrictAction struct {
	action         types.Action
	nodeDataMap    map[string]*nodeData
	logger         *logger.Logger
	done           bool
	processedCount int64
	inited         bool
}

func (a *defaultStrictAction) CallHttp(ctx *types.HttpRequestContext) *types.HttpResponse {
	if h, ok := a.action.(types.HttpAction); ok {
		return h.CallHttp(ctx)
	}

	var dataList []types.ActionDataItem

	err := json.Unmarshal(ctx.Body, &dataList)

	if err != nil {
		return &types.HttpResponse{
			Status: 500,
			Body:   []byte(err.Error()),
		}
	}

	batchCtx := &types.BatchRequestContext{
		ActionNodeContext: ctx.ActionNodeContext,
		Inputs:            types.ActionDataBatch(dataList).ToActionDataMap(),
	}

	//var result []types.ActionDataItem

	result := a.CallBatch(batchCtx)

	resultBytes, err := json.Marshal(result)

	if err != nil {
		return &types.HttpResponse{
			Status: 500,
			Body:   []byte(err.Error()),
		}
	}

	return &types.HttpResponse{
		Status: 200,
		Body:   resultBytes,
	}
}

func (a *defaultStrictAction) CallPullStream(ctx *types.StreamRequestContext) error {
	streamAction := a.action.(types.PullStreamAction)
	if streamAction != nil {
		return streamAction.CallPullStream(ctx)
	}

	for {
		res := a.CallBatch(&ctx.BatchRequestContext)
		if res == nil || len(res) < 1 {
			ctx.Done()
			break
		}
		err := ctx.Push(res)
		if err != nil {
			return err
		}
		if ctx.HasDone() {
			break
		}
	}

	return nil
}

func (a *defaultStrictAction) Status(ctx *types.ActionNodeContext) *types.ActionStatus {
	if sp, ok := a.action.(types.StatusProvider); ok {
		return sp.Status(ctx)
	}

	count := 0

	for _, nodeData0 := range a.nodeDataMap {
		count += len(nodeData0.actionDataMapQueue) + len(nodeData0.messageCache)
	}

	res := &types.ActionStatus{
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

func (a *defaultStrictAction) Init(ctx *types.InitContext) {

	if i, ok := a.action.(types.ActionInit); ok {
		i.Init(ctx)
	}
	//a.action.Init(ctx)
	a.nodeDataMap = map[string]*nodeData{}

	a.inited = true
}

func (a *defaultStrictAction) RegisterNode(ctx *types.NodeRegisterContext) error {

	if !a.inited {
		a.Init(&types.InitContext{ActionContext: ctx.ActionContext})
	}

	nd := &nodeData{
		nodeId:             ctx.NodeId,
		nextActions:        ctx.NextActions,
		messageCache:       map[int64]map[string]*services.ActionData{},
		actionOptions:      ctx.ActionOptions,
		actionDataMapQueue: make(chan types.ServiceActionDataMap),
		serviceClients:     map[string]*action_client.RemoteActionClient{},
		pushMessageClient:  map[string]services.ActionsService_PushMessageClient{},
	}
	a.nodeDataMap[ctx.NodeId] = nd
	go a.sendOutputs(nd)

	if rn, ok := a.action.(types.RegisterNode); ok {
		return rn.RegisterNode(ctx)
	}

	return nil
}

func (a *defaultStrictAction) Call(ctx *types.PieceRequestContext) types.ActionDataItem {
	a.processedCount++

	action := a.action.(types.PieceProcess)

	if action != nil {
		return action.Call(ctx)
	}

	return ctx.Input
}

func (a *defaultStrictAction) CallBatch(ctx *types.BatchRequestContext) types.ActionDataBatch {
	if ba, ok := a.action.(types.BatchAction); ok {
		a.processedCount += int64(ctx.Inputs.Count())
		return ba.CallBatch(ctx)
	}

	//ctx.Logger.Info("CallBatch")

	requestContext := &types.PieceRequestContext{
		ActionNodeContext: ctx.ActionNodeContext,
		//NodeId: ctx.NodeId,
		//ActionOptions: ctx.ActionOptions,
		//Store: ctx.Store,
	}

	//ctx.Logger.Warning(ctx.Inputs)

	if len(ctx.Inputs) < 1 && len(ctx.OutputNames) > 0 {
		//	数据源
		batchSize := ctx.ActionOptions.GetIntOrDefault(BatchSizeName, 1)
		batches := ctx.ActionOptions.GetIntOrDefault(BatchesName, 1)

		batch := ctx.Store.GetIntOrDefault(BatchName, 0)
		//ctx.Logger.Warning(batchSize)
		//ctx.Logger.Warning(batches)
		//ctx.Store.PutValue(BatchName, batch)

		result := make(types.ActionDataBatch, batchSize)
		for i := 0; i < batchSize; i++ {
			result[i] = a.Call(requestContext)
		}
		batch++
		//ctx.Logger.Warning(batch)
		if batch >= batches {
			ctx.Done()
		}

		return result
	} else {
		return ctx.Inputs.Map(func(item types.ActionDataItem) types.ActionDataItem {
			requestContext.Input = item
			return a.Call(requestContext)
		})
	}
}

func (a *defaultStrictAction) OnMessage(ctx *types.PushMessageContext) error {

	if pmp, ok := a.action.(types.PushMessageProcess); ok {
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
		inputDataMap := types.TransActionDataMap(dataCache, ctx.ActionNodeContext.Info.Inputs)

		batchCtx := &types.BatchRequestContext{
			ActionNodeContext: ctx.ActionNodeContext,
			//NodeId:            ctx.NodeId,
			Inputs: inputDataMap,
			//ActionOptions:     nodeData0.actionOptions,
		}

		result := a.CallBatch(batchCtx)

		outputs := types.ToServiceActionDataMap(actionData.Id, traceId, result, ctx.ActionNodeContext.Info.Outputs)
		nodeData0.actionDataMapQueue <- outputs

		if batchCtx.HasDone() {
			a.done = true
		}
	}

	return nil
}

func NewStrictAction(action types.Action) types.StrictAction {

	return &defaultStrictAction{
		action: action,
		logger: logger.GetDefaultLogger(),
	}
}
