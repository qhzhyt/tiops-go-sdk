package actions

import (
	"github.com/qhzhyt/tiops-go-sdk/common/models"
	"github.com/qhzhyt/tiops-go-sdk/service"
)

type BuildinAction struct {
	Processor func(service.ActionDataBatch, service.ActionOptions) service.ActionDataBatch
	ActionInfo *models.ActionInfo
}

var BuildinActionMap = map[string]*BuildinAction{

}