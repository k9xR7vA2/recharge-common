package constant

const (
	SystemName         = "SaasAdmin"
	USDTRate           = "usdt_rate"           //费率
	RateExpireTime     = 5 * 60                //费率过期时间
	GoogleVerifyAction = "view-wallet-address" // 谷歌验证器操作

	PhoneRegexChinaStrict = `^1(3[0-9]|4[01456879]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\d{8}$`
	SystemConfigs         = "system:db:configs" //系统配置
	PaymentHost           = "payment_host"      //系统配置

	SystemQueueEventLog   = "system_queue_event_log" //系统队列异常错误日志
	ExtraRetentionSeconds = 5 * 60                   // 订单过期后的额外保留时间 5分钟 (秒)

	TenantDefaultPasswd = "tenantHub@123"
	TraceIdContent      = "trace_id" //上下文的trace_id
	TraceIDKey          = "X-Trace-ID"
	// google
	GoogleIssuer      = "WuKong"
	GoogleAccountName = "WuKong@gmail.com"

	SupplierMobileOrderTablePrefix = "supplier_order_mobile" // 供货商话费订单表前缀
	MerchantMobileOrderTablePrefix = "merchant_order_mobile" // 商户话费订单表前缀

	SupplierOrderCachePrefix = "supplier:orders:cache"

	// AgentAccountLogs 日志
	AgentAccountLogs  = "agent_account_logs"
	TenantAccountLogs = "tenant_account_logs"

	PlatformRule = "platform_rules"
	Cookies      = "cookies"

	// ProductCodeRegRule 产品编码验证规则
	ProductCodeRegRule = "^[a-zA-Z][a-zA-Z_]*$"
	// AccountRegRule 账户验证规则
	AccountRegRule = "[A-Za-z]+"
	// RebateRegRule 费率验证规则
	RebateRegRule = "^(0?\\.[1-9]\\d?|[1-9]\\d?(\\.\\d{1,2})?|100(\\.00?)?)$$"
)

const (
	SupplierTotalOrder   = "supplier_total_order"
	SupplierSuccessOrder = "supplier_success_order"
	MerchantTotalOrder   = "merchant_total_order"
	MerchantSuccessOrder = "merchant_success_order"

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
)
