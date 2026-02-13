package constant

type CookieStatus int

const (
	CookieStatusNormal  CookieStatus = iota + 1 // CookieStatusNormal 正常
	CookieStatusPaused                          // CookieStatusPaused 暂停
	CookieStatusInvalid                         // CookieStatusInvalid 失效
	CookieStatusBlocked                         // CookieStatusBlocked 封禁
)

// String 返回Cookie状态的字符串描述
func (s CookieStatus) String() string {
	switch s {
	case CookieStatusNormal:
		return "正常"
	case CookieStatusPaused:
		return "暂停"
	case CookieStatusInvalid:
		return "失效"
	case CookieStatusBlocked:
		return "封禁"
	default:
		return "未知状态"
	}
}

func (s CookieStatus) IsValid() bool {
	return s == CookieStatusNormal || s == CookieStatusPaused || s == CookieStatusInvalid || s == CookieStatusBlocked
}
