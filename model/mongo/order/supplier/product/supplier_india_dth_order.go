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
}

func (o *SupplierIndiaDTHOrder) GetAppID() string { return o.AppID }
func (o *SupplierIndiaDTHOrder) GetOperator() constant.IndiaDthOperatorType {
	return o.Operator
}
