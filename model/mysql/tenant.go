package mysql

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"time"
)

type Tenant struct {
	TenantID      uint            `json:"tenant_id,omitempty" gorm:"autoIncrement:true;primaryKey;column:tenant_id;type:int unsigned;not null" `          // 租户ID
	AgentID       int             `json:"agent_id" gorm:"column:agent_id;type:int;not null;default:0;comment:'代理ID'" `                                    // 代理ID
	TenantName    string          `json:"tenant_name,omitempty" gorm:"column:tenant_name;type:varchar(255);not null;comment:'租户名称'" `                     // 租户名称
	TenantAccount string          `json:"tenant_account,omitempty" gorm:"column:tenant_account;unique;type:varchar(255);not null;comment:'租户账号'" `        // 租户账号
	RebateSwitch  int             `json:"rebate_switch,omitempty" gorm:"column:rebate_switch;type:tinyint;not null;comment:'返佣开关,1关闭，2开启'" `              // 返佣开关,1关闭，2开启
	Rebate        decimal.Decimal `json:"rebate,omitempty" gorm:"column:rebate;type:decimal(10,2);not null;comment:'代理佣金汇率'" `                            // 代理佣金汇率
	RentBalance   decimal.Decimal `json:"rent_balance,omitempty" gorm:"column:rent_balance;type:decimal(12,2);not null;default:0.00;comment:'租金账户余额'" `   // 租金账户余额
	IsCredit      int             `json:"is_credit,omitempty" gorm:"column:is_credit;type:tinyint;not null;default:1;comment:'是否是授信账户1不是，2是'" `           // 是否是授信账户
	CreditBalance decimal.Decimal `json:"credit_balance,omitempty" gorm:"column:credit_balance;type:decimal(10,2);not null;default:0.00;comment:'授信额度'" ` // 授信额度
	UsedCredit    decimal.Decimal `json:"used_credit" gorm:"column:used_credit;type:decimal(12,2);not null;default:0.00;comment:'已用授信额度'"`                // 已用授信额度
	TotalBalance  decimal.Decimal `json:"total_balance,omitempty" gorm:"column:total_balance;type:decimal(12,2);not null;default:0.00;comment:'租户总收益'" `  // 租户总收益
	BusinessType  datatypes.JSON  `json:"business_type,omitempty" gorm:"column:business_type;type:json;comment:'业务类型'"  `                                 //业务类型

	CreatedAt time.Time    `json:"created_at"`                                                                                  // 创建时间
	UpdatedAt time.Time    `json:"updated_at"`                                                                                  // 更新时间
	IsTest    int          `json:"is_test,omitempty" gorm:"column:is_test;type:tinyint;not null;default:1;comment:'状态1是，2不是'" ` // 是否是测试账户1是，2不是
	Users     []TenantUser `json:"users,omitempty" gorm:"foreignKey:TenantID;references:TenantID" `                             // 关联到租户的用户
	User      TenantUser   `json:"user,omitempty" gorm:"foreignKey:TenantID;references:TenantID;comment:租户信息"  `                // 租户信息
	//MobileProduct []product.MobileProduct `json:"mobile_products,omitempty" gorm:"many2many:as_tenant_products;joinForeignKey:TenantID;joinReferences:ProductID"  `
}

// TableName get sql table name.获取数据库表名
func (Tenant) TableName() string {
	return "as_tenants"
}

func (t *Tenant) IsBusinessEnabled(businessType string) bool {
	if t.BusinessType == nil {
		return false
	}
	var types []string
	if err := json.Unmarshal(t.BusinessType, &types); err != nil {
		return false
	}
	for _, bt := range types {
		if bt == businessType {
			return true
		}
	}
	return false
}
