package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	actionTypes "tiops/action/types"
	"tiops/common/logger"
	"tiops/common/services"
	"tiops/common/stores"
)

func (a *actionServer) GetActionStatus(ctx context.Context, request *services.ActionStatusRequest) (*services.ActionStatus, error) {
	actionName := request.ActionName

	if a.actions[actionName] != nil {
		result := a.actions[actionName].Status(
			a.actionNodeContextMap[request.NodeId])
		return &services.ActionStatus{
			ProcessedCount:        result.ProcessedCount,
			RestCount: result.RestCount,
			Done:     result.Done,
			Message: result.Message,
			Extra:        result.Extra,
		}, nil
	} else {
		return nil, errors.New("Action " + actionName + " not found")
	}
}

func (a *actionServer) CallHttpAction(ctx context.Context, request *services.HttpRequest) (*services.HttpResponse, error) {

	//httpContext

	actionName := request.ActionName

	//defer func() {
	//	if err0 := recover(); err0 != nil {
	//		switch err1 := err0.(type) {
	//		case error:
	//			err = err1
	//			a.Logger.Error(errorLog(actionName, request.Inputs, err1))
	//		}
	//	} else {
	//		a.Logger.Info(successLog(actionName, request.Inputs, res.Outputs))
	//	}
	//}()

	if a.actions[actionName] != nil {

		if a.actionNodeContextMap[request.ContextId] == nil {
			a.actionNodeContextMap[request.ContextId] = &actionTypes.ActionNodeContext{
				ActionContext:   nil,
				Store:           stores.NewActionNodeStore(request.Id),
				NodeId:          request.ContextId,
				ActionOptions:   request.Query,
			}
		}

		//actionNodeContext :=

		result := a.actions[actionName].CallHttp(
			&actionTypes.HttpRequestContext{
				ActionNodeContext: a.actionNodeContextMap[request.ContextId],
				Method:        request.Method,
				Path:          request.Path,
				Header:        services.ServiceHeadersToHttpHeader(request.Headers),
				Query:         request.Query,
				Body:          request.Body,
			})
		return &services.HttpResponse{
			Code:        result.Status,
			ContentType: result.ContentType,
			Headers:     services.HttpHeaderToServiceHeaders(result.Header),
			Body:        result.Body,
		}, nil
	} else {
		return nil, errors.New("Action " + actionName + " not found")
	}

}

func (a *actionServer) CallAction(ctx context.Context, request *services.ActionRequest) (res *services.ActionResponse, err error) {
	actionName := request.ActionName

	defer func() {
		if err0 := recover(); err0 != nil {
			switch err1 := err0.(type) {
			case error:
				err = err1
				a.Logger.Error(errorLog(actionName, request.Inputs, err1))
			}
		} else {
			a.Logger.Info(successLog(actionName, request.Inputs, res.Outputs))
		}
	}()

	inputDataMap := actionTypes.TransActionDataMap(request.Inputs, a.actionInfoMap[actionName].Inputs)
	actionNodeContext := a.actionNodeContextMap[request.NodeId]

	var result actionTypes.ActionDataBatch
	if a.actions[actionName] != nil {
		result = a.actions[actionName].CallBatch(
			&actionTypes.BatchRequestContext{
				ActionNodeContext: actionNodeContext,
				Inputs:            inputDataMap,
			})
	} else {
		return nil, errors.New("Action " + actionName + " not found")
	}

	outputs := actionTypes.ToServiceActionDataMap(request.Id, request.TraceId, result, a.actionInfoMap[actionName].Outputs)

	return &services.ActionResponse{Id: request.Id, Outputs: outputs, Done: actionNodeContext.HasDone(), TraceId: request.TraceId}, nil
}

func (a *actionServer) CallActionPullStream(request *services.ActionRequest, server services.ActionsService_CallActionPullStreamServer) (err error) {
	actionName := request.ActionName
	defer func() {
		if err0 := recover(); err0 != nil {
			switch err1 := err0.(type) {
			case error:
				err = err1
				a.Logger.Error(errorLog(actionName, request.Inputs, err1))
			}
		}
	}()

	inputDataMap := actionTypes.TransActionDataMap(request.Inputs, a.actionInfoMap[actionName].Inputs)
	actionNodeContext := a.actionNodeContextMap[request.NodeId]

	if a.actions[actionName] != nil {
		return a.actions[actionName].CallPullStream(&actionTypes.StreamRequestContext{
			BatchRequestContext: actionTypes.BatchRequestContext{
				ActionNodeContext: actionNodeContext,
				Inputs:            inputDataMap,
			},
			Push: func(data actionTypes.ActionDataBatch) error {
				outputs := actionTypes.ToServiceActionDataMap(request.Id, request.TraceId, data, a.actionInfoMap[actionName].Outputs)

				a.Logger.Info(successLog(actionName, request.Inputs, outputs))

				res := &services.ActionResponse{Id: request.Id, Outputs: outputs, Done: actionNodeContext.HasDone(), TraceId: request.TraceId}

				return server.Send(res)
			},
		})
	}

	return errors.New("Action " + actionName + " not found")

}

func (a *actionServer) RegisterActionNode(ctx context.Context, request *services.RegisterActionNodeRequest) (res *services.StatusResponse, err error) {

	actionName := request.ActionName

	defer func() {
		optionsData, _ := json.Marshal(request.ActionOptions)
		if err0 := recover(); err0 != nil {
			switch err1 := err0.(type) {
			case error:
				err = err1
			}
			a.Logger.Error(fmt.Sprint("Register node ", request.NodeId[:8], " for action ", actionName, " with options: ", string(optionsData), " get error: ", err0))
		} else {
			a.Logger.Info(fmt.Sprint("Register node ", request.NodeId[:8], " for action ", actionName, " with options: ", string(optionsData)))
		}
	}()

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
			actionContext := &actionTypes.ActionContext{Logger: logger.NewActionLogger(actionInfo.XId), Info: actionInfo, InputNames: inputs, OutputNames: outputs}
			a.actionContextMap[actionInfo.Name] = actionContext
			action.Init(&actionTypes.InitContext{ActionContext: actionContext})
		}

		nodeStore := stores.NewActionNodeStore(request.NodeId)

		actionNodeContext := &actionTypes.ActionNodeContext{
			ActionContext: a.actionContextMap[actionName],
			//done:          false,
			ActionOptions:   request.ActionOptions,
			NodeId:          request.NodeId,
			Store:           nodeStore,
			InnerActionInfo: request.InnerActionInfo,
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
		return nil, errors.New("Action " + actionName + " not found")
	}

	return &services.StatusResponse{Status: services.Status_Ok}, nil
}
