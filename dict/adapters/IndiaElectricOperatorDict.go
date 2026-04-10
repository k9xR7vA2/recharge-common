package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

type IndiaElectricOperatorDict struct{}

func (d *IndiaElectricOperatorDict) GetKey() string {
	return "india_electric_operator"
}

func (d *IndiaElectricOperatorDict) GetName() string {
	return "印度电费运营商"
}

func (d *IndiaElectricOperatorDict) GetOptions() []types.DictOption {
	ts := constant.GetAllElectricOperators()
	options := make([]types.DictOption, 0, len(ts))
	for _, t := range ts {
		options = append(options, types.DictOption{
			Label: t.Label,
			Value: t.Value,
		})
	}
	return options
}
