package action_server

import (
	"tiops/common/logger"
	"tiops/common/services"
)

type ActionContext struct {
	Logger *logger.Logger
	done   bool
}

type RequestContext struct {
	*ActionContext
	NodeId        string
	Inputs        map[string]interface{}
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

type ActionResult ActionDataItemMap
type BatchResult ActionDataItemsMap
type ActionOptions map[string]string

type ServiceActionDataMap map[string]*services.ActionData

type ActionFunction func(requestContext *RequestContext) ActionResult

//type

type ActionApplication interface {
	Init(ctx *InitContext)
	RegisterNode(ctx *NodeRegisterContext) error
	PieceProcess
	BatchProcess
}

type PieceProcess interface {
	Call(ctx *RequestContext) ActionResult
}

type BatchProcess interface {
	CallBatch(ctx *BatchRequestContext) BatchResult
}
