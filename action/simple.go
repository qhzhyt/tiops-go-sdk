package action

import "tiops/action/types"

type simpleItemAction struct {
	itemFunc types.ItemProcessFunc
}

func (a *simpleItemAction) Call(ctx *types.PieceRequestContext) types.ActionDataItem {
	return a.itemFunc(ctx.Input)
}

func NewSimpleItemAction(processFunc types.ItemProcessFunc) *simpleItemAction {
	return &simpleItemAction{
		itemFunc: processFunc,
	}
}
