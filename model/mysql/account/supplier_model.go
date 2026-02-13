package account

import (
	"github.com/shopspring/decimal"
	"github.com/small-cat1/recharge-common/model/mysql/tenant"
	"gorm.io/gorm"
	"time"
)

// Supplier [...]
type Supplier struct {
	SupplierID      int             `json:"supplier_id" gorm:"autoIncrement:true;primaryKey;column:supplier_id;type:int unsigned;not null" `
	TenantID        uint            `json:"tenant_id" gorm:"primaryKey;column:tenant_id;type:int unsigned;not null;comment:'租户ID'" `               // 租户ID
	SupplierName    string          `json:"supplier_name" gorm:"column:supplier_name;type:varchar(255);not null;comment:'供货商名称'" `                 // 供货商名称
	SupplierAccount string          `json:"supplier_account" gorm:"column:supplier_account;unique;type:varchar(255);not null;comment:'供货商账号'" `    // 供货商账号
	SupplierPasswd  string          `json:"-" gorm:"column:supplier_passwd;type:varchar(500);not null;comment:'供货商密码'" `                           // 供货商密码
	Status          int             `json:"status" gorm:"column:status;type:tinyint;not null;comment:'状态1开启，2关闭'" `                                // 状态1开启，2关闭
	AppID           string          `json:"app_id" gorm:"column:app_id;type:varchar(500);not null;comment:'商户appid'" `                             // 商户appid
	AppSecret       string          `json:"app_secret" gorm:"column:app_secret;type:varchar(500);not null;comment:'商户密钥'" `                        // 商户密钥
	PrepaidBalance  decimal.Decimal `json:"prepaid_balance" gorm:"column:prepaid_balance;type:decimal(12,2);not null;default:0.00;comment:'预付额度'"` //预付额度
	CreditLimit     decimal.Decimal `json:"credit_limit" gorm:"column:credit_limit;type:decimal(12,2);not null;default:0.00;comment:'授信额度'"`       // 授信额度
	UsedCredit      decimal.Decimal `json:"used_credit" gorm:"column:used_credit;type:decimal(12,2);not null;default:0.00;comment:'已用授信额度'"`       // 已用授信额度
	CreditStatus    int             `json:"credit_status" gorm:"column:credit_status;type:tinyint;not null;comment:'状态1正常，2冻结'" `                  // 状态1开启，2关闭
	Weight          int             `json:"weight" gorm:"column:weight;type:int;not null;comment:'进单权重'" `                                         // 进单权重
	CreatedAt       time.Time       `json:"created_at"`                                                                                            // 创建时间
	UpdatedAt       time.Time       `json:"updated_at"`                                                                                            // 更新时间
	DeletedAt       gorm.DeletedAt  `gorm:"index" json:"-"`                                                                                        // 删除时间

	Tenant tenant.Tenant `gorm:"foreignKey:TenantID;references:TenantID" json:"tenant,omitempty"`
}

// TableName get sql table name.获取数据库表名
func (Supplier) TableName() string {
	return "as_suppliers"
}
