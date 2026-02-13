package mysql

import (
	"gorm.io/gorm"
	"time"
)

// TRC20Transaction 交易记录表模型
type TRC20Transaction struct {
	ID              uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	TxID            string         `json:"tx_id" gorm:"column:tx_id;type:varchar(100);not null;uniqueIndex;comment:交易ID"`
	BlockNumber     int64          `json:"block_number" gorm:"column:block_number;comment:区块高度"`
	Timestamp       int64          `json:"timestamp" gorm:"column:timestamp;not null;index;comment:交易时间戳"`
	FromAddress     string         `json:"from_address" gorm:"column:from_address;type:varchar(100);not null;index;comment:发送地址"`
	ToAddress       string         `json:"to_address" gorm:"column:to_address;type:varchar(100);not null;index;comment:接收地址"`
	ContractAddress string         `json:"contract_address" gorm:"column:contract_address;type:varchar(100);not null;comment:合约地址"`
	Value           string         `json:"value" gorm:"column:value;type:varchar(100);not null;comment:原始金额"`
	FormattedValue  string         `json:"formatted_value" gorm:"column:formatted_value;type:varchar(100);not null;comment:格式化后的金额"`
	TokenName       string         `json:"token_name" gorm:"column:token_name;type:varchar(100);comment:代币名称"`
	TokenSymbol     string         `json:"token_symbol" gorm:"column:token_symbol;type:varchar(50);comment:代币符号"`
	TokenDecimal    int            `json:"token_decimal" gorm:"column:token_decimal;comment:代币精度"`
	FormattedTime   string         `json:"formatted_time" gorm:"column:formatted_time;type:varchar(50);comment:格式化的时间"`
	Status          int            `json:"status" gorm:"column:status;type:tinyint;default:0;index;comment:处理状态: 0-待处理, 1-处理中, 2-已确认, 3-处理失败, 4-已拒绝"`
	Confirmations   int64          `json:"confirmations" gorm:"column:confirmations;default:0;comment:确认数"`
	ProcessTime     *time.Time     `json:"process_time" gorm:"column:process_time;comment:处理时间"`
	FailReason      string         `json:"fail_reason" gorm:"column:fail_reason;type:varchar(255);comment:失败原因"`
	CreatedAt       time.Time      `json:"created_at" gorm:"column:created_at;autoCreateTime;comment:创建时间"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;comment:更新时间"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 设置表名
func (TRC20Transaction) TableName() string {
	return "trc20_transactions"
}
