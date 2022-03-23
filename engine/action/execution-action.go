package action

import (
	"context"
	"time"
	actionClient "tiops/common/action-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/logger"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/engine/types"
)

type ExecutionAction struct {
	//RemoteServiceAction
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
		client:              a.client,
		node:                a.node,
		nodeInfo:            a.nodeInfo,
		innerActionInfo:     a.innerActionInfo,
		executionActionInfo: a.executionActionInfo,
	}
}

func (a *ExecutionAction) Init(node *types.Node) error {
	// 初始化时再创建 action 客户端
	serviceName := getServiceNameByAction(node.Info, a.innerActionInfo)

	//if nodeInfo.StandAlone {
	//	serviceName = tiopsConfigs.StandAloneActionServiceName(info.Name, nodeInfo.Id)
	//}

	if _actionClientMap[serviceName] == nil {
		_actionClientMap[serviceName] = actionClient.NewRemoteActionClient(serviceName, tiopsConfigs.ActionServerPort)
	}

	//logger.Info("node:", node)
	//logger.Info("innerActionInfo:", a.innerActionInfo)
	//logger.Info("executorActionInfo:", a.executionActionInfo)
	//logger.Info("service name:", serviceName)

	a.client = actionClient.NewRemoteActionClient(serviceName, tiopsConfigs.ActionServerPort)

	a.node = node
	//node.Outputs
	var allNextActions []*services.NextActions
	for s, connections := range node.Outputs {
		nextActions := &services.NextActions{OutputName: s, Actions: []*services.ServiceAndAction{}}
		allNextActions = append(allNextActions, nextActions)
		for _, connection := range connections {
			targetNode := connection.TargetNode
			var inputName string
			for in, inputConnections := range targetNode.Inputs {
				for _, inputConnection := range inputConnections {
					if inputConnection == connection {
						inputName = in
					}
				}
			}
			//fmt.Println(targetNode.Inputs)
			nextActions.Actions = append(nextActions.Actions, &services.ServiceAndAction{
				Service:   getServiceName(targetNode.Info),
				Action:    targetNode.Info.ActionName,
				NodeId:    targetNode.ID,
				InputName: inputName,
			})
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()
	_, err := a.client.RegisterActionNode(ctx, &services.RegisterActionNodeRequest{
		ActionName:      a.executionActionInfo.Name,
		NodeId:          node.Info.Id,
		ActionOptions:   node.Info.ActionOptions,
		NextActions:     allNextActions,
		InnerActionInfo: a.innerActionInfo,
		ActionInfo:      a.executionActionInfo,
	})
	logger.Error(err)
	return err
}

func (a *ExecutionAction) Info() *models.ActionInfo {
	return a.innerActionInfo
}

func (a *ExecutionAction) Type() types.ActionType {
	return RemoteService
}

func (a *ExecutionAction) Call(request *types.ActionRequest) (*types.ActionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()
	res, err := a.client.CallAction(ctx, &services.ActionRequest{Id: request.ID, ActionName: a.executionActionInfo.Name, NodeId: a.nodeInfo.Id, Inputs: request.Inputs})

	if err != nil {
		return nil, err
	}

	return &types.ActionResponse{ID: res.Id, Outputs: res.Outputs, Done: res.Done}, nil
}

func NewExecutionAction(actionInfo, executionActionInfo *models.ActionInfo, nodeInfo *models.WorkflowNodeInfo) types.Action {
	return &ExecutionAction{
		nodeInfo:            nodeInfo,
		innerActionInfo:     actionInfo,
		executionActionInfo: executionActionInfo,
	}
}
