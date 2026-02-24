package constant

type SupOrderStatus int

// 1等待充值 - 在订单池等待匹配
// 2充值中   - 已匹配上，正在充值处理
// 3成功    - 充值成功
// 4失败    - 充值失败
// 5未使用   - 匹配超时，等待回调供货商
// 6黑名单   - 账号被识别为黑名单
// 7撤单    - 订单被撤销
// SupOrderStatusWaiting 供货商订单充值状态 (1等待充值，2充值中，3成功，4失败，5未使用，6账户黑名单 7撤单
const (
	SupOrderStatusWaiting     SupOrderStatus = 1 + iota //等待充值
	SupOrderStatusProcessing                            //充值中
	SupOrderStatusSuccess                               //充值成功
	SupOrderStatusFailed                                //充值失败
	SupOrderStatusUnused                                //订单未使用
	SupOrderStatusBlacklisted                           //账户黑名单
	SupOrderStatusCancelled                             //已撤销
)

// String 实现 Stringer 接口，返回状态的字符串描述
func (s SupOrderStatus) String() string {
	switch s {
	case SupOrderStatusWaiting:
		return "等待充值"
	case SupOrderStatusProcessing:
		return "充值中"
	case SupOrderStatusSuccess:
		return "充值成功"
	case SupOrderStatusFailed:
		return "充值失败"
	case SupOrderStatusUnused:
		return "未使用"
	case SupOrderStatusBlacklisted:
		return "账户黑名单"
	case SupOrderStatusCancelled:
		return "已撤销"
	default:
		return "未知状态"
	}
}

func (s SupOrderStatus) Code() string {
	switch s {
	case SupOrderStatusWaiting:
		return "waiting"
	case SupOrderStatusProcessing:
		return "processing"
	case SupOrderStatusSuccess:
		return "success"
	case SupOrderStatusFailed:
		return "failed"
	case SupOrderStatusUnused:
		return "unused"
	case SupOrderStatusBlacklisted:
		return "blacklisted"
	case SupOrderStatusCancelled:
		return "cancelled"
	default:
		return "unknown"
	}
}

// IsValid 检查状态值是否有效
func (s SupOrderStatus) IsValid() bool {
	switch s {
	case SupOrderStatusWaiting, SupOrderStatusProcessing, SupOrderStatusSuccess,
		SupOrderStatusFailed, SupOrderStatusUnused, SupOrderStatusBlacklisted, SupOrderStatusCancelled:
		return true
	}
	return false
}

// CanCancel 检查当前状态是否可以撤销
func (s SupOrderStatus) CanCancel() bool {
	switch s {
	case SupOrderStatusWaiting:
		return true
	default:
		return false
	}
}

// IsFinalStatus 检查是否为最终状态
func (s SupOrderStatus) IsFinalStatus() bool {
	switch s {
	case SupOrderStatusSuccess, SupOrderStatusFailed,
		SupOrderStatusCancelled, SupOrderStatusBlacklisted:
		return true
	default:
		return false
	}
}

// IsProcessing 检查是否处于处理中状态
func (s SupOrderStatus) IsProcessing() bool {
	return s == SupOrderStatusProcessing
}
