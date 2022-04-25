package service

import (
	actionTypes "tiops/action/types"
	"tiops/common/logger"
	engineTypes "tiops/engine/types"
)

type actionServerCtl struct {
	actionServer *actionServer
}

func ActionServer() *actionServerCtl {
	return &actionServerCtl{actionServer: _actionServer}
}

func (a *actionServerCtl) Register(name string, action interface{}) *actionServerCtl {
	//a.actionServer.Register(name, action)

	switch action0 := action.(type) {
	case engineTypes.WorkflowEngine:
		logger.Info("Register WorkflowEngine " + name)
		a.actionServer.RegisterEngine(name, action0)
	case actionTypes.PieceProcessFunction:
		logger.Info("Register PieceProcessFunction " + name)
		a.actionServer.RegisterAction(name, actionTypes.PieceFunctionAction(action0))
	case actionTypes.BatchProcessFunction:
		logger.Info("Register BatchProcessFunction " + name)
		a.actionServer.RegisterAction(name, actionTypes.BatchFunctionAction(action0))
	case func(ctx *actionTypes.PieceRequestContext) actionTypes.ActionDataItem:
		logger.Info("Register func(ctx *actionTypes.PieceRequestContext) actionTypes.ActionDataItem " + name)
		a.actionServer.RegisterAction(name, actionTypes.PieceFunctionAction(action0))
	case func(ctx *actionTypes.BatchRequestContext) actionTypes.ActionDataBatch:
		logger.Info("Register func(ctx *actionTypes.BatchRequestContext) actionTypes.ActionDataBatch " + name)
		a.actionServer.RegisterAction(name, actionTypes.BatchFunctionAction(action0))
	case actionTypes.Action:
		logger.Info("Register Action " + name)
		a.actionServer.RegisterAction(name, action0)
	default:
		panic(IsNotActionOrEngineError(action0))
	}

	return a
}

//func (a *actionServerCtl) RegisterFunction(name string, function actionTypes.ActionFunction) *actionServerCtl {
//	return a.Register(name, function)
//}

func (a *actionServerCtl) Start() {
	a.actionServer.Start()
}

//func (a *actionServerCtl) RegisterEngine(name string, engine engineTypes.WorkflowEngine) *actionServerCtl {
//	a.actionServer.RegisterEngine(name, engine)
//	return a
//}
