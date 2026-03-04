package merchant

import (
	"context"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/stats/base"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// RecordOrderCreated 下单成功时调用（tenant_api）
// 只累加 total_orders 和 total_amount
// 层级3/4/5 同时写入 channel_type（静态属性，用 HSET 而非 HINCRBY）
func RecordOrderCreated(ctx context.Context, rdb redis.UniversalClient, input RecordInput) error {
	if err := validateInput(input); err != nil {
		return fmt.Errorf("stats/merchant.RecordOrderCreated: %w", err)
	}

	hour := hourStr(time.Unix(input.OrderTime, 0))
	amtFen, err := base.AmountToFen(input.Amount)
	if err != nil {
		return fmt.Errorf("stats/merchant.RecordOrderCreated: %w", err)
	}

	keys := allKeysForInput(input.TenantID, input.BusinessType, input.MerchantID, input.ChannelCode, input.Amount, hour)
	chTypeStr := strconv.Itoa(input.ChannelType)

	pipe := rdb.Pipeline()
	for i, k := range keys {
		pipe.HIncrBy(ctx, k, fieldTotal, 1)
		pipe.HIncrBy(ctx, k, fieldTotalAmount, amtFen)
		// 层级3(index=2)、层级4(index=3)、层级5(index=4) 写入通道类型
		if i >= 2 {
			pipe.HSet(ctx, k, fieldChannelType, chTypeStr)
		}
		pipe.Expire(ctx, k, keyTTL)
	}
	if _, err = pipe.Exec(ctx); err != nil {
		return fmt.Errorf("stats/merchant.RecordOrderCreated: pipeline exec: %w", err)
	}
	return nil
}

// RecordOrderResult 订单到终态时调用（tenant_notify）
// 成功：累加 success_orders 和 success_amount
// 失败：不写（fail = total - success）
func RecordOrderResult(ctx context.Context, rdb redis.UniversalClient, input RecordInput) error {
	if !input.IsSuccess {
		return nil
	}
	if err := validateInput(input); err != nil {
		return fmt.Errorf("stats/merchant.RecordOrderResult: %w", err)
	}

	hour := hourStr(time.Unix(input.OrderTime, 0))
	amtFen, err := base.AmountToFen(input.Amount)
	if err != nil {
		return fmt.Errorf("stats/merchant.RecordOrderResult: %w", err)
	}

	keys := allKeysForInput(input.TenantID, input.BusinessType, input.MerchantID, input.ChannelCode, input.Amount, hour)

	pipe := rdb.Pipeline()
	for _, k := range keys {
		pipe.HIncrBy(ctx, k, fieldSuccess, 1)
		pipe.HIncrBy(ctx, k, fieldSuccessAmount, amtFen)
		// TTL 不刷新，下单时已设置
	}
	if _, err = pipe.Exec(ctx); err != nil {
		return fmt.Errorf("stats/merchant.RecordOrderResult: pipeline exec: %w", err)
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
	if input.MerchantID == 0 {
		return fmt.Errorf("merchant_id is required")
	}
	if input.ChannelCode == "" {
		return fmt.Errorf("channel_code is required")
	}
	if input.Amount == "" {
		return fmt.Errorf("amount is required")
	}
	if input.OrderTime <= 0 {
		return fmt.Errorf("order_time is required")
	}
	return nil
}
