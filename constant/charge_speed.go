package constant

import (
	"fmt"
	"strconv"
)

type ChargeSpeed int //充值类型1快充，2慢充，3先快后慢

const (
	Fast ChargeSpeed = iota + 1 //快充
	Slow                        //慢充
)

// String 实现 Stringer 接口，返回类型名称的字符串描述
func (c ChargeSpeed) String() string {
	switch c {
	case Fast:
		return "快充"
	case Slow:
		return "慢充"
	default:
		return "未知充值类型"
	}
}

// Code  实现 Stringer 接口，返回类型编码的字符串描述
func (c ChargeSpeed) Code() string {
	switch c {
	case Fast:
		return "fast"
	case Slow:
		return "slow"
	default:
		return "Unknown"
	}
}

// IsValid
func (c ChargeSpeed) IsValid() bool {
	return c == Fast || c == Slow
}

// ParseChargeType 将字符串解析为 ChargeType
func ParseChargeType(s string) (ChargeSpeed, error) {
	// 尝试解析数字
	if num, err := strconv.Atoi(s); err == nil {
		switch ChargeSpeed(num) {
		case Fast, Slow:
			return ChargeSpeed(num), nil
		default:
			return 0, fmt.Errorf("invalid charge type number: %d", num)
		}
	}

	// 尝试解析字符串编码
	switch s {
	case "fast", "FAST":
		return Fast, nil
	case "slow", "SLOW":
		return Slow, nil
	default:
		return 0, fmt.Errorf("invalid charge type code: %s", s)
	}
}
