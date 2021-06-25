package types

import "tiops/common/models"

// Node 工作流节点
type Node struct {
	ID              string
	Action          Action
	Inputs          InputConnectionsMap
	Outputs         OutputConnectionsMap
	Info            *models.WorkflowNodeInfo
	HasVarInputOnly bool
}
