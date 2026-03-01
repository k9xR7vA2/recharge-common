package payloads

import (
	"github.com/small-cat1/recharge-common/constant"
)

type MerchantOrderCallbackTask struct {
	PayTime         int64                 `json:"pay_time"`
	NotifyStatus    constant.NotifyStatus `json:"status"`
	SystemOrderSn   string                `json:"system_order_sn"`
	Amount          string                `json:"amount"`
	NotifyUrl       string                `json:"notify_url"`
	MerchantOrderSn string                `json:"merchant_order_sn"`
	MerchantAppID   string                `json:"merchant_app_id"`
}
