package config

const (
	ManifestPath                  = "manifest.yaml"
	WorkflowEngineLabel           = "workflow-engine"
	WorkflowEngineImageBase       = "tiops-system/workflow-engine:latest"
	WorkflowActionServerImageBase = "tiops-project-images/%s:latest"
	WorkflowActionServerLabelBase = "workflow-action-server-%s"
	WorkflowIdLabel               = "workflow-id"
	ApiServerGrpcPort             = 5000
	ApiServerHttpPort             = 8000
	DefaultTiopsSystemNamespace       = "tiops-system"
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
	EnvNameEngineName   = "ENGINE_NAME"
	EnvNameInMainEngine = "IN_MAIN_ENGINE"
	EnvNameTiopsSystemNamespace = "TIOPS_SYSTEM_NAMESPACE"
)

const (
	AppTypeWorkflowEngine = "workflow-engine"
	AppTypeActionServer   = "action-server"
)

const (
	BuildinEngineImage = "tiops-buildin-projects/workflow-engines:latest"
	DefaultEngineName  = "default"
)
