package config

const (
	WorkflowEngineLabel           = "workflow-engine"
	WorkflowEngineImageBase       = "tiops-system/workflow-engine:latest"
	WorkflowActionServerImageBase = "tiops-project-images/%s:latest"
	WorkflowActionServerLabelBase = "workflow-action-server-%s"
	LogIdEnvName                  = "LOG_ID"
	ProjectIdEnvName              = "PROJECT_ID"
	WorkflowIdEnvName             = "WORKFLOW_ID"
	WorkflowIdLabel               = "workflow-id"
	WorkflowTypeEnvName           = "WORKFLOW_TYPE"
	ExecutionIdEnvName            = "EXECUTION_ID"
	LogLevelEnvName               = "LOG_LEVEL"
	ApiServerHost                 = "tiops-api-server.tiops-system"
	ApiServerGrpcPort             = 5000
	ApiServerHttpPort             = 8000
	TiopsSystemK8sNamespace       = "tiops-system"
)
