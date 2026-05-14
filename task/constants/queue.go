package constants

import (
	"fmt"
	"github.com/k9xR7vA2/recharge-common/constant"
)

// ==================== 队列名称 ====================

const (
	// tenant_notify 专用
	TnCallbackHighQueue      = "tn:callback-high-queue"
	TnQueryHighQueue         = "tn:query-high-queue"
	tnExpirationQueuePattern = "tn:expiration:%s"
	TnPreCodeQueue           = "tn:pre-code-queue"
	// SaasAdmin 专用
	SaasPaymentExpiredQueue = "saas:payment-expired-queue"

	// tenant_server 专用
	TsPermissionUpdateQueue = "ts:permission-update-queue"

	ExpiredBuffer              = 10
	GlobalQueueMaxRetry        = 3
	ProcessingSupplierOrderSet = "processing_supplier_orders"
)

// ==================== 各服务队列配置 ====================

// GetSaasAdminQueueConfig SaasAdmin 只监听自己的队列
func GetSaasAdminQueueConfig() map[string]int {
	return map[string]int{
		SaasPaymentExpiredQueue: 5,
	}
}

// GetTenantServerQueueConfig tenant_server 只监听权限更新队列
func GetTenantServerQueueConfig() map[string]int {
	return map[string]int{
		TsPermissionUpdateQueue: 5,
	}
}

// GetTenantNotifyQueueConfig tenant_notify 监听查单、回调、动态过期队列
func GetTenantNotifyQueueConfig() map[string]int {
	queues := map[string]int{
		TnCallbackHighQueue: 9,
		TnQueryHighQueue:    7,
		TnPreCodeQueue:      3,
	}
	for _, biz := range constant.GetAllBusinessTypes() {
		queues[TnExpirationQueue(biz.Value.String())] = 5
	}
	return queues
}

func TnExpirationQueue(bizType string) string {
	return fmt.Sprintf(tnExpirationQueuePattern, bizType)
}
