package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

// RateTypeDict 费率类型字典
type RateTypeDict struct{}

func (d *RateTypeDict) GetKey() string {
	return "rate_type"
}

func (d *RateTypeDict) GetName() string {
	return "费率类型"
}

func (d *RateTypeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.RateTypePercentage.String(),
			Value: constant.RateTypePercentage.Val(),
			Code:  constant.RateTypePercentage.Code(),
		},
		{
			Label: constant.RateTypePerMille.String(),
			Value: constant.RateTypePerMille.Val(),
			Code:  constant.RateTypePerMille.Code(),
		},
		{
			Label: constant.RateTypeFixed.String(),
			Value: constant.RateTypeFixed.Val(),
			Code:  constant.RateTypeFixed.Code(),
		},
	}
}
