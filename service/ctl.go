package service

import (
	actionTypes "github.com/qhzhyt/tiops-go-sdk/action/types"
	engineTypes "github.com/qhzhyt/tiops-go-sdk/engine/types"
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
	case actionTypes.Action:
		a.actionServer.RegisterAction(name, action0)
	case engineTypes.WorkflowEngine:
		a.actionServer.RegisterEngine(name, action0)
	case actionTypes.ActionFunction:
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
