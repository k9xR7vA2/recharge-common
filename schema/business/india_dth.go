package business

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/schema"
)

var IndiaDTH = schema.BusinessSchema{
	BusinessType: constant.IndiaDTH,
	Fields: []schema.RawField{
		// ── product scene ──
		{
			Key:      "operator",
			Label:    "支持运营商",
			Type:     schema.FieldTypeCheckbox,
			Scene:    schema.SceneProduct,
			Required: true,
			DictKey:  "india_dth_operator",
			TagType:  "success",
		},
		{
			Key:      "valid_time",
			Label:    "订单有效期",
			Type:     schema.FieldTypeInputNumber,
			Scene:    schema.SceneProduct,
			Required: true,
			Min:      schema.Ptr(1),
			Unit:     "秒",
			TagType:  "info",
		},
		// ── account scene ──
		{
			Key:      "operator",
			Label:    "运营商",
			Type:     schema.FieldTypeSelect,
			Scene:    schema.SceneAccount,
			Required: true,
			DictKey:  "india_dth_operator",
		},
	},
}
