package utils

import (
	"os"
	"time"
)

func SleepAndExit(duration time.Duration, code int)  {
	time.Sleep(duration)
	os.Exit(code)
}
