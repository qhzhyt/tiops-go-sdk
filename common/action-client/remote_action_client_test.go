package action_client

import (
	"context"
	"testing"
	"tiops/common/services"
)

func TestNewRemoteActionClient(t *testing.T) {
	client := NewRemoteActionClient("localhost", 5555)

	res, err := client.CallAction(context.TODO(), &services.ActionRequest{
		Id:     "",
		ActionName:   "hello",
		Inputs: map[string]*services.ActionData{
			"name": {Count: 2,
				Type: services.DataType_List,
				ValueType: "str",
				Data: [][]byte{[]byte("bob"), []byte("jack")},
			},
		},
	})
	t.Error(res, err)
}
