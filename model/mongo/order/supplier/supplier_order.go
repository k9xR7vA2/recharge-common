package supplier

import (
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/model/mongo/order"
)

// 供货商基础订单
type SupplierOrderBase struct {
	order.OrderBase `bson:",inline"`
	SupplierID      uint   `bson:"supplier_id" json:"supplier_id"`             // 供货商ID
	SupplierName    string `bson:"supplier_name" json:"supplier_name"`         // 供货商名称
	SupplierOrderSn string `bson:"supplier_order_sn" json:"supplier_order_sn"` // 供货商订单号

	ProductID   uint   `bson:"product_id" json:"product_id"`     // 产品iD
	ProductName string `bson:"product_name" json:"product_name"` // 产品名称
	ProductCode string `bson:"product_code" json:"product_code"` // 产品编码

	RechargeStatus constant.SupOrderStatus `bson:"recharge_status" json:"recharge_status"` // 订单状态(1等待充值，2充值中，3成功，4失败，5未使用，6账户黑名单,7撤单
	ExpiredAt      int64                   `bson:"expired_at" json:"expired_at"`           // 过期时间
	ValidTime      uint                    `bson:"valid_time" json:"valid_time"`           // 订单有效期时长（秒）
	RechargeLogs   []RechargeLog           `json:"RechargeLogs" bson:"recharge_logs"`
}

// 充值日志结构
type RechargeLog struct {
	Round       int    `json:"round" bson:"round"`             // 第几轮处理（从1开始）
	Operation   string `json:"operation" bson:"operation"`     // 操作类型：create/match/recharge/query/success/fail
	Message     string `json:"message" bson:"message"`         // 描述信息
	Status      int    `json:"status" bson:"status"`           // 1-进行中 2-成功 3-失败
	RequestURL  string `json:"request_url" bson:"request_url"` // 请求地址
	RequestBody string `json:"request_body" bson:"request_body"`
	Response    string `json:"response" bson:"response"`
	CreatedAt   int64  `json:"created_at" bson:"created_at"`
}

func (o SupplierOrderBase) GetSupplierID() uint     { return o.SupplierID }   // 供货商ID
func (o SupplierOrderBase) GetSupplierName() string { return o.SupplierName } // 供货商名称
func (o SupplierOrderBase) GetProductID() uint      { return o.ProductID }    //产品ID
func (o SupplierOrderBase) GetProductName() string  { return o.ProductName }  //产品名称
func (o SupplierOrderBase) GetProductCode() string  { return o.ProductCode }  //产品编码
func (o SupplierOrderBase) GetRechargeStatus() constant.SupOrderStatus {
	return o.RechargeStatus
} //供货商订单充值状态
