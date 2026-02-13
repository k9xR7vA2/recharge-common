package constant

// 全局账户状态
type GlobalAccountStatus int

const (
	AccountStatusActive   GlobalAccountStatus = 1 + iota //正常
	AccountStatusDisabled                                //禁用
)

func (s GlobalAccountStatus) Text() string {
	switch s {
	case AccountStatusActive:
		return "正常"
	case AccountStatusDisabled:
		return "禁用"
	default:
		return "未知状态"
	}
}

func (s GlobalAccountStatus) Code() string {
	switch s {
	case AccountStatusActive:
		return "active"
	case AccountStatusDisabled:
		return "disabled"
	default:
		return "unknown"
	}
}

// IsValid 检查通知状态值是否有效
func (s GlobalAccountStatus) IsValid() bool {
	return s == AccountStatusActive || s == AccountStatusDisabled
}
