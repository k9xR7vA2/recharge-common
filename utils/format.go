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

func FromMillis(ts int64) time.Time {
	return time.UnixMilli(ts)
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
