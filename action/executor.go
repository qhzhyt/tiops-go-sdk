package action

import (
	"tiops/action/types"
)

// executor 执行器
type executor struct {
	funcMap      map[string]types.ItemProcessFunc
	executorFunc types.ExecutorFunc
}

func (a *executor) RegisterNode(ctx *types.NodeRegisterContext) error {

	f, err := a.executorFunc(ctx.ActionNodeContext)

	if err != nil {
		return err
	}

	a.funcMap[ctx.NodeId] = f

	return nil
}

func (a *executor) Call(ctx *types.PieceRequestContext) types.ActionDataItem {

	f := a.funcMap[ctx.NodeId]

	if f == nil {
		var err error
		f, err = a.executorFunc(ctx.ActionNodeContext)

		if err != nil {
			return types.ActionDataItem{"err": err}
		}
	}

	return f(ctx.Input)
}

func NewSimpleExecutor(executorFunc types.ExecutorFunc) *executor {
	return &executor{
		funcMap:      map[string]types.ItemProcessFunc{},
		executorFunc: executorFunc,
	}
}
