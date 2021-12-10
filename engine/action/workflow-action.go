package action

import (
	"tiops/common/models"
	"tiops/engine/types"
)

type WorkflowAction struct {
	RemoteServiceAction
}

func NewWorkflowAction(info *models.ActionInfo, nodeInfo *models.WorkflowNodeInfo) types.Action {
	return &WorkflowAction{RemoteServiceAction{
		info:     info,
		nodeInfo: nodeInfo,
	}}
}
