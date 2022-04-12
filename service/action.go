package service

import (
	"context"
	"errors"
	"fmt"
	actionTypes "github.com/qhzhyt/tiops-go-sdk/action/types"
	services "tiops/common/services"
	"tiops/common/stores"
)

func (a *actionServer) GetActionStatus(ctx context.Context, request *services.ActionStatusRequest) (*services.ActionStatus, error) {
	panic("implement me")
}

func (a *actionServer) CallHttpAction(ctx context.Context, request *services.HttpRequest) (*services.HttpResponse, error) {
	panic("implement me")
}

func (a *actionServer) CallAction(ctx context.Context, request *services.ActionRequest) (*services.ActionResponse, error) {
	actionName := request.ActionName
	a.Logger.Info(inputLog(actionName, request.Inputs))
	inputDataMap := actionTypes.TransActionDataMap(request.Inputs, a.actionInfoMap[actionName].Inputs)
	actionNodeContext := a.actionNodeContextMap[request.NodeId]

	var result actionTypes.ActionDataBatch
	if a.actions[actionName] != nil {
		result = a.actions[actionName].CallBatch(
			&actionTypes.BatchRequestContext{
				ActionNodeContext: actionNodeContext,
				//NodeId:            request.NodeId,
				Inputs: inputDataMap,
				//Store:             a.nodeStores[request.NodeId],
				//ActionOptions:     a.getActionOptions(request.NodeId),
			})
	} else {
		return nil, errors.New("action " + actionName + " not found")
	}

	outputs := actionTypes.ToServiceActionDataMap(request.Id, request.TraceId, result, a.actionInfoMap[actionName].Outputs)
	a.Logger.Info(outputLog(actionName, outputs))

	return &services.ActionResponse{Id: request.Id, Outputs: outputs, Done: actionNodeContext.HasDone(), TraceId: request.TraceId}, nil
}

func (a *actionServer) CallActionPullStream(request *services.ActionRequest, server services.ActionsService_CallActionPullStreamServer) error {
	actionName := request.ActionName
	a.Logger.Info(inputLog(actionName, request.Inputs))
	inputDataMap := actionTypes.TransActionDataMap(request.Inputs, a.actionInfoMap[actionName].Inputs)
	actionNodeContext := a.actionNodeContextMap[request.NodeId]

	//var result ActionDataBatch

	if a.actions[actionName] != nil {
		return a.actions[actionName].CallPullStream(&actionTypes.StreamRequestContext{
			BatchRequestContext: actionTypes.BatchRequestContext{
				ActionNodeContext: actionNodeContext,
				//Store:         a.nodeStores[request.NodeId],
				//NodeId:        request.NodeId,
				Inputs: inputDataMap,
				//ActionOptions: a.getActionOptions(request.NodeId),
				//done: false,
			},
			Push: func(data actionTypes.ActionDataBatch) error {
				outputs := actionTypes.ToServiceActionDataMap(request.Id, request.TraceId, data, a.actionInfoMap[actionName].Outputs)
				a.Logger.Info(outputLog(actionName, outputs))

				res := &services.ActionResponse{Id: request.Id, Outputs: outputs, Done: actionNodeContext.HasDone(), TraceId: request.TraceId}

				return server.Send(res)
			},
		})
	}

	return errors.New("action " + actionName + " not found")

}

func (a *actionServer) RegisterActionNode(ctx context.Context, request *services.RegisterActionNodeRequest) (*services.StatusResponse, error) {

	actionName := request.ActionName

	//a.actionNodeOptionsMap[request.NodeId] = request.ActionOptions
	a.Logger.Info(fmt.Sprint("register node ", request.NodeId, " with options ", request.ActionOptions))

	//a.nodeStores[request.NodeId] = nodeStore

	if action := a.actions[actionName]; action != nil {

		if _, ok := a.actionInfoMap[actionName]; !ok {
			actionInfo := request.ActionInfo

			a.actionInfoMap[actionInfo.Name] = actionInfo
			outputs := make([]string, len(actionInfo.Outputs))
			for i, output := range actionInfo.Outputs {
				outputs[i] = output.Name
			}
			inputs := make([]string, len(actionInfo.Inputs))
			for i, input := range actionInfo.Inputs {
				inputs[i] = input.Name
			}
			actionContext := &actionTypes.ActionContext{Logger: a.Logger, Info: actionInfo, InputNames: inputs, OutputNames: outputs}
			a.actionContextMap[actionInfo.Name] = actionContext
			action.Init(&actionTypes.InitContext{ActionContext: actionContext})
		}

		nodeStore := stores.NewActionNodeStore(request.NodeId)

		actionNodeContext := &actionTypes.ActionNodeContext{
			ActionContext: a.actionContextMap[actionName],
			//done:          false,
			ActionOptions: request.ActionOptions,
			NodeId:        request.NodeId,
			Store:         nodeStore,
		}

		a.actionNodeContextMap[request.NodeId] = actionNodeContext

		err := action.RegisterNode(&actionTypes.NodeRegisterContext{
			ActionNodeContext: actionNodeContext,
			NextActions:       request.NextActions,
		},
		)
		if err != nil {
			return &services.StatusResponse{Status: services.Status_Failed}, err
		}
	} else {
		return nil, errors.New("action " + actionName + " not found")
	}

	return &services.StatusResponse{Status: services.Status_Ok}, nil
}
