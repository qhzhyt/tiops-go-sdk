// @Title  action.go
// @Description  组件相关数据结构与接口
// @Create  heyitong  2022/6/17 18:50
// @Update  heyitong  2022/6/17 18:50

package types

// Action 基础处理组件，用户不实现任何接口仍可将组件注册到服务中，但接口将全部使用默认实现
type Action interface{}

// RegisterNode 注册节点接口，用于将
type RegisterNode interface {
	RegisterNode(ctx *NodeRegisterContext) error
}

// ActionInit 组件初始化接口
type ActionInit interface {
	Init(ctx *InitContext)
}

//type ActionApplication interface {
//	//PieceProcess
//}

type StatusProvider interface {
	Status(ctx *ActionNodeContext) *ActionStatus
}

type PieceProcess interface {
	Call(ctx *PieceRequestContext) ActionDataItem
}

type BatchProcess interface {
	CallBatch(ctx *BatchRequestContext) ActionDataBatch
}

type HttpProcess interface {
	CallHttp(ctx *HttpRequestContext) *HttpResponse
}

type PullStreamProcess interface {
	CallPullStream(ctx *StreamRequestContext) error
}

type PushMessageProcess interface {
	OnMessage(ctx *PushMessageContext) error
}

type BatchAction BatchProcess

type PieceAction PieceProcess

type HttpAction HttpProcess

type PullStreamAction PullStreamProcess

type StrictAction interface {
	PieceProcess
	BatchProcess
	ActionInit
	PullStreamProcess
	PushMessageProcess
	StatusProvider
	RegisterNode
	HttpProcess
}
