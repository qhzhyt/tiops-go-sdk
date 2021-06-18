package utils

import (
	"testing"
	"time"
)

func TestParseCrontabText(t *testing.T) {

	s, err := ParseCrontabText("0 0 0 * * ?")
	t.Log(err)
	if err == nil {
		t.Log(s.Next(time.Now()))
	}
}
