package types

import "github.com/k9xR7vA2/recharge-common/constant"

type FieldType string
type SceneType string

const (
	FieldTypeRadio       FieldType = "radio"
	FieldTypeCheckbox    FieldType = "checkbox"
	FieldTypeSelect      FieldType = "select"
	FieldTypeInput       FieldType = "input"
	FieldTypeInputNumber FieldType = "input-number"
)

const (
	SceneProduct SceneType = "product"
	SceneAccount SceneType = "account"
)

type SchemaField struct {
	Key      string         `json:"key"`
	Label    string         `json:"label"`
	Type     FieldType      `json:"type"`
	Scene    SceneType      `json:"scene"`
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
	Scene    SceneType
	Required bool
	DictKey  string
	Options  []SchemaOption
	Min      *float64
	Max      *float64
	Unit     string
	TagType  string
}

type BusinessSchema struct {
	BusinessType constant.BusinessType // 用 string 避免再引入 constant，或直接用 constant
	Fields       []RawField
}

func Ptr(v float64) *float64 { return &v }
