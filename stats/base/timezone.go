package base

import (
	"fmt"
	"time"
)

// HourRange 一段时间范围内的 UTC 小时列表
type HourRange struct {
	Hours    []string // Redis key 用，如 ["2026030416", "2026030417", ...]
	HoursInt []int64  // MongoDB 查询用，如 [2026030416, 2026030417, ...]
	DateUTC  string   // 该范围对应的 UTC 日期（取开始小时所在 UTC 日期）
	MinHour  int64    // MongoDB $gte 用
	MaxHour  int64    // MongoDB $lte 用
}

// TodayHourRange 根据时区获取今日对应的 UTC 小时列表
// timezone 如 "Asia/Shanghai"、"Asia/Kolkata"
func TodayHourRange(timezone string) (HourRange, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return HourRange{}, fmt.Errorf("stats/timezone: invalid timezone %q: %w", timezone, err)
	}
	now := time.Now().In(loc)
	return buildHourRange(now.Year(), int(now.Month()), now.Day(), loc), nil
}

// DateHourRange 根据时区和指定日期字符串获取对应的 UTC 小时列表（历史查询）
// dateStr 格式 "2026-03-04"（用户本地日期）
func DateHourRange(dateStr, timezone string) (HourRange, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return HourRange{}, fmt.Errorf("stats/timezone: invalid timezone %q: %w", timezone, err)
	}
	t, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		return HourRange{}, fmt.Errorf("stats/timezone: invalid date %q: %w", dateStr, err)
	}
	return buildHourRange(t.Year(), int(t.Month()), t.Day(), loc), nil
}

// buildHourRange 构建一天内从本地 00:00 到 23:59 对应的 UTC 小时切片
// 跨天是正常的（如 UTC+8 的今日 = UTC 昨日16点到今日15点）
func buildHourRange(year, month, day int, loc *time.Location) HourRange {
	// 本地当天 00:00:00 → UTC
	start := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc).UTC()
	// 本地当天 23:59:59 → UTC
	end := time.Date(year, time.Month(month), day, 23, 59, 59, 0, loc).UTC()

	var hours []string
	var hoursInt []int64
	cur := start
	for !cur.After(end) {
		h := cur.Format("2006010215")
		var hi int64
		fmt.Sscanf(h, "%d", &hi)
		hours = append(hours, h)
		hoursInt = append(hoursInt, hi)
		cur = cur.Add(time.Hour)
	}

	var minH, maxH int64
	if len(hoursInt) > 0 {
		minH = hoursInt[0]
		maxH = hoursInt[len(hoursInt)-1]
	}

	return HourRange{
		Hours:    hours,
		HoursInt: hoursInt,
		DateUTC:  start.Format("2006-01-02"),
		MinHour:  minH,
		MaxHour:  maxH,
	}
}

// IsTodayInTimezone 判断某个 Unix 时间戳是否在指定时区的"今天"
func IsTodayInTimezone(unixSec int64, timezone string) (bool, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return false, err
	}
	t := time.Unix(unixSec, 0).In(loc)
	now := time.Now().In(loc)
	return t.Year() == now.Year() && t.YearDay() == now.YearDay(), nil
}
