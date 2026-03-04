package supplier

import (
	"context"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/stats/base"
	"time"

	"github.com/redis/go-redis/v9"
)

// RecordInput 写入统计的通用参数
type RecordInput struct {
	TenantID     uint
	BusinessType string // "mobile" / "india_mobile"
	SupplierID   uint
	ProductCode  string
	Amount       string // "50"
	IsSuccess    bool   // RecordOrderResult 时有效，RecordOrderCreated 时忽略
	OrderTime    int64  // Unix 秒，用于定位 UTC 小时
}

// RecordOrderCreated 下单成功时调用（tenant_api）
// 只累加 total_orders 和 total_amount，不影响 success 字段
// 统计失败不阻断主流程，调用方用 _ = 忽略错误即可
func RecordOrderCreated(ctx context.Context, rdb redis.UniversalClient, input RecordInput) error {
	if err := validateInput(input); err != nil {
		return fmt.Errorf("stats/writer.RecordOrderCreated: %w", err)
	}

	hour := hourStr(time.Unix(input.OrderTime, 0))
	amtFen, err := base.AmountToFen(input.Amount)
	if err != nil {
		return fmt.Errorf("stats/writer.RecordOrderCreated: %w", err)
	}

	keys := allKeysForInput(input.TenantID, input.BusinessType, input.SupplierID, input.ProductCode, input.Amount, hour)

	pipe := rdb.Pipeline()
	for _, k := range keys {
		pipe.HIncrBy(ctx, k, fieldTotal, 1)
		pipe.HIncrBy(ctx, k, fieldTotalAmount, amtFen)
		pipe.Expire(ctx, k, keyTTL)
	}
	if _, err = pipe.Exec(ctx); err != nil {
		return fmt.Errorf("stats/writer.RecordOrderCreated: redis pipeline exec: %w", err)
	}
	return nil
}

// RecordOrderResult 订单到达终态时调用（tenant_notify）
// 成功：累加 success_orders 和 success_amount
// 失败：不写任何字段（fail = total - success 由查询时计算）
func RecordOrderResult(ctx context.Context, rdb redis.UniversalClient, input RecordInput) error {
	if !input.IsSuccess {
		// 失败不需要写 Redis，total 已在下单时计入
		return nil
	}
	if err := validateInput(input); err != nil {
		return fmt.Errorf("stats/writer.RecordOrderResult: %w", err)
	}

	hour := hourStr(time.Unix(input.OrderTime, 0))
	amtFen, err := base.AmountToFen(input.Amount)
	if err != nil {
		return fmt.Errorf("stats/writer.RecordOrderResult: %w", err)
	}

	keys := allKeysForInput(input.TenantID, input.BusinessType, input.SupplierID, input.ProductCode, input.Amount, hour)

	pipe := rdb.Pipeline()
	for _, k := range keys {
		pipe.HIncrBy(ctx, k, fieldSuccess, 1)
		pipe.HIncrBy(ctx, k, fieldSuccessAmount, amtFen)
		// 不刷新 TTL：key 在下单时已经设置过，这里只更新数值
	}
	if _, err = pipe.Exec(ctx); err != nil {
		return fmt.Errorf("stats/writer.RecordOrderResult: redis pipeline exec: %w", err)
	}
	return nil
}

func validateInput(input RecordInput) error {
	if input.TenantID == 0 {
		return fmt.Errorf("tenant_id is required")
	}
	if input.BusinessType == "" {
		return fmt.Errorf("business_type is required")
	}
	if input.SupplierID == 0 {
		return fmt.Errorf("supplier_id is required")
	}
	if input.ProductCode == "" {
		return fmt.Errorf("product_code is required")
	}
	if input.Amount == "" {
		return fmt.Errorf("amount is required")
	}
	if input.OrderTime <= 0 {
		return fmt.Errorf("order_time is required")
	}
	return nil
}
