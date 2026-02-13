package constant

// BalanceOperation 定义操作类型枚举
type BalanceOperation int

const (
	AddBalance      BalanceOperation = iota + 1 //加款
	SubtractBalance                             //扣款
)

func (b BalanceOperation) Value() int {
	switch b {
	case AddBalance:
		return 1
	case SubtractBalance:
		return 2
	default:
		return 0
	}
}

func (b BalanceOperation) Text() string {
	switch b {
	case AddBalance:
		return "加款"
	case SubtractBalance:
		return "扣款"
	default:
		return "unknown"
	}
}

func (b BalanceOperation) Code() string {
	switch b {
	case AddBalance:
		return "add"
	case SubtractBalance:
		return "subtract"
	default:
		return "unknown"
	}
}

func (b BalanceOperation) IsValid() bool {
	return b >= AddBalance && b <= SubtractBalance
}
