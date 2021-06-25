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
	Info() *models.ActionInfo
	Type() ActionType
	Copy() Action
}
