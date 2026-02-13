package constant

type RateType uint

const (
	RateTypePercentage RateType = iota + 1 // 百分比
	RateTypePerMille                       // 千分比
	RateTypeFixed                          // 单笔固定
)

func (r RateType) String() string {
	switch r {
	case RateTypePercentage:
		return "百分比"
	case RateTypePerMille:
		return "千分比"
	case RateTypeFixed:
		return "单笔"
	default:
		return "未知类型"
	}
}

func (r RateType) Val() int {
	switch r {
	case RateTypePercentage:
		return 1
	case RateTypePerMille:
		return 2
	case RateTypeFixed:
		return 3
	default:
		return 0
	}
}

func (r RateType) Code() string {
	switch r {
	case RateTypePercentage:
		return "percentage"
	case RateTypePerMille:
		return "permille"
	case RateTypeFixed:
		return "fixed"
	default:
		return "unknown"
	}
}

// Symbol 返回费率符号
func (r RateType) Symbol() string {
	switch r {
	case RateTypePercentage:
		return "%"
	case RateTypePerMille:
		return "‰"
	case RateTypeFixed:
		return "元"
	default:
		return ""
	}
}

func (r RateType) IsValid() bool {
	return r >= RateTypePercentage && r <= RateTypeFixed
}

// CalculateFee 计算手续费（rate是整数）
// amount: 交易金额（单位：分）
// rate: 费率整数值（百分比存5表示5%，千分比存2表示千2，单笔存1000表示10元）
// 返回: 手续费（单位：分）
func (r RateType) CalculateFee(amount int64, rate int) int64 {
	switch r {
	case RateTypePercentage:
		// 5% = 5/100，金额以分为单位
		return amount * int64(rate) / 100
	case RateTypePerMille:
		// 千2 = 2/1000
		return amount * int64(rate) / 1000
	case RateTypeFixed:
		// 固定费用，rate直接就是分
		return int64(rate)
	default:
		return 0
	}
}

// GetMaxRate 获取该类型的最大费率值
func (r RateType) GetMaxRate() int {
	switch r {
	case RateTypePercentage:
		return 100
	case RateTypePerMille:
		return 1000
	case RateTypeFixed:
		return 99999900 // 999999元 = 99999900分
	default:
		return 0
	}
}
