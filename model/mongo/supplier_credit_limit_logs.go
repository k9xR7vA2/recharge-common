package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

// SupplierCreditLimitLogs SupplierCreditLimitChange  供货商信用额度变更记录表 (MongoDB)
type SupplierCreditLimitLogs struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`            // 变更ID
	TenantID     uint               `bson:"tenant_id" json:"tenant_id"`         // 租户ID(新增字段)
	SupplierID   uint               `bson:"supplier_id" json:"supplier_id"`     // 供应商ID
	SupplierName string             `bson:"supplier_name" json:"supplier_name"` // 供货商名称
	OldLimit     string             `bson:"old_limit" json:"old_limit"`         // 原授信额度
	NewLimit     string             `bson:"new_limit" json:"new_limit"`         // 新授信额度
	UsedAmount   string             `bson:"used_amount" json:"used_amount"`     // 已用额度
	ChangeType   int                `bson:"change_type" json:"change_type"`     // 变更类型
	ChangeReason string             `bson:"change_reason" json:"change_reason"` // 变更原因
	Operator     string             `bson:"operator" json:"operator"`           // 操作人
	CreatedAt    string             `bson:"created_at" json:"created_at"`       // 创建时间
}

func (SupplierCreditLimitLogs) CollName() string {
	return "supplier_credit_limit_logs"
}
