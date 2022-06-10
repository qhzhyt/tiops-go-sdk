package action

import (
	"tiops/buildin/actions"
	"tiops/common/models"
	"tiops/engine/types"
)

func New(actionInfo *models.ActionInfo) types.Action {

	//actionInfo := actionInfos[nodeInfo.ActionId]

	if actionInfo.ExecutorInfo != nil {
		return NewExecutionAction(actionInfo)
	}

	switch actionInfo.Source {
	case models.ActionSource_Buildin:

		return actions.NewBuildinAction(actionInfo)
	case models.ActionSource_FromService:
		fallthrough
	default:
		switch actionInfo.Type {
		case models.ActionType_ExecutorAction:
			fallthrough
		case models.ActionType_ServiceAction:
			return NewRemoteServiceAction(actionInfo)
		case models.ActionType_BuildInAction:
			fallthrough
		case models.ActionType_CodeAction:
			fallthrough
		case models.ActionType_CustomAction:
			return NewCodeAction(actionInfo)
		case models.ActionType_EngineAction:
			return NewRemoteServiceAction(actionInfo)
		case models.ActionType_WorkflowAction:
			return NewWorkflowAction(actionInfo)
		default:
			return NewRemoteServiceAction(actionInfo)
		}
	}

	return nil
}

func newAction() {

}
