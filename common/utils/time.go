// @Title  time.go
// @Description  时间相关函数
// @Create  heyitong  2022/6/23 16:02
// @Update  heyitong  2022/6/23 16:02

package utils

import "time"

func CurrentTimeClear() string {
	return time.Now().Format("20060102150405")
}
