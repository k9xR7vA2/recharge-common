package schema

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

// SchemaField 对外暴露的字段描述（填充了 options）
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
	TagType  string         `json:"tag_type,omitempty"` // el-tag 颜色
}

type SchemaOption struct {
	Label string `json:"label"`
	Value any    `json:"value"`
}

// RawField 内部定义，不填充 options
type RawField struct {
	Key      string
	Label    string
	Type     FieldType
	Scene    SceneType
	Required bool
	DictKey  string         // 有则运行时从字典取 options
	Options  []SchemaOption // 无 DictKey 时直接用
	Min      *float64
	Max      *float64
	Unit     string
	TagType  string
}

func Ptr(v float64) *float64 { return &v }
