package constant

type PaymentOrderStatus int

const (
	PaymentOrderStatusPending   PaymentOrderStatus = iota + 1 // PaymentOrderStatusPending 待支付
	PaymentOrderStatusPaid                                    // PaymentOrderStatusPaid 已支付
	PaymentOrderStatusCancelled                               // PaymentOrderStatusCancelled 已取消
	PaymentOrderStatusExpired                                 // PaymentOrderStatusExpired 已过期
)

// String 返回订单状态的字符串描述
func (s PaymentOrderStatus) String() string {
	switch s {
	case PaymentOrderStatusPending:
		return "待支付"
	case PaymentOrderStatusPaid:
		return "已支付"
	case PaymentOrderStatusCancelled:
		return "已取消"
	case PaymentOrderStatusExpired:
		return "已过期"
	default:
		return "未知状态"
	}
}
