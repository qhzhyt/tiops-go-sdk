package types

import (
	"fmt"
	"sync"
	"time"
	apiClient "tiops/common/api-client"
	"tiops/common/logger"
	"tiops/common/models"
)

const RegisterActionRetryTimes = 120

// Workflow 工作流
type Workflow struct {
	Name    string
	ID      string
	Nodes   map[string]*Node
	Actions map[string]Action
	//Projects   map[string]*models.ProjectInfo
	InputNode  *Node
	OutputNode *Node
	Variables  map[string]Variable
	//Packages   map[string]*Package
	Logger      *logger.Logger
	ApiClient   *apiClient.APIClient
	info        *models.WorkflowInfo
	ActionInfos map[string]*models.ActionInfo
}

func (w *Workflow) Info() *models.WorkflowInfo {
	return w.info
}

func (w *Workflow) GetAction(aId string) Action {
	//return w.Packages[pId].GetAction(aName)
	return w.Actions[aId]
}

func (w *Workflow) GetActionInfo(aId string) *models.ActionInfo {
	//return w.Packages[pId].GetAction(aName)
	return w.ActionInfos[aId]
}

func (w *Workflow) RegisterActionNodes() bool {
	result := true
	wg := &sync.WaitGroup{}
	for _, node := range w.Nodes {
		wg.Add(1)
		node0 := node
		go func() {
			defer wg.Done()
			retryTimes := 0
			var err error
			for retryTimes < RegisterActionRetryTimes {
				if err = node0.Action.Init(node0); err != nil {
					retryTimes++
					time.Sleep(time.Second * 2)
				} else {
					break
				}
				if retryTimes%10 == 5 {
					w.Logger.Warning(fmt.Sprintf("register node %s for action %s error, %s, retry %d", node0.Info.Id, node0.Info.ActionName, err, retryTimes))
				}
			}
			if retryTimes == RegisterActionRetryTimes {
				w.Logger.Error(fmt.Sprintf("register node %s for action %s failed, %s", node0.Info.Id, node0.Info.ActionName, err))
				result = false
			}
			//w.Logger.Error(fmt.Sprintf("after register node %s for action %s", node0.Info.Id, node0.Info.ActionName))
		}()
	}
	wg.Wait()

	//w.Logger.Info("after RegisterActionNodes")

	return result
}

func (w *Workflow) Init() {
	for _, node := range w.Nodes {
		node.HasVarInputOnly = node.Inputs.HasVarOnly()
	}
}

func NewWorkflow(info *models.WorkflowInfo) *Workflow {
	return &Workflow{
		ID:          info.XId,
		Name:        info.Name,
		Logger:      logger.GetDefaultLogger(),
		Nodes:       map[string]*Node{},
		Actions:     map[string]Action{},
		ActionInfos: map[string]*models.ActionInfo{},
		Variables:   map[string]Variable{},
		info:        info,
	}
	//sync.NewCond(&sync.Mutex{})
}
