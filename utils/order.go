package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func GenerateOrderIDWithRedis(ctx context.Context, redisClient redis.UniversalClient, prefix string) (string, error) {
	// 生成一个基于日期的键，确保不同日期的订单号分开计数
	dateKey := time.Now().Format("20060102")
	key := fmt.Sprintf("order:id:counter:%s", dateKey)
	// 使用Redis的INCR命令原子递增
	counter, err := redisClient.Incr(ctx, key).Result()
	if err != nil {
		return "", err
	}
	// 确保计数器键有过期时间（例如保留7天）
	redisClient.Expire(ctx, key, 7*24*time.Hour)
	// 生成格式化的订单号
	timestamp := time.Now().Format("20060102150405")
	orderID := fmt.Sprintf("%s%s%08d", prefix, timestamp, counter)
	return orderID, nil
}

// utils/order.go
func ParseOrderDate(systemOrderSn string, prefixLen int) (string, error) {
	// 订单号格式: prefix(prefixLen位) + 20060102150405(14位) + counter(8位)
	if len(systemOrderSn) < prefixLen+14 {
		return "", fmt.Errorf("invalid order sn: %s", systemOrderSn)
	}
	// 取日期部分 yyyyMMdd
	return systemOrderSn[prefixLen : prefixLen+8], nil
}
