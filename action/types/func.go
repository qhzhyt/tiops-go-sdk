// @Title  func.go
// @Description  处理组件函数相关数据结构定义
// @Create  heyitong  2022/6/23 15:22
// @Update  heyitong  2022/6/23 15:22

package types

// ItemProcessFunc 单条数据处理函数
type ItemProcessFunc func(item ActionDataItem) ActionDataItem

// ExecutorFunc 执行器函数
type ExecutorFunc func(ctx *ActionNodeContext) (ItemProcessFunc, error)

//type ActionFunction func(requestContext *PieceRequestContext) ActionDataItem

// PieceProcessFunction 单条数据处理函数
type PieceProcessFunction func(requestContext *PieceRequestContext) ActionDataItem

// BatchProcessFunction 批量数据处理函数
type BatchProcessFunction func(requestContext *BatchRequestContext) ActionDataBatch

// PieceFunctionAction 单条数据处理函数
type PieceFunctionAction PieceProcessFunction

// BatchFunctionAction 批量数据处理函数组件
type BatchFunctionAction BatchProcessFunction

//func (af ActionFunction) Init(ctx *InitContext) {
//
//}
//
//func (af ActionFunction) RegisterNode(ctx *NodeRegisterContext) error {
//	return nil
//}

func (pf PieceFunctionAction) Call(ctx *PieceRequestContext) ActionDataItem {
	return pf(ctx)
}

func (bf BatchFunctionAction) CallBatch(ctx *BatchRequestContext) ActionDataBatch {
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
