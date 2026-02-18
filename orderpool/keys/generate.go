package keys

import (
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/orderpool/entities"
)

// 关键前缀定义
const (
	// 角色前缀
	RoleSupplier = "supplier" // 供货商
	RoleMerchant = "merchant" // 商户

	// 数据类型前缀
	TypeOrder      = "order"       // 订单信息
	TypePool       = "pool"        // 池子
	TypeStats      = "stats"       // 统计信息
	TypePayment    = "payment"     // 支付参数
	TypeEvents     = "events"      //事件
	TypeOrderCache = "order:cache" //订单缓存

	KeySeparator = ":" // 分隔符
)

// RedisKeysGenerate key生成器
type RedisKeysGenerate struct {
	tenantId     uint
	role         string // 角色：supplier/merchant
	businessType string // 业务类型：mobile/game/oil
}

func NewRedisKeysGenerate(tenantId uint, role, businessType string) *RedisKeysGenerate {
	return &RedisKeysGenerate{
		tenantId:     tenantId,
		role:         role,
		businessType: businessType,
	}
}

// KeyBuilder 基础方法
func (rg *RedisKeysGenerate) baseBuilder() *KeyBuilder {
	return NewKeyBuilder(rg.tenantId).
		Add(rg.role).
		Add(rg.businessType)
}

// GenerateMobilePoolKey  话费订单池(Stream)
// Key: tenant:{tenantID}:{role}:{businessType}:pool:{priority}:{amount}:{carrier}:{chargeType}:{region}:{province}
// Entry:
// - order_sn: 订单编号
// - expire_at: 过期时间戳
// - retry_count: 重试次数
// - create_time: 入池时间
func (rg *RedisKeysGenerate) GenerateMobilePoolKey(Priority string, poolArgs entities.MobilePoolArgs) string {
	kb := rg.baseBuilder().
		Add(TypePool).
		Add(Priority).
		Add(poolArgs.Amount).
		Add(poolArgs.Carrier).
		Add(poolArgs.ChargeSpeed).
		Add(poolArgs.Region)
	if poolArgs.Region == constant.Province.Code() {
		kb.Add(poolArgs.Province)
	}
	return kb.Build()
}

// OrderKey  订单信息 (Hash)
// Key: tenant:{tenantID}:{role}:{businessType}:order:{orderSN}
// Fields:
//   - info: 订单完整信息（JSON）
//   - status: 订单状态（1=等待, 2=处理中, 3=成功, 4=失败, 9=撤销）
//   - create_time: 创建时间
//   - update_time: 更新时间
//   - expire_at: 过期时间戳
//
// TTL: 与订单有效期一致
func (rg *RedisKeysGenerate) OrderKey(orderSn string) string {
	return rg.baseBuilder().
		Add(TypeOrder).
		Add(orderSn).
		Build()
}

// StatsKey 统计信息 (Hash)
// Key: tenant:{tenantID}:{role}:{businessType}:stats
// Fields:
//   - amount:{value}: 金额维度计数
//   - carrier:{value}: 运营商维度计数
//   - charge_type:{value}: 充值类型维度计数
//   - area:{value}: 区域维度计数
//   - total_orders: 总订单数
//   - processing_orders: 处理中订单数
//   - pool_orders: 池中订单数
//   - priority:high: 高优先级订单数
//   - priority:normal: 普通优先级订单数
func (rg *RedisKeysGenerate) StatsKey() string {
	return rg.baseBuilder().
		Add(TypeStats).
		Build()
}

// EventKey 事件流 (Stream)
// Key: tenant:{tenantID}:{role}:{businessType}:order:{orderSN}:events
// Entry:
//   - order_sn: 订单编号
//   - event: 事件类型 (created, processing, completed, failed, expired, canceled)
//   - timestamp: 时间戳
//   - details: 事件详情（JSON）
func (rg *RedisKeysGenerate) EventKey(orderSn string) string {
	return rg.baseBuilder().
		Add(orderSn).
		Add(TypeEvents).
		Build()
}

// PaymentKey 生成支付参数的 key
// 格式: tenant_{id}:merchant:mobile:payment:{order_sn}
func (rg *RedisKeysGenerate) PaymentKey(orderSn string) string {
	return rg.baseBuilder().
		Add(TypePayment).
		Add(orderSn).
		Build()
}
