// @Title  common.go
// @Description  处理流程变量工厂
// @Create  heyitong  2022/6/23 17:29
// @Update  heyitong  2022/6/23 17:29

package variable

import "tiops/engine/types"

const (
	ConstantType = "constant"
)

// New 处理流程变量工厂函数
func New(name, _type, valueType, value string) types.Variable {
	switch _type {
	case ConstantType:
		return NewConstantVariable(name, valueType, value)
	default:
		return NewConstantVariable(name, valueType, value)
	}
}
