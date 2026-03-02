package dict

import (
	"github.com/k9xR7vA2/recharge-common/dict/adapters"
	"github.com/k9xR7vA2/recharge-common/dict/types"
	"sync"
)

var (
	// GlobalManager 全局字典管理器
	GlobalManager *Manager
	once          sync.Once
)

// InitDictManager  初始化全局字典管理器（在 main.go 或 router 初始化时调用）
func InitDictManager() {
	once.Do(func() {
		GlobalManager = NewManager()
		// 注册所有字典
		registerAllDicts()
	})
}

// registerAllDicts 注册所有字典
func registerAllDicts() {
	// 业务类型
	GlobalManager.Register(&adapters.BusinessTypeDict{})
	// 运营商类型
	GlobalManager.Register(&adapters.CarrierTypeDict{})
	// 区域范围
	GlobalManager.Register(&adapters.AreaScopeDict{})
	// 印度运营商
	GlobalManager.Register(&adapters.IndiaCarrierTypeDict{})
	// 充值速度
	GlobalManager.Register(&adapters.ChargeSpeedDict{})
	// 设备类型
	GlobalManager.Register(&adapters.DeviceTypeDict{})
	//支付方式
	GlobalManager.Register(&adapters.PaymentTypeDict{})
	//支付方法
	GlobalManager.Register(&adapters.PaymentMethodDict{})
	//费率类型
	GlobalManager.Register(&adapters.RateTypeDict{})
	//地区
	GlobalManager.Register(&adapters.ProvinceDict{})
	// 优先级
	GlobalManager.Register(&adapters.PriorityDict{})
	//交易类型
	GlobalManager.Register(&adapters.TradeTypeDict{})

	// 账户相关字典
	GlobalManager.Register(&adapters.GlobalAccountStatusDict{})
	GlobalManager.Register(&adapters.AccountTypeDict{})
	GlobalManager.Register(&adapters.BalanceOperationDict{})
	// 通知状态字典
	GlobalManager.Register(&adapters.GlobalNotifyStatusDict{})
	// 通道相关字典
	GlobalManager.Register(&adapters.ChannelMatchRuleDict{})
	GlobalManager.Register(&adapters.ChannelTypeDict{})
	// 商户订单状态
	GlobalManager.Register(&adapters.MerOrderMainStatDict{}) // 商户订单主状态
	GlobalManager.Register(&adapters.MerOrderSubStatDict{})  // 商户订单子状态

	// 供货商订单状态
	GlobalManager.Register(&adapters.SupOrderPoolStatDict{}) // 供货商订单池状态
	GlobalManager.Register(&adapters.SupOrderStatusDict{})   // 供货商订单状态
	// TODO: 在这里继续注册其他字典
	// GlobalManager.Register(&YourNewDict{})
}

// GetDict 便捷方法：获取单个字典
func GetDict(key string) *types.DictResponse {
	if GlobalManager == nil {
		InitDictManager()
	}
	return GlobalManager.Get(key)
}

// GetAllDicts 便捷方法：获取所有字典
func GetAllDicts() map[string]*types.DictResponse {
	if GlobalManager == nil {
		InitDictManager()
	}
	return GlobalManager.GetAll()
}

// GetMultipleDicts 便捷方法：批量获取字典
func GetMultipleDicts(keys []string) map[string]*types.DictResponse {
	if GlobalManager == nil {
		InitDictManager()
	}
	return GlobalManager.GetMultiple(keys)
}
