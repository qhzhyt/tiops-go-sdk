package types


type WorkflowEngine interface {
	Exec(workflow *Workflow)
	Stop()
}
