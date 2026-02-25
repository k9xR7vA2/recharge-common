package options

import "github.com/small-cat1/recharge-common/orderpool/entities"

// ---------- 分组参数结构体 ----------

// RedisKeys 所有 Redis key 归为一组
type RedisKeys struct {
	OrderKey string
	PoolKey  string
}

// OrderInfo 订单标识信息归为一组
type OrderInfo struct {
	OrderSn         string
	SupplierOrderSN string
	Priority        string
}

// TimeArgs 时间相关参数归为一组
type TimeArgs struct {
	ValidTime int   // 有效期(秒)
	ExpiredAt int64 // 优先级分数(过期时间戳)
}

// MobileHandlerOptions 处理器选项
type MobileHandlerOptions struct {
	priority string //池子优先级
	orderKey string //订单信息的key
	poolKey  string //订单池的key

	orderData       string // ARGV[1] 订单数据(JSON)
	validTime       int    // ARGV[2] 订单有效期(秒)
	expiredAt       int64  // ARGV[3] 订单优先级分数(过期时间)
	orderSn         string // 订单号
	supplierOrderSN string //
	poolArgs        entities.MobilePoolArgs
}

func (h MobileHandlerOptions) GetPriority() string {
	return h.priority
}

func (h MobileHandlerOptions) GetSupplierOrderSnArg() string {
	return h.supplierOrderSN
}
func (h MobileHandlerOptions) GetPoolKey() string {
	return h.poolKey
}

func (h MobileHandlerOptions) GetOrderKey() string {
	return h.orderKey
}

func (h MobileHandlerOptions) GetOrderDataArg() string {
	return h.orderData
}

// GetValidTimeArg 有效期(秒)
func (h MobileHandlerOptions) GetValidTimeArg() int {
	return h.validTime
}

// GetExpiredAtArg GetExpiredAt 订单过期时间
func (h MobileHandlerOptions) GetExpiredAtArg() int64 {
	return h.expiredAt
}

func (h MobileHandlerOptions) GetSystemOrderSnArg() string {
	return h.orderSn
}

func (h MobileHandlerOptions) GetPoolArgs() entities.MobilePoolArgs {
	return h.poolArgs
}
