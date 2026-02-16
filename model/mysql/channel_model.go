package mysql

import (
	"encoding/json"
	"github.com/small-cat1/recharge-common/constant"
	"gorm.io/datatypes"
	"time"
)

type Channel struct {
	ID            uint                          `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	InterfaceID   uint                          `json:"interface_id" gorm:"column:interface_id;type:int unsigned;not null"`
	ChannelCode   string                        `json:"channel_code" gorm:"column:channel_code;type:varchar(50);not null;uniqueIndex:uk_channel_code;comment:通道编码"`
	ChannelName   string                        `json:"channel_name" gorm:"column:channel_name;type:varchar(100);not null;comment:通道名称"`
	BusinessType  constant.BusinessType         `json:"business_type" gorm:"column:business_type;type:varchar(100);not null;comment:业务类型"` // 业务类型：话费/游戏等
	Status        constant.TenantBusinessStatus `json:"status" gorm:"column:status;type:tinyint;not null;default:1;comment:状态: 1-启用2-禁用"`
	Payment       constant.PaymentType          `json:"payment" gorm:"column:payment;type:tinyint;not null;default:1;comment:支付方式"`
	Device        constant.DeviceType           `json:"device" gorm:"column:device;type:tinyint;not null;default:1;comment:设备"`
	Amount        datatypes.JSON                `json:"amount" gorm:"column:amount;type:json;not null;comment:金额"`
	PaymentMethod string                        `json:"payment_method" gorm:"column:payment_method;type:varchar(100);not null;comment:支付方法"`
	Attributes    datatypes.JSON                `json:"attributes" gorm:"column:attributes;type:json;not null;comment:JSON格式的特定属性"` // JSON格式的特定属性

	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Interface AsInterface `json:"interface,omitempty" gorm:"foreignKey:InterfaceID;references:InterfaceID" `

	// 添加多对多关系
	Tenants []Tenant `json:"tenants,omitempty" gorm:"many2many:as_tenant_channels;foreignKey:ID;joinForeignKey:ChannelID;References:TenantID;joinReferences:TenantID"`

	TenantChannels []TenantChannel `json:"tenant_channels,omitempty" gorm:"foreignKey:ChannelID;references:ID"`
}

// TableName 指定表名
func (Channel) TableName() string {
	return "as_channels"
}

func (g Channel) GetAmounts() []int {
	var amounts []int
	_ = json.Unmarshal(g.Amount, &amounts)
	return amounts
}

type MobileAttrs struct {
	Carrier      constant.CarrierType `json:"carrier"`
	AreaCode     constant.AreaScope   `json:"area_code"`
	ProvinceCode int                  `json:"province_code"`
	ChargeSpeed  constant.ChargeSpeed `json:"charge_speed"`
}

func (g Channel) GetMobileAttrs() (*MobileAttrs, error) {
	var attr MobileAttrs
	err := json.Unmarshal(g.Attributes, &attr)
	if err != nil {
		return nil, err
	}
	return &attr, nil
}

type IndiaMobileAttrs struct {
	Carrier     constant.IndiaCarrierType `json:"carrier"`
	ChargeSpeed constant.ChargeSpeed      `json:"charge_speed"`
}

func (g Channel) GetIndiaMobileAttrs() (*IndiaMobileAttrs, error) {
	var attr IndiaMobileAttrs
	err := json.Unmarshal(g.Attributes, &attr)
	if err != nil {
		return nil, err
	}
	return &attr, nil
}
