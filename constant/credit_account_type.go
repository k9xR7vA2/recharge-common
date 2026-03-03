package constant

// CreditAccountHandleType 信用账户相关（仅供应商）
type CreditAccountHandleType int

const (
	BusinessTypeCreditInit    CreditAccountHandleType = iota + 10 // 信用账户初始化
	BusinessTypeCreditDeposit                                     // 信用账户充值 额度扣减(订单支付)
	BusinessTypeCreditDeduct                                      // 信用账户扣减
	BusinessTypeCreditSettle                                      //授信结清
)
