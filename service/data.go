package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"tiops/common/models"
	"tiops/common/services"
)

func repeat(item interface{}, count int) []interface{} {
	result := make([]interface{}, count)
	for i := 0; i < count; i++ {
		result[i] = item
	}
	return result
}

type ActionData interface {
	List() []interface{}
	Count() int
	Foreach(func(item interface{}))
	Map(func(item interface{}) interface{}) []interface{}
	RawData() *services.ActionData
	Item(i int) interface{}
}

//type ActionDataItemMap map[string]interface{}

type ActionDataItemsMap map[string][]interface{}

type ConstantData struct {
	count     int
	valueType string
	value     interface{}
	rawData   *services.ActionData
}

func (c *ConstantData) Item(i int) interface{} {
	return c.value
}

func (c *ConstantData) RawData() *services.ActionData {
	return c.rawData
}

func (c *ConstantData) List() []interface{} {
	return repeat(c.value, c.count)
}

func (c *ConstantData) Count() int {
	return c.count
}

func (c *ConstantData) Foreach(f func(item interface{})) {
	for i := 0; i < c.count; i++ {
		f(c.value)
	}
}

func (c *ConstantData) Map(f func(item interface{}) interface{}) []interface{} {
	result := make([]interface{}, c.count)
	for i := 0; i < c.count; i++ {
		result[i] = f(c.value)
	}
	return result
}

type ListData struct {
	rawData    *services.ActionData
	rawDataLen int
	count      int
	valueType  string
}

func (l *ListData) Item(i int) interface{} {
	if i < l.rawDataLen {
		if stringTypes[l.valueType] {
			return string(l.rawData.Data[i])
		}
		return l.rawData.Data[i]
	}
	return nil
}

func (l *ListData) RawData() *services.ActionData {

	return l.rawData
}

func (l *ListData) List() []interface{} {
	result := make([]interface{}, l.count)
	if stringTypes[l.valueType] {
		for i := 0; i < len(l.rawData.Data); i++ {
			result[i] = string(l.rawData.Data[i])
		}
	} else {
		for i := 0; i < len(l.rawData.Data); i++ {
			result[i] = l.rawData.Data[i]
		}
	}

	return result
}

func (l *ListData) Count() int {
	return l.count
}

func (l *ListData) Foreach(f func(item interface{})) {
	for i := 0; i < l.count; i++ {
		if i < l.rawDataLen {
			f(l.rawData.Data[i])
		} else {
			f(nil)
		}
	}
}

func (l *ListData) Map(f func(item interface{}) interface{}) []interface{} {
	result := make([]interface{}, l.count)
	for i := 0; i < l.count; i++ {
		if i < l.rawDataLen {
			result[i] = f(l.rawData.Data[i])
		} else {
			result[i] = f(nil)
		}
	}
	return result
}

func NewConstantData(data *services.ActionData) *ConstantData {

	result := &ConstantData{
		value:     data.Data[0],
		rawData:   data,
		count:     int(data.Count),
		valueType: strings.ToLower(data.ValueType),
	}
	if stringTypes[result.valueType] {
		result.value = string(data.Data[0])
	}

	return result
}

func NewListData(data *services.ActionData) *ListData {
	return &ListData{
		rawData:    data,
		rawDataLen: len(data.Data),
		count:      int(data.Count),
		valueType:  strings.ToLower(data.ValueType),
	}
}

func TransActionData(data *services.ActionData) ActionData {
	switch data.Type {
	case services.DataType_Constant:
		return NewConstantData(data)
	default:
		return NewListData(data)
	}
}

type ActionDataMap map[string]ActionData

func TransActionDataMap(dataMap map[string]*services.ActionData, inputs []*models.Parameter, ) ActionDataMap {
	result := ActionDataMap{}

	for _, input := range inputs {
		srcData := dataMap[input.Name]
		if srcData != nil {
			srcData.ValueType = input.Type
			result[input.Name] = TransActionData(srcData)
		}
	}

	//for k, v := range dataMap {
	//	result[k] = TransActionData(v)
	//}
	return result
}

func (m ActionDataMap) Item(i int) ActionDataItem {
	result := map[string]interface{}{}
	for k, v := range m {
		result[k] = v.Item(i)
	}
	return result
}

func (m ActionDataMap) Map(f func(item ActionDataItem) ActionDataItem) ActionDataBatch {
	count := m.Count()
	result := make(ActionDataBatch, count)
	for i := 0; i < count; i++ {
		result[i] = f(m.Item(i))
	}
	return result
}

func (m ActionDataMap) MapTrans(f func(map[string]interface{}) map[string]interface{}, keys []string) map[string][]interface{} {
	result := map[string][]interface{}{}
	count := m.Count()

	for _, key := range keys {
		result[key] = make([]interface{}, count)
	}
	var resultItem map[string]interface{}
	for i := 0; i < count; i++ {
		resultItem = f(m.Item(i))
		for _, key := range keys {
			result[key][i] = resultItem[key]
		}
	}
	return result
}

func (m ActionDataMap) Foreach(f func(map[string]interface{})) {
	count := m.Count()
	for i := 0; i < count; i++ {
		f(m.Item(i))
	}
}

func (m ActionDataMap) Count() int {
	count := 0
	for _, v := range m {
		if v.Count() > count {
			count = v.Count()
		}
	}
	return count
}

func (m ActionDataMap) Items() chan map[string]interface{} {
	ch := make(chan map[string]interface{})

	go func() {
		count := m.Count()

		for i := 0; i < count; i++ {
			ch <- m.Item(i)
		}

		close(ch)
	}()

	return ch
}

//type ActionDataMap struct {
//	actionDataMap
//}
//
//func init() {
//	a := ActionDataMap{}
//	for k, v := range actionDataMap(a) {
//
//	}
//}

var (
	stringTypes = map[string]bool{
		"str":    true,
		"string": true,
		"text":   true,
		"json":   true,
	}
)

func ToServiceActionDataMap(id string, traceId int64, source ActionDataBatch, outputs []*models.Parameter) ServiceActionDataMap {
	result := ServiceActionDataMap{}
	for _, output := range outputs {
		count := len(source)
		items := make([][]byte, count)
		//dataCol := source[output.Name]
		name := output.Name
		if stringTypes[strings.ToLower(output.Type)] {
			for i := 0; i < count; i++ {
				items[i] = []byte(source[i][name].(string))
			}
		} else if strings.ToLower(output.Type) == "bytes" {
			for i := 0; i < count; i++ {
				items[i] = source[i][name].([]byte)
			}
		} else {
			for i := 0; i < count; i++ {
				res, err := json.Marshal(source[i][name])
				if err != nil {
					items[i] = []byte(fmt.Sprint(source[i][name]))
				} else {
					items[i] = res
				}
			}
		}
		result[output.Name] = &services.ActionData{
			Id: id,
			Count:     int32(count),
			Data:      items,
			ValueType: output.Type,
			TraceId: traceId,
		}
	}
	return result
}
