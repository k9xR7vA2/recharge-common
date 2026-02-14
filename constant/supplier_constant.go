package constant

// 供货商 状态常量定义
const (
	SupplierContextKey          = "supplierContext"                 //供货商信息上下文内容
	SupplierOrderApiTraceLog    = "supplier_order_api_trace_log"    //供货商下单api接口链路日志
	SupplierOrderCancelTraceLog = "supplier_order_cancel_trace_log" //供货商撤单接口链路日志
	//  供应商状态
	SupplierStatusActive   = 1 // 正常
	SupplierStatusInactive = 2 // 禁用

	// 信用账户状态
	SupplierCreditStatusActive = 1 // 正常
	SupplierCreditStatusFrozen = 2 // 冻结

	AsSupplierCreditLimitChange  = "as_supplier_credit_limit_change" // 供货商账户交易记录表
	AsSupplierPrepaidTransaction = "as_supplier_prepaid_transaction" // 供货商信用额度变更记录表

	// 支付方式
	PaymentTypePrepaid = 1 // 预付款
	PaymentTypeCredit  = 2 // 信用额度

	// 交易状态
	TransactionStatusSuccess = 1 // 成功
	TransactionStatusFailed  = 2 // 失败

	// 额度变更类型
	SupplierCreditChangeTypeInit     = 1 // 初始化
	SupplierCreditChangeTypeIncrease = 2 // 提额
	SupplierCreditChangeTypeDecrease = 3 // 降额
)
