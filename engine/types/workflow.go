package types

import (
	"fmt"
	"time"
	apiClient "tiops/common/api-client"
	"tiops/common/logger"
	"tiops/common/models"
)

const RegisterActionRetryTimes = 30

// Workflow 工作流
type Workflow struct {
	Name       string
	ID         string
	Nodes      map[string]*Node
	Actions    map[string]Action
	InputNode  *Node
	OutputNode *Node
	Variables  map[string]Variable
	//Packages   map[string]*Package
	Logger    *logger.Logger
	ApiClient *apiClient.APIClient
	info *models.WorkflowInfo
}

func (w *Workflow) Info() *models.WorkflowInfo {
	return w.info
}

func (w *Workflow) GetAction(aId string) Action {
	//return w.Packages[pId].GetAction(aName)
	return w.Actions[aId]
}

func (w *Workflow) RegisterActionNodes() bool {
	result := true
	for _, node := range w.Nodes {
		retryTimes := 0
		var err error
		for retryTimes < RegisterActionRetryTimes {
			if err = node.Action.Init(node); err != nil {
				retryTimes++
				time.Sleep(time.Second * 2)
			} else {
				break
			}
		}
		if retryTimes == RegisterActionRetryTimes {
			w.Logger.Error(fmt.Sprintf("register node %s for action %s failed, %s", node.Info.Id, node.Info.ActionName, err))
			result = false
		}
	}
	return result
}

func (w *Workflow) Init() {
	for _, node := range w.Nodes {
		node.HasVarInputOnly = node.Inputs.HasVarOnly()
	}
}
