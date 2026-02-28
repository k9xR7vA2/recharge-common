package constant

type MerchantTypes int

const (
	AccountTypeMerchant MerchantTypes = 1 //商户
	AccountTypeSupplier MerchantTypes = 2 //供货商
)

const (
	SupplierTotalOrder   = "supplier_total_order"
	SupplierSuccessOrder = "supplier_success_order"
	MerchantTotalOrder   = "merchant_total_order"
	MerchantSuccessOrder = "merchant_success_order"

	// 业务类型
	BusinessTypeDeposit       = 1 // 预付账户加款
	BusinessTypePrepaidDeduct = 2 // 预付账户扣款
	BusinessTypeOrderDeduct   = 3 // 订单扣款（优先预付，不足用信用）
	// 结算相关
	BusinessTypeSettlementAdmin = 21 // 后台结算

	ErrRecordNotFound = "记录不存在"
)

//merchant_constant

const (
	MerchantContextKey       = "merchantContext"              //商户上下文信息key
	MerchantOrderApiTraceLog = "merchant_order_api_trace_log" //商户下单api接口链路日志
	PaymentTraceLog          = "payment_trace_log"            //付款api接口链路日志
	CompositeChannelPrefix   = "9"                            //混合通道的前缀
)

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

const (

	// google

	// redis
	MerchantOrderPrefix      = "merchant:order:" // 单个订单下单缓存前缀
	SupplierOrderCachePrefix = "supplier:orders:cache"
	MerchantOrderCachePrefix = "merchant:orders:cache"
	SystemConfigHashKey      = "system:configs" // 系统配置Hash键

	OrderTypeMerchant = 1 // 商户订单
	OrderTypeSupplier = 2 // 供货商订单

	EntityTypeAllMerchants    = "all_merchants"
	EntityTypeMerchant        = "merchant"
	EntityTypeMerchantChannel = "merchant_channel"
	EntityTypeAllChannels     = "all_channels"
	EntityTypeChannel         = "channel"

	EntityTypeAllSuppliers    = "all_suppliers"
	EntityTypeSupplier        = "supplier"
	EntityTypeAllProducts     = "all_products"
	EntityTypeProduct         = "product"
	EntityTypeSupplierProduct = "supplier_product"

	//mongo coll
	AgentAccountLogs         = "agent_account_logs"
	TenantAccountLogs        = "tenant_account_logs"
	SupplierOrderTablePrefix = "supplier_order_"         // 供货商订单表前缀
	MerchantOrderTablePrefix = "merchant_order_"         // 商户订单表前缀
	QueryLogCollectionPrefix = "query_log_"              //查单日志订单表前缀
	AccountDailySummaries    = "account_daily_summaries" //租户账户日志归集表
	PlatformRule             = "platform_rules"
	Cookies                  = "cookies"

	MerchantOrderCacheFiled       = "order_json"       //商户缓存订单在hash里的字段
	MerchantOrderCacheStatusFiled = "order_sub_status" //商户缓存订单 状态在hash里的字段
)

func MerchantOrderCacheKey(orderSn string) string {
	return MerchantOrderPrefix + orderSn
}
