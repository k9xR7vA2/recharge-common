package business

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/schema/types" // ← 改这里
)

var IndiaDTH = types.BusinessSchema{
	BusinessType: constant.IndiaDTH,
	Fields: []types.RawField{
		{
			Key:      "operator",
			Label:    "支持运营商",
			Type:     types.FieldTypeCheckbox, // ← schema.xxx 全改 types.xxx
			Scene:    types.SceneProduct,
			Required: true,
			DictKey:  "india_dth_operator",
			TagType:  "success",
		},
		{
			Key:      "valid_time",
			Label:    "订单有效期",
			Type:     types.FieldTypeInputNumber,
			Scene:    types.SceneProduct,
			Required: true,
			Min:      types.Ptr(1),
			Unit:     "秒",
			TagType:  "info",
		},
		{
			Key:      "operator",
			Label:    "运营商",
			Type:     types.FieldTypeSelect,
			Scene:    types.SceneAccount,
			Required: true,
			DictKey:  "india_dth_operator",
		},
	},
}
