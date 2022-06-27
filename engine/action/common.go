// @Title  common.go
// @Description  action工厂，根据models.ActionInfo创建types.Action
// @Create  heyitong  2022/6/23 17:09
// @Update  heyitong  2022/6/23 17:09

package action

import (
	"tiops/buildin/actions"
	"tiops/common/models"
	"tiops/engine/types"
)

// New Action工厂函数，根据models.ActionInfo创建types.Action
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
