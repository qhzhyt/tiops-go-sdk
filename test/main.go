package main

func main() {
	//items, err := action.NewCommandAction("jq -c", action.CommandActionOptions{BatchMode: action.DelimiterSplit, Delimiter: '\n'}).ProcessItems(types.ActionDataBatch{
	//	types.ActionDataItem{"in": []byte("{\"a\": 666}")},
	//	types.ActionDataItem{"in": []byte("{\"a\": 666}")},
	//	types.ActionDataItem{"in": []byte("{\"a\": 666}")},
	//})
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	for _, item := range items {
	//		fmt.Print(string(item["out"].([]byte)))
	//	}
	//}

	//jp := jsonpath.MustCompile("$")
	//
	//data := map[string]interface{}{
	//	"id": 12345,
	//	"aaa": map[string]interface{}{
	//		"ppp": 123456,
	//	},
	//}
	//
	//lookup, err := jp.Lookup(data)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(lookup)
	//}

}
