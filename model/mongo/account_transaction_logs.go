package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccountTransactionLog struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	TenantID        uint               `bson:"tenant_id" json:"tenant_id" `              // 租户ID
	TransactionID   string             `bson:"transaction_id" json:"transaction_id"`     // 交易唯一ID
	AccountType     int                `bson:"account_type" json:"account_type"`         // 1-商户 2-供货商
	AccountName     string             `bson:"account_name" json:"account_name"`         // (商户/供货商)名称
	AccountID       uint               `bson:"account_id" json:"account_id"`             // 关联(商户/供货商账户ID
	TransactionType int                `bson:"transaction_type" json:"transaction_type"` // 交易类型 1加款，2扣款
	BusinessType    int                `bson:"business_type" json:"business_type"`       // 业务类型 具体的业务
	BalanceType     int                `bson:"balance_type" json:"balance_type"`         // 余额账户类型 1-预付款账户 2-信用账户 影响账户类型
	Amount          string             `bson:"amount" json:"amount"`                     // 交易金额
	BeforeBalance   string             `bson:"before_balance" json:"before_balance"`     // 交易前余额
	AfterBalance    string             `bson:"after_balance" json:"after_balance"`       // 交易后余额
	OrderNo         string             `bson:"order_no" json:"order_no"`                 // 关联订单号
	Operator        string             `bson:"operator" json:"operator" `                // 操作人
	CreatedAt       string             `bson:"created_at" json:"created_at"`
}

func (AccountTransactionLog) CollName() string {
	return "account_transaction_log"
}
