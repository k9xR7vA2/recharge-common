package options

import (
	"github.com/small-cat1/recharge-common/orderpool/entities"
	"strconv"
)

// ---------- Fetch 专属分组结构体 ----------

// TenantInfo 租户身份信息（可被其他 Builder 共用）
type TenantInfo struct {
	TenantID     uint
	RoleType     string
	BusinessType string
}

// FetchRedisKeys Fetch 场景下的 Redis key 分组
type FetchRedisKeys struct {
	StatsKey              string
	HighPriorityPoolKey   string
	NormalPriorityPoolKey string
}

// 处理器选项
type FetchMobileHandlerOptions struct {
	tenantID              uint
	roleType              string
	statsKey              string
	businessType          string
	poolKey               string //订单池的key
	poolArgs              entities.MobilePoolArgs
	highPriorityPoolKey   string // 新增
	normalPriorityPoolKey string // 新增
}

func (o FetchMobileHandlerOptions) GetTenantId() uint {
	return o.tenantID
}
func (o FetchMobileHandlerOptions) GetRoleType() string {
	return o.roleType
}
func (o FetchMobileHandlerOptions) GetBusinessType() string {
	return o.businessType
}
func (o FetchMobileHandlerOptions) GetPoolKey() string {
	return o.poolKey
}
func (o FetchMobileHandlerOptions) GetPoolArgs() entities.MobilePoolArgs {
	return o.poolArgs
}
func (o FetchMobileHandlerOptions) GetStatsKey() string {
	return o.statsKey
}
func (o FetchMobileHandlerOptions) GetHighPriorityPoolKey() string {
	return o.highPriorityPoolKey
}
func (o FetchMobileHandlerOptions) GetNormalPriorityPoolKey() string {
	return o.normalPriorityPoolKey
}

func (o FetchMobileHandlerOptions) GetEventsKey(orderSn string) string {
	tenantIdStr := strconv.Itoa(int(o.GetTenantId()))
	return "tenant:" + tenantIdStr + ":" + o.GetRoleType() + ":" + o.GetBusinessType() + ":order:" + orderSn + ":events"
}
