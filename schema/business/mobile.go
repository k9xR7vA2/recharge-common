package business

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/schema/types"
)

var Mobile = types.BusinessSchema{
	BusinessType: constant.MobileBusiness,
	// API模式，无 AccountFields
	ProductFields: []types.RawField{
		{
			Key:      "carrier",
			Label:    "运营商",
			Type:     types.FieldTypeRadio,
			Required: true,
			DictKey:  "carrier_type",
			TagType:  "success",
		},
		{
			Key:      "charge_speed",
			Label:    "充值速度",
			Type:     types.FieldTypeRadio,
			Required: true,
			DictKey:  "charge_speed",
			TagType:  "warning",
		},
		{
			Key:      "area_code",
			Label:    "区域范围",
			Type:     types.FieldTypeRadio,
			Required: true,
			DictKey:  "area_code",
			TagType:  "primary",
		},
		{
			Key:      "is_check_isp",
			Label:    "携号转网检测",
			Type:     types.FieldTypeRadio,
			Required: true,
			Options: []types.SchemaOption{
				{Label: "是", Value: 1},
				{Label: "否", Value: 2},
			},
			TagType: "warning",
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
}
