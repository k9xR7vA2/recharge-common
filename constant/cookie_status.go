package constant

type CookieStatus int

const (
	CookieStatusNormal    CookieStatus = 1 // 正常可用
	CookieStatusCooldown  CookieStatus = 2 // 冷却中
	CookieStatusSuspected CookieStatus = 3 // 疑似风控
	CookieStatusSuspended CookieStatus = 4 // 封控沉淀
	CookieStatusBanned    CookieStatus = 5 // 已废弃（永久封禁）
	CookieStatusExpired   CookieStatus = 6 // 已过期
)

func (s CookieStatus) String() string {
	switch s {
	case CookieStatusNormal:
		return "正常"
	case CookieStatusCooldown:
		return "冷却中"
	case CookieStatusSuspected:
		return "疑似风控"
	case CookieStatusSuspended:
		return "封控沉淀"
	case CookieStatusBanned:
		return "已废弃"
	case CookieStatusExpired:
		return "已过期"
	default:
		return "未知"
	}
}

func (s CookieStatus) IsValid() bool {
	switch s {
	case CookieStatusNormal, CookieStatusCooldown, CookieStatusSuspected,
		CookieStatusSuspended, CookieStatusBanned, CookieStatusExpired:
		return true
	}
	return false
}

func (s CookieStatus) Code() string {
	switch s {
	case CookieStatusNormal:
		return "normal"
	case CookieStatusCooldown:
		return "cooldown"
	case CookieStatusSuspected:
		return "suspected"
	case CookieStatusSuspended:
		return "suspended"
	case CookieStatusBanned:
		return "banned"
	case CookieStatusExpired:
		return "expired"
	default:
		return "unknown"
	}
}
