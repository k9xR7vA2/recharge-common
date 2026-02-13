package product

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// ProductSku 产品SKU表/套餐表
type ProductSku struct {
	ID          uint            `json:"id" gorm:"primarykey;comment:主键ID"`
	ProductID   uint            `json:"productId" gorm:"column:product_id;not null;index:idx_product_id;comment:关联产品ID(SPU)"`
	PlanID      string          `json:"planId" gorm:"column:plan_id;type:varchar(50);not null;uniqueIndex:uk_plan_id;comment:套餐编号(三方接口plan_id)"`
	Amount      decimal.Decimal `json:"amount" gorm:"column:amount;type:decimal(10,2);not null;index:idx_amount;comment:面额"`
	Currency    string          `json:"currency" gorm:"column:currency;type:varchar(10);default:INR;comment:货币"`
	Description string          `json:"description" gorm:"column:description;type:text;comment:套餐说明"`
	Status      int             `json:"status" gorm:"column:status;type:tinyint;default:1;comment:状态:1启用,2禁用"`
	CreatedAt   *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime(3);comment:创建时间"`
	UpdatedAt   *time.Time      `json:"updatedAt" gorm:"column:updated_at;type:datetime(3);comment:更新时间"`
	DeletedAt   gorm.DeletedAt  `json:"-" gorm:"index;comment:删除时间"`
	Product     Product         `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID"`
}

// TableName 指定表名
func (ProductSku) TableName() string {
	return "as_product_skus"
}
