package payloads

import (
	"github.com/small-cat1/recharge-common/constant"
	"time"
)

type MerchantOrderCallbackTask struct {
	TenantID        uint                  `json:"tenant_id"`
	OrderDate       time.Time             `json:"order_date"`
	NotifyStatus    constant.NotifyStatus `json:"status"`
	BusinessType    constant.BusinessType `json:"business_type"`
	SystemOrderSn   string                `json:"system_order_sn"`
	Amount          string                `json:"amount"`
	NotifyUrl       string                `json:"notify_url"`
	MerchantOrderSn string                `json:"merchant_order_sn"`
	MerchantAppID   string                `json:"merchant_app_id"`
}
