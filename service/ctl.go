// @Title  ctl.go
// @Description ActionServer封装，为用户提供调用接口
// @Create  heyitong  2022/6/17 17:05
// @Update  heyitong  2022/6/17 17:05

package service

import (
	actionTypes "tiops/action/types"
	engineTypes "tiops/engine/types"
)

type actionServerCtl struct {
	actionServer *actionServer
}

// ActionServer 获取包装后的_actionServer
func ActionServer() *actionServerCtl {
	return &actionServerCtl{actionServer: _actionServer}
}

// Register 将各类组件注册到ActionServer中，该函数会根据用户实现的接口注册为相应类型的组件
func (a *actionServerCtl) Register(name string, action interface{}) *actionServerCtl {
	//a.actionServer.Register(name, action)

	switch action0 := action.(type) {
	case engineTypes.WorkflowEngine:
		//logger.Info("Register WorkflowEngine " + name)
		a.actionServer.RegisterEngine(name, action0)
	case actionTypes.PieceProcessFunction:
		//logger.Info("Register PieceProcessFunction " + name)
		a.actionServer.RegisterAction(name, actionTypes.PieceFunctionAction(action0))
	case actionTypes.BatchProcessFunction:
		//logger.Info("Register BatchProcessFunction " + name)
		a.actionServer.RegisterAction(name, actionTypes.BatchFunctionAction(action0))
	case func(ctx *actionTypes.PieceRequestContext) actionTypes.ActionDataItem:
		//logger.Info("Register func(ctx *actionTypes.PieceRequestContext) actionTypes.ActionDataItem " + name)
		a.actionServer.RegisterAction(name, actionTypes.PieceFunctionAction(action0))
	case func(ctx *actionTypes.BatchRequestContext) actionTypes.ActionDataBatch:
		//logger.Info("Register func(ctx *actionTypes.BatchRequestContext) actionTypes.ActionDataBatch " + name)
		a.actionServer.RegisterAction(name, actionTypes.BatchFunctionAction(action0))
	case actionTypes.Action:
		//logger.Info("Register Action " + name)
		a.actionServer.RegisterAction(name, action0)
	default:
		panic(IsNotActionOrEngineError(action0))
	}

	return a
}

//func (a *actionServerCtl) RegisterFunction(name string, function actionTypes.ActionFunction) *actionServerCtl {
//	return a.Register(name, function)
//}

// Start 启动ActionServer
func (a *actionServerCtl) Start() {
	a.actionServer.Start()
}

//func (a *actionServerCtl) RegisterEngine(name string, engine engineTypes.WorkflowEngine) *actionServerCtl {
//	a.actionServer.RegisterEngine(name, engine)
//	return a
//}
