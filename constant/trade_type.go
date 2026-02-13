package constant

type TradeType int

const (
	BackgroundOperations TradeType = iota + 1 //后台操作
	OrderSettlement                           //订单结算
)

func (b TradeType) Value() int {
	switch b {
	case BackgroundOperations:
		return 1
	case OrderSettlement:
		return 2
	default:
		return 0
	}
}

// Text Code 返回交易类型的编码
func (b TradeType) Text() string {
	switch b {
	case BackgroundOperations:
		return "后台操作"
	case OrderSettlement:
		return "订单结算"
	default:
		return "unknown"
	}
}

// Code 返回交易类型的编码
func (b TradeType) Code() string {
	switch b {
	case BackgroundOperations:
		return "background"
	case OrderSettlement:
		return "settlement"
	default:
		return "unknown"
	}
}

func (b TradeType) IsValid() bool {
	return b >= BackgroundOperations && b <= OrderSettlement
}
