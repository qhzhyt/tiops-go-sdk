package langs

import (
	_ "embed"
	python3 "github.com/go-python/cpy3"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/engine/types"
)

//go:embed code/python3_init.py
var python3InitCode string

//var mainModule *python3.PyObject
var mainModuleDict *python3.PyObject
var callFunc *python3.PyObject
var callBatchFunc *python3.PyObject

type ProcessFunction func(*services.ActionData, ) (*services.ActionData, error)

func ProcessFunctionFromPythonFunc(pyFunc *python3.PyObject) ProcessFunction {
	return func(data *services.ActionData) (*services.ActionData, error) {
		return data, nil
	}
}

//func ProcessFunctionFromPythonBatchFunc(pyFunc *python3.PyObject) ProcessFunction {
//	return func(data *services.ActionData) (*services.ActionData, error) {
//
//	}
//}

func initPython3() {
	python3.Py_Initialize()
	python3.PyRun_SimpleString(python3InitCode)
	mainModuleDict = python3.PyModule_GetDict(python3.PyImport_GetModule(python3.PyUnicode_FromString("__main__")))
	callFunc = python3.PyDict_GetItemString(mainModuleDict, "call_func")
	callBatchFunc = python3.PyDict_GetItemString(mainModuleDict, "call_batch_func")
}

type Python3Action struct {
	PythonFunc      *python3.PyObject
	PythonBatchFunc *python3.PyObject
}

func (p *Python3Action) Init(node *types.Node) error {
	panic("implement me")
}

func (p *Python3Action) Call(request *types.ActionRequest) (*types.ActionResponse, error) {
	panic("implement me")
	//request.Inputs
}

func (p *Python3Action) Info() *models.ActionInfo {
	panic("implement me")
}

func (p *Python3Action) Type() types.ActionType {
	panic("implement me")
}

func (p *Python3Action) Copy() types.Action {
	panic("implement me")
}

func NewPython3Action(actionInfo *models.ActionInfo) *Python3Action {
	action := &Python3Action{}

	if mainModuleDict == nil {
		initPython3()
	}

	python3.PyRun_SimpleString(actionInfo.Code)

	//python3.PyRun_SimpleString()

	action.PythonFunc = callFunc.CallFunctionObjArgs(python3.PyDict_GetItemString(mainModuleDict, actionInfo.Func))

	//python3.Bytes

	return action
}
