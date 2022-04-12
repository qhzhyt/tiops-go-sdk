package types

type ActionFunction func(requestContext *PieceRequestContext) ActionDataItem

//func (af ActionFunction) Init(ctx *InitContext) {
//
//}
//
//func (af ActionFunction) RegisterNode(ctx *NodeRegisterContext) error {
//	return nil
//}

func (af ActionFunction) Call(ctx *PieceRequestContext) ActionDataItem {
	return af(ctx)
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
