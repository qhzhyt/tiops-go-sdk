package action_server

import (
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
	NodeId        string
	Input         ActionDataItem
	ActionOptions ActionOptions
}

type BatchRequestContext struct {
	*ActionContext
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
	NodeId        string
	ActionOptions ActionOptions
}

type ActionDataItem map[string]interface{}

type ActionDataBatch []ActionDataItem

//type ActionResult ActionDataItemMap
//type BatchResult ActionDataBatch

type ActionOptions map[string]string

type ServiceActionDataMap map[string]*services.ActionData

//type

type ActionApplication interface {
	Init(ctx *InitContext)
	RegisterNode(ctx *NodeRegisterContext) error
	PieceProcess
	BatchProcess
}

type ActionItemFunc func(item ActionDataItem) ActionDataItem

type Action ActionApplication

type PieceProcess interface {
	Call(ctx *RequestContext) ActionDataItem
}

type BatchProcess interface {
	CallBatch(ctx *BatchRequestContext) ActionDataBatch
}
