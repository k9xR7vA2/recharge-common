package mysql

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/small-cat1/recharge-common/constant"
	"gorm.io/datatypes"
	"time"
)

// SupplierProduct [...]
type SupplierProduct struct {
	BusinessType   constant.BusinessType         `json:"business_type" gorm:"not null;comment:业务类型"`
	SupplierID     uint                          `gorm:"primaryKey;column:supplier_id;type:int unsigned;not null" json:"supplier_id"`
	ProductID      uint                          `gorm:"primaryKey;column:product_id;type:int unsigned;not null" json:"product_id"`
	TenantID       uint                          `gorm:"column:tenant_id;type:int unsigned;not null" json:"tenant_id"`
	SettlementRate decimal.Decimal               `gorm:"column:settlement_rate;type:decimal(10,2) unsigned;not null;comment:'手续费'" json:"settlement_rate"` // 手续费
	Amounts        datatypes.JSON                `gorm:"column:amounts;type:json;not null;comment:'开通金额,逗号分隔'" json:"amounts"`                             // 开通的金额配置
	Status         constant.TenantBusinessStatus `gorm:"column:status;type:tinyint;not null;default:1;comment:状态 1:启用 2:禁用,3未绑定" json:"status"`
	CreatedAt      *time.Time                    `json:"created_at"`
	UpdatedAt      *time.Time                    `json:"updated_at"`

	Tenant   Tenant   `gorm:"-;foreignKey:TenantID;references:TenantID" json:"tenant,omitempty"`
	Supplier Supplier `gorm:"-;foreignKey:SupplierID;references:TenantID" json:"supplier,omitempty"`
}

// TableName get sql table name.获取数据库表名
func (SupplierProduct) TableName() string {
	return "as_supplier_products"
}

func (s SupplierProduct) GetAmounts() []uint {
	var amounts []uint
	_ = json.Unmarshal(s.Amounts, &amounts)
	return amounts
}
