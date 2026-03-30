package schema

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict"
	"github.com/k9xR7vA2/recharge-common/schema/types"
)

// GetProductSchema 平台后台产品表单用：ProductFields + SystemFields
func GetProductSchema(businessType constant.BusinessType) []types.SchemaField {
	initRegistry()
	bs := registry[businessType]
	all := append(bs.ProductFields, bs.SystemFields...)
	return buildFields(all)
}

// GetPublicSchema 租户/供货商产品展示用：只有 ProductFields
func GetPublicSchema(businessType constant.BusinessType) []types.SchemaField {
	initRegistry()
	return buildFields(registry[businessType].ProductFields)
}

// GetAccountSchema 账号表单用：AccountFields，API模式返回 nil
func GetAccountSchema(businessType constant.BusinessType) []types.SchemaField {
	initRegistry()
	if !businessType.IsAccountMode() {
		return nil
	}
	return buildFields(registry[businessType].AccountFields)
}

func buildFields(raws []types.RawField) []types.SchemaField {
	result := make([]types.SchemaField, 0, len(raws))
	for _, r := range raws {
		field := types.SchemaField{
			Key:      r.Key,
			Label:    r.Label,
			Type:     r.Type,
			Required: r.Required,
			Options:  r.Options,
			Min:      r.Min,
			Max:      r.Max,
			Unit:     r.Unit,
			TagType:  r.TagType,
		}
		if r.DictKey != "" {
			if d := dict.GetDict(r.DictKey); d != nil {
				opts := make([]types.SchemaOption, len(d.Options))
				for i, o := range d.Options {
					opts[i] = types.SchemaOption{Label: o.Label, Value: o.Value}
				}
				field.Options = opts
			}
		}
		result = append(result, field)
	}
	return result
}
