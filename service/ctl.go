package service

import (
	actionTypes "tiops/action/types"
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
		a.actionServer.RegisterEngine(name, action0)
	case actionTypes.ActionFunction:
		a.actionServer.RegisterAction(name, action0)
	case actionTypes.Action:
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
