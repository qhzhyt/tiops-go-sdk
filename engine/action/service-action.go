package action

import (
	"context"
	"time"
	actionClient "tiops/common/action-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/engine/types"
)

type RemoteServiceAction struct {
	info   *models.ActionInfo
	client *actionClient.RemoteActionClient
	node   *types.Node
	nodeInfo *models.WorkflowNodeInfo
}

const RemoteService = "RemoteService"

//func (a *RemoteServiceAction) Call(request *ActionRequest) (*ActionResponse, error) {
//	panic("implement me")
//}

func (a *RemoteServiceAction) Copy() types.Action {
	return &RemoteServiceAction{
		info:   a.info,
		client: a.client,
		node:   nil,
		nodeInfo: a.nodeInfo,
	}
}

func (a *RemoteServiceAction) Init(node *types.Node) error {
	a.node = node
	//node.Outputs
	var allNextActions []*services.NextActions

	for s, connections := range node.Outputs {
		nextActions := &services.NextActions{OutputName: s, Actions: []*services.ServiceAndAction{}}
		allNextActions = append(allNextActions, nextActions)
		for _, connection := range connections {
			targetNode := connection.TargetNode
			nextActions.Actions = append(nextActions.Actions, &services.ServiceAndAction{
				Service: getServiceName(targetNode.Info),
				Action:  targetNode.Info.ActionName,
				NodeId: targetNode.ID,
			})
		}
	}

	_, err := a.client.RegisterActionNode(context.TODO(), &services.RegisterActionNodeRequest{
		ActionName:    a.info.Name,
		NodeId:        node.Info.Id,
		ActionOptions: node.Info.ActionOptions,
		NextActions: allNextActions,
	})
	return err
}

func (a *RemoteServiceAction) Info() *models.ActionInfo {
	return a.info
}

func (a *RemoteServiceAction) Type() types.ActionType {
	return RemoteService
}

func (a *RemoteServiceAction) Call(request *types.ActionRequest) (*types.ActionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()
	res, err := a.client.CallAction(ctx, &services.ActionRequest{Id: request.ID, ActionName: a.info.Name, NodeId: a.nodeInfo.Id, Inputs: request.Inputs})

	if err != nil {
		return nil, err
	}
	return &types.ActionResponse{ID: res.Id, Outputs: res.Outputs, Done: res.Done}, nil
}

var _actionClientMap = map[string]*actionClient.RemoteActionClient{}

func getServiceName(nodeInfo *models.WorkflowNodeInfo)  string{
	serviceName := tiopsConfigs.ActionServiceName(nodeInfo.ProjectId)

	if nodeInfo.StandAlone {
		serviceName = tiopsConfigs.StandAloneActionServiceName(nodeInfo.ActionName, nodeInfo.Id)
	}

	return serviceName
}

func NewRemoteServiceAction(info *models.ActionInfo, nodeInfo *models.WorkflowNodeInfo) types.Action {

	serviceName := getServiceName(nodeInfo)

	//if nodeInfo.StandAlone {
	//	serviceName = tiopsConfigs.StandAloneActionServiceName(info.Name, nodeInfo.Id)
	//}


	if _actionClientMap[serviceName] == nil{
		_actionClientMap[serviceName] = actionClient.NewRemoteActionClient(serviceName, tiopsConfigs.ActionServerPort)
	}

	return &RemoteServiceAction{info: info, client: _actionClientMap[serviceName], nodeInfo: nodeInfo}
}
