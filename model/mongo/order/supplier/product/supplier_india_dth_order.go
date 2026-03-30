package product

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/model/mongo/order/supplier"
)

// SupplierIndiaDTHOrder 印度DTH订单
// DTH 与话费差异：无充值速度、无携号转网，运营商为 DTH 运营商
type SupplierIndiaDTHOrder struct {
	supplier.SupplierOrderBase `bson:",inline"`
	Operator                   constant.IndiaDthOperatorType `bson:"operator" json:"operator"`     // DTH 运营商
	AccountID                  uint                          `bson:"account_id" json:"account_id"` // 来源账号ID，用于失败时释放额度
	// 拆单关联字段
	SourceOrderSn string `bson:"source_order_sn" json:"source_order_sn"` // 原始供货商单号，回调用这个
	IsSplit       bool   `bson:"is_split" json:"is_split"`               // 是否拆单子单
}

func (o *SupplierIndiaDTHOrder) GetAppID() string { return o.AppID }
func (o *SupplierIndiaDTHOrder) GetOperator() constant.IndiaDthOperatorType {
	return o.Operator
}
func (o *SupplierIndiaDTHOrder) GetChargeSpeed() constant.ChargeSpeed {
	return constant.Fast
}
