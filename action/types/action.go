package types

type Action interface{}

type RegisterNode interface {
	RegisterNode(ctx *NodeRegisterContext) error
}

type ActionInit interface {
	Init(ctx *InitContext)
}

//type ActionApplication interface {
//	//PieceProcess
//}

type StatusProvider interface {
	Status(ctx *ActionNodeContext) *ActionStatus
}

type ItemProcessFunc func(item ActionDataItem) ActionDataItem

type PieceProcess interface {
	Call(ctx *PieceRequestContext) ActionDataItem
}

type BatchProcess interface {
	CallBatch(ctx *BatchRequestContext) ActionDataBatch
}

type PullStreamProcess interface {
	CallPullStream(ctx *StreamRequestContext) error
}

type PushMessageProcess interface {
	OnMessage(ctx *PushMessageContext) error
}

type BatchAction BatchProcess

type PieceAction PieceProcess

type PullStreamAction PullStreamProcess

type StrictAction interface {
	PieceProcess
	BatchProcess
	ActionInit
	PullStreamProcess
	PushMessageProcess
	StatusProvider
	RegisterNode
}
