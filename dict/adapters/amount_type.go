package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

type AmountTypeDict struct{}

func (d *AmountTypeDict) GetKey() string {
	return "amount_type"
}

func (d *AmountTypeDict) GetName() string {
	return "金额类型"
}

func (d *AmountTypeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.AmountTypeFixed.Label(),
			Value: constant.AmountTypeFixed.Val(),
		},
		{
			Label: constant.AmountTypeRange.Label(),
			Value: constant.AmountTypeRange.Val(),
		},
		{
			Label: constant.AmountTypeDynamic.Label(),
			Value: constant.AmountTypeDynamic.Val(),
		},
		{
			Label: constant.AmountTypeBill.Label(),
			Value: constant.AmountTypeBill.Val(),
		},
	}
}
