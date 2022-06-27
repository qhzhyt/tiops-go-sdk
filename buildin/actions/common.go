// @Title  common.go
// @Description  内置组件封装为通用组件
// @Create  heyitong  2022/6/23 15:34
// @Update  heyitong  2022/6/23 15:34

package actions

import (
	"encoding/json"
	actionTypes "tiops/action/types"
	"tiops/common/logger"
	"tiops/common/models"
	"tiops/engine/types"
	engineTypes "tiops/engine/types"
)

type BuildinActionBatchProcessor func(ctx *actionTypes.BatchRequestContext) (actionTypes.ActionDataBatch, error)

type BuildinActionPullStreamProcessor func(ctx *actionTypes.StreamRequestContext) error

// BuildinAction 内置组件数据结构定义
type BuildinAction struct {
	*actionTypes.ActionContext
	CallBatchFunc      BuildinActionBatchProcessor
	CallPullStreamFunc BuildinActionPullStreamProcessor
	ActionInfo         *models.ActionInfo
	NodeInfo           *models.WorkflowNodeInfo
	nodeCtx            *actionTypes.ActionNodeContext
}

func (a *BuildinAction) CallDuplexStream(callback func(res *types.ActionResponse, err error) bool) (func(request *types.ActionRequest) error, error) {
	//TODO implement me
	panic("implement me")
}

func (a *BuildinAction) GetRequiredResources(node *types.Node, stage int32) (*models.WorkflowResources, error) {
	return nil, nil
}

func (a *BuildinAction) Init(node *engineTypes.Node) error {
	a.NodeInfo = node.Info
	a.nodeCtx = &actionTypes.ActionNodeContext{
		ActionContext: a.ActionContext,
		Store:         nil,
		NodeId:        node.ID,
		ActionOptions: node.Info.ActionOptions,
	}
	return nil
}

func (a *BuildinAction) Call(request *engineTypes.ActionRequest) (*engineTypes.ActionResponse, error) {
	inputDataMap := actionTypes.TransActionDataMap(request.Inputs, a.ActionInfo.Inputs)
	//a.Logger.Error(inputDataMap)
	result, err := a.CallBatchFunc(
		&actionTypes.BatchRequestContext{
			ActionNodeContext: a.nodeCtx,
			Inputs:            inputDataMap,
		})
	if err != nil {
		return nil, err
	}
	//a.Logger.Error(a.ActionInfo)
	outputs := actionTypes.ToServiceActionDataMap(request.ID, request.TraceIds, result, a.ActionInfo.Outputs)
	//a.Logger.Error(outputs)
	return &engineTypes.ActionResponse{
		ID:      request.ID,
		Outputs: outputs,
		Request: request,
		Done:    a.nodeCtx.HasDone(),
	}, nil
}

func (a *BuildinAction) CallPullStream(request *engineTypes.ActionRequest, callback func(res *engineTypes.ActionResponse, err error) bool) error {
	inputDataMap := actionTypes.TransActionDataMap(request.Inputs, a.ActionInfo.Inputs)

	ctx := &actionTypes.StreamRequestContext{
		BatchRequestContext: actionTypes.BatchRequestContext{
			ActionNodeContext: a.nodeCtx,
			Inputs:            inputDataMap,
		},
	}

	ctx.Push = func(data actionTypes.ActionDataBatch) error {
		outputs := actionTypes.ToServiceActionDataMap(request.ID, request.TraceIds, data, a.ActionInfo.Outputs)
		callback(&engineTypes.ActionResponse{
			ID:      request.ID,
			Outputs: outputs,
			Request: request,
			Done:    ctx.HasDone(),
		}, nil)
		return nil
	}

	err := a.CallPullStreamFunc(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (a *BuildinAction) Info() *models.ActionInfo {
	return a.ActionInfo
}

func (a *BuildinAction) Control(ctrl engineTypes.ActionControl, args map[string]string) error {
	panic("implement me")
}

func (a *BuildinAction) Type() engineTypes.ActionType {
	panic("implement me")
}

func (a *BuildinAction) Copy() engineTypes.Action {
	return NewBuildinAction(a.ActionInfo)
}

func NewBuildinAction(actionInfo *models.ActionInfo) engineTypes.Action {

	var outputNames, inputNames []string

	for _, output := range actionInfo.Outputs {
		outputNames = append(outputNames, output.Name)
	}
	for _, input := range actionInfo.Inputs {
		inputNames = append(inputNames, input.Name)
	}

	return &BuildinAction{
		CallBatchFunc:      batchProcessorMap[actionInfo.XId],
		CallPullStreamFunc: pullStreamProcessorMap[actionInfo.XId],
		ActionInfo:         actionInfo,
		ActionContext: &actionTypes.ActionContext{
			Logger:      logger.NewActionLogger(actionInfo.XId),
			Info:        actionInfo,
			InputNames:  inputNames,
			OutputNames: outputNames,
		},
	}
}

func JsonProcess(input string, processor func(map[string]interface{}) (map[string]interface{}, error)) (string, error) {
	inputMap := map[string]interface{}{}

	err := json.Unmarshal([]byte(input), &inputMap)

	if err != nil {
		return "", err
	}

	outputMap, err := processor(inputMap)

	if err != nil {
		return "", err
	}

	outputBytes, err := json.Marshal(outputMap)

	if err != nil {
		return "", err
	}

	return string(outputBytes), nil

}
