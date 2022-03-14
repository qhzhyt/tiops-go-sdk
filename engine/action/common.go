package action

import (
	"tiops/common/models"
	"tiops/engine/types"
)

func New(nodeInfo *models.WorkflowNodeInfo, actionInfos map[string]*models.ActionInfo) types.Action{
	actionInfo := actionInfos[nodeInfo.ActionId]
	switch actionInfo.Type {
	case models.ActionType_ServiceAction:
		return NewRemoteServiceAction(actionInfo, nodeInfo)
	case models.ActionType_BuildInAction:
		fallthrough
	case models.ActionType_CodeAction:
		fallthrough
	case models.ActionType_CustomAction:
		if nodeInfo.ActionExecutor != "" {
			return NewExecutionAction(actionInfo, actionInfos[nodeInfo.ActionExecutor], nodeInfo)
		}
		return NewCodeAction(actionInfo, nodeInfo)
	case models.ActionType_EngineAction:
		return NewRemoteServiceAction(actionInfo, nodeInfo)
	case models.ActionType_WorkflowAction:
		return NewWorkflowAction(actionInfo, nodeInfo)
	default:
	}
	return nil
}

func newAction() {

}
