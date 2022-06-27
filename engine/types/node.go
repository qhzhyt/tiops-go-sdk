// @Title  node.go
// @Description  处理流程节点定义
// @Create  heyitong  2022/6/23 17:23
// @Update  heyitong  2022/6/23 17:23

package types

import (
	"tiops/common/models"
)

// Node 处理流程节点
type Node struct {
	ID              string
	Action          Action
	Inputs          InputConnectionsMap
	Outputs         OutputConnectionsMap
	Info            *models.WorkflowNodeInfo
	ActionOptions   map[string]string
	ActionExecutor  *models.ActionInfo
	SubNodes        map[string][]*Node
	HasVarInputOnly bool
}

func (n *Node) GetRequiredResources(stage int32) (*models.WorkflowResources, error) {
	return n.Action.GetRequiredResources(n, stage)
}
