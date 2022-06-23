// @Title  build.go
// @Description  处理流程构建
// @Create  heyitong  2022/6/23 17:30
// @Update  heyitong  2022/6/23 17:30

package workflow

import (
	apiClient "tiops/common/api-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/logger"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/utils"
	"tiops/engine/action"
	"tiops/engine/types"
	"tiops/engine/variable"
)

var _apiClient = apiClient.NewAPIClient(tiopsConfigs.ApiServerHost, tiopsConfigs.ApiServerGrpcPort)

func loadActionInfos(wi *models.WorkflowInfo, client *apiClient.APIClient) (map[string]*models.ActionInfo, error) {
	nodeInfos := wi.Spec.Nodes

	actionNameSet := map[string]bool{}
	for _, nodeInfo := range nodeInfos {
		actionNameSet[nodeInfo.ActionId] = true
		if nodeInfo.ActionExecutor != "" {
			actionNameSet[nodeInfo.ActionExecutor] = true
		}
	}

	if wi.Engine != "" {
		actionNameSet[wi.Engine] = true
	}

	actionIds := make([]string, 0, len(actionNameSet))
	for name, _ := range actionNameSet {
		actionIds = append(actionIds, name)
	}
	actionInfos, err := client.GetActionListByIds(actionIds)
	result := map[string]*models.ActionInfo{}
	for _, actionInfo := range actionInfos {
		result[actionInfo.XId] = actionInfo
	}
	for _, actionInfo := range actionInfos {
		actionInfo.ExecutorInfo = result[actionInfo.Executor]
	}

	if wi.Engine != "" {
		wi.EngineInfo = result[wi.Engine]
	}

	result, err = loadSubFlowEngineInfos(result, client)

	if err != nil {
		return nil, err
	}

	//logger.Info(actionInfos)
	return result, err
}

func loadSubFlowEngineInfos(_actionInfos map[string]*models.ActionInfo, client *apiClient.APIClient) (map[string]*models.ActionInfo, error) {
	actionNameSet := map[string]bool{}
	for _, actionInfo := range _actionInfos {
		if actionInfo.Type == models.ActionType_WorkflowAction && actionInfo.Func != "" {
			actionNameSet[actionInfo.Func] = true
		}
	}

	if len(actionNameSet) > 0 {
		actionIds := make([]string, 0, len(actionNameSet))
		for name, _ := range actionNameSet {
			actionIds = append(actionIds, name)
		}
		actionInfos, err := client.GetActionListByIds(actionIds)
		//result := map[string]*models.ActionInfo{}
		for _, actionInfo := range actionInfos {
			_actionInfos[actionInfo.XId] = actionInfo
		}
		for _, actionInfo := range _actionInfos {
			if actionInfo.Type == models.ActionType_WorkflowAction && actionInfo.Func != "" {
				actionInfo.EngineInfo = _actionInfos[actionInfo.Func]
			}
		}
		return _actionInfos, err
	}

	//logger.Info(actionInfos)
	return _actionInfos, nil
}

// buildWorkflow 根据 models.WorkflowInfo 构建相应的 Workflow 对象
func buildWorkflow(wi *models.WorkflowInfo, client *apiClient.APIClient) (*types.Workflow, error) {

	wf := types.NewWorkflow(wi)

	wf.ApiClient = client

	actionInfos, err := loadActionInfos(wi, client)

	if err != nil {
		return nil, err
	}

	wf.EngineInfo = wi.EngineInfo

	//engineInfos, err := loadSubFlowEngineInfos(actionInfos, client)

	//if err != nil {
	//	return nil, err
	//}

	//for xId, info := range engineInfos {
	//	actionInfos[xId] = info
	//}

	wf.ActionInfos = actionInfos
	//marshal, _ := json.Marshal(actionInfos)
	//wf.Logger.Error(string(marshal))

	for _, nodeInfo := range wi.Spec.Nodes {
		wf.Actions[nodeInfo.Id] = action.New(actionInfos[nodeInfo.ActionId])
	}

	//time.Sleep(time.Second * 6)

	spec := wi.Spec
	//fmt.Println(wi.Constants)
	nodeInfos := map[string]*models.WorkflowNodeInfo{}
	nodes := map[string]*types.Node{}
	for _, _variable := range spec.Variables {
		wf.Variables[_variable.Name] = variable.New(_variable.Name, _variable.Type, _variable.ValueType, _variable.Value)
	}

	inputNode := &types.Node{ID: types.InputNodeId, Outputs: types.OutputConnectionsMap{}, Info: &models.WorkflowNodeInfo{
		Id: types.InputNodeId, ActionId: types.InputNodeId, ProjectId: types.InputNodeId, ActionName: types.InputNodeId, Inputs: map[string]*models.WorkflowConnections{}, StandAlone: false,
	}}

	//packages := map[string]*Package{}
	if spec.Inputs != nil && len(spec.Inputs) > 0 {
		nodes[types.InputNodeId] = inputNode
		for _, input := range spec.Inputs {
			inputNode.Outputs[input.Name] = []*types.Connection{}
		}
		wf.InputNode = inputNode
	}
	// 遍历节点信息列表，生成节点对象列表
	for _, nodeInfo := range spec.Nodes {
		nodeInfos[nodeInfo.Id] = nodeInfo
		node := &types.Node{
			ID:             nodeInfo.Id,
			Action:         wf.GetAction(nodeInfo.Id),
			Inputs:         types.InputConnectionsMap{},
			Outputs:        types.OutputConnectionsMap{},
			Info:           nodeInfo,
			SubNodes:       map[string][]*types.Node{},
			ActionExecutor: actionInfos[nodeInfo.ActionExecutor],
		}

		//fmt.Println(node)
		if node.Action.Info().Inputs != nil {
			for _, input := range node.Action.Info().Inputs {
				node.Inputs[input.Name] = []*types.Connection{}
			}
		}

		if node.Action.Info().Outputs != nil {
			for _, output := range node.Action.Info().Outputs {
				node.Outputs[output.Name] = []*types.Connection{}
			}
		}
		nodes[nodeInfo.Id] = node
	}
	// 遍历节点输入列表，构建节点连接关系
	for _, nodeInfo := range spec.Nodes {
		node := nodes[nodeInfo.Id]

		for name, connections := range nodeInfo.SubActions {
			node.SubNodes[name] = make([]*types.Node, len(connections.InputInfos))
			for i, inputInfo := range connections.InputInfos {
				node.SubNodes[name][i] = nodes[inputInfo.NodeId]
			}
		}

		for name, input := range nodeInfo.Inputs {
			for _, inputInfo := range input.InputInfos {
				//inputInfo.NodeId
				if inputInfo.NodeId == types.VariableNodeId {
					connection := &types.Connection{TargetNode: node, Variable: wf.Variables[inputInfo.OutputName]}
					node.Inputs[name] = append(node.Inputs[name], connection)
				} else {
					preNode := nodes[inputInfo.NodeId]
					connectionInfo := &types.ConnectionInfo{
						OutputID:   preNode.ID,
						OutputName: inputInfo.OutputName,
						InputID:    node.ID,
						InputName:  name,
					}
					connection := &types.Connection{
						SourceNode: preNode,
						TargetNode: node,
						Info:       connectionInfo,
						DataChan:   make(chan *services.ActionData, 100000),
					}
					node.Inputs[name] = append(node.Inputs[name], connection)
					preNode.Outputs[inputInfo.OutputName] = append(preNode.Outputs[inputInfo.OutputName], connection)
				}
			}
		}

	}

	if spec.Outputs != nil && len(spec.Outputs) > 0 {
		outputNodeInfo := &models.WorkflowNodeInfo{
			Id:         types.OutputNodeId,
			ActionId:   types.OutputNodeId,
			ProjectId:  types.OutputNodeId,
			ActionName: types.OutputNodeId,
			Inputs:     map[string]*models.WorkflowConnections{},
		}

		wf.OutputNode = &types.Node{ID: types.OutputNodeId, Inputs: types.InputConnectionsMap{}, Info: outputNodeInfo}
		for _, input := range spec.Outputs {

			outputNodeInfo.Inputs[input.Name] = input

			for _, inputInfo := range input.InputInfos {
				if inputInfo.NodeId == types.VariableNodeId {
					connection := &types.Connection{TargetNode: wf.OutputNode, Variable: wf.Variables[inputInfo.OutputName]}
					wf.OutputNode.Inputs[input.Name] = append(wf.OutputNode.Inputs[input.Name], connection)
				} else {
					preNode := nodes[inputInfo.NodeId]
					connection := &types.Connection{SourceNode: preNode, TargetNode: wf.OutputNode, DataChan: make(chan *services.ActionData, 100000)}
					wf.OutputNode.Inputs[input.Name] = append(wf.OutputNode.Inputs[input.Name], connection)
					preNode.Outputs[inputInfo.OutputName] = append(preNode.Outputs[inputInfo.OutputName], connection)
				}
			}
		}
	}

	wf.Nodes = nodes

	wf.Init()

	return wf, nil
}

// GetWorkflowInfoByID 根据ID获取WorkflowInfo
func GetWorkflowInfoByID(id string) (*models.WorkflowInfo, error) {
	wfi, err := _apiClient.GetWorkflowById(id)
	if err != nil {
		return nil, err
	}
	return wfi, nil
}

var (
	workflowMap = map[string]*types.Workflow{}
)

// GetWorkflowByID 根据ID构建Workflow
func GetWorkflowByID(id string) (*types.Workflow, error) {

	if wf, ok := workflowMap[id]; ok {
		return wf, nil
	}

	wf, err := New(id)

	if err != nil {
		return nil, err
	}

	workflowMap[id] = wf

	return wf, nil
}

// New 根据ID获取WorkflowInfo
func New(id string) (*types.Workflow, error) {
	workflowId := id
	//global.Logger.Info("connecting to api server")

	//client := apiClient.NewAPIClient(tiopsConfigs.ApiServerHost, tiopsConfigs.ApiServerGrpcPort)

	_logger := logger.GetDefaultLogger()

	//global.Logger.Info("connected to api server")
	//global.Logger.Info("getting workflow info")
	wfi, err := _apiClient.GetWorkflowById(workflowId)

	if err != nil {
		return nil, err
	}

	_logger.Debug("Get workflow info success:", workflowId)

	_, err = _apiClient.CreateExecutionRecord(&models.ExecutionRecord{
		XId:         tiopsConfigs.ExecutionId,
		ExecutionId: tiopsConfigs.ExecutionId,
		ProcessRecords: []*models.ProcessRecord{
			{
				RecordTime: utils.CurrentTimeStampMS(),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return buildWorkflow(wfi, _apiClient)
}

var currentWorkflow *types.Workflow

// Current 当前处理流程
func Current() (*types.Workflow, error) {
	if currentWorkflow != nil {
		return currentWorkflow, nil
	}
	wf, err := New(tiopsConfigs.WorkflowId)
	currentWorkflow = wf
	return wf, err
}

// Context 当前流程引擎上下文
func Context() *types.EngineContext {
	return &types.EngineContext{
		//WorkflowType: tiopsConfigs.WorkflowType,
		//WorkflowId:   tiopsConfigs.WorkflowId,
		//ExecutionId:  tiopsConfigs.ExecutionId,
		//RecordId:     tiopsConfigs.ExecutionId,
		Logger:    logger.GetDefaultLogger(),
		APIClient: _apiClient,
	}
}
