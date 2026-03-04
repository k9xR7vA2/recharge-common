package supplier

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RecordOrderResult 记录一笔供货商订单终态到 Redis 统计
//
// 调用时机：tenant_notify 查单到终态（成功或失败）后调用
// 一次 Pipeline 同步写入 5 个层级的 key，减少网络 RTT
func RecordOrderResult(ctx context.Context, rdb redis.UniversalClient, input RecordInput) error {
	if err := validateInput(input); err != nil {
		return fmt.Errorf("stats/writer: %w", err)
	}

	orderTime := time.Unix(input.OrderTime, 0)
	hour := hourStr(orderTime)

	amtFen, err := amountToFen(input.Amount)
	if err != nil {
		return fmt.Errorf("stats/writer: %w", err)
	}

	keys := allKeysForInput(
		input.TenantID,
		input.BusinessType,
		input.SupplierID,
		input.ProductCode,
		input.Amount,
		hour,
	)

	pipe := rdb.Pipeline()
	for _, k := range keys {
		pipe.HIncrBy(ctx, k, fieldTotal, 1)
		pipe.HIncrBy(ctx, k, fieldTotalAmount, amtFen)

		if input.IsSuccess {
			pipe.HIncrBy(ctx, k, fieldSuccess, 1)
			pipe.HIncrBy(ctx, k, fieldSuccessAmount, amtFen)
		}
		// 失败不写额外字段，查询时 fail = total - success

		pipe.Expire(ctx, k, keyTTL)
	}

	if _, err = pipe.Exec(ctx); err != nil {
		return fmt.Errorf("stats/writer: redis pipeline exec: %w", err)
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
