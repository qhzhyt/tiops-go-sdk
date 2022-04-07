package utils

import "time"

func CurrentTimeClear() string {
	return time.Now().Format("20060102150405")
}
