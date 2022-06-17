// @Title  errors.go
// @Description  ActionService相关错误
// @Create  heyitong  2022/6/17 17:30
// @Update  heyitong  2022/6/17 17:30

package service

import (
	"errors"
	"fmt"
)

var NoMoreDataError = errors.New("no more data")

func IsNotActionOrEngineError(i interface{}) error {
	return errors.New(fmt.Sprintf("%v is not an action or an engine!", i))
}
