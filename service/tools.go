package service

import (
	"fmt"
	"tiops/common/services"
)

func inputLog(name string, inputData map[string]*services.ActionData) string {
	counts := map[string]int32{}
	for k, v := range inputData {
		counts[k] = v.Count
	}
	return fmt.Sprintf("Call %s with inputs: %v", name, counts)
}

func outputLog(name string, outputData map[string]*services.ActionData) string {
	counts := map[string]int32{}
	for k, v := range outputData {
		counts[k] = v.Count
	}
	return fmt.Sprintf("Call %s get results: %v", name, counts)
}