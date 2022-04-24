package types

type ItemProcessFunc func(item ActionDataItem) ActionDataItem

type ExecutorFunc func(ctx *ActionNodeContext) (ItemProcessFunc, error)

//type ActionFunction func(requestContext *PieceRequestContext) ActionDataItem

type PieceProcessFunction func(requestContext *PieceRequestContext) ActionDataItem
//func (af ActionFunction) Init(ctx *InitContext) {
//
//}
//
//func (af ActionFunction) RegisterNode(ctx *NodeRegisterContext) error {
//	return nil
//}

func (pf PieceProcessFunction) Call(ctx *PieceRequestContext) ActionDataItem {
	return pf(ctx)
}

type BatchProcessFunction func(requestContext *BatchRequestContext) ActionDataBatch

func (bf BatchProcessFunction) CallBatch(ctx *BatchRequestContext) ActionDataBatch {
	return bf(ctx)
}
//func (af ActionFunction) CallBatch(ctx *BatchRequestContext) ActionDataBatch {
//	/*return ctx.Inputs.MapTrans(func(m map[string]interface{}) map[string]interface{} {
//
//	}, ctx.Info.Inputs)*/
//
//	requestContext := &PieceRequestContext{ActionNodeContext: ctx.ActionNodeContext, NodeId: ctx.NodeId, ActionOptions: ctx.ActionOptions, Store: ctx.Store,}
//
//	return ctx.Inputs.Map(func(item ActionDataItem) ActionDataItem {
//		requestContext.Input = item
//		return af.Call(requestContext)
//	})
//}
