package mysql

import (
	"github.com/shopspring/decimal"
	"github.com/small-cat1/recharge-common/constant"
	"time"
)

type PaymentOrder struct {
	BaseModel
	OrderNo         string                      `gorm:"column:order_no;type:varchar(64);uniqueIndex:idx_order_no;not null;comment:订单号" json:"order_no"`
	UserID          uint                        `gorm:"column:user_id;type:bigint(20) unsigned;index:idx_user_id;not null;comment:用户ID" json:"user_id"`
	TenantID        uint                        `gorm:"column:tenant_id;type:bigint(20) unsigned;index:idx_tenant_id;not null;comment:租户ID" json:"tenant_id"`
	Amount          decimal.Decimal             `gorm:"column:amount;type:decimal(10,2);not null;comment:充值金额(RMB)" json:"amount"`
	UsDtAmount      decimal.Decimal             `gorm:"column:usdt_amount;type:decimal(20,6);not null;comment:USDT金额" json:"usdt_amount"`
	ExchangeRate    decimal.Decimal             `gorm:"column:exchange_rate;type:decimal(10,2);not null;comment:兑换汇率" json:"exchange_rate"`
	PayAddress      *string                     `gorm:"column:pay_address;type:varchar(128);not null;comment:TRC20付款地址" json:"pay_address"`
	Address         string                      `gorm:"column:address;type:varchar(128);not null;comment:TRC20收款地址" json:"address"`
	TransactionHash *string                     `gorm:"column:transaction_hash;type:varchar(128);comment:交易哈希" json:"transaction_hash"`
	Status          constant.PaymentOrderStatus `gorm:"column:status;type:tinyint(4);index:idx_status;default:1;not null;comment:订单状态: 1待支付 2已支付 3已取消 4已过期" json:"status"`
	PayAt           *time.Time                  `gorm:"column:pay_at;type:datetime;comment:支付时间" json:"pay_at"`
	ExpireAt        time.Time                   `gorm:"column:expire_at;type:datetime;index:idx_expire_at;not null;comment:过期时间" json:"expire_at"`

	//Use
	TenantUse TenantUser `json:"user,omitempty" gorm:"foreignKey:ID;references:UserID"`
	Tenant    Tenant     `json:"tenant,omitempty" gorm:"foreignKey:TenantID;references:TenantID"`
}

func (PaymentOrder) TableName() string {
	return "payment_orders"
}
