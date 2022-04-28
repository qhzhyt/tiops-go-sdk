package action

import (
	"tiops/common/models"
	"tiops/engine/types"
)

type WorkflowAction struct {
	RemoteServiceAction
}

func NewWorkflowAction(info *models.ActionInfo) types.Action {
	return &WorkflowAction{RemoteServiceAction{
		info:     info,
	}}
}
