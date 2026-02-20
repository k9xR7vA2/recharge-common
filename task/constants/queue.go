package constants

import "github.com/small-cat1/recharge-common/constant"

const (
	CallbackHighQueue = "callback-high-queue"
	QueryHighQueue    = "query-high-queue"
	GenerationQueue   = "generation-queue"
	TracingQueue      = "tracing-queue"
	ExpirationQueue   = "expiration-queue"

	GlobalQueueMaxRetry        = 3
	ProcessingSupplierOrderSet = "processing_supplier_orders"

	SupplierOrderCallbackTask    = "supplier:order:callback"  // 供货商订单回调
	TenantSystemPermissionUpdate = "tenant:permission:update" //租户系统权限更新队列
	PaymentOrderExpiredTask      = "payment:order:expired"    //系统充值订单过期队列
)

// GetBaseQueueConfig 固定队列权重，业务项目按需追加动态队列
func GetBaseQueueConfig() map[string]int {
	return map[string]int{
		CallbackHighQueue: 9,
		QueryHighQueue:    7,
		GenerationQueue:   3,
		TracingQueue:      1,
	}
}

// GetQueueConfigWithExpiration 含动态过期队列（tenant_api、tenant_notify 用）
func GetQueueConfigWithExpiration() map[string]int {
	queues := GetBaseQueueConfig()
	for _, biz := range constant.GetAllBusinessTypes() {
		queues["expiration:"+biz.Value.String()] = 5
	}
	return queues
}
