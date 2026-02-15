package channel

import (
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/model/mysql/tenant"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

// TenantCompositeChannel 租户组合通道表
type TenantCompositeChannel struct {
	ID          uint64                        `json:"id" gorm:"primary_key;column:id;type:bigint unsigned;not null;auto_increment"`
	Payment     constant.PaymentType          `json:"payment" gorm:"column:payment;type:tinyint;not null;default:1;comment:支付方式"`
	TenantID    uint                          `json:"tenant_id" gorm:"column:tenant_id;type:int unsigned;not null;comment:租户ID"`
	ChannelCode string                        `json:"channel_code" gorm:"column:channel_code;type:varchar(50);not null;comment:通道编码(2xx)"`
	ChannelName string                        `json:"channel_name" gorm:"column:channel_name;type:varchar(100);not null;comment:通道名称"`
	Status      constant.TenantBusinessStatus `json:"status" gorm:"column:status;type:tinyint;not null;default:1;comment:状态: 1-启用 2-禁用"`
	RuleType    constant.ChannelMatchRule     `json:"rule_type" gorm:"column:rule_type;type:varchar(20);not null;comment:规则类型：WEIGHT-权重模式 AMOUNT_MAPPING-金额映射 MIXED-混合权重"`
	RuleConfig  datatypes.JSON                `json:"rule_config" gorm:"column:rule_config;type:json;not null;comment:规则配置JSON"`
	Remark      string                        `json:"remark" gorm:"column:remark;type:varchar(255);comment:备注说明"`
	CreatedAt   time.Time                     `json:"created_at"`
	UpdatedAt   time.Time                     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt                `gorm:"index" json:"-"` // 删除时间

	Tenant tenant.Tenant `gorm:"foreignKey:TenantID;references:TenantID" json:"tenant,omitempty"`
}

// TableName 表名
func (t *TenantCompositeChannel) TableName() string {
	return "as_tenant_composite_channels"
}
