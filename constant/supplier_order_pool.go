package constant

import "strconv"

// SupOrderPoolStat  表示订单在池中的状态
type SupOrderPoolStat int

const (
	SupOrderPoolWaiting          SupOrderPoolStat = iota + 1 // 等待处理
	SupOrderPoolProcessing                                   // 正在处理
	SupOrderPoolCompleted                                    // 处理完成
	SupOrderPoolAwaitingCallback                             // 等待回调
	SupOrderPoolFinalized                                    // 回调完成，订单流程结束
	SupOrderPoolExpired                                      // 订单已过期
	SupOrderPoolCanceled                                     // 订单已取消
)

// String 实现 fmt.Stringer 接口，返回状态的字符串描述
func (s SupOrderPoolStat) String() string {
	switch s {
	case SupOrderPoolWaiting:
		return "等待处理"
	case SupOrderPoolProcessing:
		return "正在处理"
	case SupOrderPoolCompleted:
		return "处理完成"
	case SupOrderPoolAwaitingCallback:
		return "等待回调"
	case SupOrderPoolFinalized:
		return "已完成"
	case SupOrderPoolExpired:
		return "已过期"
	case SupOrderPoolCanceled:
		return "已取消"
	default:
		return "未知状态"
	}
}

// Code 返回供货商订单池状态的编码
func (s SupOrderPoolStat) Code() string {
	switch s {
	case SupOrderPoolWaiting:
		return "waiting"
	case SupOrderPoolProcessing:
		return "processing"
	case SupOrderPoolCompleted:
		return "completed"
	case SupOrderPoolAwaitingCallback:
		return "awaiting_callback"
	case SupOrderPoolFinalized:
		return "finalized"
	case SupOrderPoolExpired:
		return "expired"
	case SupOrderPoolCanceled:
		return "canceled"
	default:
		return "unknown"
	}
}

func (s SupOrderPoolStat) Val() int {
	switch s {
	case SupOrderPoolWaiting:
		return 1
	case SupOrderPoolProcessing:
		return 2
	case SupOrderPoolCompleted:
		return 3
	case SupOrderPoolAwaitingCallback:
		return 4
	case SupOrderPoolFinalized:
		return 5
	case SupOrderPoolExpired:
		return 6
	case SupOrderPoolCanceled:
		return 7
	default:
		return 0
	}
}

// CanCancel 检查当前状态是否可以撤销
func (s SupOrderPoolStat) CanCancel() bool {
	switch s {
	case SupOrderPoolWaiting:
		return true
	default:
		return false
	}
}

func (s SupOrderPoolStat) MarshalBinary() ([]byte, error) {
	// 将状态转换为字节
	return []byte(strconv.Itoa(int(s))), nil
}

// UnmarshalBinary 实现 encoding.BinaryUnmarshaler 接口
func (s *SupOrderPoolStat) UnmarshalBinary(data []byte) error {
	// 从字节恢复状态
	i, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	*s = SupOrderPoolStat(i)
	return nil
}
