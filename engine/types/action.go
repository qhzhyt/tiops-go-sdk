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

type ActionRequest struct {
	ID     string
	Inputs map[string]*services.ActionData
}

type ActionResponse struct {
	ID      string
	Done    bool
	Outputs map[string]*services.ActionData
}

type Action interface {
	Init(node *Node) error
	Call(request *ActionRequest) (*ActionResponse, error)
	CallPullStream(request *ActionRequest, callback func(res *ActionResponse, err error) bool) error
	Info() *models.ActionInfo
	Control(ctrl ActionControl, args map[string]string) error
	Type() ActionType
	Copy() Action
}
