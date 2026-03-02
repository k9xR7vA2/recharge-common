package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
	"strconv"
)

// ProvinceDict 省份字典
type ProvinceDict struct{}

func (d *ProvinceDict) GetKey() string {
	return "province"
}

func (d *ProvinceDict) GetName() string {
	return "省份"
}

func (d *ProvinceDict) GetOptions() []types.DictOption {
	options := make([]types.DictOption, 0, len(constant.ProvinceList))
	for _, province := range constant.ProvinceList {
		options = append(options, types.DictOption{
			Label: province.Label,
			Value: province.Value,
			Code:  strconv.Itoa(province.Value), // 使用 value 作为 code
		})
	}
	return options
}
