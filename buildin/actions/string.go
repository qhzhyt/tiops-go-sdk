package actions

import (
	"tiops/action/types"
	actionTypes "tiops/action/types"
	"tiops/buildin"
)

// LogString 打印输入至日志
func LogString(ctx *actionTypes.BatchRequestContext) (actionTypes.ActionDataBatch, error) {
	ctx.Inputs.Foreach(func(item types.ActionDataItem) {
		ctx.Logger.Info(item[buildin.Input])
	})
	return types.ActionDataBatch{}, nil
}

// StringSplit 字符串分割
func StringSplit(ctx *actionTypes.BatchRequestContext) (actionTypes.ActionDataBatch, error) {
	ctx.Inputs.Foreach(func(item types.ActionDataItem) {
		ctx.Logger.Info(item[buildin.Input])
	})
	return types.ActionDataBatch{}, nil
}
