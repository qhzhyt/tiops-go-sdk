package service

import (
	"errors"
	"fmt"
)

var NoMoreDataError = errors.New("no more data")

func IsNotActionOrEngineError(i interface{}) error {
	return errors.New(fmt.Sprintf("%v is not an action or an engine!", i))
}
