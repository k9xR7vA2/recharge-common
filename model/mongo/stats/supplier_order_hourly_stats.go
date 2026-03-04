package stats

import "go.mongodb.org/mongo-driver/bson/primitive"

// CollSupplierOrderHourlyStats 集合名
const CollSupplierOrderHourlyStats = "supplier_order_hourly_stats"

// StatLevel 统计维度层级
type StatLevel int8

const (
	StatLevelTenant      StatLevel = 1 // 层级1：租户汇总
	StatLevelSupplier    StatLevel = 2 // 层级2：供货商
	StatLevelProduct     StatLevel = 3 // 层级3：供货商 + 产品
	StatLevelAmount      StatLevel = 4 // 层级4：供货商 + 产品 + 面额
	StatLevelProductOnly StatLevel = 5 // 层级5：仅产品（跨供货商，supplier_id=0）
)

// SupplierOrderHourlyStat 供货商订单小时统计文档
// 存放在租户独立库 tenant_{tenantID}.supplier_order_hourly_stats
//
// Upsert 唯一键：
//
//	{ tenant_id, business_type, stat_level, supplier_id, product_code, amount, hour_utc }
//
// 失败订单数 = total_orders - success_orders，不单独存储
// 若后期需要区分失败原因，可在此处扩展 fail_reason_stats 字段
type SupplierOrderHourlyStat struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"  json:"id"`
	TenantID     uint               `bson:"tenant_id"      json:"tenant_id"`
	BusinessType string             `bson:"business_type"  json:"business_type"` // "mobile" / "india_mobile"

	// 维度字段
	// 层级1: supplier_id=0, product_code="", amount=""
	// 层级2: product_code="", amount=""
	// 层级3: amount=""
	// 层级4: 全部有值
	// 层级5: supplier_id=0, amount=""（仅 product_code 有值，跨供货商）
	StatLevel   StatLevel `bson:"stat_level"    json:"stat_level"`
	SupplierID  uint      `bson:"supplier_id"   json:"supplier_id"`
	ProductCode string    `bson:"product_code"  json:"product_code"`
	Amount      string    `bson:"amount"        json:"amount"` // "50"、"100" 等，无则为 ""

	// 时间维度
	HourUTC int64  `bson:"hour_utc"  json:"hour_utc"` // 2026030408
	DateUTC string `bson:"date_utc"  json:"date_utc"` // "2026-03-04"

	// 统计数据（金额单位：分）
	// FailOrders = TotalOrders - SuccessOrders，不单独存储
	TotalOrders   int64 `bson:"total_orders"    json:"total_orders"`
	SuccessOrders int64 `bson:"success_orders"  json:"success_orders"`
	SuccessAmount int64 `bson:"success_amount"  json:"success_amount"` // 分
	TotalAmount   int64 `bson:"total_amount"    json:"total_amount"`   // 分

	FlushAt   int64 `bson:"flush_at"    json:"flush_at"`
	CreatedAt int64 `bson:"created_at"  json:"created_at"`
	UpdatedAt int64 `bson:"updated_at"  json:"updated_at"`
}

// FailOrders 失败订单数（计算属性，不存 DB）
func (s SupplierOrderHourlyStat) FailOrders() int64 {
	return s.TotalOrders - s.SuccessOrders
}

// IndexModels 建议的索引
//
//	// 唯一索引（upsert 依据）
//	{ tenant_id, business_type, stat_level, supplier_id, product_code, amount, hour_utc } unique
//	// 层级1：租户汇总按天
//	{ tenant_id, stat_level, business_type, hour_utc }
//	// 层级2：供货商维度
//	{ tenant_id, supplier_id, stat_level, hour_utc }
//	// 层级3/4：供货商+产品维度
//	{ tenant_id, supplier_id, product_code, hour_utc }
//	// 层级5：仅产品维度
//	{ tenant_id, product_code, stat_level, hour_utc }
//	// 按日聚合
//	{ tenant_id, date_utc, stat_level }
