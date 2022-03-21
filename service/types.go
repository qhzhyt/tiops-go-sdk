package service

import (
	"github.com/qhzhyt/tiops-go-sdk/common/stores"
	"tiops/common/logger"
	"tiops/common/models"
	"tiops/common/services"
)

type ActionContext struct {
	Logger      *logger.Logger
	Info        *models.ActionInfo
	InputNames  []string
	OutputNames []string
	done        bool
}

type RequestContext struct {
	*ActionContext
	Store         stores.DataStore
	NodeId        string
	Input         ActionDataItem
	ActionOptions ActionOptions
}

type BatchRequestContext struct {
	*ActionContext
	Store         stores.DataStore
	NodeId        string
	Inputs        ActionDataMap
	ActionOptions ActionOptions
	done          bool
}

type InitContext struct {
	*ActionContext
}

type NodeRegisterContext struct {
	*ActionContext
	Store         stores.DataStore
	NodeId        string
	ActionOptions ActionOptions
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
	PieceProcess
}

type StatusProvider interface {
	Status() *services.ActionStatus
}

type ActionItemFunc func(item ActionDataItem) ActionDataItem

type Action ActionApplication

type BatchAction interface {
	Action
	BatchProcess
}

type StrictAction interface {
	BatchAction
	ActionInit
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
	*ActionContext
	NodeId        string
	MessageHeader ActionOptions
	MessageData   []byte
}

type PushMessageProcess interface {
	OnMessage(ctx *PushMessageContext) error
}
