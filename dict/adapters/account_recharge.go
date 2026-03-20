package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

// AccountRechargeStatusDict 充值账号状态字典
type AccountRechargeStatusDict struct{}

func (d *AccountRechargeStatusDict) GetKey() string {
	return "account_recharge_status"
}

func (d *AccountRechargeStatusDict) GetName() string {
	return "充值账号状态"
}

func (d *AccountRechargeStatusDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.AccountRechargePending.Label(),
			Value: int(constant.AccountRechargePending),
			Code:  constant.AccountRechargePending.Code(),
		},
		{
			Label: constant.AccountRechargeApproved.Label(),
			Value: int(constant.AccountRechargeApproved),
			Code:  constant.AccountRechargeApproved.Code(),
		},
		{
			Label: constant.AccountRechargeCharging.Label(),
			Value: int(constant.AccountRechargeCharging),
			Code:  constant.AccountRechargeCharging.Code(),
		},
		{
			Label: constant.AccountRechargeFinished.Label(),
			Value: int(constant.AccountRechargeFinished),
			Code:  constant.AccountRechargeFinished.Code(),
		},
		{
			Label: constant.AccountRechargeRejected.Label(),
			Value: int(constant.AccountRechargeRejected),
			Code:  constant.AccountRechargeRejected.Code(),
		},
	}
}
