package business

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/schema/types"
)

var IndiaDTH = types.BusinessSchema{
	BusinessType: constant.IndiaDTH,
	// 账号池模式
	ProductFields: []types.RawField{
		{
			Key:      "operator",
			Label:    "支持运营商",
			Type:     types.FieldTypeCheckbox,
			Required: true,
			DictKey:  "india_dth_operator",
			TagType:  "success",
		},
	},
	SystemFields: []types.RawField{
		{
			Key:      "valid_time",
			Label:    "订单有效期",
			Type:     types.FieldTypeInputNumber,
			Required: true,
			Min:      types.Ptr(1),
			Unit:     "秒",
			TagType:  "info",
		},
	},
	AccountFields: []types.RawField{
		{
			Key:      "operator",
			Label:    "运营商",
			Type:     types.FieldTypeSelect,
			Required: true,
			DictKey:  "india_dth_operator",
		},
	},
}
