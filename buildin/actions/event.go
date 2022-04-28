package actions

import (
	"context"
	actionTypes "tiops/action/types"
	"tiops/buildin"
	tiopsConfigs "tiops/common/config"
	"tiops/common/services"
)

func EventTrigger(ctx *actionTypes.StreamRequestContext) error {
	workspace := tiopsConfigs.WorkspaceId
	if ctx.ActionOptions.GetString(buildin.Workspaced) == buildin.False {
		workspace = ""
	}

	eventClient, err := apiClient.WatchEvent(context.TODO(), &services.WatchEventRequest{
		Workspace: workspace,
		Name:      ctx.ActionOptions.GetString(buildin.EventName),
	})
	if err != nil {
		return err
	}

	for true {
		event, err := eventClient.Recv()
		if err != nil {
			ctx.Logger.Error(err)
		} else {
			err = ctx.Push(actionTypes.ActionDataBatch{map[string]interface{}{buildin.EventArg: string(event.Data)}})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func EventSender(ctx *actionTypes.BatchRequestContext) (actionTypes.ActionDataBatch, error) {

	workspace := tiopsConfigs.WorkspaceId
	if ctx.ActionOptions.GetString(buildin.Workspaced) == buildin.False {
		workspace = ""
	}

	ctx.Inputs.Foreach(func(item actionTypes.ActionDataItem) {
		_, err := apiClient.PushEvent(context.TODO(), &services.Event{
			Source:    tiopsConfigs.WorkflowId,
			Workspace: workspace,
			Name:      ctx.ActionOptions.GetString(buildin.EventName),
			Data:      []byte(item[buildin.EventArg].(string)),
		})
		if err != nil {
			ctx.Logger.Error(err)
		}
	})
	return actionTypes.ActionDataBatch{}, nil
}

func StaticEventSender(ctx *actionTypes.BatchRequestContext) (actionTypes.ActionDataBatch, error) {

	workspace := tiopsConfigs.WorkspaceId
	if ctx.ActionOptions.GetString(buildin.Workspaced) == buildin.False {
		workspace = ""
	}
	event := ctx.ActionOptions.GetString(buildin.EventData)

	ctx.Inputs.Foreach(func(item actionTypes.ActionDataItem) {
		_, err := apiClient.PushEvent(context.TODO(), &services.Event{
			Source:    tiopsConfigs.WorkflowId,
			Workspace: workspace,
			Name:      ctx.ActionOptions.GetString(buildin.EventName),
			Data:      []byte(event),
		})
		if err != nil {
			ctx.Logger.Error(err)
		}
	})
	return actionTypes.ActionDataBatch{}, nil
}
