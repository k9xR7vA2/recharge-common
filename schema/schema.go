package schema

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict"
)

func GetSchema(businessType constant.BusinessType) []SchemaField {
	initRegistry()
	return buildFields(registry[businessType])
}

func GetSchemaByScene(businessType constant.BusinessType, scene SceneType) []SchemaField {
	initRegistry()
	var result []SchemaField
	for _, f := range buildFields(registry[businessType]) {
		if f.Scene == scene {
			result = append(result, f)
		}
	}
	return result
}

func buildFields(raws []RawField) []SchemaField {
	result := make([]SchemaField, 0, len(raws))
	for _, r := range raws {
		field := SchemaField{
			Key:      r.Key,
			Label:    r.Label,
			Type:     r.Type,
			Scene:    r.Scene,
			Required: r.Required,
			Options:  r.Options,
			Min:      r.Min,
			Max:      r.Max,
			Unit:     r.Unit,
			TagType:  r.TagType,
		}
		// 有 DictKey 的从字典动态取 options
		if r.DictKey != "" {
			if d := dict.GetDict(r.DictKey); d != nil {
				opts := make([]SchemaOption, len(d.Options))
				for i, o := range d.Options {
					opts[i] = SchemaOption{Label: o.Label, Value: o.Value}
				}
				field.Options = opts
			}
		}
		result = append(result, field)
	}
	return result
}

//```
//
//---
//
//## 新增业务类型只需两步
//```
//1. 新建 schema/business/xxx.go，定义字段
//2. registry.go 的列表里加一行
