package config

const (
	ManifestPath                  = "manifest.yaml"
	WorkflowEngineLabel           = "workflow-engine"
	WorkflowEngineImageBase       = "tiops-system/workflow-engine:latest"
	WorkflowActionServerImageBase = "tiops-project-images/%s:latest"
	WorkflowActionServerLabelBase = "workflow-action-server-%s"
	WorkflowIdLabel               = "workflow-id"
	ApiServerHost                 = "tiops-api-server.tiops-system"
	ApiServerGrpcPort             = 5000
	ApiServerHttpPort             = 8000
	TiopsSystemK8sNamespace       = "tiops-system"
	ActionServerPort              = 5555
)

const (
	EnvNameLogId        = "LOG_ID"
	EnvNameProjectId    = "PROJECT_ID"
	EnvNameWorkflowId   = "WORKFLOW_ID"
	EnvNameWorkflowType = "WORKFLOW_TYPE"
	EnvNameExecutionId  = "EXECUTION_ID"
	EnvNameAppType      = "APP_TYPE"
	EnvNameLogLevel     = "LOG_LEVEL"
)

const (
	AppTypeWorkflowEngine = "workflow-engine"
	AppTypeActionServer   = "action-server"
)
