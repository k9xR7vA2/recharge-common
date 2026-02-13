package constant

// GlobalNotifyStatus 回调通知状态
type GlobalNotifyStatus int

const (
	NotifyStatusPreparing  GlobalNotifyStatus = 1 + iota //准备通知
	NotifyStatusProcessing                               //通知中
	NotifyStatusSuccess                                  //通知成功
	NotifyStatusError                                    //通知异常
	NotifyStatusTimeout                                  //通知超时
)

// String 实现 Stringer 接口，返回通知状态的字符串描述
func (s GlobalNotifyStatus) String() string {
	switch s {
	case NotifyStatusPreparing:
		return "准备通知"
	case NotifyStatusProcessing:
		return "通知中"
	case NotifyStatusSuccess:
		return "通知成功"
	case NotifyStatusError:
		return "通知异常"
	case NotifyStatusTimeout:
		return "通知超时"
	default:
		return "未知状态"
	}
}

// Code 返回通知状态的编码
func (s GlobalNotifyStatus) Code() string {
	switch s {
	case NotifyStatusPreparing:
		return "preparing"
	case NotifyStatusProcessing:
		return "processing"
	case NotifyStatusSuccess:
		return "success"
	case NotifyStatusError:
		return "error"
	case NotifyStatusTimeout:
		return "timeout"
	default:
		return "unknown"
	}
}

// IsValid 检查通知状态值是否有效
func (s GlobalNotifyStatus) IsValid() bool {
	switch s {
	case NotifyStatusPreparing, NotifyStatusProcessing, NotifyStatusSuccess,
		NotifyStatusError, NotifyStatusTimeout:
		return true
	}
	return false
}
