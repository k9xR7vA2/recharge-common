package constant

import (
	"fmt"
	"strconv"
)

// AreaScope  区域类型
type AreaScope int //区域类型1全国National，2分省Province

const (
	National AreaScope = iota + 1 // 全国
	Province                      // 分省
)

// String 实现 Stringer 接口，返回区域范围的字符串描述
func (s AreaScope) String() string {
	switch s {
	case National:
		return "全国"
	case Province:
		return "分省"
	default:
		return "未知区域范围"
	}
}

// Code 返回区域范围的编码
func (s AreaScope) Code() string {
	switch s {
	case National:
		return "national"
	case Province:
		return "province"
	default:
		return "Unknown"
	}
}

// IsValid 检查是否为有效的区域范围
func (s AreaScope) IsValid() bool {
	return s == National || s == Province
}

// ParseAreaScope 将字符串解析为 AreaScope
func ParseAreaScope(s string) (AreaScope, error) {
	// 尝试解析数字
	if num, err := strconv.Atoi(s); err == nil {
		switch AreaScope(num) {
		case National, Province:
			return AreaScope(num), nil
		default:
			return 0, fmt.Errorf("invalid region scope number: %d", num)
		}
	}

	// 尝试解析字符串编码
	switch s {
	case "national", "NATIONAL", "全国":
		return National, nil
	case "province", "PROVINCE", "分省":
		return Province, nil
	default:
		return 0, fmt.Errorf("invalid region scope code: %s", s)
	}
}
