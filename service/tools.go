// @Title  tools.go
// @Description  ActionService相关工具方法
// @Create  heyitong  2022/6/17 17:31
// @Update  heyitong  2022/6/17 17:31

package service

import (
	"fmt"
	"tiops/common/services"
)

// successLog 生成成功调用日志
func successLog(name string, inputData map[string]*services.ActionData, outputData map[string]*services.ActionData) string {
	inputCounts := map[string]int64{}
	for k, v := range inputData {
		inputCounts[k] = v.Count
	}

	outputCounts := map[string]int64{}
	for k, v := range outputData {
		outputCounts[k] = v.Count
	}

	return fmt.Sprintf("Call %s with inputs: %v, get outputs: %v", name, inputCounts, outputCounts)
}

// errorLog 生成调用失败错误日志
func errorLog(name string, inputData map[string]*services.ActionData, err error) string {
	inputCounts := map[string]int64{}
	for k, v := range inputData {
		inputCounts[k] = v.Count
	}
	return fmt.Sprintf("Call %s with inputs: %v, get error: %v", name, inputCounts, err)
}
