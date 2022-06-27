// @Title  list.go
// @Description  内置组件列表
// @Create  heyitong  2022/6/23 15:29
// @Update  heyitong  2022/6/23 15:29

package buildin

import (
	"fmt"
	"strings"
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
	JsonSetFieldID   = "buildin---json-set-field"
	JsonPathLookupID = "buildin---json-path-lookup"
)

const (
	EventName         = "eventName"
	Workspaced        = "workspaced"
	Input             = "in"
	Output            = "out"
	FieldName         = "fieldName"
	Value             = "value"
	FieldType         = "fieldType"
	TimestampType     = "timestampType"
	EventData         = "eventData"
	EventArg          = "evt"
	Req               = "req"
	M                 = "m"
	B                 = "b"
	Q                 = "q"
	H                 = "h"
	Path              = "path"
	DefaultPath       = "test"
	All               = "*"
	QueryName         = "queryName"
	HeaderName        = "headerName"
	HttpWorkspace     = "__http__"
	DefaultEngineName = "default"
	JsonPath          = "jsonPath"
)

const (
	True                 = "true"
	False                = "false"
	String               = "string"
	Bytes                = "bytes"
	Json                 = "json"
	Select               = "select"
	DefaultEventName     = "test"
	DefaultEventData     = "test"
	ID                   = "id"
	UUID                 = "uuid"
	DateTime             = "dateTime"
	List                 = "list"
	Object               = "object"
	Integer              = "integer"
	Float                = "float"
	RandomInt            = "randomInt"
	RandomString         = "randomString"
	Timestamp            = "timestamp"
	Second               = "second"
	Millisecond          = "millisecond"
	TimeFormat           = "timeFormat"
	RandomCharset        = "randomCharset"
	DefaultRandomCharset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	MinInt               = "minInt"
	MaxInt               = "maxInt"
)

var (
	ActionList = []*models.ActionInfo{
		{
			XId:         EventTriggerID,
			Name:        "event-trigger",
			DisplayName: "事件触发器",
			IsPublic:    true,
			CreatedBy:   "__system__",
			Engine:      DefaultEngineName,
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
			Engine:      DefaultEngineName,
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
			Engine:      DefaultEngineName,
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
			Engine:      DefaultEngineName,
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
		{
			XId:         JsonSetFieldID,
			Name:        "json-set-field",
			DisplayName: "设置JSON字段",
			IsPublic:    true,
			CreatedBy:   "__system__",
			Engine:      DefaultEngineName,
			CreatedTime: time.Date(2022, 6, 24, 9, 31, 00, 0, time.Local).UnixNano() / 1e6,
			Source:      models.ActionSource_Buildin,
			Type:        models.ActionType_BuildInAction,
			CallMode:    models.CallMode_OnceCall,
			Tags:        []string{"json"},
			Description: "为JSON设置字段",
			Inputs: []*models.Parameter{
				{Name: Input, Type: Json, Default: "", Description: "输入数据"},
			},
			Outputs: []*models.Parameter{
				{Name: Output, Type: Json, Default: "", Description: "输出数据"},
			},
			Options: []*models.ActionOption{
				{Name: FieldName, Type: String, Default: ID, Description: "字段名", Required: true},
				{Name: FieldType, Type: Select, Default: UUID, Description: "字段类型", Required: true,
					Options: []*models.ActionOptionItem{
						{Label: "UUID", Value: UUID},
						{Label: "格式化时间", Value: DateTime},
						{Label: "时间戳", Value: Timestamp},
						{Label: "随机整数", Value: RandomInt},
						{Label: "随机字符串", Value: RandomString},
						{Label: "整数", Value: Integer},
						{Label: "浮点数", Value: Float},
						{Label: "列表", Value: List},
						{Label: "对象", Value: Object},
					},
				},
				{Name: Value, Type: Json, Default: "0", Description: "设置值",
					Show: showIfFieldIn(FieldType, []string{Integer, Float, List, Object}),
				},
				{Name: TimestampType, Type: Select, Default: Millisecond, Description: "时间戳类型",
					Show: showIfFieldIs(FieldType, Timestamp),
					Options: []*models.ActionOptionItem{
						{Label: "秒", Value: Second},
						{Label: "毫秒", Value: Millisecond},
					},
				},
				{Name: TimeFormat, Type: String, Default: time.RFC3339, Description: "时间格式",
					Show: showIfFieldIs(FieldType, DateTime),
				},
				{Name: RandomCharset, Type: String, Default: DefaultRandomCharset, Description: "随机字符集",
					Show: showIfFieldIs(FieldType, RandomString),
				},
				{Name: MinInt, Type: Integer, Default: "0", Description: "最小数",
					Show: showIfFieldIs(FieldType, RandomInt),
				},
				{Name: MaxInt, Type: Integer, Default: "65535", Description: "最大数",
					Show: showIfFieldIs(FieldType, RandomInt),
				},
			},
		},
		{
			XId:         JsonPathLookupID,
			Name:        "json-path-lookup",
			DisplayName: "JsonPath提取",
			IsPublic:    true,
			CreatedBy:   "__system__",
			Engine:      DefaultEngineName,
			CreatedTime: time.Date(2022, 6, 27, 17, 35, 00, 0, time.Local).UnixNano() / 1e6,
			Source:      models.ActionSource_Buildin,
			Type:        models.ActionType_BuildInAction,
			CallMode:    models.CallMode_OnceCall,
			Tags:        []string{"json"},
			Description: "根据JsonPath提取Json内容",
			Inputs: []*models.Parameter{
				{Name: Input, Type: Json, Default: "", Description: "输入数据"},
			},
			Outputs: []*models.Parameter{
				{Name: Output, Type: Json, Default: "", Description: "输出数据"},
			},
			Options: []*models.ActionOption{
				{Name: JsonPath, Type: String, Default: "$", Description: "JsonPath", Required: true},
			},
		},
	}
)

func showIfFieldIs(field string, value string) string {
	return fmt.Sprintf("(configs) => configs.%s === \"%s\"", field, value)
}

func showIfFieldIn(field string, values []string) string {
	return fmt.Sprintf("(configs) => [\"%s\"].indexOf(configs.%s) > -1", strings.Join(values, "\", \""), field)
}
