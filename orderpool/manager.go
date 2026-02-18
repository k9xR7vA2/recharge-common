package orderpool

import (
	"github.com/redis/go-redis/v9"
)

type OrderPoolManager struct {
	mobileHandlerManager *MobileOrderPool
}

func NewOrderPoolManager(redisClient redis.UniversalClient) *OrderPoolManager {
	return &OrderPoolManager{
		mobileHandlerManager: NewMobileOrderPool(redisClient), //话费订单池处理器
	}
}

func (m *OrderPoolManager) GetOrderPoolOptionsFactory() IOrderPoolOptionsFactory {
	return NewOrderPoolOptionsFactory()
}

// GetMobileOrderPool 话费业务订单池
func (m *OrderPoolManager) GetMobileOrderPool() *MobileOrderPool {
	return m.mobileHandlerManager
}
