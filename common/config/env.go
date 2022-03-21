package config

import "os"

var (
	WorkflowId           = os.Getenv(EnvNameWorkflowId)
	WorkspaceId          = os.Getenv(EnvNameWorkspaceId)
	SubWorkflowId        = os.Getenv(EnvNameSubWorkflowId)
	LogId                = os.Getenv(EnvNameLogId)
	LogLevel             = os.Getenv(EnvNameLogLevel)
	WorkflowType         = os.Getenv(EnvNameWorkflowType)
	ExecutionId          = os.Getenv(EnvNameExecutionId)
	JobId                = os.Getenv(EnvNameJobId)
	AppType              = os.Getenv(EnvNameAppType)
	ProjectId            = os.Getenv(EnvNameProjectId)
	EngineName           = os.Getenv(EnvNameEngineName)
	inMainEngine         = os.Getenv(EnvNameInMainEngine)
	runEngine            = os.Getenv(EnvNameRunEngine)
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

func InSubMainEngine() bool {
	return SubWorkflowId != ""
}

func RunEngineAtStartup() bool {
	return runEngine != ""
}
