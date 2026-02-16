package traceConstant

// span错误节点
// 供货商span的节点标识

const (
	// 基础信息查询错误
	GetSupplierFailed        = "get_supplier_failed"         // 获取供货商信息失败
	GetTenantFailed          = "get_tenant_failed"           // 获取供货商租户失败
	GetSupplierInfoFailed    = "get_supplier_info_failed"    // 获取供货商上下文失败
	GetSupplierProductFailed = "get_supplier_product_failed" //获取供货商产品信息失败
	ProductVerifyFailed      = "product_verify_failed"       // 产品信息验证失败

	// domain领域层
	GetTenantProductFailed         = "get_tenant_product_failed"          //获取租户产品信息失败
	VerifyOrderUniqueErr           = "verify_order_unique_error"          //检查订单是否存在错误
	GenerateSystemOrderSnFailed    = "generate_system_order_sn_failed"    //生成系统订单号错误
	FindCashAndPaymentConfigFailed = "find_cash_payment_config_failed"    //获取系统收银和支付配置参数失败
	GenerateCashUrlWithSignFailed  = "generate_cash_url_with_sign_failed" //生成包含加密参数的收银台URL错误
	CreateOrderCountFailed         = "create_order_count_failed"          //创建订单统计失败

	// 订单相关错误
	QuerySupplierOrderFailed = "query_supplier_order_failed" // 查询供货商订单失败
	GenerateOrderSnFailed    = "generate_order_sn_failed"    // 生成订单编号失败
	//repo 仓储
	AddOrderToPoolFailed            = "add_order_to_pool_failed"          // 添加订单至订单池失败
	OrderDuplicateError             = "OrderDuplicateError"               //唯一索引冲突错误
	CreateSupplierOrderFailed       = "create_supplier_order_failed"      // 创建供货商订单失败
	FailedToExecSupplierOrderScript = "failed_exec_supplier_order_script" //执行供货商订单入池lua脚本失败

	SupplierOrderExpireQueueWriteError     = "supplier_order_expire_queue_write_error"       // 写入供货商订单过期队列失败
	ClearFailSupplierOrderExpireQueueError = "clear_faile_supplier_order_expire_queue_error" // 清理失败的供货商订单过期队列错误
	//execute cancel script error
)

// 商户链路日志标识
const (
	GetMerchantFailed    = "get_merchant_failed"     // 获取商户信息失败
	GetChannelInfoFailed = "get_channel_info_failed" // 获取通道信息失败

	QueryMerchantOrderFailed          = "query_merchant_order_failed"              // 查询商户订单失败
	GetMerchantInfoFailed             = "get_merchant_info_failed"                 // 获取商户上下文失败
	MerchantOrderMatchOrderFailed     = "merchant_order_match_order_failed"        //  商户下单配单失败
	CreateMerchantOrderFailed         = "create_merchant_order_failed"             // 创建商户订单失败
	MatchSuccessExecUpdateOrderFailed = "match_success_exec_update_order_failed"   // 配单成功 执行更细订单失败
	ExecUpdateOrderFailed             = "exec_update_order_failed"                 // 执行更新订单失败
	MerOrderExpireQueueWriteError     = "merchant_order_expire_queue_writer_error" // 推送商户订单到延迟队列失败
)
