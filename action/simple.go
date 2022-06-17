// @Title  simple.go
// @Description  单函数处理组件封装
// @Create  heyitong  2022/6/17 17:39
// @Update  heyitong  2022/6/17 17:39

package action

import "tiops/action/types"

// simpleItemAction 简单处理组件
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
