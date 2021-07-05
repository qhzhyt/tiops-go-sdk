package types

import (
	apiClient "tiops/common/api-client"
	"tiops/common/logger"
)

type WorkflowContext struct {
	WorkflowType string
	WorkflowId string
	ExecutionId string
	RecordId string
	*logger.Logger
	*apiClient.APIClient
}
