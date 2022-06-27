// @Title  workflow-action.go
// @Description  子流程组件
// @Create  heyitong  2022/6/23 17:13
// @Update  heyitong  2022/6/23 17:13

package action

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	actionClient "tiops/common/action-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/logger"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/utils"
	"tiops/engine/types"
)

// WorkflowAction 子流程组件
type WorkflowAction struct {
	//RemoteServiceAction
	info         *models.ActionInfo
	engineInfo   *models.ActionInfo
	node         *types.Node
	subflow      *types.Workflow
	engineClient *actionClient.RemoteActionClient
	engineName   string
	subflowId    string
}

func (a *WorkflowAction) CallDuplexStream(callback func(res *types.ActionResponse, err error) bool) (func(request *types.ActionRequest) error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()

	streamClient, err := a.engineClient.CallEngine(ctx)

	if err != nil {
		return nil, err
	}

	var lastRequests = map[int64]*types.ActionRequest{}
	var lastRequestsLocker = sync.Mutex{}

	sender := func(request *types.ActionRequest) error {
		lastRequestsLocker.Lock()
		defer lastRequestsLocker.Unlock()
		lastRequests[utils.Int64ListHash(request.TraceIds)] = request
		return streamClient.Send(&services.ActionRequest{
			Id:         request.ID,
			NodeId:     a.node.ID,
			ActionName: a.engineName,
			Inputs:     request.Inputs,
			TraceIds:   request.TraceIds,
		})
	}

	go func() {
		for {
			res, err := streamClient.Recv()
			if err != nil {
				callback(nil, err)
				continue
			}
			if res == nil {
				continue
			}

			lastRequestsLocker.Lock()

			traceHash := utils.Int64ListHash(res.TraceIds)

			req := lastRequests[traceHash]

			if req == nil {
				req = &types.ActionRequest{}
			} else {
				delete(lastRequests, traceHash)
			}

			lastRequestsLocker.Unlock()

			if callback(&types.ActionResponse{
				ID:      res.Id,
				Request: req,
				Outputs: res.Outputs,
				Done:    res.Done,
			}, nil) {
				break
			}
			if res.Done {
				break
			}
		}
	}()

	return sender, nil
}

func (a *WorkflowAction) Init(node *types.Node) error {

	logger.Info(node.ID)
	logger.Info(a.engineName)

	a.node = node

	serviceName := tiopsConfigs.StandAloneActionServiceName(a.info.Name, node.ID) // config.ActionServiceName(actionInfo.ProjectId)

	a.engineClient = actionClient.NewRemoteActionClient(serviceName, tiopsConfigs.ActionServerPort)

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

	_, err := a.engineClient.RunEngine(ctx, &services.RunEngineRequest{
		EngineName:    a.engineName,
		NodeId:        node.ID,
		WorkflowId:    a.subflowId,
		ActionOptions: node.Info.ActionOptions,
		NextActions:   allNextActions,
		ActionInfo:    a.info,
	})

	return err
}

func (a *WorkflowAction) Call(request *types.ActionRequest) (*types.ActionResponse, error) {
	//a.CallPullStream(request, func(res *types.ActionResponse, err error) bool {
	//	res
	//})
	return nil, errors.New("workflow action not support Call method")
}

func (a *WorkflowAction) CallPullStream(request *types.ActionRequest, callback func(res *types.ActionResponse, err error) bool) error {
	return nil
}

func (a *WorkflowAction) Info() *models.ActionInfo {
	return a.info
}

func (a *WorkflowAction) Control(ctrl types.ActionControl, args map[string]string) error {
	switch ctrl {
	case types.ActionControlStream:
		ctx, _ := context.WithTimeout(context.Background(), time.Second*60)
		p, err := a.engineClient.PushMessage(ctx)
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

func (a *WorkflowAction) Type() types.ActionType {
	return "workflow-action"
}

func (a *WorkflowAction) Copy() types.Action {
	return NewWorkflowAction(a.info)
}

func (a *WorkflowAction) GetRequiredResources(n *types.Node, stage int32) (*models.WorkflowResources, error) {

	// 	subflow, err := workflow.New(info.Path)

	nodeInfo := n.Info

	serviceName := tiopsConfigs.StandAloneActionServiceName(nodeInfo.ActionName, n.ID) // config.ActionServiceName(actionInfo.ProjectId)

	if a.engineClient == nil {
		a.engineClient = actionClient.NewRemoteActionClient(serviceName, tiopsConfigs.ActionServerPort)
	}

	if stage == 0 {
		var apps []*models.K8SApp
		actionInfo := a.info

		//engineImage := tiopsConfigs.BuildinEngineImage
		//
		//if a.engineInfo != nil {
		//	engineImage =
		//}

		app := &models.K8SApp{
			Name:        serviceName,
			ActionId:    actionInfo.XId,
			Replica:     1,
			ServiceMode: models.ServiceMode_One,
			//WorkContainers: []*models.K8SContainer{
			//	{
			//		Name:  serviceName,
			//		Image: tiopsConfigs.DefaultEngineName,
			//	},
			//},
		}

		apps = append(apps, app)
		return &models.WorkflowResources{
			Apps: apps,
		}, nil
	} else if stage > 0 {
		resources, err := a.engineClient.GetRequiredResources(context.TODO(),
			&services.RequiredResourcesRequest{
				EngineName: a.engineName,
				WorkflowId: a.subflowId,
				Stage:      stage - 1,
			})
		if err != nil {
			return nil, err
		}
		return resources, err
	}
	return nil, nil
}

func NewWorkflowAction(info *models.ActionInfo) types.Action {
	engineName := tiopsConfigs.DefaultEngineName
	if info.EngineInfo != nil {
		engineName = info.EngineInfo.Name
	}
	logger.Info(info.Path)
	return &WorkflowAction{
		info:       info,
		engineInfo: info.EngineInfo,
		engineName: engineName,
		subflowId:  info.Func,
	}
}
