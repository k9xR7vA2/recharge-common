package constant

// NeedCkStatus 下单接口是否需要CK的状态
type NeedCkStatus int //下单接口是否需要CK的状态 1需要，2不需要

const (
	NeedCk NeedCkStatus = iota + 1 // 1需要
	NoneCk                         // 2不需要
)

// String 实现 Stringer 接口，返回区域范围的字符串描述
func (s NeedCkStatus) String() string {
	switch s {
	case NeedCk:
		return "需要"
	case NoneCk:
		return "不需要"
	default:
		return "未知区域范围"
	}
}

// Code 返回区域范围的编码
func (s NeedCkStatus) Code() int {
	switch s {
	case NeedCk:
		return 1
	case NoneCk:
		return 2
	default:
		return 0
	}
}

// IsValid 检查是否为有效的区域范围
func (s NeedCkStatus) IsValid() bool {
	return s == NeedCk || s == NoneCk
}
