package constant

import (
	"fmt"
	"strconv"
)

type CarrierType int //运营商类型1联通，2移动，3电信

const (
	CUCC CarrierType = iota + 1 //中国联通
	CMCC                        // 中国移动
	CTCC                        // 中国电信
)

// String 实现 Stringer 接口，返回运营商名称的字符串描述
func (s CarrierType) String() string {
	switch s {
	case CUCC:
		return "联通"
	case CMCC:
		return "移动"
	case CTCC:
		return "电信"
	default:
		return "未知运营商"
	}
}

func (s CarrierType) Int() int {
	switch s {
	case CUCC:
		return 1
	case CMCC:
		return 2
	case CTCC:
		return 3
	default:
		return 0
	}
}

// Code  实现 Stringer 接口，返回运营商编码的字符串描述
func (s CarrierType) Code() string {
	switch s {
	case CUCC:
		return "cucc"
	case CMCC:
		return "cmcc"
	case CTCC:
		return "ctcc"
	default:
		return "Unknown"
	}
}

func (s CarrierType) IsValid() bool {
	return s >= CUCC && s <= CTCC
}

// ParseCarrier 将字符串解析为 CarrierType
func ParseCarrier(s string) (CarrierType, error) {
	// 尝试解析数字
	if num, err := strconv.Atoi(s); err == nil {
		switch CarrierType(num) {
		case CUCC, CMCC, CTCC:
			return CarrierType(num), nil
		default:
			return 0, fmt.Errorf("invalid carrier number: %d", num)
		}
	}
	// 尝试解析字符串编码
	switch s {
	case "cucc", "CUCC":
		return CUCC, nil
	case "cmcc", "CMCC":
		return CMCC, nil
	case "ctcc", "CTCC":
		return CTCC, nil
	default:
		return 0, fmt.Errorf("invalid carrier code: %s", s)
	}
}
