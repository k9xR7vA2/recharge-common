package types

import "github.com/k9xR7vA2/recharge-common/constant"

type FieldType string

const (
	FieldTypeRadio       FieldType = "radio"
	FieldTypeCheckbox    FieldType = "checkbox"
	FieldTypeSelect      FieldType = "select"
	FieldTypeInput       FieldType = "input"
	FieldTypeInputNumber FieldType = "input-number"
)

type SchemaField struct {
	Key      string         `json:"key"`
	Label    string         `json:"label"`
	Type     FieldType      `json:"type"`
	Required bool           `json:"required"`
	Options  []SchemaOption `json:"options,omitempty"`
	Min      *float64       `json:"min,omitempty"`
	Max      *float64       `json:"max,omitempty"`
	Unit     string         `json:"unit,omitempty"`
	TagType  string         `json:"tag_type,omitempty"`
}

type SchemaOption struct {
	Label string `json:"label"`
	Value any    `json:"value"`
}

type RawField struct {
	Key      string
	Label    string
	Type     FieldType
	Required bool
	DictKey  string
	Options  []SchemaOption
	Min      *float64
	Max      *float64
	Unit     string
	TagType  string
}

// BusinessSchema 一个业务类型的完整定义
type BusinessSchema struct {
	BusinessType  constant.BusinessType
	ProductFields []RawField // 业务属性：平台+租户+供货商可见，平台可编辑
	SystemFields  []RawField // 系统配置：仅平台可见可编辑
	AccountFields []RawField // 账号属性：账号池模式专用，nil表示API模式
}

func Ptr(v float64) *float64 { return &v }
