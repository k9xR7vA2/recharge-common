package adapters

import (
	"recharge-common/constant"
	"recharge-common/dict/types"
)

// AreaScopeDict 区域范围字典
type AreaScopeDict struct{}

func (d *AreaScopeDict) GetKey() string {
	return "area_code"
}

func (d *AreaScopeDict) GetName() string {
	return "区域范围"
}

func (d *AreaScopeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.National.String(),
			Value: int(constant.National),
			Code:  constant.National.Code(),
		},
		{
			Label: constant.Province.String(),
			Value: int(constant.Province),
			Code:  constant.Province.Code(),
		},
	}
}
