package action

import (
	"github.com/qhzhyt/tiops-go-sdk/common/models"
	"github.com/qhzhyt/tiops-go-sdk/engine/types"
)

type CodeAction struct {
	info     *models.ActionInfo
	nodeInfo *models.WorkflowNodeInfo
	code     string
	lang     string
	class    string
	funcName string
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

func NewCodeAction(info *models.ActionInfo, nodeInfo *models.WorkflowNodeInfo) types.Action {

	action := &CodeAction{info: info, nodeInfo: nodeInfo}
	action.code = info.Code
	action.lang = info.Lang
	action.class = info.Class
	action.funcName = info.Func

	return action
}
