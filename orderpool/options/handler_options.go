package options

import (
	"github.com/small-cat1/recharge-common/orderpool/entities"
)

type IMobileHandlerOptions interface {
	GetPriority() string // 获取池子优先级
	GetOrderKey() string //订单信息的key
	GetPoolKey() string  //订单池的key

	GetOrderDataArg() string //订单数据(JSON)
	GetValidTimeArg() int    //有效期(秒)
	GetExpiredAtArg() int64  //订单优先级分数

	GetSystemOrderSnArg() string   // 系统订单号
	GetSupplierOrderSnArg() string // 商户订单号

	GetPoolArgs() entities.MobilePoolArgs //获取订单池参数
}

type IFetchMobileOrderOptions interface {
	GetTenantId() uint
	GetRoleType() string
	GetBusinessType() string
	GetPoolArgs() entities.MobilePoolArgs //获取订单池参数
	GetHighPriorityPoolKey() string
	GetNormalPriorityPoolKey() string
}
