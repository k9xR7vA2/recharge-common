package mysql

import (
	"github.com/gin-gonic/gin"
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/model/mysql/tenant"
	"gorm.io/gorm"
	"time"
)

// Merchant [...]
type Merchant struct {
	MerchantID      uint                         `gorm:"autoIncrement:true;primaryKey;column:merchant_id;type:int unsigned;not null" json:"merchant_id"`
	TenantID        uint                         `gorm:"primaryKey;column:tenant_id;type:int unsigned;not null;comment:'租户ID'" json:"tenant_id"`           // 租户ID
	MerchantName    string                       `gorm:"column:merchant_name;type:varchar(255);not null;comment:'商户名称'" json:"merchant_name"`              // 商户名称
	MerchantAccount string                       `gorm:"column:merchant_account;unique;type:varchar(255);not null;comment:'商户账号'" json:"merchant_account"` // 商户账号
	MerchantPasswd  string                       `gorm:"column:merchant_passwd;type:varchar(500);not null;comment:'商户密码'" json:"-"`                        // 商户密码
	Status          constant.GlobalAccountStatus `gorm:"column:status;type:tinyint;not null;comment:'状态1开启，2关闭'" json:"status"`                            // 状态1开启，2关闭
	AppID           string                       `json:"app_id" gorm:"column:app_id;type:varchar(500);not null;comment:'appid'" `                          // appid
	AppSecret       string                       `json:"app_secret" gorm:"column:app_secret;type:varchar(500);not null;comment:'密钥'" `                     // 密钥
	Balance         float64                      `gorm:"column:balance;type:decimal(12,2);not null;comment:'商户充值金额'" json:"balance"`                       // 商户充值金额
	PreAmount       float64                      `gorm:"column:pre_amount;type:decimal(12,2);not null;comment:'商户预付金额'" json:"pre_amount"`                 // 商户预付金额
	SettingAmount   float64                      `gorm:"column:setting_amount;type:decimal(12,2);not null;comment:'商户结算金额'" json:"setting_amount"`         // 商户结算金额
	CreatedAt       time.Time                    `json:"created_at"`                                                                                       // 创建时间
	UpdatedAt       time.Time                    `json:"updated_at"`                                                                                       // 更新时间
	DeletedAt       gorm.DeletedAt               `gorm:"index" json:"-"`                                                                                   // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m Merchant) TableName() string {
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

type MerchantContext struct {
	*gin.Context
	MerchantInfo *Merchant
	TenantInfo   *tenant.Tenant
}

// GetMerchant 获取供应商信息的辅助方法
func (mc *MerchantContext) GetMerchant() *Merchant {
	return mc.MerchantInfo
}

func (mc *MerchantContext) GetTenantInfo() *tenant.Tenant {
	return mc.TenantInfo
}
