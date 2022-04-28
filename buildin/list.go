package buildin

import (
	"time"
	"tiops/common/models"
)

const (
	EventTriggerID      = "buildin---event-trigger"
	StaticEventSenderID = "buildin---static-event-sender"
	EventSenderID       = "buildin---event-sender"
	HttpTriggerID       = "buildin---http-trigger"
)

const (
	EventName     = "eventName"
	Workspaced    = "workspaced"
	EventData     = "eventData"
	EventArg      = "evt"
	Req           = "req"
	M             = "m"
	B             = "b"
	Q             = "q"
	H             = "h"
	Path          = "path"
	DefaultPath   = "test"
	All           = "*"
	QueryName     = "queryName"
	HeaderName    = "headerName"
	HttpWorkspace = "__http__"
)

const (
	True             = "true"
	False            = "false"
	String           = "string"
	Bytes            = "bytes"
	Json             = "json"
	Select           = "select"
	DefaultEventName = "test"
	DefaultEventData = "test"
)

var (
	ActionList = []*models.ActionInfo{
		{
			XId:         EventTriggerID,
			Name:        "event-trigger",
			DisplayName: "事件触发器",
			IsPublic:    true,
			CreatedBy:   "__system__",
			CreatedTime: time.Date(2022, 4, 27, 21, 38, 00, 0, time.Local).UnixNano() / 1e6,
			Source:      models.ActionSource_Buildin,
			Type:        models.ActionType_BuildInAction,
			CallMode:    models.CallMode_PullStreamCall,
			Tags:        []string{"trigger"},
			Description: "系统事件触发器",
			Outputs:     []*models.Parameter{{Name: EventArg, Type: String, Default: "", Description: "事件内容"}},
			Options: []*models.ActionOption{
				{Name: EventName, Type: String, Default: DefaultEventName, Description: "监听的事件名"},
				{Name: Workspaced, Type: Select, Default: True, Description: "仅作用于工作空间", Options: []*models.ActionOptionItem{{Label: True, Value: True}, {Label: False, Value: False}}},
			},
		},
		{
			XId:         StaticEventSenderID,
			Name:        "static-event-sender",
			DisplayName: "特定事件生成器",
			IsPublic:    true,
			CreatedBy:   "__system__",
			CreatedTime: time.Date(2022, 4, 28, 12, 29, 00, 0, time.Local).UnixNano() / 1e6,
			Source:      models.ActionSource_Buildin,
			Type:        models.ActionType_BuildInAction,
			CallMode:    models.CallMode_OnceCall,
			Tags:        []string{"event-sender"},
			Description: "特定事件生成器",
			Inputs:      []*models.Parameter{{Name: "trg", Type: String, Default: "", Description: "触发"}},
			Options: []*models.ActionOption{
				{Name: EventName, Type: String, Default: DefaultEventName, Description: "触发的事件名"},
				{Name: EventData, Type: String, Default: DefaultEventData, Description: "事件数据"},
				{Name: Workspaced, Type: Select, Default: True, Description: "仅作用于工作空间", Options: []*models.ActionOptionItem{{Label: True, Value: True}, {Label: False, Value: False}}},
			},
		},
		{
			XId:         EventSenderID,
			Name:        "event-sender",
			DisplayName: "事件生成器",
			IsPublic:    true,
			CreatedBy:   "__system__",
			CreatedTime: time.Date(2022, 4, 28, 10, 17, 00, 0, time.Local).UnixNano() / 1e6,
			Source:      models.ActionSource_Buildin,
			Type:        models.ActionType_BuildInAction,
			CallMode:    models.CallMode_OnceCall,
			Tags:        []string{"event-sender"},
			Description: "事件生成器",
			Inputs:      []*models.Parameter{{Name: EventArg, Type: String, Default: "", Description: "事件内容"}},
			Options: []*models.ActionOption{
				{Name: EventName, Type: String, Default: DefaultEventName, Description: "触发的事件名"},
				{Name: Workspaced, Type: Select, Default: True, Description: "仅作用于工作空间", Options: []*models.ActionOptionItem{{Label: True, Value: True}, {Label: False, Value: False}}},
			},
		},
		{
			XId:         HttpTriggerID,
			Name:        "http-trigger",
			DisplayName: "HTTP请求触发器",
			IsPublic:    true,
			CreatedBy:   "__system__",
			CreatedTime: time.Date(2022, 4, 29, 10, 37, 00, 0, time.Local).UnixNano() / 1e6,
			Source:      models.ActionSource_Buildin,
			Type:        models.ActionType_DataSourceAction,
			CallMode:    models.CallMode_PullStreamCall,
			Tags:        []string{"trigger"},
			Description: "HTTP请求触发器",
			Outputs: []*models.Parameter{
				{Name: Req, Type: Json, Default: "", Description: "请求内容"},
				{Name: M, Type: String, Default: "", Description: "请求方法"},
				{Name: Q, Type: String, Default: "", Description: "Get参数"},
				{Name: H, Type: String, Default: "", Description: "请求头"},
				{Name: B, Type: Bytes, Default: "", Description: "请求体"},
			},
			Options: []*models.ActionOption{
				{Name: Path, Type: String, Default: DefaultPath, Description: "Http路径"},
				{Name: HeaderName, Type: String, Default: All, Description: "请求头字段名"},
				{Name: QueryName, Type: String, Default: All, Description: "Get参数字段名"},
			},
		},
	}
)
