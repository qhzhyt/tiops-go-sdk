// @Title  convert.go
// @Description  转换器
// @Create  heyitong  2022/6/23 15:56
// @Update  heyitong  2022/6/23 15:56

package utils

import "strconv"

func DecimalToInt64(str string) int64 {
	result , _ := strconv.ParseInt(str, 10, 64)
	return result
}
