package main

// #cgo pkg-config: python3
// #include <Python.h>
import "C"

import (
	python3 "github.com/go-python/cpy3"

	"tiops/common/models"
	"tiops/engine/action/langs"
)

func main() {
	pycodeGo := `
import json
import sys
# for path in sys.path:
# 	print(path)
def json_loads(string):
	print(string)
	print(json.loads(string))


`
	action := langs.NewPython3Action(&models.ActionInfo{Code: pycodeGo, Func: "json_loads"})

	action.PythonFunc.CallFunctionObjArgs(python3.PyBytes_FromString(string([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9})))
}
