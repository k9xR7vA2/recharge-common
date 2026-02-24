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

func (o *SupplierIndiaMobileOrder) GetTenantID() uint                      { return o.TenantID }
func (o *SupplierIndiaMobileOrder) GetBusinessType() constant.BusinessType { return o.BusinessType }
func (o *SupplierIndiaMobileOrder) GetSupplierOrderSn() string             { return o.SupplierOrderSn }
func (o *SupplierIndiaMobileOrder) GetSystemOrderSn() string               { return o.SystemOrderSn }
func (o *SupplierIndiaMobileOrder) GetAmount() string                      { return o.Amount }
func (o *SupplierIndiaMobileOrder) GetRechargeAccount() string             { return o.RechargeAccount }
func (o *SupplierIndiaMobileOrder) GetAppID() string                       { return o.AppID }
func (o *SupplierIndiaMobileOrder) GetNotifyURL() string                   { return o.NotifyURL }
func (o *SupplierIndiaMobileOrder) GetPlanID() string                      { return o.PlanID }
