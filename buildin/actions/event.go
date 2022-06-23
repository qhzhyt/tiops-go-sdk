// @Title  event.go
// @Description  事件相关组件实现
// @Create  heyitong  2022/6/23 15:36
// @Update  heyitong  2022/6/23 15:36

package actions

import (
	"context"
	actionTypes "tiops/action/types"
	"tiops/buildin"
	tiopsConfigs "tiops/common/config"
	"tiops/common/services"
)

// EventTrigger 事件触发器，接收到eventName对应事件后输出一条数据
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

// EventSender 事件生成器，接收一条数据后将其封装成eventName对应的事件并发送
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

// StaticEventSender 固定事件生成器，接收一条数据后发送eventName对应的事件，事件内容为eventData
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
