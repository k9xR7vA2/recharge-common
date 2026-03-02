package mysql

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"time"
)

type TenantProduct struct {
	ID           uint                          `json:"id" gorm:"primaryKey"`
	BusinessType constant.BusinessType         `json:"business_type" gorm:"not null;comment:业务类型"`
	TenantID     uint                          `json:"tenant_id" gorm:"not null;comment:租户ID"`
	ProductID    uint                          `json:"product_id" gorm:"not null;comment:产品ID"`
	Status       constant.TenantBusinessStatus `json:"status" gorm:"not null;default:1;comment:状态 1:启用 2:禁用,3未绑定"` //1:启用 2:禁用,3未绑定
	Remark       string                        `json:"remark" gorm:"size:255;comment:备注说明"`
	CreatedAt    time.Time                     `json:"created_at"`
	UpdatedAt    time.Time                     `json:"updated_at"`

	// 关联字段
	Tenant Tenant `gorm:"foreignKey:TenantID;references:TenantID" json:"tenant,omitempty"`

	Product Product `gorm:"foreignKey:ID;references:ProductID" json:"product,omitempty"`
}

func (TenantProduct) TableName() string {
	return "as_tenant_products"
}
