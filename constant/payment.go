package constant

type PaymentType int

const (
	AliPayPayment PaymentType = iota + 1 //支付宝
	WxPayment                            //微信
	UPayPayment                          //云闪付
	CUPPayment                           //银联
	UPIPayment                           // Upi支付
)

// String 方法用于将支付类型转换为可读的字符串
func (p PaymentType) String() string {
	switch p {
	case AliPayPayment:
		return "支付宝"
	case WxPayment:
		return "微信"
	case UPayPayment:
		return "云闪付"
	case CUPPayment:
		return "银联"
	case UPIPayment:
		return "UPI"
	default:
		return "未知支付方式"
	}
}

func (p PaymentType) Code() string {
	switch p {
	case AliPayPayment:
		return "alipay"
	case WxPayment:
		return "wechat"
	case UPayPayment:
		return "upay"
	case CUPPayment:
		return "cup"
	default:
		return "unknown"
	}
}

// IsValid 方法用于验证支付类型是否有效
func (p PaymentType) IsValid() bool {
	return p >= AliPayPayment && p <= CUPPayment
}
