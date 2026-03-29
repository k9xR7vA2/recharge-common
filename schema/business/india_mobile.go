package business

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/schema"
)

var IndiaMobile = schema.BusinessSchema{
	BusinessType: constant.IndiaMobile,
	Fields: []schema.RawField{
		{
			Key:      "carrier",
			Label:    "运营商",
			Type:     schema.FieldTypeRadio,
			Scene:    schema.SceneProduct,
			Required: true,
			DictKey:  "india_carrier_type",
			TagType:  "success",
		},
		{
			Key:      "charge_speed",
			Label:    "充值速度",
			Type:     schema.FieldTypeRadio,
			Scene:    schema.SceneProduct,
			Required: true,
			DictKey:  "charge_speed",
			TagType:  "warning",
		},
		{
			Key:      "has_sku",
			Label:    "是否有SKU",
			Type:     schema.FieldTypeRadio,
			Scene:    schema.SceneProduct,
			Required: true,
			Options: []schema.SchemaOption{
				{Label: "是", Value: 1},
				{Label: "否", Value: 2},
			},
			TagType: "info",
		},
		{
			Key:      "is_check_isp",
			Label:    "携号转网检测",
			Type:     schema.FieldTypeRadio,
			Scene:    schema.SceneProduct,
			Required: true,
			Options: []schema.SchemaOption{
				{Label: "是", Value: 1},
				{Label: "否", Value: 2},
			},
			TagType: "warning",
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
	},
}
