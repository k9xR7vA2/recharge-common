package traceConstant

const (
	// API模块前缀
	SupplierPrefix = "supplier-api" // 供应商API
	MerchantPrefix = "merchant-api" // 商户API

	middlewareLayer  = "middleware-layer"  //中间件
	interfaceLayer   = "interface-layer"   //接口层
	applicationLayer = "application-layer" //应用层
	domainLayer      = "domain-layer"      //领域层
	repositoryLayer  = "repository-layer"  //仓储层

	// SupplierOrderMiddlewareSpan  ===== 中间件层级 =====
	SupplierOrderMiddlewareSpan = SupplierPrefix + "." + middlewareLayer //供应商下单中间件
	MerchantOrderMiddlewareSpan = MerchantPrefix + "." + middlewareLayer //商户下单中间件

	//  ===== 接口层 =====
	SupplierOrderInterfaceSpan = SupplierPrefix + "." + interfaceLayer //供应商下单接口层
	MerchantOrderInterfaceSpan = MerchantPrefix + "." + interfaceLayer //商户下单接口层

	// SupplierOrderApplicationSpan ===== 应用层 =====
	SupplierOrderApplicationSpan = SupplierPrefix + "." + applicationLayer //供应商下单应用层
	MerchantOrderApplicationSpan = MerchantPrefix + "." + applicationLayer //商户下单应用层

	// SupplierOrderDomainSpan ===== 领域层 =====
	SupplierOrderDomainSpan = SupplierPrefix + "." + domainLayer //供应商下单领域层
	MerchantOrderDomainSpan = MerchantPrefix + "." + domainLayer //商户下单领域层

	// SupplierOrderRepositorySpan ===== 仓储层 =====
	SupplierOrderRepositorySpan = SupplierPrefix + "." + repositoryLayer //供应商下单仓储层
	MerchantOrderRepositorySpan = MerchantPrefix + "." + repositoryLayer //商户下单仓储层

)

//
