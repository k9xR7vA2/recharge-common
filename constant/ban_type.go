package constant

// BanType 封控类型
type BanType int

const (
	BanTypeSoftRisk   BanType = 1 // 软性风控（验证码/邮件验证）
	BanTypeSessionExp BanType = 2 // 会话失效（401/重定向登录）
	BanTypeOrderLimit BanType = 3 // 下单限制（能浏览不能下单）
	BanTypeTempFreeze BanType = 4 // 账号临时冻结
	BanTypePermBan    BanType = 5 // 账号永久封禁
)

func (b BanType) String() string {
	switch b {
	case BanTypeSoftRisk:
		return "软性风控"
	case BanTypeSessionExp:
		return "会话失效"
	case BanTypeOrderLimit:
		return "下单限制"
	case BanTypeTempFreeze:
		return "临时冻结"
	case BanTypePermBan:
		return "永久封禁"
	default:
		return "未知"
	}
}
