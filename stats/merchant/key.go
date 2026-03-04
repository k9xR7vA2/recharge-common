package merchant

import (
	"fmt"
	"time"
)

const keyTTL = 72 * time.Hour

// Redis Hash 字段名
// FailOrders 不存储，查询时由 total - success 计算
const (
	fieldTotal         = "total"   // 总订单数
	fieldSuccess       = "success" // 成功订单数
	fieldSuccessAmount = "s_amt"   // 成功金额（分）
	fieldTotalAmount   = "t_amt"   // 总下单金额（分）
	fieldChannelType   = "ch_type" // 通道类型（静态属性，层级3/4/5写入，用于前端过滤）
)

func hourStr(t time.Time) string {
	return t.UTC().Format("2006010215")
}

func hourInt(t time.Time) int64 {
	var result int64
	fmt.Sscanf(hourStr(t), "%d", &result)
	return result
}

// keyL1Tenant 层级1：租户汇总
// stats:mer:{tenantID}:{bizType}:t:hr:{hour}
func keyL1Tenant(tenantID uint, bizType, hour string) string {
	return fmt.Sprintf("stats:mer:%d:%s:t:hr:%s", tenantID, bizType, hour)
}

// keyL2Merchant 层级2：商户
// stats:mer:{tenantID}:{bizType}:m:{merchantID}:hr:{hour}
func keyL2Merchant(tenantID uint, bizType string, merchantID uint, hour string) string {
	return fmt.Sprintf("stats:mer:%d:%s:m:%d:hr:%s", tenantID, bizType, merchantID, hour)
}

// keyL3Channel 层级3：商户 + 通道
// stats:mer:{tenantID}:{bizType}:m:{merchantID}:c:{channelCode}:hr:{hour}
func keyL3Channel(tenantID uint, bizType string, merchantID uint, channelCode, hour string) string {
	return fmt.Sprintf("stats:mer:%d:%s:m:%d:c:%s:hr:%s", tenantID, bizType, merchantID, channelCode, hour)
}

// keyL4Amount 层级4：商户 + 通道 + 面额
// stats:mer:{tenantID}:{bizType}:m:{merchantID}:c:{channelCode}:a:{amount}:hr:{hour}
func keyL4Amount(tenantID uint, bizType string, merchantID uint, channelCode, amount, hour string) string {
	return fmt.Sprintf("stats:mer:%d:%s:m:%d:c:%s:a:%s:hr:%s", tenantID, bizType, merchantID, channelCode, amount, hour)
}

// keyL5ChannelOnly 层级5：仅通道（跨商户，用于通道成功率排行）
// stats:mer:{tenantID}:{bizType}:c:{channelCode}:hr:{hour}
func keyL5ChannelOnly(tenantID uint, bizType, channelCode, hour string) string {
	return fmt.Sprintf("stats:mer:%d:%s:c:%s:hr:%s", tenantID, bizType, channelCode, hour)
}

// allKeysForInput 生成全部 5 个层级的 key
func allKeysForInput(tenantID uint, bizType string, merchantID uint, channelCode, amount, hour string) [5]string {
	return [5]string{
		keyL1Tenant(tenantID, bizType, hour),
		keyL2Merchant(tenantID, bizType, merchantID, hour),
		keyL3Channel(tenantID, bizType, merchantID, channelCode, hour),
		keyL4Amount(tenantID, bizType, merchantID, channelCode, amount, hour),
		keyL5ChannelOnly(tenantID, bizType, channelCode, hour),
	}
}
