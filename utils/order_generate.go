package utils

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// OrderIDGenerator 订单号生成器
type OrderIDGenerator struct {
	// 用于确保计数器递增的线程安全
	mu sync.Mutex
	// 存储当前日期的计数器
	counters map[string]int
	// 配置选项
	prefix         string
	counterDigits  int
	counterExpires time.Duration
}

// NewOrderIDGenerator 创建新的订单号生成器 admin_server,tenant_server使用
func NewOrderIDGenerator(prefix string, counterDigits int) *OrderIDGenerator {
	return &OrderIDGenerator{
		counters:       make(map[string]int),
		prefix:         prefix,
		counterDigits:  counterDigits,
		counterExpires: 7 * 24 * time.Hour, // 默认7天过期
	}
}

// GenerateOrderID 生成订单号
func (g *OrderIDGenerator) GenerateOrderID(ctx context.Context) (string, error) {
	// 检查上下文是否已取消
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	// 生成日期键
	now := time.Now()
	dateKey := now.Format("20060102")

	// 递增计数器
	g.counters[dateKey]++
	counter := g.counters[dateKey]

	// 清理过期的计数器（实际应用中可以定期执行）
	g.cleanExpiredCounters(now)

	// 生成订单号
	timestamp := now.Format("20060102150405")
	orderID := fmt.Sprintf("%s%s%0*d", g.prefix, timestamp, g.counterDigits, counter)

	return orderID, nil
}

// 清理过期的计数器
func (g *OrderIDGenerator) cleanExpiredCounters(now time.Time) {
	threshold := now.Add(-g.counterExpires)
	thresholdKey := threshold.Format("20060102")

	for dateKey := range g.counters {
		if dateKey < thresholdKey {
			delete(g.counters, dateKey)
		}
	}
}
