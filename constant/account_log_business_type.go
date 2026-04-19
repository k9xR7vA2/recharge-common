package constant

// CreditAccountHandleType 流水日志操作类型
type AccountLogBusinessType int

const (
	BusinessTypeDeposit       AccountLogBusinessType = 1 // 预付账户加款
	BusinessTypePrepaidDeduct AccountLogBusinessType = 2 // 预付账户扣款
	BusinessTypeOrderDeduct   AccountLogBusinessType = 3 // 订单扣款
	BusinessTypeRepay         AccountLogBusinessType = 4 // 供货商回款
)

func (t AccountLogBusinessType) Label() string {
	switch t {
	case BusinessTypeDeposit:
		return "预付账户加款"
	case BusinessTypePrepaidDeduct:
		return "预付账户扣款"
	case BusinessTypeOrderDeduct:
		return "订单扣款"
	case BusinessTypeRepay:
		return "供货商回款"
	default:
		return "未知类型"
	}
}

func (t AccountLogBusinessType) Code() string {
	switch t {
	case BusinessTypeDeposit:
		return "deposit"
	case BusinessTypePrepaidDeduct:
		return "prepaid_deduct"
	case BusinessTypeOrderDeduct:
		return "order_deduct"
	case BusinessTypeRepay:
		return "repay"
	default:
		return "unknown"
	}
}

func (t AccountLogBusinessType) Val() int {
	return int(t)
}

func (t AccountLogBusinessType) IsValid() bool {
	switch t {
	case BusinessTypeDeposit,
		BusinessTypePrepaidDeduct,
		BusinessTypeOrderDeduct,
		BusinessTypeRepay:
		return true
	default:
		return false
	}
}
