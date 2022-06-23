// @Title  context.go
// @Description  流程引擎上下文
// @Create  heyitong  2022/6/23 17:17
// @Update  heyitong  2022/6/23 17:17

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
