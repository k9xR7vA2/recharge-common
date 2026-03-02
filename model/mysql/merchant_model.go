package mysql

import (
	"github.com/shopspring/decimal"
	"github.com/k9xR7vA2/recharge-common/constant"
	"gorm.io/gorm"
	"time"
)

// Merchant  [...]
type Merchant struct {
	MerchantID      uint                         `gorm:"autoIncrement:true;primaryKey;column:merchant_id;type:int unsigned;not null" json:"merchant_id"`
	TenantID        uint                         `gorm:"primaryKey;column:tenant_id;type:int unsigned;not null;comment:'租户ID'" json:"tenant_id"`           // 租户ID
	MerchantName    string                       `gorm:"column:merchant_name;type:varchar(255);not null;comment:'商户名称'" json:"merchant_name"`              // 商户名称
	MerchantAccount string                       `gorm:"column:merchant_account;unique;type:varchar(255);not null;comment:'商户账号'" json:"merchant_account"` // 商户账号
	MerchantPasswd  string                       `gorm:"column:merchant_passwd;type:varchar(500);not null;comment:'商户密码'" json:"-"`                        // 商户密码
	Status          constant.GlobalAccountStatus `gorm:"column:status;type:tinyint;not null;comment:'状态1开启，2关闭'" json:"status"`                            // 状态1开启，2关闭
	AppID           string                       `json:"app_id" gorm:"column:app_id;type:varchar(500);not null;comment:'商户appid'" `                        // 商户appid
	AppSecret       string                       `json:"app_secret" gorm:"column:app_secret;type:varchar(500);not null;comment:'商户密钥'" `                   // 商户密钥
	Balance         decimal.Decimal              `gorm:"column:balance;type:decimal(12,2);not null;comment:'商户充值金额'" json:"balance"`                       // 商户充值金额
	PreAmount       decimal.Decimal              `gorm:"column:pre_amount;type:decimal(12,2);not null;comment:'商户预付金额'" json:"pre_amount"`                 // 商户预付金额
	CreatedAt       time.Time                    `json:"created_at"`                                                                                       // 创建时间
	UpdatedAt       time.Time                    `json:"updated_at"`                                                                                       // 更新时间
	DeletedAt       gorm.DeletedAt               `gorm:"index" json:"-"`                                                                                   // 删除时间
	Tenant          Tenant                       `gorm:"foreignKey:tenant_id;references:tenant_id"`
}

// TableName get sql table name.获取数据库表名
func (Merchant) TableName() string {
	return "as_merchants"
}

func (m Merchant) GetTenantID() uint {
	return m.TenantID
}

func (m *Merchant) GetID() uint {
	return m.MerchantID
}

func (m *Merchant) GetAppID() string {
	return m.AppID
}

func (m *Merchant) GetName() string {
	return m.MerchantName
}

func (m *Merchant) GetStatus() constant.GlobalAccountStatus {
	return m.Status
}
