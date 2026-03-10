package orderpool

import "strings"

type PoolKeyItem struct {
	Priority    string
	Amount      string
	Carrier     string
	ChargeSpeed string
	Count       int64
}

func ParsePoolKey(key string, length int64) (PoolKeyItem, bool) {
	parts := strings.Split(key, ":")
	if len(parts) < 8 {
		return PoolKeyItem{}, false
	}
	return PoolKeyItem{
		Priority:    parts[4],
		Amount:      parts[5],
		Carrier:     parts[6],
		ChargeSpeed: parts[7],
		Count:       length,
	}, true
}
