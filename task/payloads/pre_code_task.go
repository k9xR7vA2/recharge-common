package payloads

// PreCodeTask 产码任务负载
type PreCodeTask struct {
	TenantID          uint   `json:"tenant_id"`
	ChannelCode       string `json:"channel_code"` // 租户通道编码
	Amount            uint   `json:"amount"`
	Count             int    `json:"count"`               // 需要产几张
	TenantChannelCode string `json:"tenant_channel_code"` // 租户通道编码（冗余方便日志）
}
