package adapters

import (
	"recharge-common/constant"
	"recharge-common/dict/types"
)

// GlobalAccountStatusDict 账户状态字典
type GlobalAccountStatusDict struct{}

func (d *GlobalAccountStatusDict) GetKey() string {
	return "account_status"
}

func (d *GlobalAccountStatusDict) GetName() string {
	return "账户状态"
}

func (d *GlobalAccountStatusDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.AccountStatusActive.Text(),
			Value: int(constant.AccountStatusActive),
			Code:  constant.AccountStatusActive.Code(),
		},
		{
			Label: constant.AccountStatusDisabled.Text(),
			Value: int(constant.AccountStatusDisabled),
			Code:  constant.AccountStatusDisabled.Code(),
		},
	}
}

// AccountTypeDict 账户类型字典
type AccountTypeDict struct{}

func (d *AccountTypeDict) GetKey() string {
	return "account_type"
}

func (d *AccountTypeDict) GetName() string {
	return "账户类型"
}

func (d *AccountTypeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.RentalAccount.Text(),
			Value: constant.RentalAccount.Value(),
			Code:  constant.RentalAccount.Code(),
		},
		{
			Label: constant.CreditAccount.Text(),
			Value: constant.CreditAccount.Value(),
			Code:  constant.CreditAccount.Code(),
		},
	}
}

// BalanceOperationDict 余额操作字典
type BalanceOperationDict struct{}

func (d *BalanceOperationDict) GetKey() string {
	return "balance_operation"
}

func (d *BalanceOperationDict) GetName() string {
	return "余额操作"
}

func (d *BalanceOperationDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.AddBalance.Text(),
			Value: constant.AddBalance.Value(),
			Code:  constant.AddBalance.Code(),
		},
		{
			Label: constant.SubtractBalance.Text(),
			Value: constant.SubtractBalance.Value(),
			Code:  constant.SubtractBalance.Code(),
		},
	}
}
