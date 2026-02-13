package adapters

import (
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/dict/types"
)

// ChargeSpeedDict 充值速度字典
type ChargeSpeedDict struct{}

func (d *ChargeSpeedDict) GetKey() string {
	return "charge_speed"
}

func (d *ChargeSpeedDict) GetName() string {
	return "充值速度"
}

func (d *ChargeSpeedDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.Fast.String(),
			Value: int(constant.Fast),
			Code:  constant.Fast.Code(),
		},
		{
			Label: constant.Slow.String(),
			Value: int(constant.Slow),
			Code:  constant.Slow.Code(),
		},
	}
}

// DeviceTypeDict 设备类型字典
type DeviceTypeDict struct{}

func (d *DeviceTypeDict) GetKey() string {
	return "device_type"
}

func (d *DeviceTypeDict) GetName() string {
	return "设备类型"
}

func (d *DeviceTypeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.AndroidDevice.String(),
			Value: int(constant.AndroidDevice),
			Code:  constant.AndroidDevice.Code(),
		},
		{
			Label: constant.IOSDevice.String(),
			Value: int(constant.IOSDevice),
			Code:  constant.IOSDevice.Code(),
		},
		{
			Label: constant.BothDevice.String(),
			Value: int(constant.BothDevice),
			Code:  constant.BothDevice.Code(),
		},
	}
}

// PaymentTypeDict 支付类型字典
type PaymentTypeDict struct{}

func (d *PaymentTypeDict) GetKey() string {
	return "payment_type"
}

func (d *PaymentTypeDict) GetName() string {
	return "支付类型"
}

func (d *PaymentTypeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.AliPayPayment.String(),
			Value: int(constant.AliPayPayment),
			Code:  constant.AliPayPayment.Code(),
		},
		{
			Label: constant.WxPayment.String(),
			Value: int(constant.WxPayment),
			Code:  constant.WxPayment.Code(),
		},
		{
			Label: constant.UPayPayment.String(),
			Value: int(constant.UPayPayment),
			Code:  constant.UPayPayment.Code(),
		},
		{
			Label: constant.CUPPayment.String(),
			Value: int(constant.CUPPayment),
			Code:  constant.UPayPayment.Code(),
		},
		{
			Label: constant.UPIPayment.String(),
			Value: int(constant.UPIPayment),
			Code:  constant.UPIPayment.Code(),
		},
	}
}
