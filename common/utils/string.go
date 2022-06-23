// @Title  string.go
// @Description  字符串处理函数
// @Create  heyitong  2022/6/23 16:02
// @Update  heyitong  2022/6/23 16:02

package utils

import (
	"path"
	"strings"
)

func ContainsString(list []string, target string) bool {
	for _, str := range list{
		if str == target {
			return true
		}
	}
	return false
}

func PathClean(str string) string {
	for strings.Contains(str, "../") {
		str = strings.ReplaceAll(str, "../", "")
	}
	if str == ".." {
		str = ""
	}
	str = path.Clean(str)
	return str
}