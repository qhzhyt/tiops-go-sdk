package engines

import (
	"fmt"
	"tiops/common/models"
	"tiops/common/services"
)

func itemCount(inputData map[string]*services.ActionData) int32 {
	count := int32(0)
	for _, data := range inputData {
		if data != nil && data.Count > count {
			count = data.Count
		}
	}

	return count
}

func inputLog(info *models.ActionInfo, requestId string, inputData map[string]*services.ActionData) string {
	infos := map[string]map[string]interface{}{}
	for k, v := range inputData {
		if v != nil {
			infos[k] = map[string]interface{}{
				"id":    v.Id,
				"count": v.Count,
			}
		} else {
			infos[k] = map[string]interface{}{
				"id":    "no data",
				"count": "0",
			}
		}
	}
	return fmt.Sprintf("Action request %s to %s with inputs: %v", requestId, info.Name, infos)
}

func errorLog(info *models.ActionInfo, err error, inputData map[string]*services.ActionData) string {
	infos := map[string]map[string]interface{}{}
	for k, v := range inputData {
		if v != nil {
			infos[k] = map[string]interface{}{
				"id":      v.Id,
				"count":   v.Count,
				"example": v.Data[0],
			}
		} else {
			infos[k] = map[string]interface{}{
				"id":    "no data",
				"count": "0",
			}
		}
	}
	return fmt.Sprintf("Call action %s get error %s with inputs: %v", info.Name, err.Error(), infos)
}

func outputLog(info *models.ActionInfo, responseId string, outputData map[string]*services.ActionData) string {
	infos := map[string]map[string]interface{}{}
	for k, v := range outputData {
		infos[k] = map[string]interface{}{
			"id":    v.Id,
			"count": v.Count,
		}
	}
	return fmt.Sprintf("Action reponse %s from %s with outputs: %v", responseId, info.Name, infos)
}