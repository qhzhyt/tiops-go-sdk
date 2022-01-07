package config

import "os"

var (
	WorkflowId   = os.Getenv(EnvNameWorkflowId)
	LogId        = os.Getenv(EnvNameLogId)
	LogLevel     = os.Getenv(EnvNameLogLevel)
	WorkflowType = os.Getenv(EnvNameWorkflowType)
	ExecutionID  = os.Getenv(EnvNameExecutionId)
	AppType      = os.Getenv(EnvNameAppType)
	ProjectId    = os.Getenv(EnvNameProjectId)
	EngineName   = os.Getenv(EnvNameEngineName)
	inMainEngine = os.Getenv(EnvNameInMainEngine)
	TiopsSystemNamespace = os.Getenv(EnvNameTiopsSystemNamespace)
)

func init() {
	if TiopsSystemNamespace == "" {
		TiopsSystemNamespace = DefaultTiopsSystemNamespace
	}
}

func InMainEngine() bool {
	return inMainEngine != ""
}

