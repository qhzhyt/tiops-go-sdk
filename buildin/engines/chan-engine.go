// @Title  chan-engine.go
// @Description  基于channel的简易流程引擎
// @Create  heyitong  2022/6/23 15:30
// @Update  heyitong  2022/6/23 15:30

package engines

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
	tiopsConfigs "tiops/common/config"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/utils"
	"tiops/engine/types"
)

// basicChanEngine 基于channel的简易流程引擎
type basicChanEngine struct {
	wg sync.WaitGroup
	*types.EngineContext
	recordManager *types.ExecutionRecordManager
	running       bool
	ready         bool
	inputNode     *types.Node
	outputNode    *types.Node
}

func (w *basicChanEngine) ProcessData(requests chan *types.ActionRequest) (chan *types.ActionResponse, error) {
	go func() {
		for request := range requests {
			//w.Logger.Info("Receive ", request)
			for name, connections := range w.inputNode.Outputs {
				for _, connection := range connections {
					d := request.Inputs[name]
					if d == nil {
						continue
					}
					connection.DataChan <- d
				}
			}
		}
	}()

	res := make(chan *types.ActionResponse, 1000)

	go func() {
		for true {
			outputData := map[string]*services.ActionData{}
			var done bool
			count := int64(0)
			var traceIds []int64
			for k, inputs := range w.outputNode.Inputs {
				//fmt.Println(node.Info.ActionName, node.Inputs, node.Outputs)
				if len(inputs) > 0 {
					data, isOK := inputs.SelectInput()
					if !isOK {
						done = true
						break
					}
					outputData[k] = data

					if data == nil {
						continue
					}

					traceIds = append(traceIds, data.TraceIds...)

					if data.Count > count {
						count = data.Count
					}
				}
			}

			res <- &types.ActionResponse{
				Done:     done,
				Outputs:  outputData,
				Count:    count,
				TraceIds: utils.Int64ListDeduplicate(traceIds),
			}
		}
	}()

	return res, nil
}

/*func (w *basicChanEngine) ProcessData(request *types.ActionRequest, outputCallback func(*types.ActionResponse) error) error {
	go func() {
		for name, connections := range w.inputNode.Outputs {
			for _, connection := range connections {
				connection.DataChan <- request.Inputs[name]
			}
		}
	}()

	outputData := map[string]*services.ActionData{}
	var done bool
	count := int64(0)
	for k, inputs := range w.outputNode.Inputs {
		//fmt.Println(node.Info.ActionName, node.Inputs, node.Outputs)
		if len(inputs) > 0 {
			data, isOK := inputs.SelectInput()
			if !isOK {
				done = true
				break
			}
			outputData[k] = data

			if data.Count > count {
				count = data.Count
			}
		}
	}

	err := outputCallback(&types.ActionResponse{
		ID:      request.ID,
		Done:    done,
		Outputs: outputData,
		Count:   count,
	})

	if err != nil || done {
		return err
	}

	for !done {
		outputData = map[string]*services.ActionData{}
		count = int64(0)
		first := true
		for k, inputs := range w.outputNode.Inputs {
			//fmt.Println(node.Info.ActionName, node.Inputs, node.Outputs)
			var data *services.ActionData
			var isOK bool
			if len(inputs) > 0 {
				if first {
					data, isOK = inputs.SelectInputOnce()
					if !isOK {
						return nil
					}
				} else {
					data, isOK = inputs.SelectInput()
					if !isOK {
						return nil
					}
				}

				first = false

				outputData[k] = data

				if data.Count > count {
					count = data.Count
				}
			}
		}

		err = outputCallback(&types.ActionResponse{
			ID:      request.ID,
			Done:    done,
			Outputs: outputData,
			Count:   count,
		})

		if err != nil || done {
			return err
		}

	}

	return nil
}
*/

func (w *basicChanEngine) Status() (code types.EngineStatusCode, msg string) {
	if w.running || !w.ready {
		return types.EngineStatusBusy, "busy"
	}
	return types.EngineStatusIdle, "idle"
}

func (w *basicChanEngine) ExecutionRecord() *models.ExecutionRecord {
	recordsData, _ := json.Marshal(w.recordManager.Records())
	return &models.ExecutionRecord{
		//XId:            w.ExecutionId,
		//ExecutionId:    w.ExecutionId,
		ProcessRecords: nil,
		StatusRecords: []*models.ExecutionStatusRecord{
			{ConfigNames: []string{"cumulativeItem", "backlogItems", "processRate"}, Data: string(recordsData)},
		},
	}
}

func (w *basicChanEngine) WaitForResources(workflow *types.Workflow) error {
	if !workflow.RegisterActionNodes() {
		utils.SleepAndExit(time.Second*6, -1)
	}
	w.recordManager.Start()

	return nil
}

func (w *basicChanEngine) RequiredResources(workflowInfo *types.Workflow, stage int32) (*models.WorkflowResources, error) {

	var apps []*models.K8SApp
	nodes := workflowInfo.Nodes
	for _, node := range nodes {
		//w.Logger.Info(stage, " ", node)

		resources, err := node.GetRequiredResources(stage)

		//w.Logger.Info(resources)
		//time.Sleep(time.Second)
		if err != nil {
			w.Logger.Error(err.Error())
			return nil, err
		}
		if resources != nil && len(resources.Apps) > 0 {
			apps = append(apps, resources.Apps...)
		}
	}

	return &models.WorkflowResources{
		Apps: apps,
	}, nil
}

func (w *basicChanEngine) Init(ctx *types.EngineContext) {
	w.EngineContext = ctx
	w.recordManager = types.NewExecutionRecordManager(10*time.Second, ctx)
}

func (w *basicChanEngine) processResponse(node *types.Node, res *types.ActionResponse, err error) bool {

	if res == nil {
		w.Logger.Error(err)
		return false
	}

	req := res.Request
	inputData := req.Inputs

	processRecord := req.Record

	requestId := req.ID

	if err != nil {
		w.Logger.Error(errorLog(node.Action.Info(), err, inputData))
		return false
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

	processRecord.ElapsedTime = processRecord.EndTime - processRecord.StartTime
	processRecord.ProcessRate = float32(processRecord.ItemCount) / float32(processRecord.ElapsedTime) * 1000

	w.recordManager.AddProcessRecord(processRecord)

	w.Logger.Debug(outputLog(node.Action.Info(), requestId, res.Outputs))
	//time.Sleep(time.Second * 10)
	for k, outputs := range node.Outputs {
		for _, output := range outputs {
			//w.Logger.Debug(output.Info)
			//time.Sleep(time.Second * 2)
			output.DataChan <- res.Outputs[k]
		}
	}

	return res.Done
}

func (w *basicChanEngine) ExecNodeWithInput(node *types.Node) {
	done := false

	doneWg := &sync.WaitGroup{}

	actionInfo := node.Action.Info()

	if actionInfo.Type == models.ActionType_WorkflowAction {
		actionInfo.CallMode = models.CallMode_DuplexStreamCall
	}

	duplexStreamSender := func(request *types.ActionRequest) error {
		return nil
	}

	if actionInfo.CallMode == models.CallMode_DuplexStreamCall {
		var err error
		doneWg.Add(1)
		duplexStreamSender, err = node.Action.CallDuplexStream(func(res *types.ActionResponse, err error) bool {
			if err != nil {
				w.Logger.Error(err)
				return false
			}

			//w.Logger.Debug(res)

			w.processResponse(node, res, err)

			if res.Done || done {
				done = true
				doneWg.Done()
			}

			return done
		})

		if err != nil {
			w.Logger.Error(err.Error())
			utils.SleepAndExit(time.Second*6, -1)
		}
	}

	for !done {
		maxBacklog := 0
		inputData := map[string]*services.ActionData{}
		var traceIds []int64
		for k, inputs := range node.Inputs {
			//fmt.Println(node.Info.ActionName, node.Inputs, node.Outputs)
			if len(inputs) > 0 {
				data, isOK := inputs.SelectInput()
				if !isOK {
					done = true
					break
				}
				inputData[k] = data
				if data == nil {
					break
				}
				//w.Logger.Error("data, ", data)
				//time.Sleep(time.Second * 6)
				traceIds = append(traceIds, data.TraceIds...)
				if inputs.BacklogCount() > maxBacklog {
					maxBacklog = inputs.BacklogCount()
				}
			}
		}

		//w.Logger.Info(inputData)

		if !done {
			requestId := fmt.Sprintf("%x", utils.SnowflakeID())
			w.Logger.Debug(inputLog(node.Action.Info(), requestId, inputData))

			processRecord := &models.ProcessRecord{
				StartTime:  utils.CurrentTimeStampMS(),
				NodeId:     node.ID,
				ActionName: node.Action.Info().Name,
				ItemCount:  itemCount(inputData),
				BatchCount: 1,
				RecordId:   requestId,
				//ExecutionId:    w.ExecutionId,
				BacklogBatches: int64(maxBacklog),
			}

			if actionInfo.CallMode == models.CallMode_OnceCall {

			}

			actionRequest := &types.ActionRequest{
				ID:       requestId,
				Inputs:   inputData,
				Record:   processRecord,
				TraceIds: utils.Int64ListDeduplicate(traceIds),
			}

			switch actionInfo.CallMode {
			case models.CallMode_PullStreamCall:
				err := node.Action.CallPullStream(actionRequest, func(res *types.ActionResponse, err error) bool {
					done = w.processResponse(node, res, err)
					if done || (res != nil && res.Done) {
						return false
					}
					return true
				})
				if err != nil {
					w.Logger.Error(err)
				}
			case models.CallMode_DuplexStreamCall:
				err := duplexStreamSender(actionRequest)
				if err != nil {
					w.Logger.Error(err)
				}
			default:
				res, err := node.Action.Call(actionRequest)
				done = w.processResponse(node, res, err)
			}

		}

		if done {
			doneWg.Wait()
			for _, outputs := range node.Outputs {
				for _, output := range outputs {
					//w.Logger.Debug(actionInfo.Name, " ", "done")
					output.Done()
					//time.Sleep(time.Second)
				}
			}
		}
	}
	if !node.HasVarInputOnly {
		w.wg.Done()
	}
}

func (w *basicChanEngine) Exec(workflow *types.Workflow) error {
	w.running = true
	defer func() {
		w.running = false
	}()
	w.recordManager.Start()

	w.inputNode = workflow.InputNode
	w.outputNode = workflow.OutputNode

	waitCount := 0
	for _, node := range workflow.Nodes {
		if !node.HasVarInputOnly {
			w.wg.Add(1)
			waitCount++
		}
		go w.ExecNodeWithInput(node)
	}

	if workflow.OutputNode != nil && tiopsConfigs.InMainEngine() {
		for k, outputs := range workflow.OutputNode.Inputs {
			for _, output := range outputs {
				w.Logger.Debug(fmt.Sprintln("Output", k, <-output.DataChan))
				//fmt.Println(k, <-output.DataChan)
			}
		}
	}

	w.wg.Wait()

	time.Sleep(time.Second * 6)

	return nil
}

func (w *basicChanEngine) Stop() error {
	panic("implement me")
}

func NewBasicChanEngine() types.WorkflowEngine {
	return &basicChanEngine{}
}
