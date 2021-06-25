package variable

import "tiops/engine/types"

const (
	ConstantType = "constant"
)

func New(name, _type, valueType, value string) types.Variable {
	switch _type {
	case ConstantType:
		return NewConstantVariable(name, valueType, value)
	default:
		return NewConstantVariable(name, valueType, value)
	}
}
