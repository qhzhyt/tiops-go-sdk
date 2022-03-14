package types

import "tiops/common/models"

// Node 工作流节点
type Node struct {
	ID              string
	Action          Action
	Inputs          InputConnectionsMap
	Outputs         OutputConnectionsMap
	Info            *models.WorkflowNodeInfo
	ActionExecutor  *models.ActionInfo
	SubNodes        map[string][]*Node
	HasVarInputOnly bool
}
