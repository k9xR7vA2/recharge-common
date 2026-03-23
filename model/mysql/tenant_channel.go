package mysql

import (
	constant "github.com/k9xR7vA2/recharge-common/constant"
	"gorm.io/datatypes"
	"time"
)

// TenantChannel 租户通道关联
type TenantChannel struct {
	ID                uint                          `json:"id"`
	BusinessType      constant.BusinessType         `json:"business_type" gorm:"column:business_type;not null;comment:业务类型"`
	TenantID          uint                          `json:"tenant_id" gorm:"column:tenant_id;not null;comment:租户ID"`
	ChannelID         uint                          `json:"channel_id" gorm:"column:channel_id;not null;comment:通道ID"`
	TenantChannelCode string                        `json:"tenant_channel_code" gorm:"column:tenant_channel_code;type:varchar(50);not null;uniqueIndex:uk_tenant_channel_code;comment:租户通道编码"`
	SelectedAmounts   datatypes.JSON                `json:"selected_amounts"    gorm:"column:selected_amounts;type:json;comment:租户选中的金额"` // 新增
	Status            constant.TenantBusinessStatus `json:"status" gorm:"column:status;not null;default:1;comment:状态 1:启用 2:禁用,3未绑定"`     //1:启用 2:禁用,3未绑定
	Remark            string                        `json:"remark" gorm:"column:remark;type:varchar(255);comment:备注说明"`
	CreatedAt         time.Time                     `json:"created_at"`
	UpdatedAt         time.Time                     `json:"updated_at" `

	// 关联字段
	Tenant  Tenant  `gorm:"foreignKey:TenantID;references:TenantID" json:"tenant,omitempty"`
	Channel Channel `gorm:"foreignKey:ChannelID;references:ID"      json:"channel,omitempty"`
}

// TableName 指定表名
func (TenantChannel) TableName() string {
	return "as_tenant_channels"
}
