package action

import (
	"testing"
	"tiops/action/types"
)

func TestCmd(t *testing.T) {
	t.Error(string(NewCommandAction("jq", CommandActionOptions{}).Call(&types.PieceRequestContext{
		Input: types.ActionDataItem{"in": []byte("{\"a\": 666}")},
	})["out"].([]byte)))
}
