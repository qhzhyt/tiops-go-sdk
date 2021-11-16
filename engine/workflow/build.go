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

func loadActionInfos(nodeInfos []*models.WorkflowNodeInfo, client *apiClient.APIClient) (map[string]*models.ActionInfo, error) {
	//return w.Packages[pId].GetAction(aName)
	actionNameSet := map[string]bool{}
	for _, nodeInfo := range nodeInfos {
		actionNameSet[nodeInfo.ActionId] = true
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
	return result, err
}

// buildWorkflow 根据 models.WorkflowInfo 构建相应的 Workflow 对象
func buildWorkflow(wi *models.WorkflowInfo, client *apiClient.APIClient) (*types.Workflow, error) {

	//wi.Spec.Nodes

	wf := types.NewWorkflow(wi)

	wf.ApiClient = client

	actionInfos, err := loadActionInfos(wi.Spec.Nodes, client)

	if err != nil {
		return nil, err
	}

	for _, nodeInfo := range wi.Spec.Nodes {
		wf.Actions[nodeInfo.Id] = action.New(actionInfos[nodeInfo.ActionId], nodeInfo)
	}

	spec := wi.Spec
	//fmt.Println(wi.Constants)
	nodeInfos := map[string]*models.WorkflowNodeInfo{}
	nodes := map[string]*types.Node{}
	for _, _variable := range spec.Variables {
		wf.Variables[_variable.Name] = variable.New(_variable.Name, _variable.Type, _variable.ValueType, _variable.Value)
	}

	inputNode := &types.Node{ID: types.InputNodeId, Outputs: types.OutputConnectionsMap{}}
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
			ID:      nodeInfo.Id,
			Action:  wf.GetAction(nodeInfo.Id).Copy(),
			Inputs:  types.InputConnectionsMap{},
			Outputs: types.OutputConnectionsMap{},
			Info:    nodeInfo,
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
	for _, nodeInfo := range spec.Nodes {
		node := nodes[nodeInfo.Id]
		for name, input := range nodeInfo.Inputs {
			for _, inputInfo := range input.InputInfos {
				//inputInfo.NodeId
				if inputInfo.NodeId == types.VariableNodeId {
					connection := &types.Connection{TargetNode: node, Variable: wf.Variables[inputInfo.OutputName]}
					node.Inputs[name] = append(node.Inputs[name], connection)
				} else {
					preNode := nodes[inputInfo.NodeId]
					connection := &types.Connection{SourceNode: preNode, TargetNode: node, DataChan: make(chan *services.ActionData, 100000)}
					node.Inputs[name] = append(node.Inputs[name], connection)
					preNode.Outputs[inputInfo.OutputName] = append(preNode.Outputs[inputInfo.OutputName], connection)
				}
			}
		}
	}

	for _, nodeInfo := range spec.Nodes {
		node := nodes[nodeInfo.Id]
		toDelete := true
		for _, input := range node.Inputs {
			if len(input) > 0 {
				toDelete = false
				break
			}
		}
		if toDelete {
			for _, output := range node.Outputs {
				if len(output) > 0 {
					toDelete = false
					break
				}
			}
		}
		if toDelete {
			delete(nodes, nodeInfo.Id)
		}
	}

	if spec.Outputs != nil && len(spec.Outputs) > 0 {
		wf.OutputNode = &types.Node{ID: types.OutputNodeId, Inputs: types.InputConnectionsMap{}}
		for _, input := range spec.Outputs {
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

	_logger.Info("get workflow info success")

	_, err = _apiClient.CreateExecutionRecord(&models.ExecutionRecord{
		XId:         tiopsConfigs.ExecutionID,
		ExecutionId: tiopsConfigs.ExecutionID,
		ProcessRecords: []*models.ProcessRecord{
			{
				RecordTime: utils.CurrentTimeStampMS(),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	//workflowSpec := wfi.Spec
	//packages := map[string]*workflow.Package{}
	//for _, info := range workflowSpec.Nodes {
	//	if packages[info.ProjectId] == nil {
	//		packages[info.ProjectId] = workflow.BuildPackage(client.GetProjectByID(info.ProjectId))
	//	}
	//}

	return buildWorkflow(wfi, _apiClient)
}

func Current() (*types.Workflow, error) {
	return New(tiopsConfigs.WorkflowId)
}

func Context() *types.WorkflowContext {
	return &types.WorkflowContext{
		WorkflowType: tiopsConfigs.WorkflowType,
		WorkflowId:   tiopsConfigs.WorkflowId,
		ExecutionId:  tiopsConfigs.ExecutionID,
		RecordId:     tiopsConfigs.ExecutionID,
		Logger:       logger.GetDefaultLogger(),
		APIClient:    _apiClient,
	}
}
