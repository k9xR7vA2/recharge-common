package payloads

import "time"

type PaymentOrderExpiredTask struct {
	OrderNo   string        `json:"order_no"`
	TenantID  uint          `json:"tenant_id"`
	DelayTime time.Duration `json:"delay_time"`
}
