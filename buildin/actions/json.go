// @Title  json.go
// @Description  JSON相关组件实现
// @Create  heyitong  2022/6/23 15:44
// @Update  heyitong  2022/6/23 15:44

package actions

import (
	"github.com/google/uuid"
	"time"
	"tiops/action/types"
	actionTypes "tiops/action/types"
	"tiops/buildin"
)

// JsonSetField 设置JSON字段
func JsonSetField(ctx *actionTypes.BatchRequestContext) (actionTypes.ActionDataBatch, error) {

	fieldName := ctx.ActionOptions.GetStringOrDefault(buildin.FieldName, buildin.ID)
	fieldType := ctx.ActionOptions.GetStringOrDefault(buildin.FieldType, buildin.UUID)
	timeFormat := ctx.ActionOptions.GetStringOrDefault(buildin.TimeFormat, time.RFC3339)

	return ctx.Inputs.MapExistError(func(item types.ActionDataItem) (types.ActionDataItem, error) {
		input := item[buildin.Input]

		var value interface{}
		var err error

		switch fieldType {
		case buildin.Object:
			value, err = ctx.ActionOptions.GetMap(buildin.Value)
			if err != nil {
				return nil, err
			}
		case buildin.List:
			value, err = ctx.ActionOptions.GetList(buildin.Value)
			if err != nil {
				return nil, err
			}
		case buildin.Integer:
			value = ctx.ActionOptions.GetInt(buildin.Value)
		case buildin.Float:
			value = ctx.ActionOptions.GetFloat(buildin.Value)
		case buildin.String:
			value = ctx.ActionOptions.GetString(buildin.Value)
		case buildin.Timestamp:
			if ctx.ActionOptions.GetString(buildin.TimestampType) == buildin.Second {
				value = time.Now().Second()
			} else {
				value = time.Now().UnixNano() / 1e6
			}
		case buildin.DateTime:
			value = time.Now().Format(timeFormat)
		case buildin.UUID:
			fallthrough
		default:
			value = uuid.New().String()
		}

		output, err := JsonProcess(input.(string), func(in map[string]interface{}) (map[string]interface{}, error) {
			in[fieldName] = value
			return in, nil
		})

		if err != nil {
			return nil, err
		}

		return types.ActionDataItem{
			buildin.Output: output,
		}, nil

	})

}
