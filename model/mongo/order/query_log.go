package order

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderQueryLog 查单日志
type OrderQueryLog struct {
	ID                    primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	TenantID              uint                   `bson:"tenant_id" json:"tenant_id"`
	BusinessType          constant.BusinessType  `bson:"business_type" json:"business_type"`
	SupplierSystemOrderSn string                 `bson:"supplier_system_order_sn" json:"supplier_system_order_sn"` // 供货商系统订单号
	MerchantSystemOrderSn string                 `bson:"merchant_system_order_sn" json:"merchant_system_order_sn"` // 商户系统订单号
	UpstreamOrderSn       string                 `bson:"upstream_order_sn" json:"upstream_order_sn"`               // 三方网关订单号
	QueryUrl              string                 `bson:"query_url" json:"query_url"`                               // 查单地址
	Request               map[string]interface{} `bson:"request" json:"request"`                                   // 请求参数
	Response              map[string]interface{} `bson:"response" json:"response"`                                 // 响应内容
	HttpStatus            int                    `bson:"http_status" json:"http_status"`                           // HTTP状态码
	CreatedAt             int64                  `bson:"created_at" json:"created_at"`                             // 日志时间
}
