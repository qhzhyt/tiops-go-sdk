package types

import "tiops/common/models"

type WorkflowEngine interface {
	RequiredResources(workflow *Workflow) *models.WorkflowResources
	WaitForResources(workflow *Workflow)
	ExecutionRecord() *models.ExecutionRecord
	Exec(workflow *Workflow)
	Stop()
	Init(ctx *WorkflowContext)
}
