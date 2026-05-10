package constant

type CodeStatus string

const (
	CodeStatusAvailable CodeStatus = "available"
	CodeStatusUsed      CodeStatus = "used"
	CodeStatusExpired   CodeStatus = "expired"
)

func (s CodeStatus) Val() string {
	return string(s)
}

func (s CodeStatus) Label() string {
	switch s {
	case CodeStatusAvailable:
		return "可用"
	case CodeStatusUsed:
		return "已使用"
	case CodeStatusExpired:
		return "已过期"
	default:
		return "未知状态"
	}
}

func (s CodeStatus) IsValid() bool {
	return s == CodeStatusAvailable || s == CodeStatusUsed || s == CodeStatusExpired
}
