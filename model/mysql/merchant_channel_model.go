package mysql

import (
	constant "github.com/k9xR7vA2/recharge-common/constant"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"time"
)

type MerchantChannel struct {
	TenantID       uint                          `gorm:"uniqueIndex:uk_tenant_merchant_channel;index:idx_tenant_id;column:tenant_id;type:int unsigned;not null;comment:'租户ID'" json:"tenant_id"` // 租户ID
	BusinessType   constant.BusinessType         `json:"business_type" gorm:"column:business_type;type:varchar(100);not null;comment:业务类型"`
	MerchantID     uint                          `gorm:"uniqueIndex:uk_tenant_merchant_channel;index:idx_merchant_id;index:idx_channel_code_merchant;column:merchant_id;type:int unsigned;not null;comment:'商户ID'" json:"merchant_id"` // 商户ID
	ChannelType    constant.ChannelType          `gorm:"uniqueIndex:uk_tenant_merchant_channel;column:channel_type;type:tinyint;not null;comment:'通道类型: 1-基础通道 2-组合通道'" json:"channel_type"`                                           // 通道类型: 1-基础通道 2-组合通道
	ChannelID      int                           `gorm:"uniqueIndex:uk_tenant_merchant_channel;column:channel_id;type:int;not null;comment:'通道ID/通道编码'" json:"channel_id"`                                                             // 通道ID/通道编码
	ChannelCode    string                        `gorm:"index:idx_channel_code_merchant;column:channel_code;type:varchar(50);not null;comment:'通道编码(2xx)'" json:"channel_code"`                                                        // 通道编码(2xx)
	SettlementRate decimal.Decimal               `gorm:"column:settlement_rate;type:decimal(10,2);not null;default:0.00;comment:'结算费率'" json:"settlement_rate"`                                                                        // 结算费率
	Amounts        datatypes.JSON                `gorm:"column:amounts;type:json;not null;comment:'金额'" json:"amounts"`                                                                                                                // 金额
	Status         constant.TenantBusinessStatus `gorm:"index:idx_status;column:status;type:tinyint;not null;default:1;comment:'状态 1:启用 2:禁用,3未绑定'" json:"status"`                                                                     // 状态 1:启用 2:禁用,3未绑定
	CreatedAt      *time.Time                    `gorm:"column:created_at;type:datetime;default:null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      *time.Time                    `gorm:"column:updated_at;type:datetime;default:null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *MerchantChannel) TableName() string {
	return "as_merchant_channels"
}
