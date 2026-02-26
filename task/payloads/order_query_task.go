// payloads/order_query_task.go
package payloads

import (
	"github.com/small-cat1/recharge-common/constant"
	"time"
)

type OrderQueryTask struct {
	TenantID     uint                  `json:"tenant_id"`
	BusinessType constant.BusinessType `json:"business_type"`

	// 商户订单
	MerchantSystemOrderSn string `json:"merchant_system_order_sn"`
	MerchantOrderSn       string `json:"merchant_order_sn"`
	MerchantNotifyURL     string `json:"merchant_notify_url"`
	MerchantAppID         string `json:"merchant_app_id"`

	// 供货商订单
	SupplierID            uint      `json:"supplier_id"`
	SupplierName          string    `json:"supplier_name"`
	SupplierOrderSn       string    `json:"supplier_order_sn"`
	SupplierSystemOrderSn string    `json:"supplier_system_order_sn"`
	SupplierNotifyURL     string    `json:"supplier_notify_url"`
	SupplierAppID         string    `json:"supplier_app_id"`
	SupplierOrderAmount   string    `json:"supplier_order_amount"`
	SupplierOrderCreateAt time.Time `json:"supplier_order_create_at"`

	// 三方查单信息
	TradeNo         string `json:"trade_no"`   // 三方交易号
	ChannelID       uint   `json:"channel_id"` // 通道ID，worker 用来查最新的接口配置
	RechargeAccount string `json:"recharge_account"`
	Amount          string `json:"amount"`

	// 查单策略（从 Interface 快照）
	QueryUrl             string `json:"query_url"`
	QueryInterval        int    `json:"query_interval"`         // 查单间隔（秒）
	QueryMaxTimes        int    `json:"query_max_times"`        // 最大查单次数
	PaySeconds           int    `json:"pay_seconds"`            // 支付有效期（秒）
	OfficialSerialNumber string `json:"official_serial_number"` // 官方流水号
	UpstreamOrderSn      string `json:"upstream_order_sn"`      //三方网关订单号

	// 运行时
	QueryCount int   `json:"query_count"` // 已查单次数
	CreatedAt  int64 `json:"created_at"`  // 任务创建时间
}
