package constant

type AmountType uint

const (
	AmountTypeFixed   AmountType = 1 + iota // 固定金额
	AmountTypeRange                         // 区间金额
	AmountTypeDynamic                       // 动态金额（套餐）
	AmountTypeBill                          // 账单金额
)

func (t AmountType) Label() string {
	switch t {
	case AmountTypeFixed:
		return "固定金额"
	case AmountTypeRange:
		return "区间金额"
	case AmountTypeDynamic:
		return "动态金额"
	case AmountTypeBill:
		return "账单金额"
	default:
		return "未知"
	}
}

func (t AmountType) Val() uint {
	return uint(t)
}

func (t AmountType) Code() string {
	switch t {
	case AmountTypeFixed:
		return "fixed"
	case AmountTypeRange:
		return "range"
	case AmountTypeDynamic:
		return "dynamic"
	case AmountTypeBill:
		return "bill"
	default:
		return "unknown"
	}
}

func (t AmountType) IsValid() bool {
	return t >= AmountTypeFixed && t <= AmountTypeBill
}

// IsBillMode 账单模式，录入时不填金额
func (t AmountType) IsBillMode() bool {
	return t == AmountTypeBill
}

// IsDynamicMode 动态金额，金额由套餐决定
func (t AmountType) IsDynamicMode() bool {
	return t == AmountTypeDynamic
}
