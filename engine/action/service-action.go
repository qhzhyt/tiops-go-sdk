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
}

const ActionTypeRemoteService = "RemoteService"

//func (a *RemoteServiceAction) Call(request *ActionRequest) (*ActionResponse, error) {
//	panic("implement me")
//}

func (a *RemoteServiceAction) Copy() types.Action {
	return &RemoteServiceAction{
		info:   a.info,
		client: a.client,
		node:   nil,
	}
}

func (a *RemoteServiceAction) Init(node *types.Node) error {
	a.node = node
	_, err := a.client.RegisterActionNode(context.TODO(), &services.RegisterActionNodeRequest{
		ActionName:    a.info.Name,
		NodeId:        node.Info.Id,
		ActionOptions: node.Info.ActionOptions,
	})
	return err
}

func (a *RemoteServiceAction) Info() *models.ActionInfo {
	return a.info
}

func (a *RemoteServiceAction) Type() types.ActionType {
	return ActionTypeRemoteService
}

func (a *RemoteServiceAction) Call(request *types.ActionRequest) (*types.ActionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)

	//var deadlineMs = flag.Int("deadline_ms", 20*1000, "Default deadline in milliseconds.")
	//ctx, cancel := context.WithTimeout(ctx, time.Duration(*deadlineMs) * time.Millisecond)
	//fmt.Println("Call", a.info.Name, arguments)

	defer cancel()
	res, err := a.client.CallAction(ctx, &services.ActionRequest{Id: request.ID, ActionName: a.info.Name, NodeId: a.node.Info.Id, Inputs: request.Inputs})

	if err != nil {
		return nil, err
	}

	return &types.ActionResponse{ID: res.Id, Outputs: res.Outputs, Done: res.Done}, nil

	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//fmt.Println("result", a.info.Name, res.Outputs)
	//return res.Outputs, err
}

var _actionClientMap = map[string]*actionClient.RemoteActionClient{}

func NewRemoteServiceAction(info *models.ActionInfo) types.Action {

	if _actionClientMap[info.ProjectId] == nil{
		_actionClientMap[info.ProjectId] = actionClient.NewRemoteActionClient(tiopsConfigs.ActionServiceName(info.ProjectId), tiopsConfigs.ActionServerPort)
	}
	return &RemoteServiceAction{info: info, client: _actionClientMap[info.ProjectId]}
}
