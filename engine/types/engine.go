package types

import "tiops/common/models"

type EngineStatusCode int

const (
	EngineStatusNotFound EngineStatusCode = -1
	EngineStatusIdle = iota
	EngineStatusBusy
)

type WorkflowEngine interface {
	RequiredResources(workflow *Workflow, stage int) *models.WorkflowResources
	WaitForResources(workflow *Workflow)
	ExecutionRecord() *models.ExecutionRecord
	Exec(workflow *Workflow)
	Stop()
	Init(ctx *WorkflowContext)
	Status() (EngineStatusCode int, msg string)
}
