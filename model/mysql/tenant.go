package mysql

import (
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

type TenantView struct {
	TenantID      uint            `json:"tenant_id" gorm:"column:tenant_id"`
	BusinessType  datatypes.JSON  `json:"business_type" gorm:"column:business_type"`
	AgentID       int             `gorm:"column:agent_id;type:int;not null;default:0;comment:'代理ID'" json:"agent_id"`                          // 代理ID
	RebateSwitch  int             `gorm:"column:rebate_switch;type:tinyint;not null;comment:'返佣开关,1关闭，2开启'" json:"rebate_switch"`              // 返佣开关,1关闭，2开启
	Rebate        decimal.Decimal `gorm:"column:rebate;type:decimal(10,2);not null;comment:'代理佣金汇率'" json:"rebate"`                            // 代理佣金汇率
	RentBalance   decimal.Decimal `gorm:"column:rent_balance;type:decimal(12,2);not null;default:0.00;comment:'租金账户余额'" json:"rent_balance"`   // 租金账户余额
	IsCredit      int             `gorm:"column:is_credit;type:tinyint;not null;default:1;comment:'是否是授信账户1不是，2是'" json:"is_credit"`           // 是否是授信账户
	CreditBalance decimal.Decimal `gorm:"column:credit_balance;type:decimal(10,2);not null;default:0.00;comment:'授信额度'" json:"credit_balance"` // 授信额度
	TotalBalance  decimal.Decimal `gorm:"column:total_balance;type:decimal(12,2);not null;default:0.00;comment:'租户总收益'" json:"total_balance"`  // 租户总收益
	IsTest        int             `gorm:"column:is_test;type:tinyint;not null;default:1;comment:'状态1是，2不是'" json:"is_test"`                    // 是否是测试账户 1是，2不是
}

// TableName 指定表名
func (TenantView) TableName() string {
	return "as_tenants"
}
