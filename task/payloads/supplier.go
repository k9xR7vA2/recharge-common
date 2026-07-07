package payloads

import (
	"fmt"
	"github.com/k9xR7vA2/recharge-common/constant"
	"time"
)

type SupplierOrderExpirationTask struct {
	DelayTime    time.Duration         `json:"delay_time"`
	OrderInfo    string                `json:"order_info"`
	BusinessType constant.BusinessType `json:"business_type"`
}

func (s SupplierOrderExpirationTask) Validate() error {
	if !s.BusinessType.IsValid() {
		return fmt.Errorf("business_type is required")
	}
	if s.OrderInfo == "" {
		return fmt.Errorf("order_info is required")
	}
	if s.DelayTime <= 0 {
		return fmt.Errorf("delay_time must be positive")
	}
	return nil
}

type SupplierOrderCallbackTask struct {
	SupplierOrderSn string                `json:"supplier_order_sn"`
	SystemOrderSn   string                `json:"system_order_sn"`
	Amount          string                `json:"amount"`
	NotifyUrl       string                `json:"notify_url"`
	SupplierAppID   string                `json:"supplier_app_id"`
	SupplierID      uint                  `json:"supplier_id"`
	TenantID        uint                  `json:"tenant_id"`
	OrderDate       int64                 `json:"order_date"` //下单日期，20060102
	NotifyStatus    constant.NotifyStatus `json:"status"`
	BusinessType    constant.BusinessType `json:"business_type"`
	UpstreamOrderNo string                `json:"upstream_order_no"`
	UpstreamTxnID   string                `json:"upstream_txn_id"`
}

type SupplierAccountLogTask struct {
	AccountID     uint   `json:"account_id"`
	TenantID      uint   `json:"tenant_id"`
	BusinessType  string `json:"business_type"`
	SystemOrderSn string `json:"system_order_sn"`
	Operation     string `json:"operation"` // lock / release / charge_success
	Amount        string `json:"amount"`
	LockedBefore  string `json:"locked_before"`
	LockedAfter   string `json:"locked_after"`
	ChargedBefore string `json:"charged_before"`
	ChargedAfter  string `json:"charged_after"`
	Remark        string `json:"remark,omitempty"`
	CreatedAt     int64  `json:"created_at"`
}
