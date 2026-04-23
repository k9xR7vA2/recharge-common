package mysql

import (
	"encoding/json"
	"fmt"
	constant "github.com/k9xR7vA2/recharge-common/constant"
	"gorm.io/datatypes"
	"time"
)

// Product 实体实现
type Product struct {
	ID           uint                          `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ProductCode  string                        `json:"product_code" gorm:"column:product_code;unique;type:varchar(50);not null;comment:产品编码"`
	ProductName  string                        `json:"product_name" gorm:"column:product_name;type:varchar(100);not null;comment:产品名称"`
	BusinessType constant.BusinessType         `json:"business_type" gorm:"column:business_type;type:varchar(100);not null;comment:业务类型"`
	Status       constant.TenantBusinessStatus `json:"status" gorm:"column:status;type:tinyint;not null;default:1;comment:状态:1正常,2下架"`
	HasSku       int                           `json:"has_sku" gorm:"column:has_sku;type:tinyint;not null;default:2;comment:是否有SKU: 1有, 2没有"`
	ValidTime    uint                          `json:"valid_time" gorm:"column:valid_time;type:int unsigned;not null;default:0;comment:订单有效期(秒)"`
	AmountType   constant.AmountType           `json:"amount_type" gorm:"not null;default:1;comment:金额类型: 1固定, 2区间, 3动态金额（套餐）4账单金额"`
	Amount       datatypes.JSON                `json:"amount" gorm:"column:amount;type:json;not null;comment:金额"`
	// 三套 schema
	ProductSchema datatypes.JSON `json:"product_schema" gorm:"column:product_schema;type:json;comment:产品属性schema+值"`
	SystemSchema  datatypes.JSON `json:"system_schema"  gorm:"column:system_schema;type:json;comment:系统配置schema+值"`
	UserSchema    datatypes.JSON `json:"user_schema"    gorm:"column:user_schema;type:json;comment:供货商下单字段定义"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`

	TenantProducts []TenantProduct `json:"tenant_products,omitempty" gorm:"foreignKey:ProductID;references:ID"`
	ProductSkus    []ProductSku    `json:"product_skus,omitempty" gorm:"foreignKey:ProductID;references:ID"`
}

func (Product) TableName() string {
	return "as_products"
}

func (g Product) GetID() uint {
	return g.ID
}

func (g Product) GetProductName() string {
	return g.ProductName
}

func (g Product) GetProductCode() string {
	return g.ProductCode
}

func (g Product) GetType() constant.BusinessType {
	return g.BusinessType
}

func (g Product) GetStatus() constant.TenantBusinessStatus {
	return g.Status
}

// SchemaFieldType 字段类型
type SchemaFieldType string

const (
	FieldTypeInput       SchemaFieldType = "input"
	FieldTypeInputNumber SchemaFieldType = "input-number"
	FieldTypeSelect      SchemaFieldType = "select"
	FieldTypeCheckbox    SchemaFieldType = "checkbox"
	FieldTypeRadio       SchemaFieldType = "radio"
	FieldTypeSwitch      SchemaFieldType = "switch"
)

// SchemaFieldOption select/checkbox/radio 的选项
type SchemaFieldOption struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

// SchemaField 通用字段定义（ProductSchema 和 SystemSchema 共用）
type SchemaField struct {
	Key      string              `json:"key"`
	Label    string              `json:"label"`
	Type     SchemaFieldType     `json:"type"`
	Required bool                `json:"required"`
	Options  []SchemaFieldOption `json:"options,omitempty"`
}

// UserSchemaField 供货商下单字段，多一个 options_from
type UserSchemaField struct {
	Key         string              `json:"key"`
	Label       string              `json:"label"`
	Type        SchemaFieldType     `json:"type"`
	Required    bool                `json:"required"`
	Options     []SchemaFieldOption `json:"options,omitempty"`
	OptionsFrom string              `json:"options_from,omitempty"` // 引用 product.xxx
}

// ProductSchemaJSON 产品属性 schema + 值
type ProductSchemaJSON struct {
	Schema     []SchemaField          `json:"schema"`
	Attributes map[string]interface{} `json:"attributes"`
}

// SystemSchemaJSON 系统配置 schema + 值
type SystemSchemaJSON struct {
	Schema     []SchemaField          `json:"schema"`
	Attributes map[string]interface{} `json:"attributes"`
}

// UserSchemaJSON 供货商下单字段定义，无 attributes
type UserSchemaJSON struct {
	Schema []UserSchemaField `json:"schema"`
}

// ----- 序列化/反序列化工具方法 -----

func ParseProductSchema(data []byte) (*ProductSchemaJSON, error) {
	var s ProductSchemaJSON
	return &s, json.Unmarshal(data, &s)
}

func ParseSystemSchema(data []byte) (*SystemSchemaJSON, error) {
	var s SystemSchemaJSON
	return &s, json.Unmarshal(data, &s)
}

func ParseUserSchema(data []byte) (*UserSchemaJSON, error) {
	var s UserSchemaJSON
	return &s, json.Unmarshal(data, &s)
}

// ResolveUserSchemaOptions 把 options_from 引用替换成真实选项
// 供货商下单渲染表单时调用
func ResolveUserSchemaOptions(userSchema *UserSchemaJSON, productSchema *ProductSchemaJSON) {
	if userSchema == nil || productSchema == nil {
		return
	}
	for i, field := range userSchema.Schema {
		if field.OptionsFrom == "" {
			continue
		}
		// options_from 格式: "product.operator_ids"
		// 从 productSchema.Attributes 里取对应的值，转成 Options
		attrKey := extractAttrKey(field.OptionsFrom)
		if attrKey == "" {
			continue
		}
		attrVal, ok := productSchema.Attributes[attrKey]
		if !ok {
			continue
		}
		userSchema.Schema[i].Options = toOptions(attrVal)
	}
}

// extractAttrKey 从 "product.operator_ids" 提取 "operator_ids"
func extractAttrKey(optionsFrom string) string {
	const prefix = "product."
	if len(optionsFrom) > len(prefix) && optionsFrom[:len(prefix)] == prefix {
		return optionsFrom[len(prefix):]
	}
	return ""
}

// toOptions 把 attributes 里的值转成 []SchemaFieldOption
// 支持 []interface{} 和 map 两种格式
func toOptions(val interface{}) []SchemaFieldOption {
	switch v := val.(type) {
	case []interface{}:
		opts := make([]SchemaFieldOption, 0, len(v))
		for _, item := range v {
			opts = append(opts, SchemaFieldOption{
				Label: fmt.Sprintf("%v", item),
				Value: item,
			})
		}
		return opts
	}
	return nil
}
