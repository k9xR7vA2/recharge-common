package account_log

import (
	"SaasAdmin/internal/common/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// TenantBalanceLog 转成mongo
type TenantBalanceLog struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TenantId   uint               `json:"tenant_id" bson:"tenant_id"`
	TenantName string             `json:"tenant_name" bson:"tenant_name"`

	// 账户信息
	TradeAccount constant.AccountType      `json:"trade_account" bson:"trade_account"`
	Action       constant.BalanceOperation `json:"action" bson:"action"`
	TradeType    constant.TradeType        `json:"trade_type" bson:"trade_type"`
	// 金额相关
	BeforeBalance string `json:"before_balance" bson:"before_balance"`
	ChangeAmount  string `json:"change_amount" bson:"change_amount"`
	AfterBalance  string `json:"after_balance" bson:"after_balance"`

	// 订单相关
	DealOrderSn     string `json:"deal_order_sn" bson:"deal_order_sn"`
	RechargeOrderSn string `json:"recharge_order_sn" bson:"recharge_order_sn"`

	// 操作信息
	ActionUser string    `json:"action_user" bson:"action_user"`
	Remark     string    `json:"remark" bson:"remark"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}

func (TenantBalanceLog) TableName() string {
	return "as_tenant_balance_logs"
}
