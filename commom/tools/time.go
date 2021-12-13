package tools

import (
	"js_statistics/constant"
	"time"
)

// 过去7天
func GetLastWeekTimeScope(now time.Time) (beginAt, endAt time.Time) {
	return now.AddDate(0, 0, -7), now
}

// 本月
func GetThisMonthTimeScope(now time.Time) (beginAt, endAt time.Time) {
	beginAt = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	return beginAt, now
}

// 近一个月
func GetLatestMonthTimeScope(now time.Time) (beginAt, endAt time.Time) {
	return now.AddDate(0, -1, 0), now
}

// 上月
func GetLastMonthTimeScope(now time.Time) (time.Time, time.Time) {
	end := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	begin := end.AddDate(0, -1, 0)
	return begin, end
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
	return now.Before(offTime) && now.After(onTime), nil
}

// 时间区间补点
func DayIterator(beginAt, endAt time.Time) []string {
	dates := make([]string, 0)
	dates = append(dates, beginAt.Format(constant.DateFormat))
	tIter := beginAt.AddDate(0, 0, 1)
	for {
		if tIter.After(endAt) {
			break
		}
		dates = append(dates, tIter.Format(constant.DateFormat))
		tIter = tIter.AddDate(0, 0, 1)
	}
	return dates
}
