package constant

// MerOrderMainStat 订单主状态
type MerOrderMainStat int

const (
	MerOrderMainStatCreated    MerOrderMainStat = 10 // 下单成功
	MerOrderMainStatProcessing MerOrderMainStat = 20 // 支付中
	MerOrderMainStatSuccess    MerOrderMainStat = 30 // 支付成功
	MerOrderMainStatFailed     MerOrderMainStat = 40 // 支付失败
	MerOrderMainStatTimeout    MerOrderMainStat = 50 // 超时支付
	MerOrderMainStatException  MerOrderMainStat = 60 // 订单异常
)

// String 实现主状态的字符串描述
func (s MerOrderMainStat) String() string {
	switch s {
	case MerOrderMainStatCreated:
		return "下单成功"
	case MerOrderMainStatProcessing:
		return "支付中"
	case MerOrderMainStatSuccess:
		return "支付成功"
	case MerOrderMainStatFailed:
		return "支付失败"
	case MerOrderMainStatTimeout:
		return "超时支付"
	case MerOrderMainStatException:
		return "订单异常"
	default:
		return "未知状态"
	}
}

// IsValid 检查状态值是否有效
func (s MerOrderMainStat) IsValid() bool {
	switch s {
	case MerOrderMainStatCreated, MerOrderMainStatProcessing, MerOrderMainStatSuccess,
		MerOrderMainStatFailed, MerOrderMainStatTimeout, MerOrderMainStatException:
		return true
	}
	return false
}

// Code 返回商户订单主状态的编码
func (s MerOrderMainStat) Code() string {
	switch s {
	case MerOrderMainStatCreated:
		return "created"
	case MerOrderMainStatProcessing:
		return "processing"
	case MerOrderMainStatSuccess:
		return "success"
	case MerOrderMainStatFailed:
		return "failed"
	case MerOrderMainStatTimeout:
		return "timeout"
	case MerOrderMainStatException:
		return "exception"
	default:
		return "unknown"
	}
}

// MerOrderSubStat 订单子状态，在订单池里面的流转状态
type MerOrderSubStat int

const (
	// 下单成功(10)的子状态
	MerOrderSubStatInitiated    MerOrderSubStat = 11 // 创建完成
	MerOrderSubStatMatching     MerOrderSubStat = 12 // 配单中
	MerOrderSubStatMatchFailed  MerOrderSubStat = 13 // 配单失败
	MerOrderSubStatMatchSuccess MerOrderSubStat = 14 // 配单成功

	// 支付中(20)的子状态
	MerOrderSubStatAuthorizing     MerOrderSubStat = 21 // 授权中
	MerOrderSubStatAuthSuccess     MerOrderSubStat = 22 // 授权成功
	MerOrderSubStatPayMatching     MerOrderSubStat = 23 // 配单中
	MerOrderSubStatPayMatchFailed  MerOrderSubStat = 24 // 配单失败
	MerOrderSubStatPayMatchSuccess MerOrderSubStat = 25 // 配单成功
	MerOrderSubStatCodeGenSuccess  MerOrderSubStat = 26 // 产码成功
	MerOrderSubStatCodeGenFailed   MerOrderSubStat = 27 // 产码失败
	MerOrderSubStatPaying          MerOrderSubStat = 28 // 支付中
	MerOrderSubStatQuerying        MerOrderSubStat = 29 // 查单中

	// 支付成功(30)的子状态
	MerOrderSubStatCompleted MerOrderSubStat = 31 // 支付完成

	// 支付失败(40)的子状态
	MerOrderSubStatUnpaid   MerOrderSubStat = 41 // 未付款
	MerOrderSubStatRefunded MerOrderSubStat = 42 // 已退款

	// 超时支付(50)的子状态
	MerOrderSubStatPayTimeout   MerOrderSubStat = 51 // 支付超时
	MerOrderSubStatOrderTimeout MerOrderSubStat = 52 // 订单超时
	MerOrderSubStatAuthTimeout  MerOrderSubStat = 53 // 授权超时

	// 订单异常(60)的子状态
	MerOrderSubStatBlacklist   MerOrderSubStat = 61 // 黑名单订单
	MerOrderSubStatRiskControl MerOrderSubStat = 62 // 风控拦截
	MerOrderSubStatSysError    MerOrderSubStat = 63 // 系统错误
)

// String 实现子状态的字符串描述
func (s MerOrderSubStat) String() string {
	switch s {
	case MerOrderSubStatInitiated:
		return "创建完成"
	case MerOrderSubStatMatching:
		return "配单中"
	case MerOrderSubStatMatchFailed:
		return "配单失败"
	case MerOrderSubStatMatchSuccess:
		return "配单成功"
	case MerOrderSubStatAuthorizing:
		return "授权中"
	case MerOrderSubStatAuthSuccess:
		return "授权成功"
	case MerOrderSubStatPayMatching:
		return "配单中"
	case MerOrderSubStatPayMatchFailed:
		return "配单失败"
	case MerOrderSubStatPayMatchSuccess:
		return "配单成功"
	case MerOrderSubStatCodeGenSuccess:
		return "产码成功"
	case MerOrderSubStatCodeGenFailed:
		return "产码失败"
	case MerOrderSubStatPaying:
		return "支付中"
	case MerOrderSubStatQuerying:
		return "查单中"
	case MerOrderSubStatCompleted:
		return "支付完成"
	case MerOrderSubStatUnpaid:
		return "未付款"
	case MerOrderSubStatRefunded:
		return "已退款"
	case MerOrderSubStatPayTimeout:
		return "支付超时"
	case MerOrderSubStatOrderTimeout:
		return "订单超时"
	case MerOrderSubStatAuthTimeout:
		return "授权超时"
	case MerOrderSubStatBlacklist:
		return "黑名单订单"
	case MerOrderSubStatRiskControl:
		return "风控拦截"
	case MerOrderSubStatSysError:
		return "系统错误"
	default:
		return "未知子状态"
	}
}

// Int 返回商户订单子状态的编码
func (s MerOrderSubStat) Int() int {
	return int(s)
}

// Code 返回商户订单子状态的编码
func (s MerOrderSubStat) Code() string {
	switch s {
	// 下单成功(10)的子状态
	case MerOrderSubStatInitiated:
		return "initiated"
	case MerOrderSubStatMatching:
		return "matching"
	case MerOrderSubStatMatchFailed:
		return "match_failed"
	case MerOrderSubStatMatchSuccess:
		return "match_success"
	// 支付中(20)的子状态
	case MerOrderSubStatAuthorizing:
		return "authorizing"
	case MerOrderSubStatAuthSuccess:
		return "auth_success"
	case MerOrderSubStatPayMatching:
		return "pay_matching"
	case MerOrderSubStatPayMatchFailed:
		return "pay_match_failed"
	case MerOrderSubStatPayMatchSuccess:
		return "pay_match_success"
	case MerOrderSubStatCodeGenSuccess:
		return "code_gen_success"
	case MerOrderSubStatCodeGenFailed:
		return "code_gen_failed"
	case MerOrderSubStatPaying:
		return "paying"
	case MerOrderSubStatQuerying:
		return "querying"
	// 支付成功(30)的子状态
	case MerOrderSubStatCompleted:
		return "completed"
	// 支付失败(40)的子状态
	case MerOrderSubStatUnpaid:
		return "unpaid"
	case MerOrderSubStatRefunded:
		return "refunded"
	// 超时支付(50)的子状态
	case MerOrderSubStatPayTimeout:
		return "pay_timeout"
	case MerOrderSubStatOrderTimeout:
		return "order_timeout"
	case MerOrderSubStatAuthTimeout:
		return "auth_timeout"
	// 订单异常(60)的子状态
	case MerOrderSubStatBlacklist:
		return "blacklist"
	case MerOrderSubStatRiskControl:
		return "risk_control"
	case MerOrderSubStatSysError:
		return "sys_error"
	default:
		return "unknown"
	}
}

// GetMerOrderMainStat 根据子状态获取主状态
func (s MerOrderSubStat) GetMerOrderMainStat() MerOrderMainStat {
	switch {
	case s >= 11 && s <= 14:
		return MerOrderMainStatCreated
	case s >= 21 && s <= 29:
		return MerOrderMainStatProcessing
	case s == 31:
		return MerOrderMainStatSuccess
	case s >= 41 && s <= 42:
		return MerOrderMainStatFailed
	case s >= 51 && s <= 53:
		return MerOrderMainStatTimeout
	case s >= 61 && s <= 63:
		return MerOrderMainStatException
	default:
		return 0
	}
}

// OrderStatus 订单状态结构体
type OrderStatus struct {
	MerOrderMainStat MerOrderMainStat
	MerOrderSubStat  MerOrderSubStat
}

// NewOrderStatus 创建新的订单状态
func NewOrderStatus(main MerOrderMainStat, sub MerOrderSubStat) OrderStatus {
	return OrderStatus{
		MerOrderMainStat: main,
		MerOrderSubStat:  sub,
	}
}

// IsValid 检查状态组合是否有效
func (o OrderStatus) IsValid() bool {
	return o.MerOrderSubStat.GetMerOrderMainStat() == o.MerOrderMainStat
}

// IsFinalStatus 检查是否为最终状态
func (o OrderStatus) IsFinalStatus() bool {
	switch o.MerOrderMainStat {
	case MerOrderMainStatSuccess, MerOrderMainStatFailed, MerOrderMainStatTimeout, MerOrderMainStatException:
		return true
	default:
		return false
	}
}

// IsProcessing 检查是否处于处理中状态
func (o OrderStatus) IsProcessing() bool {
	switch o.MerOrderSubStat {
	case MerOrderSubStatMatching, MerOrderSubStatAuthorizing, MerOrderSubStatPayMatching,
		MerOrderSubStatPaying, MerOrderSubStatQuerying:
		return true
	default:
		return false
	}
}

func IsProcessingSubStat(subStat int) bool {
	switch MerOrderSubStat(subStat) {
	case MerOrderSubStatMatching, MerOrderSubStatAuthorizing, MerOrderSubStatPayMatching,
		MerOrderSubStatPaying, MerOrderSubStatQuerying:
		return true
	default:
		return false
	}
}

// CanCancel 检查是否可以取消
func (o OrderStatus) CanCancel() bool {
	// 只有在下单成功和支付中的某些阶段才能取消
	switch o.MerOrderMainStat {
	case MerOrderMainStatCreated:
		return true
	case MerOrderMainStatProcessing:
		// 只有在授权前的阶段才能取消
		return o.MerOrderSubStat <= MerOrderSubStatAuthorizing
	default:
		return false
	}
}

// CanRetry 检查是否可以重试
func (o OrderStatus) CanRetry() bool {
	switch o.MerOrderSubStat {
	case MerOrderSubStatMatchFailed, MerOrderSubStatPayMatchFailed, MerOrderSubStatCodeGenFailed,
		MerOrderSubStatUnpaid:
		return true
	default:
		return false
	}
}
