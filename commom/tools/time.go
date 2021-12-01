package tools

import (
	"fmt"
	"js_statistics/constant"
	"time"
)

// 过去7天
func GetLastWeekTimeScope(now time.Time) (beginAt, endAt string) {
	return now.AddDate(0, 0, -7).Format(constant.DateFormat), now.Format(constant.DateFormat)
}

// 本月
func GetThisMonthTimeScope(now time.Time) (beginAt, endAt string) {
	beginAt = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local).Format(constant.DateFormat)
	return beginAt, now.Format(constant.DateFormat)
}

// 上月
func GetLastMonthTimeScope(now time.Time) (string, string) {
	end := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	begin := end.AddDate(0, -1, 0)
	return begin.Format(constant.DateFormat), end.Format(constant.DateFormat)
}

// 判断跳转的时间区间是否符合
func IsInRedirectOnOff(on, off string) (bool, error) {
	now := time.Now()
	ont, err := time.Parse("15:04:05", on)
	offT, _err := time.Parse("15:04:05", off)
	if err != nil && _err != nil {
		return false, err
	}
	onTime := time.Date(now.Year(), now.Month(), now.Day(), ont.Hour(), ont.Minute(), ont.Second(), 0, time.Local)
	offTime := time.Date(now.Year(), now.Month(), now.Day(), offT.Hour(), offT.Minute(), offT.Second(), 0, time.Local)
	fmt.Println(now, onTime, offTime)
	fmt.Println(now.Before(offTime) && now.After(onTime))
	fmt.Println(now.Before(onTime) && now.After(offTime))
	return now.Before(offTime) && now.After(onTime), nil
}
