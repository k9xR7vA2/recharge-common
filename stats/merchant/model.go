package merchant

import (
	"github.com/k9xR7vA2/recharge-common/stats/base"
)

// ---- 写入侧 ----

// RecordInput 记录一笔商户订单到统计的输入参数
type RecordInput struct {
	TenantID     uint
	BusinessType string
	MerchantID   uint
	ChannelCode  string
	ChannelType  int    // 1=基础通道 2=混合通道
	Amount       string // "50"
	IsSuccess    bool
	OrderTime    int64 // Unix 秒
}

// ---- 读取侧查询参数 ----

// TodayQuery 今日统计基础查询参数
type TodayQuery struct {
	TenantID     uint
	BusinessType string
	Timezone     string
}

// MerchantQuery 指定商户查询
type MerchantQuery struct {
	TodayQuery
	MerchantID uint
}

// ChannelQuery 指定商户+通道查询（层级3）
type ChannelQuery struct {
	MerchantQuery
	ChannelCode string
}

// ChannelOnlyQuery 仅通道查询（层级5，跨商户）
type ChannelOnlyQuery struct {
	TodayQuery
	ChannelCode string
}

// HistoryQuery 历史统计查询
type HistoryQuery struct {
	TenantID     uint
	BusinessType string
	Timezone     string
	Date         string       // "2026-03-04"（用户本地日期）
	StatLevel    MerStatLevel // 1/2/3/4/5
	MerchantID   uint         // 层级2/3/4时必填
	ChannelCode  string       // 层级3/4/5时必填
	ChannelType  int          // 层级5时可选，用于过滤通道类型（0=不过滤）
	Amount       string       // 层级4时必填
}

// MerStatLevel 对外暴露层级常量（与 mongo/stats 包保持一致的值）
type MerStatLevel = int8

const (
	StatLevelTenant      MerStatLevel = 1
	StatLevelMerchant    MerStatLevel = 2
	StatLevelChannel     MerStatLevel = 3
	StatLevelAmount      MerStatLevel = 4
	StatLevelChannelOnly MerStatLevel = 5
)

// ---- 读取侧返回值 ----

// StatNumbers 统一统计数字结构
type StatNumbers struct {
	TotalOrders   int64
	SuccessOrders int64
	FailOrders    int64   // = TotalOrders - SuccessOrders
	SuccessAmount float64 // 元
	TotalAmount   float64 // 元
	SuccessRate   float64 // 百分比 0~100
}

// TenantTotalResult 层级1：租户今日汇总
type TenantTotalResult struct {
	StatNumbers
}

// MerchantStatResult 层级2：商户今日统计
type MerchantStatResult struct {
	MerchantID uint
	StatNumbers
}

// ChannelStatResult 层级3：商户+通道 / 层级5：仅通道 今日统计
type ChannelStatResult struct {
	MerchantID  uint // 层级5时为0
	ChannelCode string
	ChannelType int // 1=基础 2=混合
	StatNumbers
}

// AmountStatResult 层级4：商户+通道+面额 今日统计
type AmountStatResult struct {
	MerchantID  uint
	ChannelCode string
	ChannelType int
	Amount      string
	StatNumbers
}

// ---- 内部聚合 ----

type hashAgg struct {
	total         int64
	success       int64
	successAmount int64
	totalAmount   int64
	channelType   int // 通道类型，静态属性
}

func (a *hashAgg) merge(vals map[string]string) {
	a.total += base.ParseInt(vals[fieldTotal])
	a.success += base.ParseInt(vals[fieldSuccess])
	a.successAmount += base.ParseInt(vals[fieldSuccessAmount])
	a.totalAmount += base.ParseInt(vals[fieldTotalAmount])
	// channelType 取第一个非零值即可（同 key 下所有订单的 channel_type 相同）
	if a.channelType == 0 {
		a.channelType = int(base.ParseInt(vals[fieldChannelType]))
	}
}

func (a *hashAgg) toStatNumbers() StatNumbers {
	s := StatNumbers{
		TotalOrders:   a.total,
		SuccessOrders: a.success,
		FailOrders:    a.total - a.success,
		SuccessAmount: base.FenToYuan(a.successAmount),
		TotalAmount:   base.FenToYuan(a.totalAmount),
	}
	if a.total > 0 {
		s.SuccessRate = float64(a.success) / float64(a.total) * 100
	}
	return s
}
