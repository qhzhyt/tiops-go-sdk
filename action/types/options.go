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
