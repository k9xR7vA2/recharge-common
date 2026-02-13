package utils

import (
	"errors"
	"github.com/shopspring/decimal"
	"strconv"
	"time"
)

func FormatAmount(amountStr string) decimal.Decimal {
	amount, _ := decimal.NewFromString(amountStr)
	return amount
}

func FormatTime(timestamp int64) time.Time {
	if timestamp == 0 {
		return time.Time{} // 返回 time.Time 的零值
	}
	if timestamp > 1000000000000 { // 大于1万亿，可能是毫秒级时间戳
		return time.UnixMilli(timestamp)
	}
	return time.Unix(timestamp, 0)
}

func ParseUint(s string) (uint, error) {
	if s == "" {
		return 0, errors.New("empty string")
	}
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}

func Pointer[T any](in T) (out *T) {
	return &in
}
