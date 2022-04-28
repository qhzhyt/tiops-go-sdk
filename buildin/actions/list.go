package actions

import (
	"tiops/buildin"
	apiclient "tiops/common/api-client"
	tiopsConfigs "tiops/common/config"
)

var (
	batchProcessorMap = map[string]BuildinActionBatchProcessor{
		buildin.EventSenderID: EventSender,
		buildin.StaticEventSenderID: StaticEventSender,
	}
	pullStreamProcessorMap = map[string]BuildinActionPullStreamProcessor{
		buildin.EventTriggerID: EventTrigger,
		buildin.HttpTriggerID: HttpTrigger,
	}
	apiClient = apiclient.NewAPIClient(tiopsConfigs.ApiServerHost, tiopsConfigs.ApiServerGrpcPort)
)
