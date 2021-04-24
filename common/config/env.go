package config

import "os"

var (
	WorkflowId   = os.Getenv(WorkflowIdEnvName)
	LogId        = os.Getenv(LogIdEnvName)
	LogLevel     = os.Getenv(LogLevelEnvName)
	WorkflowType = os.Getenv(WorkflowTypeEnvName)
)
