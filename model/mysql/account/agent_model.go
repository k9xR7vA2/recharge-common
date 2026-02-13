package account

import (
	"SaasAdmin/internal/common/constant"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// AsAgent [...]
type AsAgent struct {
	AgentID          uint                          `gorm:"autoIncrement:true;primaryKey;column:agent_id;type:int unsigned;not null" json:"agent_id"`
	AgentName        string                        `gorm:"column:agent_name;unique;type:varchar(255);not null;comment:'代理名称'" json:"agent_name"`        // 代理名称
	AgentPasswd      string                        `gorm:"column:agent_passwd;type:varchar(255);not null;comment:'代理密码'" json:"-"`                      // 代理密码
	Status           constant.TenantBusinessStatus `gorm:"column:status;type:tinyint;not null;comment:'状态1开启，2关闭'" json:"status"`                       // 状态1开启，2关闭
	Rebate           decimal.Decimal               `gorm:"column:rebate;type:decimal(10,2);not null;comment:'代理佣金汇率'" json:"rebate"`                    // 代理佣金汇率
	Balance          decimal.Decimal               `gorm:"column:balance;type:decimal(10,2);not null;default:0.00;comment:'代理佣金'" json:"balance"`       // 代理佣金
	WithdrawalAmount decimal.Decimal               `gorm:"column:withdrawal_amount;type:decimal(10,2);not null;comment:'总提取'" json:"withdrawal_amount"` // 总提取
	FrozenAmount     decimal.Decimal               `gorm:"column:frozen_amount;type:decimal(10,2);not null;comment:'总冻结'" json:"frozen_amount"`         // 总冻结
	Remark           string                        `gorm:"column:remark;type:text;default:null;comment:'备注'" json:"remark"`                             // 备注
	CreatedAt        time.Time                     `json:"created_at"`                                                                                  // 创建时间
	UpdatedAt        time.Time                     `json:"updated_at"`                                                                                  // 更新时间
	DeletedAt        gorm.DeletedAt                `gorm:"index" json:"-"`                                                                              // 删除时间
}

// TableName get sql table name.获取数据库表名
func (*AsAgent) TableName() string {
	return "as_agents"
}
