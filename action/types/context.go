// @Title  context.go
// @Description  组件上下文数据结构定义
// @Create  heyitong  2022/6/23 15:13
// @Update  heyitong  2022/6/23 15:13

package types

import (
	"context"
	"net/http"
	"tiops/common/logger"
	"tiops/common/models"
	"tiops/common/services"
	"tiops/common/stores"
)

// ActionContext 组件全局上下文
type ActionContext struct {
	Logger      *logger.Logger
	Info        *models.ActionInfo
	InputNames  []string
	OutputNames []string
}

// ActionNodeContext 流程节点上下文
type ActionNodeContext struct {
	*ActionContext
	Store           stores.DataStore
	NodeId          string
	ActionOptions   ActionOptions
	InnerActionInfo *models.ActionInfo
	done            bool
}

// PieceRequestContext 单条数据处理上下文
type PieceRequestContext struct {
	*ActionNodeContext
	//Store         stores.DataStore
	//NodeId        string
	Input ActionDataItem
	//ActionOptions ActionOptions
}

// BatchRequestContext 披露那个数据处理上下文
type BatchRequestContext struct {
	context *context.Context
	*ActionNodeContext
	Inputs ActionDataMap
	//ActionOptions ActionOptions
	//done bool
}

// HttpRequestContext HTTP处理上下文
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

// HttpResponse HTTP响应封装
type HttpResponse struct {
	Status      int32
	Header      http.Header
	Body        []byte
	ContentType string
	//ActionOptions ActionOptions
	//done bool
}

// StreamRequestContext 一对多流式数据处理上下文
type StreamRequestContext struct {
	BatchRequestContext
	Push func(data ActionDataBatch) error
}

/*func (s *StreamRequestContext) Push(data ActionData) error {

}*/

// InitContext 组件初始化上下文
type InitContext struct {
	*ActionContext
}

// NodeRegisterContext 流程节点注册上下文
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

// PushMessageContext 推送消息上下文
type PushMessageContext struct {
	*ActionNodeContext
	NodeId        string
	MessageHeader ActionOptions
	MessageData   []byte
}
