package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

// IndiaDthOperatorDict DTH运营商字典
type IndiaDthOperatorDict struct{}

func (d *IndiaDthOperatorDict) GetKey() string {
	return "india_dth_operator"
}

func (d *IndiaDthOperatorDict) GetName() string {
	return "DTH运营商"
}

func (d *IndiaDthOperatorDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.DthTataPlay.String(),
			Value: int(constant.DthTataPlay),
			Code:  constant.DthTataPlay.Code(),
		},
		{
			Label: constant.DthAirtel.String(),
			Value: int(constant.DthAirtel),
			Code:  constant.DthAirtel.Code(),
		},
		{
			Label: constant.DthDishTV.String(),
			Value: int(constant.DthDishTV),
			Code:  constant.DthDishTV.Code(),
		},
		{
			Label: constant.DthSunDirect.String(),
			Value: int(constant.DthSunDirect),
			Code:  constant.DthSunDirect.Code(),
		},
	}
}
