// @Title  variable.go
// @Description  处理流程变量接口定义
// @Create  heyitong  2022/6/23 17:24
// @Update  heyitong  2022/6/23 17:24


package types

import "tiops/common/services"

// Variable 处理流程变量接口
type Variable interface {
	Name() string
	Type() string
	Value() interface{}
	GetInt() int
	GetString() string
	GetBool() bool
	SetValue(string)
	ToActionArguments(count int64) *services.ActionData
	GetMap() map[string]interface{}
	//Map() map[string] interface{}
}
