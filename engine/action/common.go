package action

import (
	"tiops/common/models"
	"tiops/engine/types"
)

func New(info *models.ActionInfo) types.Action{
	switch info.Type {
	case models.ActionType_ServiceAction:
		return NewRemoteServiceAction(info)
	case models.ActionType_BuildInAction:
	case models.ActionType_CodeAction:
	case models.ActionType_EngineAction:
	case models.ActionType_WorkflowAction:
	default:

	}
	return nil
}
