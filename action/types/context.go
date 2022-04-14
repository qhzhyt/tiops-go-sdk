package types

import (
	"tiops/common/logger"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/stores"
)

type ActionContext struct {
	Logger      *logger.Logger
	Info        *models.ActionInfo
	InputNames  []string
	OutputNames []string
}

type ActionNodeContext struct {
	*ActionContext
	Store           stores.DataStore
	NodeId          string
	ActionOptions   ActionOptions
	InnerActionInfo *models.ActionInfo
	done            bool
}

type PieceRequestContext struct {
	*ActionNodeContext
	//Store         stores.DataStore
	//NodeId        string
	Input ActionDataItem
	//ActionOptions ActionOptions
}

type BatchRequestContext struct {
	*ActionNodeContext
	Inputs ActionDataMap
	//ActionOptions ActionOptions
	//done bool
}

type StreamRequestContext struct {
	BatchRequestContext
	Push func(data ActionDataBatch) error
}

/*func (s *StreamRequestContext) Push(data ActionData) error {

}*/

type InitContext struct {
	*ActionContext
}

type NodeRegisterContext struct {
	*ActionNodeContext
	//Store         stores.DataStore
	//NodeId        string
	//ActionOptions ActionOptions
	NextActions []*services.NextActions
}

func (ac *ActionNodeContext) Done() {
	ac.done = true
}

func (ac *ActionNodeContext) HasDone() bool {
	return ac.done
}

func (ac *ActionNodeContext) Copy() *ActionNodeContext {
	return &ActionNodeContext{ActionContext: ac.ActionContext}
}

type PushMessageContext struct {
	*ActionNodeContext
	NodeId        string
	MessageHeader ActionOptions
	MessageData   []byte
}
