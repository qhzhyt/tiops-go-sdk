package types

import "tiops/common/models"

type WorkflowEngine interface {
	RequiredResources(workflowInfo *models.WorkflowInfo) *models.WorkflowResources
	WaitForResources(workflow *Workflow)
	Exec(workflow *Workflow)
	Stop()
	Init(ctx *WorkflowContext)
}
