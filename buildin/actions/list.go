// @Title  list.go
// @Description  内置组件列表
// @Create  heyitong  2022/6/23 15:43
// @Update  heyitong  2022/6/23 15:43

package actions

import (
	"tiops/buildin"
	apiclient "tiops/common/api-client"
	tiopsConfigs "tiops/common/config"
)

var (
	batchProcessorMap = map[string]BuildinActionBatchProcessor{
		buildin.EventSenderID:       EventSender,
		buildin.StaticEventSenderID: StaticEventSender,
		buildin.JsonSetFieldID:      JsonSetField,
		buildin.JsonPathLookupID:    JsonPathLookup,
	}
	pullStreamProcessorMap = map[string]BuildinActionPullStreamProcessor{
		buildin.EventTriggerID: EventTrigger,
		buildin.HttpTriggerID:  HttpTrigger,
	}
	apiClient = apiclient.NewAPIClient(tiopsConfigs.ApiServerHost, tiopsConfigs.ApiServerGrpcPort)
)
