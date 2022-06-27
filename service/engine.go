// @Title  engine.go
// @Description  ActionService中流程引擎相关接口实现
// @Create  heyitong  2022/6/17 16:56
// @Update  heyitong  2022/6/17 16:56

package service

import (
	"context"
	"errors"
	"io"
	"time"
	"tiops/buildin/engines"
	tiopsConfigs "tiops/common/config"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/utils"
	engineTypes "tiops/engine/types"
	"tiops/engine/workflow"
)

// runMainEngine 运行主流程引擎
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

// getCurrentEngine 获取当前流程引擎
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

// getEngine 根据名称获取流程引擎
func (a *actionServer) getEngine(name string) engineTypes.WorkflowEngine {
	//a.Logger.Info("current engine: " + tiopsConfigs.EngineName)
	if a.engines[name] != nil {
		return a.engines[name]
	}

	return a.getCurrentEngine()
}

// afterEngineExec 处理流程执行完毕后执行的代码
func (a *actionServer) afterEngineExec() {
	record := a.getCurrentEngine().ExecutionRecord()
	a.Logger.Info(record)
	a.updateExecutionRecord(context.TODO(), record)
}

// updateExecutionRecord 更新执行记录
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

// RunEngine 执行引擎，非主流程引擎需要用户调用
func (a *actionServer) RunEngine(ctx context.Context, request *services.RunEngineRequest) (*services.StatusResponse, error) {
	engine := a.getEngine(request.EngineName)
	if engine == nil {
		return &services.StatusResponse{Status: 404, Message: "Engine " + request.EngineName + " Not Found"}, nil
	}
	a.engines[request.EngineName] = engine
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

// GetEngineStatus 获取EngineName所对应流程引擎的状态
func (a *actionServer) GetEngineStatus(ctx context.Context, request *services.EngineStatusRequest) (*services.EngineStatusResponse, error) {
	engine := a.engines[request.EngineName]
	if engine == nil {
		return &services.EngineStatusResponse{Code: -1, Message: "Engine " + request.EngineName + " Not Found"}, nil
	}
	status, msg := engine.Status()
	return &services.EngineStatusResponse{Code: int32(status), Message: msg}, nil
}

// GetExecutionRecord 获取执行记录
func (a *actionServer) GetExecutionRecord(ctx context.Context, request *services.EmptyRequest) (*models.ExecutionRecord, error) {
	record := a.getCurrentEngine().ExecutionRecord()
	record.XId = tiopsConfigs.ExecutionId
	record.ExecutionId = tiopsConfigs.ExecutionId
	go a.updateExecutionRecord(context.TODO(), record)
	return record, nil
}

// GetRequiredResources 获取处理流程所需的资源集合
func (a *actionServer) GetRequiredResources(ctx context.Context, query *services.RequiredResourcesRequest) (*models.WorkflowResources, error) {

	stage := query.Stage

	//a.Logger.Info(query)

	//cacheKey := fmt.Sprintf("%s-%d", query.WorkflowId, query.Stage)

	//if res, ok := a.requiredResourcesMap[cacheKey]; ok {
	//	return res, nil
	//}

	engine := a.getEngine(query.EngineName)
	wf, err := workflow.GetWorkflowByID(query.WorkflowId)

	//a.Logger.Info(wf)
	//a.Logger.Info("Stage", stage)

	time.Sleep(time.Second)

	if err != nil {
		return nil, err
	}
	requiredResources, err := engine.RequiredResources(wf, stage)
	if err != nil {
		return nil, err
	}
	requiredResources = workflow.ResourcesPreProcess(requiredResources, wf)
	//a.requiredResourcesMap[cacheKey] = requiredResources
	return requiredResources, nil
}

// CallEngine 调用流程引擎执行特定处理流程
func (a *actionServer) CallEngine(server services.ActionsService_CallEngineServer) error {
	//panic("implement me")
	done := false

	request, err := server.Recv()

	if err != nil {
		a.Logger.Error(err.Error())
		time.Sleep(time.Second)
		return err
	}

	engine := a.getEngine(request.ActionName)
	if engine == nil {
		return errors.New("engine " + request.ActionName + " Not Found")
	}

	reqChan := make(chan *engineTypes.ActionRequest, 3000)

	resChan, err := engine.ProcessData(reqChan)

	if err != nil {
		return err
	}

	go func() {
		for response := range resChan {
			if response.Done {
				done = true
			}

			//a.Logger.Info(response)

			err := server.Send(&services.ActionResponse{
				Id:       response.ID,
				Outputs:  response.Outputs,
				Count:    response.Count,
				Done:     response.Done,
				TraceIds: response.TraceIds,
			})

			if err != nil {
				a.Logger.Error(err.Error())
				time.Sleep(time.Second)
			}
		}
	}()

	reqChan <- &engineTypes.ActionRequest{
		ID:       request.Id,
		Inputs:   request.Inputs,
		Count:    0,
		TraceIds: request.TraceIds,
	}

	for !done {

		request, err = server.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			a.Logger.Error(err.Error())
			continue
		}

		reqChan <- &engineTypes.ActionRequest{
			ID:       request.Id,
			Inputs:   request.Inputs,
			Count:    0,
			TraceIds: request.TraceIds,
		}
	}

	return nil
}

/*func (a *actionServer) CallEngine(request *services.ActionRequest, server services.ActionsService_CallEngineServer) error {
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
*/
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
