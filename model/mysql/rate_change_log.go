package mysql

import (
	"github.com/shopspring/decimal"
	"github.com/k9xR7vA2/recharge-common/constant"
	"gorm.io/datatypes"
	"time"
)

// RateChangeLog 费率变更记录表
type RateChangeLog struct {
	ID            uint                          `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID      uint                          `gorm:"not null;index" json:"tenant_id"`
	Type          int                           `gorm:"not null;comment:1-供货商产品 2-商户通道" json:"type"`
	TargetID      uint                          `gorm:"not null;comment:(产品/通道)ID" json:"target_id"`
	TargetCode    string                        `gorm:"not null;comment:(产品/通道)编码" json:"target_code"`
	TargetName    string                        `gorm:"not null;comment:(产品/通道)名称" json:"target_name"`
	AccountID     uint                          `gorm:"not null;comment:(供货商/商户)ID" json:"account_id"`
	AccountName   string                        `gorm:"not null;comment:(供货商/商户)名称" json:"account_name"`
	OldRate       decimal.Decimal               `gorm:"type:decimal(10,2);not null" json:"old_rate"`
	NewRate       decimal.Decimal               `gorm:"type:decimal(10,2);not null" json:"new_rate"`
	OldStatus     constant.TenantBusinessStatus `gorm:"not null;comment:旧绑定状态" json:"old_status"`
	NewStatus     constant.TenantBusinessStatus `gorm:"not null;comment:新绑定状态" json:"new_status"`
	OldAmounts    datatypes.JSON                `gorm:"column:old_amounts;type:json;not null;comment:'旧开通金额,逗号分隔'" json:"old_amounts"`
	NewAmounts    datatypes.JSON                `gorm:"column:new_amounts;type:json;not null;comment:'新开通金额,逗号分隔'" json:"new_amounts"`
	OperatorID    uint                          `gorm:"not null;comment:操作人ID" json:"operator_id"`
	Operator      string                        `gorm:"type:varchar(100);not null;comment:操作人" json:"operator"`
	Remark        string                        `gorm:"type:varchar(500)" json:"remark"`
	EffectiveTime time.Time                     `gorm:"not null"`
	CreatedAt     time.Time
}

func (RateChangeLog) TableName() string {
	return "as_rate_change_logs"
}
