package types

import (
	"context"
	"net/http"
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
	context *context.Context
	*ActionNodeContext
	Inputs ActionDataMap
	//ActionOptions ActionOptions
	//done bool
}

type HttpRequestContext struct {
	*ActionNodeContext
	Method string
	Path   string
	Header http.Header
	Query  map[string]string
	Body   []byte
	//ActionOptions ActionOptions
	//done bool
}

type HttpResponse struct {
	Status      int32
	Header      http.Header
	Body        []byte
	ContentType string
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
