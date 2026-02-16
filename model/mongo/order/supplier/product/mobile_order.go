package product

import (
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/model/mongo/order/supplier"
)

// 供货商话费订单
type SupplierMobileOrder struct {
	supplier.SupplierOrderBase `bson:",inline"`
	Carrier                    constant.CarrierType `bson:"carrier" json:"carrier"`               //运营商
	ChargeSpeed                constant.ChargeSpeed `bson:"charge_speed" json:"charge_speed"`     // 充值速度
	IsPortability              uint                 `bson:"is_portability" json:"is_portability"` // 是否携号转网
	Area                       constant.AreaScope   `bson:"area" json:"area"`                     //地区
	Province                   uint                 `bson:"province" json:"province"`             //省份
}

// 运营商
func (o SupplierMobileOrder) GetCarrier() constant.CarrierType {
	return o.Carrier
}

func (o SupplierMobileOrder) GetChargeSpeed() constant.ChargeSpeed {
	return o.ChargeSpeed
}

func (o SupplierMobileOrder) GetIsPortability() uint {
	return o.IsPortability
}

func (o SupplierMobileOrder) GetValidTime() uint {
	return o.ValidTime
}

func (o SupplierMobileOrder) GetGetExpiredAt() int64 {
	return o.ExpiredAt
}

// 地区(全国/分省)
func (o SupplierMobileOrder) GetAreaCode() constant.AreaScope {
	return o.Area
}

// 省份
func (o SupplierMobileOrder) GetProvinceCode() uint {
	return o.Province
}
