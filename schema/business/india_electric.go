package business

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/schema/types"
)

var IndiaElectric = types.BusinessSchema{
	BusinessType: constant.IndiaElectric,
	ProductFields: []types.RawField{
		{
			Key:      "operator_ids",
			Label:    "运营商",
			Type:     types.FieldTypeCheckbox,
			Required: true,
			DictKey:  "india_electric_operator", // 走字典
		},
	},
	AccountFields: []types.RawField{
		{
			Key:      "operator_id",
			Label:    "运营商",
			Type:     types.FieldTypeSelect,
			Required: true,
			DictKey:  "india_electric_operator",
		},
	},
}
