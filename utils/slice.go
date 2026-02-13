package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func IsInSlice(ele interface{}, slice []string) bool {
	// 将ele转为字符串
	var eleStr string
	switch val := ele.(type) {
	case string:
		eleStr = val
	case int:
		eleStr = strconv.Itoa(val)
	case int64:
		eleStr = strconv.FormatInt(val, 10)
	case uint:
		eleStr = strconv.FormatUint(uint64(val), 10)
	case uint64:
		eleStr = strconv.FormatUint(val, 10)
	case float64:
		// 如果是整数值的float,转为整数字符串
		if float64(int(val)) == val {
			eleStr = strconv.Itoa(int(val))
		} else {
			eleStr = strconv.FormatFloat(val, 'f', -1, 64)
		}
	default:
		eleStr = fmt.Sprintf("%v", val)
	}

	// 对每个slice中的值尝试数字比较
	if eleNum, err := strconv.ParseInt(eleStr, 10, 64); err == nil {
		for _, v := range slice {
			if vNum, err := strconv.ParseInt(v, 10, 64); err == nil {
				if eleNum == vNum {
					return true
				}
			}
		}
	}

	// 再尝试字符串比较
	for _, v := range slice {
		if eleStr == strings.TrimSpace(v) {
			return true
		}
	}

	return false
}
