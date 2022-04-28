package action

import (
	"encoding/json"
	actionClient "tiops/common/action-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/models"
	"tiops/engine/types"
)

type ExecutionAction struct {
	types.Action
	client              *actionClient.RemoteActionClient
	node                *types.Node
	nodeInfo            *models.WorkflowNodeInfo
	innerActionInfo     *models.ActionInfo
	executionActionInfo *models.ActionInfo
}

func (a *ExecutionAction) Control(ctrl types.ActionControl, args map[string]string) error {
	panic("implement me")
}

func getServiceNameByAction(nodeInfo *models.WorkflowNodeInfo, executionInfo *models.ActionInfo) string {
	//serviceName := tiopsConfigs.ActionServiceName(executionInfo.ProjectId)

	//if nodeInfo.StandAlone {
	serviceName := tiopsConfigs.StandAloneActionServiceName(nodeInfo.ActionName, nodeInfo.Id)
	//}
	return serviceName
}

func (a *ExecutionAction) Copy() types.Action {
	return &ExecutionAction{
		Action: New(a.executionActionInfo),
		client:              a.client,
		node:                a.node,
		nodeInfo:            a.nodeInfo,
		innerActionInfo:     a.innerActionInfo,
		executionActionInfo: a.executionActionInfo,
	}
}

func (a *ExecutionAction) Info() *models.ActionInfo {
	return a.innerActionInfo
}

func NewExecutionAction(actionInfo *models.ActionInfo) types.Action {
	executionActionInfo := &models.ActionInfo{}
	tmp, _ := json.Marshal(actionInfo.ExecutorInfo)
	_ = json.Unmarshal(tmp, executionActionInfo)
	executionActionInfo.InnerActionInfo = actionInfo
	return &ExecutionAction{
		Action: New(executionActionInfo),
		innerActionInfo:     actionInfo,
		executionActionInfo: executionActionInfo,
	}
}
