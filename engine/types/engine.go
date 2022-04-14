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
	WaitForResources(workflow *Workflow) error
	ExecutionRecord() *models.ExecutionRecord
	Exec(workflow *Workflow) error
	Stop() error
	Init(ctx *EngineContext)
	Status() (code EngineStatusCode, msg string)
}
