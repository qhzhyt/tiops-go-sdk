package variable

import (
	"encoding/json"
	"strconv"
	"tiops/common/services"
	"tiops/engine/types"
)

type ConstantVariable struct {
	name  string
	_type string
	value string
}

func (c *ConstantVariable) GetMap() map[string]interface{} {
	a := map[string]interface{}{}
	json.Unmarshal([]byte(c.value), &a)
	return a
}

func (c *ConstantVariable) ToActionArguments(count int) *services.ActionData {
	value := c.value
	data := make([][]byte, 1)
	data[0]= []byte(value)
	return &services.ActionData{Type: services.DataType_Constant, Data: data, ValueType: c._type, Count: int32(count)}
}

func (c *ConstantVariable) SetValue(s string) {
	panic("Constant can't be modified")
}

func (c *ConstantVariable) Name() string {
	return c.name
}

func (c *ConstantVariable) Type() string {
	return c._type
}

func (c *ConstantVariable) Value() interface{} {
	//fmt.Println(c._type)
	switch c._type {
	case "int":
		return c.GetInt()
	case "bool":
		return c.GetBool()
	case "map":
		return c.GetMap()
	default:
		return c.GetString()
	}
}

func (c *ConstantVariable) GetInt() int {
	a, err := strconv.Atoi(c.value)
	if err != nil {
		return 0
	}
	return a
}

func (c *ConstantVariable) GetString() string {
	return c.value
}

func (c *ConstantVariable) GetBool() bool {
	a, err := strconv.ParseBool(c.value)
	if err != nil {
		return false
	}
	return a
}


func NewConstantVariable(name, _type, value string) types.Variable {
	return &ConstantVariable{
		name:  name,
		_type: _type,
		value: value,
	}
}
