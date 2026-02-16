package mysql

import (
	"database/sql"
	"time"
)

// TenantUsdtRechargeOrder 租户USDT充值订单结构体
type TenantUsdtRechargeOrder struct {
	ID       string `json:"id" gorm:"type:char(32);primary_key;comment:'订单唯一ID (UUID)'"`
	TenantID string `json:"tenant_id" gorm:"type:char(32);not null;index;comment:'租户ID'"`
	OrderNo  string `json:"order_no" gorm:"type:varchar(64);not null;index;comment:'订单号 (业务可见编号)'"`
	UserID   string `json:"user_id" gorm:"type:char(32);not null;index;comment:'用户ID (发起充值的用户)'"`

	// 金额信息
	Amount       float64 `json:"amount" gorm:"type:decimal(20,8);not null;comment:'USDT充值金额'"`
	ExchangeRate float64 `json:"exchange_rate" gorm:"type:decimal(20,8);comment:'充值时的汇率 (可选,如需转换为法币)'"`
	FiatAmount   float64 `json:"fiat_amount" gorm:"type:decimal(20,2);comment:'折算后的法币金额 (可选)'"`
	FiatCurrency string  `json:"fiat_currency" gorm:"type:varchar(10);comment:'法币币种 (USD, CNY等, 可选)'"`
	Fee          float64 `json:"fee" gorm:"type:decimal(20,8);default:0;comment:'手续费'"`

	// 交易信息
	BlockchainNetwork   string `json:"blockchain_network" gorm:"type:varchar(20);not null;comment:'区块链网络 (例如: TRON, ETH, BSC)'"`
	UsdtAddress         string `json:"usdt_address" gorm:"type:varchar(100);not null;index;comment:'充值目标地址'"`
	FromAddress         string `json:"from_address" gorm:"type:varchar(100);comment:'发送方地址 (可选)'"`
	TxHash              string `json:"tx_hash" gorm:"type:varchar(128);index;comment:'区块链交易哈希'"`
	TxConfirmationCount int    `json:"tx_confirmation_count" gorm:"type:int;default:0;comment:'交易确认数'"`
	BlockHeight         int64  `json:"block_height" gorm:"type:bigint;comment:'区块高度'"`

	// 状态信息
	Status     int    `json:"status" gorm:"type:tinyint;not null;default:0;index;comment:'订单状态: 0-待支付 1-支付中 2-已支付 3-已确认 4-失败 5-已取消'"`
	StatusDesc string `json:"status_desc" gorm:"type:varchar(255);comment:'状态描述 (特别是错误情况)'"`

	// 时间信息
	CreatedAt   time.Time    `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;index;comment:'订单创建时间'"`
	PaidAt      sql.NullTime `json:"paid_at" gorm:"type:timestamp;null;comment:'支付时间'"`
	ConfirmedAt sql.NullTime `json:"confirmed_at" gorm:"type:timestamp;null;comment:'确认时间'"`
	ExpiredAt   sql.NullTime `json:"expired_at" gorm:"type:timestamp;null;comment:'过期时间'"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'最后更新时间'"`

	// 业务信息
	OrderType  int    `json:"order_type" gorm:"type:tinyint;default:0;comment:'订单类型: 0-普通充值 1-套餐购买 2-服务续费等'"`
	OrderTitle string `json:"order_title" gorm:"type:varchar(128);comment:'订单标题'"`
	OrderDesc  string `json:"order_desc" gorm:"type:text;comment:'订单描述'"`
	AttachData string `json:"attach_data" gorm:"type:text;comment:'附加数据 (JSON格式存储业务相关信息)'"`

	// 审计信息
	OperatorID string `json:"operator_id" gorm:"type:varchar(32);comment:'操作人ID (后台操作时)'"`
	Remark     string `json:"remark" gorm:"type:varchar(255);comment:'备注'"`
	IPAddress  string `json:"ip_address" gorm:"type:varchar(64);comment:'用户IP地址'"`
	DeviceInfo string `json:"device_info" gorm:"type:varchar(255);comment:'设备信息'"`
}

// TableName 指定表名
func (TenantUsdtRechargeOrder) TableName() string {
	return "tenant_usdt_recharge_orders"
}

// 订单状态常量
const (
	OrderStatusPending    int = 0 // 待支付
	OrderStatusProcessing int = 1 // 支付中
	OrderStatusPaid       int = 2 // 已支付
	OrderStatusConfirmed  int = 3 // 已确认
	OrderStatusFailed     int = 4 // 失败
	OrderStatusCancelled  int = 5 // 已取消
)

// 订单类型常量
const (
	OrderTypeNormal  int = 0 // 普通充值
	OrderTypePackage int = 1 // 套餐购买
	OrderTypeRenewal int = 2 // 服务续费
)
