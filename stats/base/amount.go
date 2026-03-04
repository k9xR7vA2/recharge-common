package base

import (
	"fmt"
	"strconv"
)

// ---- 金额工具 ----

func AmountToFen(amount string) (int64, error) {
	v, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("amountToFen: invalid amount %q: %w", amount, err)
	}
	return v * 100, nil
}

func FenToYuan(fen int64) float64 {
	return float64(fen) / 100.0
}

func ParseInt(s string) int64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}
