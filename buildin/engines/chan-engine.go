package engines

/*
基于channel的简易流程引擎
*/

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"tiops/common/config"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/utils"
	"tiops/engine/record"
	"tiops/engine/types"
)

type basicChanEngine struct {
	wg sync.WaitGroup
	*types.WorkflowContext
	recordManager *record.ExecutionRecordManager
	running       bool
}

func (w *basicChanEngine) Status() (EngineStatusCode int, msg string) {
	if w.running {
		return types.EngineStatusBusy, "busy"
	}
	return types.EngineStatusIdle, "idle"
}

func (w *basicChanEngine) ExecutionRecord() *models.ExecutionRecord {
	recordsData, _ := json.Marshal(w.recordManager.Records())
	return &models.ExecutionRecord{
		XId:            w.ExecutionId,
		ExecutionId:    w.ExecutionId,
		ProcessRecords: nil,
		StatusRecords: []*models.ExecutionStatusRecord{
			{ConfigNames: []string{"cumulativeItem", "backlogItems", "processRate"}, Data: string(recordsData)},
		},
	}
}

func (w *basicChanEngine) WaitForResources(workflow *types.Workflow) {
	workflow.RegisterActionNodes()
	w.recordManager.Start()
}

func (w *basicChanEngine) RequiredResources(workflowInfo *types.Workflow, stage int) *models.WorkflowResources {
	if stage == 0 {
		var apps []*models.K8SApp
		nodes := workflowInfo.Nodes
		processedProjects := map[string]bool{}
		for _, node := range nodes {

			nodeInfo := node.Info
			actionInfo := node.Action.Info()

			if node.Info.ActionExecutor != "" {
				actionInfo = node.ActionExecutor
			}

			serviceName := config.ActionServiceName(actionInfo.ProjectId)

			app := &models.K8SApp{
				Name:      serviceName,
				ActionId: actionInfo.XId,
				Replica:     1,
				ServiceMode: models.ServiceMode_One,
			}
			if node.Info.StandAlone {
				app.Name = config.StandAloneActionServiceName(nodeInfo.ActionName, node.ID)
				apps = append(apps, app)
			} else if !processedProjects[actionInfo.ProjectId] {
				apps = append(apps, app)
				processedProjects[actionInfo.ProjectId] = true
			}
		}

		return &models.WorkflowResources{
			Apps: apps,
		}
	}
	return &models.WorkflowResources{
		Apps: nil,
	}
}

func (w *basicChanEngine) Init(ctx *types.WorkflowContext) {
	w.WorkflowContext = ctx
	w.recordManager = record.NewExecutionRecordManager(10*time.Second, ctx)
}

//func (w *basicChanEngine) selectOne(connections []*workflow.Connection, count int) (*services.ActionData, bool) {
//	var selectCase = make([]reflect.SelectCase, len(connections))
//
//	for i, connection := range connections {
//		if connection.DataChan != nil {
//			selectCase[i].Dir = reflect.SelectRecv
//			selectCase[i].Chan = reflect.ValueOf(connection.DataChan)
//		} else {
//			return connection.Variable.ToActionArguments(count), true
//		}
//	}
//	_, recv, recvOk := reflect.Select(selectCase)
//	if recvOk {
//		//fmt.Println(recv)
//		return recv.Interface().(*services.ActionData), false
//	}
//
//	return nil, false
//}

func itemCount(inputData map[string]*services.ActionData) int32 {
	count := int32(0)
	for _, data := range inputData {
		if data != nil && data.Count > count {
			count = data.Count
		}
	}

	return count
}

func inputLog(info *models.ActionInfo, requestId string, inputData map[string]*services.ActionData) string {
	infos := map[string]map[string]interface{}{}
	for k, v := range inputData {
		if v != nil {
			infos[k] = map[string]interface{}{
				"id":    v.Id,
				"count": v.Count,
			}
		} else {
			infos[k] = map[string]interface{}{
				"id":    "no data",
				"count": "0",
			}
		}
	}
	return fmt.Sprintf("Action request %s to %s with inputs: %v", requestId, info.Name, infos)
}

func outputLog(info *models.ActionInfo, responseId string, outputData map[string]*services.ActionData) string {
	infos := map[string]map[string]interface{}{}
	for k, v := range outputData {
		infos[k] = map[string]interface{}{
			"id":    v.Id,
			"count": v.Count,
		}
	}
	return fmt.Sprintf("Action reponse %s from %s with outputs: %v", responseId, info.Name, infos)
}

func (w *basicChanEngine) ExecNodeWithInput(node *types.Node) {
	done := false

	for !done {
		maxBacklog := 0
		inputData := map[string]*services.ActionData{}
		for k, inputs := range node.Inputs {
			//fmt.Println(node.Info.ActionName, node.Inputs, node.Outputs)
			if len(inputs) > 0 {
				data, isOK := inputs.SelectInput()
				if !isOK {
					done = true
					break
				}
				inputData[k] = data
				if inputs.BacklogCount() > maxBacklog {
					maxBacklog = inputs.BacklogCount()
				}
			}
		}

		if !done {
			requestId := fmt.Sprintf("%x", utils.SnowflakeID())
			w.Logger.Info(inputLog(node.Action.Info(), requestId, inputData))

			processRecord := &models.ProcessRecord{
				StartTime:      utils.CurrentTimeStampMS(),
				NodeId:         node.ID,
				ActionName:     node.Action.Info().Name,
				ItemCount:      itemCount(inputData),
				BatchCount:     1,
				RecordId:       requestId,
				ExecutionId:    w.ExecutionId,
				BacklogBatches: int32(maxBacklog),
			}

			res, err := node.Action.Call(&types.ActionRequest{Inputs: inputData, ID: requestId})
			if err != nil {
				w.Logger.Error(fmt.Sprintf("Call %s get error: %s\nCurrent inputs: %v", node.Action.Info().Name, err.Error(), inputData))
				continue
				//utils.SleepAndExit(time.Second*3, 1)
				//return
			}

			processRecord.EndTime = utils.CurrentTimeStampMS()
			if processRecord.ItemCount == 0 {
				processRecord.ItemCount = itemCount(res.Outputs)
				if processRecord.ItemCount == 0 {
					processRecord.ItemCount = 1
				}
			}
			//processRecord.BatchSize = processRecord.ItemCount

			processRecord.ElapsedTime = int32(processRecord.EndTime - processRecord.StartTime)
			processRecord.ProcessRate = float32(processRecord.ItemCount) / float32(processRecord.ElapsedTime) * 1000

			w.recordManager.AddProcessRecord(processRecord)

			w.Logger.Info(outputLog(node.Action.Info(), requestId, res.Outputs))
			//time.Sleep(time.Second * 10)
			for k, outputs := range node.Outputs {
				for _, output := range outputs {
					output.DataChan <- res.Outputs[k]
				}
			}

			if res.Done {
				done = true
			}
		}

		if done {
			for _, outputs := range node.Outputs {
				for _, output := range outputs {
					output.Done()
				}
			}
		}
	}
	if !node.HasVarInputOnly {
		w.wg.Done()
	}
}

func (w *basicChanEngine) Exec(workflow *types.Workflow) {
	w.running = true
	defer func() {
		w.running = false
	}()
	w.recordManager.Start()

	waitCount := 0
	for _, node := range workflow.Nodes {
		if !node.HasVarInputOnly {
			w.wg.Add(1)
			waitCount++
		}
		go w.ExecNodeWithInput(node)
	}

	if workflow.OutputNode != nil {
		for k, outputs := range workflow.OutputNode.Inputs {
			for _, output := range outputs {
				w.Logger.Info(fmt.Sprintln("Output", k, <-output.DataChan))
				//fmt.Println(k, <-output.DataChan)
			}
		}
	}

	w.wg.Wait()

	time.Sleep(time.Second * 6)

}

func (w *basicChanEngine) Stop() {
	panic("implement me")
}

func NewBasicChanEngine() types.WorkflowEngine {
	return &basicChanEngine{}
}
