package merchant

// AsMerchantOrder AsMerchantOrders
//type AsMerchantOrder struct {
//	ID                   uint64          `bson:"_id,omitempty" json:"id"`
//	MerchantID           uint            `bson:"merchant_id" json:"merchant_id"`                       // 商户ID
//	MerchantName         uint            `bson:"merchant_name" json:"merchant_name"`                   // 商户名称
//	TenantID             uint            `bson:"tenant_id" json:"tenant_id"`                           // 租户ID
//	TenantName           string          `bson:"tenant_name" json:"tenant_name"`                       // 租户名称
//	TenantChannelID      uint            `bson:"tenant_channel_id" json:"tenant_channel_id"`           // 租户通道ID
//	ChannelCode          string          `bson:"channel_code" json:"channel_code"`                     // 租户通道编码
//	SystemOrderSn        string          `bson:"system_order_sn" json:"system_order_sn"`               // 系统订单号
//	MerchantOrderSn      string          `bson:"merchant_order_sn" json:"merchant_order_sn"`           // 商户订单号
//	OfficialSerialNumber string          `bson:"official_serial_number" json:"official_serial_number"` // 官方流水号
//	ClientIP             string          `bson:"client_ip" json:"client_ip"`                           // 客户ip
//	Device               int             `bson:"device" json:"device"`                                 // 设备,1ios,2Android,3双端
//	Payment              int             `bson:"payment" json:"payment"`                               // 支付方式 1支付宝2微信
//	Amount               decimal.Decimal `bson:"amount" json:"amount"`                                 // 金额
//	HandingFees          decimal.Decimal `bson:"handing_fees" json:"handing_fees"`                     // 手续费
//	PayURL               string          `bson:"pay_url" json:"pay_url"`                               // 支付地址
//	QueryURL             string          `bson:"query_url" json:"query_url"`                           // 查单地址
//	OrderStatus          int             `bson:"order_status" json:"order_status"`                     // 订单状态(1等待，2代支付，3成功，4失败)
//	MakeCodeNumber       uint            `bson:"make_code_number" json:"make_code_number"`             // 产码次数
//	MatchAt              time.Time       `bson:"match_at" json:"match_at"`                             // 配单时间
//	MakeCodeAt           time.Time       `bson:"make_code_at" json:"make_code_at"`                     // 产码成功时间
//	PullAt               time.Time       `bson:"pull_at" json:"pull_at"`                               // 客户端获取链接时间
//	SuccessAt            time.Time       `bson:"success_at" json:"success_at"`                         // 到账时间
//	NotifyAt             time.Time       `bson:"notify_at" json:"notify_at"`                           // 回调时间
//	NotifyStatus         int             `bson:"notify_status" json:"notify_status"`
//	NotifyURL            string          `bson:"notify_url" json:"notify_url"`             // 回调url
//	NotifyNumbers        uint            `bson:"notify_numbers" json:"notify_numbers"`     // 回调次数
//	CreatedAt            time.Time       `bson:"created_at" json:"created_at"`             // 创建时间
//	UpdatedAt            time.Time       `bson:"updated_at" json:"updated_at"`             // 更新时间
//	IsReplenishment      int             `bson:"is_replenishment" json:"is_replenishment"` // 是否手工补单,1不是，2是
//}
//
//// TableName get sql table name.获取数据库表名
//func (AsMerchantOrder) TableName() string {
//	return "as_merchant_orders"
//}
