package adapters

import (
	"recharge-common/constant"
	"recharge-common/dict/types"
)

// CarrierTypeDict 运营商类型字典
type CarrierTypeDict struct{}

func (d *CarrierTypeDict) GetKey() string {
	return "carrier_type"
}

func (d *CarrierTypeDict) GetName() string {
	return "运营商类型"
}

func (d *CarrierTypeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.CUCC.String(),
			Value: constant.CUCC.Int(),
			Code:  constant.CUCC.Code(),
		},
		{
			Label: constant.CMCC.String(),
			Value: constant.CMCC.Int(),
			Code:  constant.CMCC.Code(),
		},
		{
			Label: constant.CTCC.String(),
			Value: constant.CTCC.Int(),
			Code:  constant.CTCC.Code(),
		},
	}
}

// IndiaCarrierTypeDict 印度运营商类型字典
type IndiaCarrierTypeDict struct{}

func (d *IndiaCarrierTypeDict) GetKey() string {
	return "india_carrier_type"
}

func (d *IndiaCarrierTypeDict) GetName() string {
	return "印度运营商类型"
}

func (d *IndiaCarrierTypeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.Airtel.String(),
			Value: constant.Airtel.Int(),
			Code:  constant.Airtel.Code(),
		},
		{
			Label: constant.Jio.String(),
			Value: constant.Jio.Int(),
			Code:  constant.Jio.Code(),
		},
	}
}
