package utils

import (
	"fmt"
	"testing"
)

func TestSnowflakeID(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Log(fmt.Sprintf("%x\n", SnowflakeID()))

	}

}
