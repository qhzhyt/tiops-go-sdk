package types

import (
	actionTypes "tiops/action/types"
	apiClient "tiops/common/api-client"
	"tiops/common/logger"
)

type EngineContext struct {
/*	WorkflowType string
	WorkflowId string
	ExecutionId string
	RecordId string*/
	Options actionTypes.ActionOptions
	*logger.Logger
	*apiClient.APIClient
}
