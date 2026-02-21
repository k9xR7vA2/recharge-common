package constants

import "fmt"

const (
	SupplierOrderCallbackTask    = "supplier:order:callback"  // 供货商订单回调
	MerchantOrderCallbackTask    = "merchant:order:callback"  // 商户订单回调
	TenantSystemPermissionUpdate = "tenant:permission:update" //租户系统权限更新队列
	PaymentOrderExpiredTask      = "payment:order:expired"    //系统充值订单过期队列

	// SupplierOrderExpireTask 过期处理任务
	supplierOrderExpirePattern = "supplier:order:%s:expire"
	expirationQueuePattern     = "expiration:%s"
)

func SupplierOrderExpireTask(bizType string) string {
	return fmt.Sprintf(supplierOrderExpirePattern, bizType)
}

func ExpirationQueue(bizType string) string {
	return fmt.Sprintf(expirationQueuePattern, bizType)
}
