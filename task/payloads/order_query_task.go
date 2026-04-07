package payloads

import (
	"github.com/k9xR7vA2/recharge-common/constant"
)

type OrderQueryTask struct {
	TenantID     uint                  `json:"tenant_id"`
	BusinessType constant.BusinessType `json:"business_type"`

	// 商户订单
	MerchantID             uint                 `json:"merchant_id"`
	ChannelCode            string               `json:"channel_code"`
	ChannelType            constant.ChannelType `json:"channel_type"`
	MerchantSystemOrderSn  string               `json:"merchant_system_order_sn"`
	MerchantOrderSn        string               `json:"merchant_order_sn"`
	MerchantNotifyURL      string               `json:"merchant_notify_url"`
	MerchantAppID          string               `json:"merchant_app_id"`
	MerchantOrderCreatedAt int64                `json:"merchant_order_created_at"`

	// 供货商订单
	SupplierID            uint   `json:"supplier_id"`
	SupplierName          string `json:"supplier_name"`
	SupplierOrderSn       string `json:"supplier_order_sn"`
	SupplierSystemOrderSn string `json:"supplier_system_order_sn"`
	SupplierNotifyURL     string `json:"supplier_notify_url"`
	SupplierAppID         string `json:"supplier_app_id"`
	SupplierOrderAmount   string `json:"supplier_order_amount"`
	SupplierOrderCreateAt int64  `json:"supplier_order_create_at"`
	// 供货商账号（DTH 充值账号释放额度用）
	SupplierAccountID uint `json:"supplier_account_id"` // 充值账号ID
	// 查单
	InterfaceID     uint   `json:"interface_id"`      // 接口ID → 实时查配置
	UpstreamOrderSn string `json:"upstream_order_sn"` // 三方网关订单号
	Amount          string `json:"amount"`
	ProductCode     string `json:"product_code"`
	// 重新入池判断所需（创建查单任务时从供货商订单中带入）
	ChargeSpeed            constant.ChargeSpeed `json:"charge_speed"`              // 充值速度：1快充 2慢充
	SupplierOrderExpiredAt int64                `json:"supplier_order_expired_at"` // 供货商订单过期时间(毫秒时间戳)
	RechargeTimes          uint                 `json:"recharge_times"`            // 当前充值次数

	// 运行时
	QueryCount int   `json:"query_count"` // 已查单次数
	CreatedAt  int64 `json:"created_at"`  // 任务创建时间
}
