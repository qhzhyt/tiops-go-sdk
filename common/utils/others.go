// @Title  others.go
// @Description  其他工具函数
// @Create  heyitong  2022/6/23 15:59
// @Update  heyitong  2022/6/23 15:59

package utils

import (
	"os"
	"time"
)

func SleepAndExit(duration time.Duration, code int)  {
	time.Sleep(duration)
	os.Exit(code)
}
