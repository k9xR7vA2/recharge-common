package payloads

import (
	"github.com/small-cat1/recharge-common/constant"
)

type MerchantOrderCallbackTask struct {
	TenantID        uint                  `json:"tenant_id"`
	BusinessType    constant.BusinessType `json:"business_type"`
	OrderTime       int64                 `json:"order_time"`
	PayTime         int64                 `json:"pay_time"`
	SystemOrderSn   string                `json:"system_order_sn"`
	Amount          string                `json:"amount"`
	NotifyUrl       string                `json:"notify_url"`
	MerchantOrderSn string                `json:"merchant_order_sn"`
	MerchantAppID   string                `json:"merchant_app_id"`
}
