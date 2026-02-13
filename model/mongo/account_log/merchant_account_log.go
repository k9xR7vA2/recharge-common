package account_log

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//产码订单日志 make_code_log
// ——————————————————————————
//商户账户余额日志 merchant_balance_log
//商户订单日志  merchant_order_log
//商户订单回调日志  merchant_order_notify_log
//————————————————————————————————
//供货商余额日志 supplier_balance_log
//供货商订单日志 supplier_order_log
//供货商订单回调日志  merchant_order_notify_log

type MerchantBalanceLog struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt       time.Time          // 创建时间
	MerchantId      int                `json:"merchant_id" gorm:"column:merchant_id;NOT NULL;comment:'商户ID'"`
	MerchantName    string             `json:"merchant_name" gorm:"column:merchant_name;NOT NULL;comment:'商户名称'"`
	Action          int                `json:"action" gorm:"not null;comment:'操作，1加款，2扣款'"`
	TradeType       int                `json:"trade_type" gorm:"column:trade_type;NOT NULL;comment:'交易类型，1预付，2商户结算，3订单结算'"`
	DealOrderSn     string             `json:"deal_order_sn" gorm:"column:deal_order_sn;NOT NULL;comment:'交易订单号'"`
	DefaultBalance  string             `json:"default_balance" gorm:"column:default_balance;default:NULL;comment:'原始金额'"`
	ChangeAmount    string             `json:"change_amount" gorm:"column:change_amount;default:NULL;comment:'变动金额'"`
	AfterBalance    string             `json:"after_balance" gorm:"column:after_balance;default:NULL;comment:'变动后金额'"`
	RechargeOrderSn string             `json:"recharge_order_sn" gorm:"column:recharge_order_sn;NOT NULL;comment:'关联充值订单'"`
	ActionUser      string             `json:"action_user" gorm:"column:action_user;default:NULL;comment:'操作人'"`
}
