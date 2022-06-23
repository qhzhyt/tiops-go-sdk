// @Title  crontab.go
// @Description  crontab工具函数
// @Create  heyitong  2022/6/23 15:56
// @Update  heyitong  2022/6/23 15:56

package utils

import (
	//"github.com/robfig/cron"
	cron "github.com/robfig/cron/v3"
	"time"
)

func ParseCrontabText(crontabText string) (cron.Schedule, error) {
	schedule, err := cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow).Parse(crontabText)

	//cron.NewParser()
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func CrontabNextTime(crontabText string, fromTime time.Time) (time.Time, error) {
	s, err := ParseCrontabText(crontabText)
	if err != nil {
		return time.Time{}, err
	}
	return s.Next(fromTime), nil
}
