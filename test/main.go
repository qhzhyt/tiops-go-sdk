package main

import (
	"fmt"
	"tiops/action"
	"tiops/action/types"
)

func main() {
	items, err := action.NewCommandAction("jq -c", action.CommandActionOptions{BatchMode: action.DelimiterSplit, Delimiter: '\n'}).ProcessItems(types.ActionDataBatch{
		types.ActionDataItem{"in": []byte("{\"a\": 666}")},
		types.ActionDataItem{"in": []byte("{\"a\": 666}")},
		types.ActionDataItem{"in": []byte("{\"a\": 666}")},
	})

	if err != nil {
		fmt.Println(err)
	} else {
		for _, item := range items {
			fmt.Print(string(item["out"].([]byte)))
		}
	}
}
