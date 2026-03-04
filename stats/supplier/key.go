package supplier

import (
	"fmt"
	"time"
)

// Redis key TTL：72小时（覆盖跨时区"昨日"查询场景）
const keyTTL = 72 * time.Hour

// Redis Hash 字段名
// FailOrders 不存储，查询时由 total - success 计算
const (
	fieldTotal         = "total"   // 总订单数
	fieldSuccess       = "success" // 成功订单数
	fieldSuccessAmount = "s_amt"   // 成功金额（分）
	fieldTotalAmount   = "t_amt"   // 总下单金额（分）
)

// hourStr 将 time.Time 转为 UTC 小时字符串，格式 "2026030408"
func hourStr(t time.Time) string {
	return t.UTC().Format("2006010215")
}

// hourInt 将 time.Time 转为 int64 格式的 UTC 小时，如 2026030408
func hourInt(t time.Time) int64 {
	var result int64
	fmt.Sscanf(hourStr(t), "%d", &result)
	return result
}

// ---- Key 生成函数（5个层级）----

// keyL1Tenant 层级1：租户汇总
// stats:sup:{tenantID}:{businessType}:t:hr:{hour}
func keyL1Tenant(tenantID uint, businessType, hour string) string {
	return fmt.Sprintf("stats:sup:%d:%s:t:hr:%s", tenantID, businessType, hour)
}

// keyL2Supplier 层级2：供货商
// stats:sup:{tenantID}:{businessType}:s:{supplierID}:hr:{hour}
func keyL2Supplier(tenantID uint, businessType string, supplierID uint, hour string) string {
	return fmt.Sprintf("stats:sup:%d:%s:s:%d:hr:%s", tenantID, businessType, supplierID, hour)
}

// keyL3Product 层级3：供货商 + 产品
// stats:sup:{tenantID}:{businessType}:s:{supplierID}:p:{productCode}:hr:{hour}
func keyL3Product(tenantID uint, businessType string, supplierID uint, productCode, hour string) string {
	return fmt.Sprintf("stats:sup:%d:%s:s:%d:p:%s:hr:%s", tenantID, businessType, supplierID, productCode, hour)
}

// keyL4Amount 层级4：供货商 + 产品 + 面额
// stats:sup:{tenantID}:{businessType}:s:{supplierID}:p:{productCode}:a:{amount}:hr:{hour}
func keyL4Amount(tenantID uint, businessType string, supplierID uint, productCode, amount, hour string) string {
	return fmt.Sprintf("stats:sup:%d:%s:s:%d:p:%s:a:%s:hr:%s", tenantID, businessType, supplierID, productCode, amount, hour)
}

// keyL5ProductOnly 层级5：仅产品维度（跨供货商聚合，不含 supplierID）
// stats:sup:{tenantID}:{businessType}:p:{productCode}:hr:{hour}
func keyL5ProductOnly(tenantID uint, businessType, productCode, hour string) string {
	return fmt.Sprintf("stats:sup:%d:%s:p:%s:hr:%s", tenantID, businessType, productCode, hour)
}

// allKeysForInput 根据写入参数生成全部 5 个层级的 key
func allKeysForInput(tenantID uint, businessType string, supplierID uint, productCode, amount, hour string) [5]string {
	return [5]string{
		keyL1Tenant(tenantID, businessType, hour),
		keyL2Supplier(tenantID, businessType, supplierID, hour),
		keyL3Product(tenantID, businessType, supplierID, productCode, hour),
		keyL4Amount(tenantID, businessType, supplierID, productCode, amount, hour),
		keyL5ProductOnly(tenantID, businessType, productCode, hour),
	}
}
