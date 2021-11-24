package tools

import (
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
