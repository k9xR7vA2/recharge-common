package orderpool

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/small-cat1/recharge-common/orderpool/keys"
	"sort"
	"strconv"
	"strings"
)

type PoolMonitor struct {
	RedisClient  redis.UniversalClient
	KeyGenerator *keys.RedisKeysGenerate
}

// NewPoolMonitor 创建订单池监控器
func NewPoolMonitor(client redis.UniversalClient, tenantID uint, role, businessType string) *PoolMonitor {
	return &PoolMonitor{
		RedisClient:  client,
		KeyGenerator: keys.NewRedisKeysGenerate(tenantID, role, businessType),
	}
}

// GetPoolStats 获取订单池统计信息
func (m *PoolMonitor) GetPoolStats(ctx context.Context) (map[string]interface{}, error) {
	statsKey := m.KeyGenerator.StatsKey()
	// 获取统计数据
	stats, err := m.RedisClient.HGetAll(ctx, statsKey).Result()
	if err != nil {
		return nil, err
	}
	// 格式化结果
	result := make(map[string]interface{})
	// 基础计数
	for _, key := range []string{"total_orders", "pool_orders", "processing_orders"} {
		if val, ok := stats[key]; ok {
			count, _ := strconv.ParseInt(val, 10, 64)
			result[key] = count
		} else {
			result[key] = 0
		}
	}
	// 维度计数
	dimensions := map[string]map[string]int64{
		"amount":      {},
		"carrier":     {},
		"charge_type": {},
		"area":        {},
		"priority":    {},
	}
	for key, val := range stats {
		for dimension := range dimensions {
			prefix := dimension + ":"
			if strings.HasPrefix(key, prefix) {
				value := strings.TrimPrefix(key, prefix)
				count, _ := strconv.ParseInt(val, 10, 64)
				dimensions[dimension][value] = count
			}
		}
	}
	for dimension, counts := range dimensions {
		result[dimension] = counts
	}
	return result, nil
}

// GetPoolDetails 获取订单池详细信息（续）
//func (m *PoolMonitor) GetPoolDetails(ctx context.Context, priority string, poolArgs *PoolArgs, limit int) ([]map[string]interface{}, error) {
//	// 获取订单池键
//	poolKey := m.KeyGenerator.PoolKey(priority, poolArgs)
//
//	// 查询Stream中的订单
//	entries, err := m.RedisClient.XRevRange(ctx, poolKey, "+", "-", int64(limit)).Result()
//	if err != nil {
//		return nil, err
//	}
//
//	// 格式化结果
//	result := make([]map[string]interface{}, 0, len(entries))
//
//	for _, entry := range entries {
//		orderSN, _ := entry.Values["order_sn"].(string)
//		expireAt, _ := entry.Values["expire_at"].(string)
//		retryCount, _ := entry.Values["retry_count"].(string)
//		createTime, _ := entry.Values["create_time"].(string)
//
//		// 获取订单详情
//		orderInfo, err := m.GetOrderInfo(ctx, orderSN)
//		if err != nil {
//			continue
//		}
//
//		item := map[string]interface{}{
//			"order_sn":    orderSN,
//			"message_id":  entry.ID,
//			"expire_at":   expireAt,
//			"retry_count": retryCount,
//			"create_time": createTime,
//			"status":      orderInfo["status"],
//			"age_seconds": time.Now().Unix() - parseUnixTime(createTime),
//			"time_left":   parseUnixTime(expireAt) - time.Now().Unix(),
//		}
//
//		// 可能需要添加额外的订单信息
//		var fullOrderInfo map[string]interface{}
//		if infoStr, ok := orderInfo["info"].(string); ok {
//			if err := json.Unmarshal([]byte(infoStr), &fullOrderInfo); err == nil {
//				item["info"] = fullOrderInfo
//			}
//		}
//
//		result = append(result, item)
//	}
//
//	return result, nil
//}

// GetOrderInfo 获取订单信息
func (m *PoolMonitor) GetOrderInfo(ctx context.Context, orderSN string) (map[string]interface{}, error) {
	orderKey := m.KeyGenerator.OrderKey(orderSN)
	result, err := m.RedisClient.HGetAll(ctx, orderKey).Result()
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("订单信息不存在")
	}

	orderInfo := make(map[string]interface{})
	for k, v := range result {
		orderInfo[k] = v
	}

	return orderInfo, nil
}

// GetOrderEvents 获取订单事件历史
func (m *PoolMonitor) GetOrderEvents(ctx context.Context, orderSN string) ([]map[string]interface{}, error) {
	eventsKey := m.KeyGenerator.EventKey(orderSN)
	// 检索所有事件
	entries, err := m.RedisClient.XRange(ctx, eventsKey, "-", "+").Result()
	if err != nil {
		return nil, err
	}

	// 筛选指定订单的事件
	result := make([]map[string]interface{}, 0)

	for _, entry := range entries {
		if entry.Values["order_sn"] == orderSN {
			event := map[string]interface{}{
				"id":        entry.ID,
				"event":     entry.Values["event"],
				"timestamp": entry.Values["timestamp"],
				"details":   entry.Values["details"],
			}
			result = append(result, event)
		}
	}

	// 按时间戳排序
	sort.Slice(result, func(i, j int) bool {
		timestampI, _ := strconv.ParseInt(result[i]["timestamp"].(string), 10, 64)
		timestampJ, _ := strconv.ParseInt(result[j]["timestamp"].(string), 10, 64)
		return timestampI < timestampJ
	})

	return result, nil
}

// 辅助函数：解析Unix时间戳
func parseUnixTime(timeStr string) int64 {
	timestamp, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return 0
	}
	return timestamp
}
