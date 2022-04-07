package action

import (
	"context"
	"fmt"
	"time"
	actionClient "tiops/common/action-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/engine/types"
)

type RemoteServiceAction struct {
	info     *models.ActionInfo
	client   *actionClient.RemoteActionClient
	node     *types.Node
	nodeInfo *models.WorkflowNodeInfo
}

func (a *RemoteServiceAction) Control(ctrl types.ActionControl, args map[string]string) error {
	switch ctrl {
	case types.ActionControlStream:
		ctx, _ := context.WithTimeout(context.Background(), time.Second*60)
		p, err := a.client.PushMessage(ctx)
		if err != nil {
			return err
		}
		fmt.Println(ctrl)
		return p.Send(&services.ActionMessage{
			NodeId:  a.node.ID,
			Type:    services.ActionMessageType_StreamCmd,
			Message: args["cmd"],
			Header: map[string]string{
				"actionName": a.info.GetName(),
			},
			Data: nil,
		})
	}
	return nil
}

const RemoteService = "RemoteService"

//func (a *RemoteServiceAction) Call(request *ActionRequest) (*ActionResponse, error) {
//	panic("implement me")
//}

func (a *RemoteServiceAction) Copy() types.Action {
	return &RemoteServiceAction{
		info:     a.info,
		client:   a.client,
		node:     nil,
		nodeInfo: a.nodeInfo,
	}
}

func (a *RemoteServiceAction) Init(node *types.Node) error {
	// 初始化时再创建 action 客户端
	serviceName := getServiceName(node.Info)

	//if nodeInfo.StandAlone {
	//	serviceName = tiopsConfigs.StandAloneActionServiceName(info.Name, nodeInfo.Id)
	//}

	if _actionClientMap[serviceName] == nil {
		_actionClientMap[serviceName] = actionClient.NewRemoteActionClient(serviceName, tiopsConfigs.ActionServerPort)
	}

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
		ActionName:    a.info.Name,
		NodeId:        node.Info.Id,
		ActionOptions: node.Info.ActionOptions,
		NextActions:   allNextActions,
		ActionInfo:    a.info,
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

func (a *RemoteServiceAction) CallStream(request *types.ActionRequest, callback func(res *types.ActionResponse, err error) bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()
	stream, err := a.client.CallActionPullStream(ctx, &services.ActionRequest{Id: request.ID, ActionName: a.info.Name, NodeId: a.nodeInfo.Id, Inputs: request.Inputs})

	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			callback(nil, err)
			continue
		}
		if res == nil {
			return nil
		}
		if !callback(&types.ActionResponse{ID: res.Id, Outputs: res.Outputs, Done: res.Done}, nil) {
			return err
		}
		if res.Done {
			break
		}
	}

	return nil
}

var _actionClientMap = map[string]*actionClient.RemoteActionClient{}

func getServiceName(nodeInfo *models.WorkflowNodeInfo) string {
	//serviceName := tiopsConfigs.ActionServiceName(nodeInfo.ProjectId)

	//if nodeInfo.StandAlone {
	serviceName := tiopsConfigs.StandAloneActionServiceName(nodeInfo.ActionName, nodeInfo.Id)
	//}

	return serviceName
}

func NewRemoteServiceAction(info *models.ActionInfo, nodeInfo *models.WorkflowNodeInfo) types.Action {
	return &RemoteServiceAction{info: info, nodeInfo: nodeInfo}
}
