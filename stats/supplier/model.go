package supplier

import (
	"fmt"
	"strconv"
)

// ---- 写入侧输入 ----

// RecordInput 记录一笔订单终态到统计的输入参数
type RecordInput struct {
	TenantID     uint
	BusinessType string // "mobile" / "india_mobile"
	SupplierID   uint
	ProductCode  string
	Amount       string // "50"（字符串面额，整数字符串）
	IsSuccess    bool
	OrderTime    int64 // 订单终态时间（Unix 秒），用于定位 UTC 小时
}

// ---- 读取侧查询参数 ----

// TodayQuery 今日统计查询基础参数
type TodayQuery struct {
	TenantID     uint
	BusinessType string
	Timezone     string // 如 "Asia/Shanghai"、"Asia/Kolkata"
}

// SupplierQuery 指定供货商查询
type SupplierQuery struct {
	TodayQuery
	SupplierID uint
}

// ProductQuery 指定供货商+产品查询（层级3）
type ProductQuery struct {
	SupplierQuery
	ProductCode string
}

// ProductOnlyQuery 仅产品查询（层级5，跨供货商）
type ProductOnlyQuery struct {
	TodayQuery
	ProductCode string
}

// HistoryQuery 历史统计查询（管理后台，按时区日期）
type HistoryQuery struct {
	TenantID     uint
	BusinessType string
	Timezone     string
	Date         string // "2026-03-04"（用户本地日期）
	StatLevel    int8   // 1/2/3/4/5
	SupplierID   uint   // 层级 2/3/4 时必填
	ProductCode  string // 层级 3/4/5 时必填
	Amount       string // 层级 4 时必填
}

// ---- 读取侧返回值 ----

// StatNumbers 统一统计数字结构
// FailOrders = TotalOrders - SuccessOrders，不单独存储，此处计算后对外暴露
type StatNumbers struct {
	TotalOrders   int64
	SuccessOrders int64
	FailOrders    int64   // = TotalOrders - SuccessOrders
	SuccessAmount float64 // 元（对外展示）
	TotalAmount   float64 // 元
	SuccessRate   float64 // 百分比，0~100
}

// TenantTotalResult 层级1：租户今日汇总
type TenantTotalResult struct {
	StatNumbers
}

// SupplierStatResult 层级2：供货商今日统计
type SupplierStatResult struct {
	SupplierID uint
	StatNumbers
}

// ProductStatResult 层级3：供货商+产品 / 层级5：仅产品 今日统计
type ProductStatResult struct {
	SupplierID  uint // 层级5时为 0
	ProductCode string
	StatNumbers
}

// AmountStatResult 层级4：供货商+产品+面额 今日统计
type AmountStatResult struct {
	SupplierID  uint
	ProductCode string
	Amount      string
	StatNumbers
}

// ---- 内部聚合工具 ----

// hashAgg 累加多个 UTC 小时的 Redis Hash 结果
type hashAgg struct {
	total         int64
	success       int64
	successAmount int64 // 分
	totalAmount   int64 // 分
}

func (a *hashAgg) merge(vals map[string]string) {
	a.total += parseInt(vals[fieldTotal])
	a.success += parseInt(vals[fieldSuccess])
	a.successAmount += parseInt(vals[fieldSuccessAmount])
	a.totalAmount += parseInt(vals[fieldTotalAmount])
}

func (a *hashAgg) toStatNumbers() StatNumbers {
	s := StatNumbers{
		TotalOrders:   a.total,
		SuccessOrders: a.success,
		FailOrders:    a.total - a.success,
		SuccessAmount: fenToYuan(a.successAmount),
		TotalAmount:   fenToYuan(a.totalAmount),
	}
	if a.total > 0 {
		s.SuccessRate = float64(a.success) / float64(a.total) * 100
	}
	return s
}

// ---- 金额工具 ----

// amountToFen 将面额字符串（"50"）转为分（5000）
func amountToFen(amount string) (int64, error) {
	v, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("amountToFen: invalid amount %q: %w", amount, err)
	}
	return v * 100, nil
}

// fenToYuan 分转元，保留2位精度
func fenToYuan(fen int64) float64 {
	return float64(fen) / 100.0
}

func parseInt(s string) int64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}
