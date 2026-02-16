package traceConstant

// 供应商相关的属性key
const (
	KeyTenantID = "tenant_id" //租户ID
	KeyUserID   = "user_id"   //(供货商/商户)ID
	KeyOrderSn  = "order_sn"  //(供货商/商户)订单号
)

const (
	// KeyErrGetSupplier 错误节点的key
	KeyInvalidSign          = "invalid_sign"           //签名不正确
	KeyInvalidAmount        = "invalid_amount"         //金额不正确
	KeyInsufficientAmount   = "insufficient_amount"    //账户金额不足
	KeyErrorBindJson        = "err_bind_json"          //json绑定错误
	KeyMoGoTransactionError = "mogo_transaction_error" //Mogo开启事务失败
)
