package actions

import (
	"tiops/action/types"
	"tiops/common/models"
)

type BuildinAction struct {
	Processor func(types.ActionDataBatch, types.ActionOptions) types.ActionDataBatch
	ActionInfo *models.ActionInfo
}

var BuildinActionMap = map[string]*BuildinAction{

}