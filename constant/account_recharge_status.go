package constant

// 充值账号状态
type AccountRechargeStatus int

const (
	AccountRechargePending  AccountRechargeStatus = 1 + iota // 待审核
	AccountRechargeApproved                                  // 审核通过
	AccountRechargeCharging                                  // 充值中
	AccountRechargeFinished                                  // 已完成
	AccountRechargeRejected                                  // 已拒绝
)

func (s AccountRechargeStatus) Label() string {
	switch s {
	case AccountRechargePending:
		return "待审核"
	case AccountRechargeApproved:
		return "审核通过"
	case AccountRechargeCharging:
		return "充值中"
	case AccountRechargeFinished:
		return "已完成"
	case AccountRechargeRejected:
		return "已拒绝"
	default:
		return "未知状态"
	}
}

func (s AccountRechargeStatus) Val() int {
	return int(s)
}

func (s AccountRechargeStatus) Code() string {
	switch s {
	case AccountRechargePending:
		return "pending"
	case AccountRechargeApproved:
		return "approved"
	case AccountRechargeCharging:
		return "charging"
	case AccountRechargeFinished:
		return "finished"
	case AccountRechargeRejected:
		return "rejected"
	default:
		return "unknown"
	}
}

func (s AccountRechargeStatus) IsValid() bool {
	return s >= AccountRechargePending && s <= AccountRechargeRejected
}
