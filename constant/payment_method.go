package constant

type PaymentMethod string

const (
	PaymentMethodH5         PaymentMethod = "h5"
	PaymentMethodUpiIntent  PaymentMethod = "upi_intent"
	PaymentMethodUpiCollect PaymentMethod = "upi_collect"
	PaymentMethodQrCode     PaymentMethod = "qr_code"
	PaymentMethodJsApi      PaymentMethod = "js_api"
)

func (p PaymentMethod) Label() string {
	switch p {
	case PaymentMethodH5:
		return "H5"
	case PaymentMethodUpiIntent:
		return "UPI Intent"
	case PaymentMethodUpiCollect:
		return "UPI Collect"
	case PaymentMethodQrCode:
		return "扫码支付"
	case PaymentMethodJsApi:
		return "JsApi"
	default:
		return string(p)
	}
}

// IsValid 方法用于验证支付类型是否有效
func (p PaymentMethod) IsValid() bool {
	switch p {
	case PaymentMethodH5, PaymentMethodUpiIntent, PaymentMethodUpiCollect, PaymentMethodQrCode, PaymentMethodJsApi:
		return true
	default:
		return false
	}
}

func (p PaymentMethod) Code() string {
	return string(p)
}

// NeedAccount 该支付方法是否需要收集账号
func (p PaymentMethod) NeedAccount() bool {
	switch p {
	case PaymentMethodUpiCollect:
		return true
	default:
		return false
	}
}

// ExtParamKeys 该支付方法需要的扩展参数字段名（直接用三方字段名）
func (p PaymentMethod) ExtParamKeys() map[string]string {
	switch p {
	case PaymentMethodUpiCollect:
		return map[string]string{
			"upiAccount": "UPI ID",
		}
	default:
		return nil
	}
}
