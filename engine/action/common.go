package action

import (
	"tiops/common/models"
	"tiops/engine/types"
)

func New(actionInfo *models.ActionInfo, info *models.WorkflowNodeInfo) types.Action{
	switch actionInfo.Type {
	case models.ActionType_ServiceAction:
		return NewRemoteServiceAction(actionInfo, info)
	case models.ActionType_BuildInAction:
	case models.ActionType_CodeAction:
	case models.ActionType_EngineAction:
	case models.ActionType_WorkflowAction:
	default:

	}
	return nil
}
