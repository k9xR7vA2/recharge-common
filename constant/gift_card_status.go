package constant

// GiftCardStatus 卡密状态
type GiftCardStatus int

const (
	GiftCardStatusUnused   GiftCardStatus = iota + 1 // 未使用
	GiftCardStatusAssigned                           // 已分配
	GiftCardStatusVerified                           // 已核销
	GiftCardStatusInvalid                            // 已失效
)

func (g GiftCardStatus) Value() int {
	switch g {
	case GiftCardStatusUnused:
		return 1
	case GiftCardStatusAssigned:
		return 2
	case GiftCardStatusVerified:
		return 3
	case GiftCardStatusInvalid:
		return 4
	default:
		return 0
	}
}

func (g GiftCardStatus) Label() string {
	switch g {
	case GiftCardStatusUnused:
		return "未使用"
	case GiftCardStatusAssigned:
		return "已分配"
	case GiftCardStatusVerified:
		return "已核销"
	case GiftCardStatusInvalid:
		return "已失效"
	default:
		return "unknown"
	}
}

func (g GiftCardStatus) Code() string {
	switch g {
	case GiftCardStatusUnused:
		return "unused"
	case GiftCardStatusAssigned:
		return "assigned"
	case GiftCardStatusVerified:
		return "verified"
	case GiftCardStatusInvalid:
		return "invalid"
	default:
		return "unknown"
	}
}

func (g GiftCardStatus) IsValid() bool {
	return g >= GiftCardStatusUnused && g <= GiftCardStatusInvalid
}
