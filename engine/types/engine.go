package types

import "tiops/common/models"

type WorkflowEngine interface {
	Exec(workflow *Workflow)
	RequiredK8sResources(workflowInfo *models.WorkflowInfo)
	Stop()
}
