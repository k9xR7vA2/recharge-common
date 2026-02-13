package merchant

import (
	"SaasAdmin/internal/common/constant"
	"SaasAdmin/internal/infrastructure/model/mongo/order"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// 商户话费订单
type MerchantOrder struct {
	order.OrderBase

	MerchantID      uint   `bson:"merchant_id" json:"merchant_id"`             // 商户ID
	MerchantName    string `bson:"merchant_name" json:"merchant_name"`         // 商户名称
	MerchantOrderSn string `bson:"merchant_order_sn" json:"merchant_order_sn"` // 商户订单号

	ChannelID   uint   `bson:"channel_id" json:"channel_id"`     // 通道ID
	ChannelCode int    `bson:"channel_code" json:"channel_code"` // 通道编码
	ChannelName string `bson:"channel_name" json:"channel_name"` // 通道名称

	OrderStatus    constant.MerOrderMainStat `bson:"order_status" json:"order_status"`         // 订单主状态(1等待支付，2支付中，3支付成功，4支付失败)
	OrderSubStatus constant.MerOrderSubStat  `bson:"order_sub_status" json:"order_sub_status"` // 订单子状态

	ClientIP string `bson:"client_ip" json:"client_ip"` // 客户ip
	Device   int    `bson:"device" json:"device"`       // 设备,1ios,2Android,3双端
	Payment  int    `bson:"payment" json:"payment"`     // 支付方式 1支付宝2微信

	MatchAt        string `bson:"match_at" json:"match_at"`                 // 配单成功时间
	MakeCodeNumber uint   `bson:"make_code_number" json:"make_code_number"` // 产码次数
	MakeCodeAt     string `bson:"make_code_at" json:"make_code_at"`         // 产码成功时间

	// 配单产码历史记录，按时间倒序
	MatchCodeHistory []*MatchCodeGroup `bson:"match_code_history" json:"match_code_history"`

	// 当前生效的配单产码组
	CurrentGroup *MatchCodeGroup `bson:"current_group" json:"current_group"`

	IsReplenishment int `bson:"is_replenishment" json:"is_replenishment"` // 是否手工补单,1不是，2是
}

// MatchCodeGroup 配单产码组合记录
type MatchCodeGroup struct {
	GroupID   primitive.ObjectID `bson:"group_id" json:"group_id"`     // 组ID
	CreatedAt time.Time          `bson:"created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"` // 更新时间

	// 配单信息
	MatchStatus           constant.MerOrderSubStat `bson:"match_status" json:"match_status"`                         // 配单状态
	SupplierID            uint                     `bson:"supplier_id" json:"supplier_id"`                           // 供应商ID
	SupplierName          string                   `bson:"supplier_name" json:"supplier_name"`                       // 供应商名称
	SystemOrderSn         string                   `bson:"system_order_sn" json:"system_order_sn"`                   // 供应商订单的系统订单号
	RechargeAccount       string                   `bson:"recharge_account" json:"recharge_account"`                 // 充值账号
	SupplierOrderCreateAt time.Time                `bson:"supplier_order_create_at" json:"supplier_order_create_at"` // 供应商订单创建时间
	MatchError            string                   `bson:"match_error" json:"match_error"`                           // 配单错误信息
	MatchedAt             time.Time                `bson:"matched_at" json:"matched_at"`                             // 配单完成时间

	// 产码信息
	CodeStatus  constant.MerOrderSubStat `bson:"code_status" json:"code_status"`   // 产码状态
	CodeValue   string                   `bson:"code_value" json:"code_value"`     // 产码值
	CodeError   string                   `bson:"code_error" json:"code_error"`     // 产码错误信息
	GeneratedAt time.Time                `bson:"generated_at" json:"generated_at"` // 产码完成时间
}

type PaymentOrderContext struct {
	*gin.Context
	PaymentOrderInfo *MerchantOrder
}

func (pc *PaymentOrderContext) GetPaymentOrder() *MerchantOrder {
	return pc.PaymentOrderInfo
}
