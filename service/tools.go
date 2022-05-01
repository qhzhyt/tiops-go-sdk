package service

import (
	"fmt"
	"tiops/common/services"
)

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

func errorLog(name string, inputData map[string]*services.ActionData, err error) string {
	inputCounts := map[string]int64{}
	for k, v := range inputData {
		inputCounts[k] = v.Count
	}
	return fmt.Sprintf("Call %s with inputs: %v, get error: %v", name, inputCounts, err)
}
