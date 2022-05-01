package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tiops/buildin/engines"
	tiopsConfigs "tiops/common/config"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/utils"
	engineTypes "tiops/engine/types"
	"tiops/engine/workflow"
)

func (a *actionServer) runMainEngine() {
	engine := a.getCurrentEngine()
	_workflow, err := workflow.Current()
	_context := workflow.Context()

	if err != nil {
		_context.Error(err)
		utils.SleepAndExit(time.Second*6, 3)
	}

	_context.Options = _workflow.Info().EngineOptions

	engine.Init(_context)

	//if requiredResources != nil {
	//	_, err := a.apiClient.CreateOrUpdateWorkflowExecution(
	//		&models.WorkflowExecution{
	//			XId:              tiopsConfigs.ExecutionId,
	//			WorkflowResource: requiredResources,
	//		})
	//	_context.Error(err)
	//}

	engine.WaitForResources(_workflow)
	//engine := engines.NewBasicChanEngine(context)
	engine.Exec(_workflow)
}

func (a *actionServer) getCurrentEngine() engineTypes.WorkflowEngine {
	//a.Logger.Info("current engine: " + tiopsConfigs.EngineName)
	if a.engines[tiopsConfigs.EngineName] != nil {
		return a.engines[tiopsConfigs.EngineName]
	}
	for _, engine := range a.engines {
		return engine
	}
	return engines.NewBasicChanEngine()
}

func (a *actionServer) getEngine(name string) engineTypes.WorkflowEngine {
	//a.Logger.Info("current engine: " + tiopsConfigs.EngineName)
	if a.engines[name] != nil {
		return a.engines[name]
	}

	return a.getCurrentEngine()
}

func (a *actionServer) afterEngineExec() {
	record := a.getCurrentEngine().ExecutionRecord()
	a.Logger.Info(record)
	a.updateExecutionRecord(context.TODO(), record)
}

func (a *actionServer) RunEngine(ctx context.Context, request *services.RunEngineRequest) (*services.StatusResponse, error) {
	engine := a.getEngine(request.EngineName)
	if engine == nil {
		return &services.StatusResponse{Status: 404, Message: "Engine " + request.EngineName + " Not Found"}, nil
	}
	_workflow, err := workflow.New(request.WorkflowId)
	if err != nil {
		return &services.StatusResponse{Status: 404, Message: err.Error()}, nil
	}
	go func() {
		_context := workflow.Context()
		engine.Init(_context)
		engine.WaitForResources(_workflow)
		engine.Exec(_workflow)
	}()
	return &services.StatusResponse{Status: 200, Message: "ok"}, nil
}

func (a *actionServer) GetEngineStatus(ctx context.Context, request *services.EngineStatusRequest) (*services.EngineStatusResponse, error) {
	engine := a.engines[request.EngineName]
	if engine == nil {
		return &services.EngineStatusResponse{Code: -1, Message: "Engine " + request.EngineName + " Not Found"}, nil
	}
	status, msg := engine.Status()
	return &services.EngineStatusResponse{Code: int32(status), Message: msg}, nil
}

func (a *actionServer) updateExecutionRecord(ctx context.Context, record *models.ExecutionRecord) {
	if !a.updatingExecutionRecord {
		a.updatingExecutionRecord = true
		_, err := a.apiClient.CreateOrUpdateExecutionRecord(ctx, record)
		if err != nil {
			a.Logger.Error(err.Error())
		}
		a.updatingExecutionRecord = false
	}
}

func (a *actionServer) GetExecutionRecord(ctx context.Context, request *services.EmptyRequest) (*models.ExecutionRecord, error) {
	record := a.getCurrentEngine().ExecutionRecord()
	go a.updateExecutionRecord(ctx, record)
	return record, nil
}

func (a *actionServer) GetRequiredResources(ctx context.Context, query *services.RequiredResourcesRequest) (*models.WorkflowResources, error) {

	stage := query.Stage

	//a.Logger.Info(query)

	cacheKey := fmt.Sprintf("%s-%d", query.WorkflowId, query.Stage)

	if res, ok := a.requiredResourcesMap[cacheKey]; ok {
		return res, nil
	}

	engine := a.getEngine(query.EngineName)
	wf, err := workflow.GetWorkflowByID(query.WorkflowId)

	if err != nil {
		return nil, err
	}
	requiredResources, err := engine.RequiredResources(wf, stage)
	if err != nil {
		return nil, err
	}
	requiredResources = workflow.ResourcesPreProcess(requiredResources, wf)
	a.requiredResourcesMap[cacheKey] = requiredResources
	return requiredResources, nil
}

func (a *actionServer) CallEngine(request *services.ActionRequest, server services.ActionsService_CallEngineServer) error {
	engine := a.getEngine(request.ActionName)
	if engine == nil {
		return errors.New("engine " + request.ActionName + " Not Found")
	}
	//_workflow, err := workflow.New(request.WorkflowId)
	//if err != nil {
	//	return err
	//}
	return engine.ProcessData(&engineTypes.ActionRequest{
		ID:     request.Id,
		Inputs: request.Inputs,
	}, func(response *engineTypes.ActionResponse) error {
		return server.Send(&services.ActionResponse{
			Id:      response.ID,
			Outputs: response.Outputs,
			Count:   response.Count,
			Done:    response.Done,
		})
	})
}

//func (a *actionServer) CallEngine(server services.ActionsService_CallEngineServer) error {
//
//	for true {
//		actionRequest, err := server.Recv()
//		if err != nil {
//			return err
//		}
//		engine := a.getEngine(actionRequest.ActionName)
//		if engine == nil {
//			return errors.New("engine " + actionRequest.ActionName + " Not Found")
//		}
//	}
//
//
//
//}
