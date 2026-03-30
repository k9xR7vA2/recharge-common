package business

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/schema/types"
)

var IndiaMobile = types.BusinessSchema{
	BusinessType: constant.IndiaMobile,
	// API模式，无 AccountFields
	// has_sku 移出属性，作为产品独立字段处理
	ProductFields: []types.RawField{
		{
			Key:      "carrier",
			Label:    "运营商",
			Type:     types.FieldTypeRadio,
			Required: true,
			DictKey:  "india_carrier_type",
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
