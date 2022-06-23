// @Title  engine.go
// @Description  流程引擎接口定义
// @Create  heyitong  2022/6/23 17:22
// @Update  heyitong  2022/6/23 17:22

package types

import "tiops/common/models"

type EngineStatusCode int

const (
	EngineStatusNotFound EngineStatusCode = -1
	EngineStatusIdle                      = iota
	EngineStatusBusy
)

// WorkflowEngine 流程引擎接口
type WorkflowEngine interface {
	RequiredResources(workflow *Workflow, stage int32) (*models.WorkflowResources, error)
	WaitForResources(workflow *Workflow) error
	ExecutionRecord() *models.ExecutionRecord
	Exec(workflow *Workflow) error
	Stop() error
	Init(ctx *EngineContext)
	Status() (code EngineStatusCode, msg string)
	ProcessData(request *ActionRequest, outputCallback func(*ActionResponse) error) error
	//Output(func(*ActionResponse) error) error
}
