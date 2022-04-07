package service

import (
	"tiops/common/logger"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/stores"
)

type ActionContext struct {
	Logger      *logger.Logger
	Info        *models.ActionInfo
	InputNames  []string
	OutputNames []string
}

type ActionNodeContext struct {
	*ActionContext
	Store         stores.DataStore
	NodeId        string
	ActionOptions ActionOptions
	done bool
}

type RequestContext struct {
	*ActionNodeContext
	//Store         stores.DataStore
	//NodeId        string
	Input         ActionDataItem
	//ActionOptions ActionOptions
}

type BatchRequestContext struct {
	*ActionNodeContext
	Inputs        ActionDataMap
	//ActionOptions ActionOptions
	done          bool
}

type StreamRequestContext struct {
	BatchRequestContext
	Push func(data ActionDataBatch) error
}

/*func (s *StreamRequestContext) Push(data ActionData) error {

}*/

type InitContext struct {
	*ActionNodeContext
}

type NodeRegisterContext struct {
	*ActionNodeContext
	//Store         stores.DataStore
	//NodeId        string
	//ActionOptions ActionOptions
	NextActions   []*services.NextActions
}

type ActionDataItem map[string]interface{}

type ActionDataBatch []ActionDataItem

//type ActionResult ActionDataItemMap
//type BatchResult ActionDataBatch

type ActionOptions map[string]string

type ServiceActionDataMap map[string]*services.ActionData

//type

type RegisterNode interface {
	RegisterNode(ctx *NodeRegisterContext) error
}

type ActionInit interface {
	Init(ctx *InitContext)
}

type ActionApplication interface {
	//PieceProcess
}

type StatusProvider interface {
	Status() *services.ActionStatus
}

type ActionItemFunc func(item ActionDataItem) ActionDataItem

type Action interface {
}

type BatchAction interface {
	Action
	BatchProcess
}

type PullStreamAction interface {
	CallPullStream(ctx *StreamRequestContext) error
}

type StrictAction interface {
	BatchAction
	ActionInit
	PullStreamAction
	PushMessageProcess
	StatusProvider
	RegisterNode
}

type PieceProcess interface {
	Call(ctx *RequestContext) ActionDataItem
}

type BatchProcess interface {
	CallBatch(ctx *BatchRequestContext) ActionDataBatch
}

type PushMessageContext struct {
	*ActionNodeContext
	NodeId        string
	MessageHeader ActionOptions
	MessageData   []byte
}

type PushMessageProcess interface {
	OnMessage(ctx *PushMessageContext) error
}
