package product

import (
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/model/mongo/order/supplier"
)

// SupplierIndiaMobileOrder 印度话费订单
type SupplierIndiaMobileOrder struct {
	supplier.SupplierOrderBase `bson:",inline"`
	Carrier                    constant.IndiaCarrierType `bson:"carrier" json:"carrier"`
	ChargeSpeed                constant.ChargeSpeed      `bson:"charge_speed" json:"charge_speed"`
	PlanID                     string                    `bson:"plan_id" json:"plan_id"` // SKU 套餐ID
}
