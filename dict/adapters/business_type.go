package adapters

import (
	"recharge-common/constant"
	"recharge-common/dict/types"
)

// BusinessTypeDict 业务类型字典
type BusinessTypeDict struct{}

func (d *BusinessTypeDict) GetKey() string {
	return "business_type"
}

func (d *BusinessTypeDict) GetName() string {
	return "业务类型"
}

func (d *BusinessTypeDict) GetOptions() []types.DictOption {
	ts := constant.GetAllBusinessTypes()
	options := make([]types.DictOption, 0, len(ts))
	for _, t := range ts {
		options = append(options, types.DictOption{
			Label: t.Label,
			Value: string(t.Value),
			Code:  string(t.Value),
		})
	}
	return options
}
