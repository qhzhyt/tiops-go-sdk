// @Title  action.go
// @Description  c相关数据结构和接口定义
// @Create  heyitong  2022/6/23 17:15
// @Update  heyitong  2022/6/23 17:15

package types

import (
	"tiops/common/models"
	"tiops/common/services"
)

type ActionType string

const (
	RemoteService = iota
	LocalCall
	RemoteApplication
)

type ActionControl int

const (
	ActionControlInit ActionControl = iota
	ActionControlStream
)

// ActionRequest 组件调用请求
type ActionRequest struct {
	ID     string
	Inputs map[string]*services.ActionData
	Count  int64
}

// ActionResponse 组件调用响应
type ActionResponse struct {
	ID      string
	Done    bool
	Outputs map[string]*services.ActionData
	Count   int64
}

// Action 子流程组件
type Action interface {
	Init(node *Node) error
	Call(request *ActionRequest) (*ActionResponse, error)
	CallPullStream(request *ActionRequest, callback func(res *ActionResponse, err error) bool) error
	Info() *models.ActionInfo
	Control(ctrl ActionControl, args map[string]string) error
	Type() ActionType
	Copy() Action
	GetRequiredResources(node *Node, stage int32) (*models.WorkflowResources, error)
}
