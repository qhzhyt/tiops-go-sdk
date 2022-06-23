// @Title  http.go
// @Description  HTTP相关组件实现
// @Create  heyitong  2022/6/23 15:42
// @Update  heyitong  2022/6/23 15:42

package actions

import (
	"context"
	"encoding/json"
	actionTypes "tiops/action/types"
	"tiops/buildin"
	"tiops/common/services"
)

// HttpTrigger HTTP触发器，接受http请求事件
func HttpTrigger(ctx *actionTypes.StreamRequestContext) error {

	eventClient, err := apiClient.WatchEvent(context.TODO(), &services.WatchEventRequest{
		Workspace: buildin.HttpWorkspace,
		Name:      ctx.ActionOptions.GetStringOrDefault(buildin.Path, buildin.DefaultPath),
	})

	if err != nil {
		return err
	}

	headerName := ctx.ActionOptions.GetStringOrDefault(buildin.HeaderName, buildin.All)
	queryName := ctx.ActionOptions.GetStringOrDefault(buildin.QueryName, buildin.All)

	for true {
		event, err := eventClient.Recv()
		if err != nil {
			ctx.Logger.Error(err)
		} else {

			httpRequest := &services.HttpRequest{}

			_ = json.Unmarshal(event.Data, httpRequest)

			h := ""

			header := services.ServiceHeadersToHttpHeader(httpRequest.Headers)

			if headerName == buildin.All {
				hBytes, _ := json.Marshal(header)
				h = string(hBytes)
			} else {
				h = header.Get(headerName)
			}

			q := ""

			if queryName == buildin.All {
				hBytes, _ := json.Marshal(httpRequest.Query)
				q = string(hBytes)
			} else {
				q = httpRequest.Query[queryName]
			}

			err = ctx.Push(actionTypes.ActionDataBatch{map[string]interface{}{
				buildin.M:   httpRequest.Method,
				buildin.Req: string(event.Data),
				buildin.Q:   q,
				buildin.H:   h,
				buildin.B:   httpRequest.Body,
			},
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
