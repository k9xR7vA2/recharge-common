package stats

import "go.mongodb.org/mongo-driver/bson/primitive"

const CollMerchantOrderHourlyStats = "merchant_order_hourly_stats"

// MerStatLevel 商户统计维度层级
type MerStatLevel int8

const (
	MerStatLevelTenant      MerStatLevel = 1 // 层级1：租户汇总（所有商户）
	MerStatLevelMerchant    MerStatLevel = 2 // 层级2：商户
	MerStatLevelChannel     MerStatLevel = 3 // 层级3：商户 + 通道
	MerStatLevelAmount      MerStatLevel = 4 // 层级4：商户 + 通道 + 面额
	MerStatLevelChannelOnly MerStatLevel = 5 // 层级5：仅通道（跨商户，用于通道成功率面板）
)

// MerchantOrderHourlyStat 商户订单小时统计文档
// 存放在租户独立库 tenant_{tenantID}.merchant_order_hourly_stats
//
// Upsert 唯一键：
//
//	{ tenant_id, business_type, stat_level, merchant_id, channel_code, amount, hour_utc }
//
// 失败订单数 = total_orders - success_orders，不单独存储
// channel_type 仅层级3/4/5有效，用于前端区分普通通道和混合通道过滤
type MerchantOrderHourlyStat struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"  json:"id"`
	TenantID     uint               `bson:"tenant_id"      json:"tenant_id"`
	BusinessType string             `bson:"business_type"  json:"business_type"`

	// 维度字段
	// 层级1: merchant_id=0, channel_code="", amount=""
	// 层级2: channel_code="", amount=""
	// 层级3: amount=""
	// 层级4: 全部有值
	// 层级5: merchant_id=0, amount=""（仅 channel_code 有值）
	StatLevel   MerStatLevel `bson:"stat_level"    json:"stat_level"`
	MerchantID  uint         `bson:"merchant_id"   json:"merchant_id"`
	ChannelCode string       `bson:"channel_code"  json:"channel_code"`
	ChannelType int          `bson:"channel_type"  json:"channel_type"` // 1=基础通道 2=混合通道，层级1/2时为0
	Amount      string       `bson:"amount"        json:"amount"`       // "50"/"100"，无则为""

	// 时间维度
	HourUTC int64  `bson:"hour_utc"  json:"hour_utc"` // 2026030408
	DateUTC string `bson:"date_utc"  json:"date_utc"` // "2026-03-04"

	// 统计数据（金额单位：分）
	TotalOrders   int64 `bson:"total_orders"    json:"total_orders"`
	SuccessOrders int64 `bson:"success_orders"  json:"success_orders"`
	SuccessAmount int64 `bson:"success_amount"  json:"success_amount"`
	TotalAmount   int64 `bson:"total_amount"    json:"total_amount"`

	FlushAt   int64 `bson:"flush_at"    json:"flush_at"`
	CreatedAt int64 `bson:"created_at"  json:"created_at"`
	UpdatedAt int64 `bson:"updated_at"  json:"updated_at"`
}

// FailOrders 失败订单数（计算属性）
func (s MerchantOrderHourlyStat) FailOrders() int64 {
	return s.TotalOrders - s.SuccessOrders
}

// 建议索引：
//
//	{ tenant_id, business_type, stat_level, merchant_id, channel_code, amount, hour_utc } unique
//	{ tenant_id, stat_level, business_type, hour_utc }
//	{ tenant_id, merchant_id, stat_level, hour_utc }
//	{ tenant_id, merchant_id, channel_code, hour_utc }
//	{ tenant_id, channel_code, stat_level, hour_utc }         // 层级5通道面板
//	{ tenant_id, channel_type, stat_level, hour_utc }         // 按通道类型过滤
//	{ tenant_id, date_utc, stat_level }
