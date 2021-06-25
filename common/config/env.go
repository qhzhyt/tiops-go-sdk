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
)
