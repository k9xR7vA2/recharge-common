package constant

// Priority 订单池优先级参数
type Priority string

const (
	HighPriority   Priority = "high"   // 高优先级
	NormalPriority Priority = "normal" //普通优先级
)

// String 实现匹配规则的字符串表示
func (p Priority) String() string {
	switch p {
	case HighPriority:
		return "high"
	case NormalPriority:
		return "normal"
	default:
		return "none"
	}
}

func (p Priority) Name() string {
	switch p {
	case HighPriority:
		return "高优先级"
	case NormalPriority:
		return "普通优先级"
	default:
		return "未知"
	}
}
