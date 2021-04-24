package utils

import "strconv"

func DecimalToInt64(str string) int64 {
	result , _ := strconv.ParseInt(str, 10, 64)
	return result
}
