package product

import (
	"encoding/json"
	constant "github.com/small-cat1/recharge-common/constant"
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
	Amount       datatypes.JSON                `json:"amount" gorm:"column:amount;type:json;not null;default:[];comment:金额"`
	ValidTime    int                           `json:"valid_time" gorm:"column:valid_time;type:int;not null;default:1;comment:订单有效期"`
	Attributes   datatypes.JSON                `json:"attributes" gorm:"column:attributes;type:json;not null;comment:业务属性"`
	CreatedAt    time.Time                     `json:"created_at"`
	UpdatedAt    time.Time                     `json:"updated_at"`

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

func (g Product) GetAmounts() []uint {
	var amounts []uint
	_ = json.Unmarshal(g.Amount, &amounts)
	return amounts
}

func (g Product) GetValidTime() int {
	return g.ValidTime
}
