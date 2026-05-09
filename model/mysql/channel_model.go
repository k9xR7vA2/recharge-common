package mysql

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"time"
)

type Channel struct {
	ID            uint                          `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ProductID     uint                          `json:"product_id" gorm:"not null;comment:产品ID"`
	InterfaceID   uint                          `json:"interface_id" gorm:"column:interface_id;type:int unsigned;not null"`
	ChannelCode   string                        `json:"channel_code" gorm:"column:channel_code;type:varchar(50);not null;uniqueIndex:uk_channel_code;comment:通道编码"`
	ChannelName   string                        `json:"channel_name" gorm:"column:channel_name;type:varchar(100);not null;comment:通道名称"`
	Lang          string                        `json:"lang" gorm:"column:lang;type:varchar(10);default:zh;comment:语言"`
	BusinessType  constant.BusinessType         `json:"business_type" gorm:"column:business_type;type:varchar(100);not null;comment:业务类型"` // 业务类型：话费/游戏等
	Status        constant.TenantBusinessStatus `json:"status" gorm:"column:status;type:tinyint;not null;default:1;comment:状态: 1-启用2-禁用"`
	Payment       constant.PaymentType          `json:"payment" gorm:"column:payment;type:tinyint;not null;default:1;comment:支付方式"`
	Device        constant.DeviceType           `json:"device" gorm:"column:device;type:tinyint;not null;default:1;comment:设备"`
	PaymentMethod constant.PaymentMethod        `json:"payment_method" gorm:"column:payment_method;type:varchar(100);not null;comment:支付方法"`

	// 预产码配置
	PreCodeEnabled        int         `json:"pre_code_enabled" gorm:"column:pre_code_enabled;type:tinyint;not null;default:0;comment:是否支持预产码 0不支持 1支持"`
	PreCodeMaxStock       int         `json:"pre_code_max_stock" gorm:"column:pre_code_max_stock;type:int;not null;default:500;comment:最大库存上限"`
	PreCodeTtlMinutes     int         `json:"pre_code_ttl_minutes" gorm:"column:pre_code_ttl_minutes;type:int;not null;default:30;comment:码有效期(分钟)"`
	PreCodeReplenishCount int         `json:"pre_code_replenish_count" gorm:"column:pre_code_replenish_count;type:int;not null;default:100;comment:每次补充数量"`
	CreatedAt             time.Time   `json:"created_at"`
	UpdatedAt             time.Time   `json:"updated_at"`
	Interface             AsInterface `json:"interface,omitempty" gorm:"foreignKey:InterfaceID;references:InterfaceID" `
	// 添加多对多关系
	Tenants        []Tenant        `json:"tenants,omitempty" gorm:"many2many:as_tenant_channels;foreignKey:ID;joinForeignKey:ChannelID;References:TenantID;joinReferences:TenantID"`
	TenantChannels []TenantChannel `json:"tenant_channels,omitempty" gorm:"foreignKey:ChannelID;references:ID"`
	Product        Product         `gorm:"foreignKey:ProductID;references:ID" json:"product,omitempty"`
}

// TableName 指定表名
func (Channel) TableName() string {
	return "as_channels"
}

type MobileAttrs struct {
	Carrier      constant.CarrierType `json:"carrier"`
	AreaCode     constant.AreaScope   `json:"area_code"`
	ProvinceCode int                  `json:"province_code"`
	ChargeSpeed  constant.ChargeSpeed `json:"charge_speed"`
}

type IndiaMobileAttrs struct {
	Carrier     constant.IndiaCarrierType `json:"carrier"`
	ChargeSpeed constant.ChargeSpeed      `json:"charge_speed"`
}
