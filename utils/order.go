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
