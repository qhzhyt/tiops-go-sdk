package types

import "tiops/common/services"

type Variable interface {
	Name() string
	Type() string
	Value() interface{}
	GetInt() int
	GetString() string
	GetBool() bool
	SetValue(string)
	ToActionArguments(count int) *services.ActionData
	GetMap() map[string]interface{}
	//Map() map[string] interface{}
}
