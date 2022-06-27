// @Title  code-action.go
// @Description  代码组件
// @Create  heyitong  2022/6/23 16:56
// @Update  heyitong  2022/6/23 16:56

package action

import (
	"tiops/common/models"
	"tiops/engine/types"
)

type CodeAction struct {
	info     *models.ActionInfo
	nodeInfo *models.WorkflowNodeInfo
	code     string
	lang     string
	class    string
	funcName string
}

func (c *CodeAction) CallDuplexStream(callback func(res *types.ActionResponse, err error) bool) (func(request *types.ActionRequest) error, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CodeAction) GetRequiredResources(node *types.Node, stage int32) (*models.WorkflowResources, error) {
	panic("implement me")
}

func (c *CodeAction) CallPullStream(request *types.ActionRequest, callback func(res *types.ActionResponse, err error) bool) error {
	panic("implement me")
}

func (c *CodeAction) Control(ctrl types.ActionControl, args map[string]string) error {
	panic("implement me")
}

func (c *CodeAction) Init(node *types.Node) error {
	panic("implement me")
}

func (c *CodeAction) Call(request *types.ActionRequest) (*types.ActionResponse, error) {
	panic("implement me")
}

func (c *CodeAction) Info() *models.ActionInfo {
	panic("implement me")
}

func (c *CodeAction) Type() types.ActionType {
	panic("implement me")
}

func (c *CodeAction) Copy() types.Action {
	panic("implement me")
}

func NewCodeAction(info *models.ActionInfo) types.Action {

	action := &CodeAction{info: info}
	action.code = info.Code
	action.lang = info.Lang
	action.class = info.Class
	action.funcName = info.Func

	return action
}
