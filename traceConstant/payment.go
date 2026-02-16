package traceConstant

// 支付API链路日志节点
const (
	DecryptPaymentTokenErr   = "decrypt_payment_token_err"   //解密支付Token失败
	GetPaymentOrderErr       = "get_payment_order_err"       //获取支付订单错误
	UnmarshalPaymentOrderErr = "unmarshal_payment_order_err" //json解析支付订单错误
)
