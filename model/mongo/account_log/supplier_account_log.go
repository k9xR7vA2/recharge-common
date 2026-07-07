package account_log

import "go.mongodb.org/mongo-driver/bson/primitive"

// SupplierAccountLog 供货商充值账号操作日志
// 覆盖拆分账号池模式（electric/india_electric/india_dth/game 等）的
// 锁定(lock)、释放(release)、充值成功(charge_success) 三类动作
type SupplierAccountLog struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AccountID     uint               `bson:"account_id" json:"account_id"`
	TenantID      uint               `bson:"tenant_id" json:"tenant_id"`
	BusinessType  string             `bson:"business_type" json:"business_type"`
	SystemOrderSn string             `bson:"system_order_sn" json:"system_order_sn"` // 触发这次变动的订单号
	Operation     string             `bson:"operation" json:"operation"`             // lock / release / charge_success
	Amount        string             `bson:"amount" json:"amount"`                   // 本次变动金额
	LockedBefore  string             `bson:"locked_before" json:"locked_before"`
	LockedAfter   string             `bson:"locked_after" json:"locked_after"`
	ChargedBefore string             `bson:"charged_before" json:"charged_before"`
	ChargedAfter  string             `bson:"charged_after" json:"charged_after"`
	Remark        string             `bson:"remark,omitempty" json:"remark,omitempty"`
	CreatedAt     int64              `bson:"created_at" json:"created_at"`
}

func (SupplierAccountLog) CollectionName() string {
	return "supplier_account_logs"
}
