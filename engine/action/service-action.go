// @Title  service-action.go
// @Description  远程服务调用组件
// @Create  heyitong  2022/6/23 17:12
// @Update  heyitong  2022/6/23 17:12

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

// RemoteServiceAction 远程服务调用组件
type RemoteServiceAction struct {
	info     *models.ActionInfo
	client   *actionClient.RemoteActionClient
	node     *types.Node
	nodeInfo *models.WorkflowNodeInfo
}

func (a *RemoteServiceAction) GetRequiredResources(n *types.Node, stage int32) (*models.WorkflowResources, error) {


	if stage == 0 {
		var apps []*models.K8SApp
		nodeInfo := n.Info
		actionInfo := a.info

		if nodeInfo.ActionExecutor != "" {
			actionInfo = n.ActionExecutor
		}
		serviceName := tiopsConfigs.StandAloneActionServiceName(nodeInfo.ActionName, n.ID) // config.ActionServiceName(actionInfo.ProjectId)

		switch actionInfo.Source {
		case models.ActionSource_Buildin:
			fallthrough
		case models.ActionSource_FromService:
			return nil, nil
		}
		app := &models.K8SApp{
			Name:        serviceName,
			ActionId:    actionInfo.XId,
			Replica:     1,
			ServiceMode: models.ServiceMode_One,
		}
		//if actionInfo.Type == models.ActionType_WorkflowAction {
		//	if actionInfo.Func == "" {
		//		app.WorkContainers = []*models.K8SContainer{{
		//			Name:  serviceName,
		//			Image: tiopsConfigs.DefaultEngineName,
		//		}}
		//	}
		//}
		apps = append(apps, app)
		return &models.WorkflowResources{
			Apps: apps,
		}, nil
	}

	return nil, nil
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
	serviceName := getServiceName(a.info, node.Info)

	a.nodeInfo = node.Info

	//if nodeInfo.StandAlone {
	//	serviceName = tiopsConfigs.StandAloneActionServiceName(info.Name, nodeInfo.Id)
	//}

	//if _actionClientMap[serviceName] == nil {
	//	_actionClientMap[serviceName] = actionClient.NewRemoteActionClient(serviceName, tiopsConfigs.ActionServerPort)
	//}

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
				Service:   getServiceName(a.info, targetNode.Info),
				Action:    targetNode.Info.ActionName,
				NodeId:    targetNode.ID,
				InputName: inputName,
			})
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()

	_, err := a.client.RegisterActionNode(ctx, &services.RegisterActionNodeRequest{
		ActionName:      a.info.Name,
		NodeId:          node.Info.Id,
		ActionOptions:   node.Info.ActionOptions,
		NextActions:     allNextActions,
		ActionInfo:      a.info,
		InnerActionInfo: a.info.InnerActionInfo,
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

func (a *RemoteServiceAction) CallPullStream(request *types.ActionRequest, callback func(res *types.ActionResponse, err error) bool) error {
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

func getServiceName(actionInfo *models.ActionInfo, nodeInfo *models.WorkflowNodeInfo) string {
	//serviceName := tiopsConfigs.ActionServiceName(nodeInfo.ProjectId)

	if actionInfo.Source == models.ActionSource_FromService {
		return tiopsConfigs.SystemActionServiceName(actionInfo.ServiceName)
	}

	//if nodeInfo.StandAlone {
	serviceName := tiopsConfigs.StandAloneActionServiceName(nodeInfo.ActionName, nodeInfo.Id)
	//}

	return serviceName
}

func NewRemoteServiceAction(info *models.ActionInfo) types.Action {
	return &RemoteServiceAction{info: info}
}
