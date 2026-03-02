package account_log

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AgentAccountLog struct {
	ID              primitive.ObjectID        `json:"id" bson:"_id,omitempty"`                              // MongoDB 默认的 ObjectID 或者保留 MySQL 中的 ID
	AgentID         uint                      `json:"agent_id" bson:"agent_id"`                             // 代理 ID
	AgentName       string                    `json:"agent_name" bson:"agent_name"`                         // 代理名称
	Action          constant.BalanceOperation `json:"action" bson:"action"`                                 // 操作类型，1加款，2扣款
	TradeType       int                       `json:"trade_type" bson:"trade_type"`                         // 交易类型，1后台操作，2订单结算，3账户提现
	DealOrderSn     string                    `json:"deal_order_sn" bson:"deal_order_sn"`                   // 交易订单号
	DefaultBalance  string                    `json:"default_balance" bson:"default_balance"`               // 原始金额
	ChangeAmount    string                    `json:"change_amount" bson:"change_amount"`                   // 变动金额
	AfterBalance    string                    `json:"after_balance" bson:"after_balance"`                   // 变动后金额
	RechargeOrderSn string                    `json:"recharge_order_sn" bson:"recharge_order_sn,omitempty"` // 关联充值订单（可选字段，允许为空）
	ActionUser      string                    `json:"action_user" bson:"action_user,omitempty"`             // 操作人（可选字段，允许为空）
	Remark          string                    `json:"remark" bson:"remark"`                                 //备注
	CreatedAt       time.Time                 `json:"created_at" bson:"created_at"`                         // 创建时间
}
