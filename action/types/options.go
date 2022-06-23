// @Title  options.go
// @Description  处理组件选项，流程节点上为相应组件配置的参数
// @Create  heyitong  2022/6/23 15:25
// @Update  heyitong  2022/6/23 15:25

package types

import (
	"encoding/json"
	"strconv"
)

// ActionOptions 处理组件选项
type ActionOptions map[string]string

func (a ActionOptions) GetString(name string) string {
	return a[name]
}

func (a ActionOptions) GetStringOrDefault(name string, _default string) string {
	return a[name]
}

func (a ActionOptions) GetInt(name string) int {
	result, err := strconv.Atoi(a[name])
	if err != nil {
		return 0
	}
	return result
}

func (a ActionOptions) GetIntOrDefault(name string, _default int) int {
	result, err := strconv.Atoi(a[name])
	if err != nil {
		return _default
	}
	return result
}

func (a ActionOptions) GetFloat(name string) float64 {
	result, err := strconv.ParseFloat(a[name], 64)
	if err != nil {
		return 0
	}
	return result
}

func (a ActionOptions) GetFloatOrDefault(name string, _default float64) float64 {
	result, err := strconv.ParseFloat(a[name], 64)
	if err != nil {
		return _default
	}
	return result
}

func (a ActionOptions) GetStringList(name string) []string {
	var result []string
	err := json.Unmarshal([]byte(a[name]), &result)
	if err != nil {
		return nil
	}
	return result
}

func (a ActionOptions) GetStringListOrDefault(name string, _default []string) []string {
	var result []string
	err := json.Unmarshal([]byte(a[name]), &result)
	if err != nil {
		return _default
	}
	return result
}

func (a ActionOptions) GetList(name string) ([]interface{}, error) {
	var result []interface{}
	err := json.Unmarshal([]byte(a[name]), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a ActionOptions) GetListOrDefault(name string, _default []interface{}) []interface{} {
	var result []interface{}
	err := json.Unmarshal([]byte(a[name]), &result)
	if err != nil {
		return _default
	}
	return result
}

func (a ActionOptions) GetMap(name string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(a[name]), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a ActionOptions) GetMapOrDefault(name string, _default map[string]interface{}) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(a[name]), &result)
	if err != nil {
		return _default
	}
	return result
}
