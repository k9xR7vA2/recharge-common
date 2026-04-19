package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

// AccountLogBusinessTypeDict 流水日志操作类型字典
type AccountLogBusinessTypeDict struct{}

func (d *AccountLogBusinessTypeDict) GetKey() string {
	return "account_log_business_type"
}

func (d *AccountLogBusinessTypeDict) GetName() string {
	return "流水日志操作类型"
}

func (d *AccountLogBusinessTypeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.BusinessTypeDeposit.Label(),
			Value: int(constant.BusinessTypeDeposit),
			Code:  constant.BusinessTypeDeposit.Code(),
		},
		{
			Label: constant.BusinessTypePrepaidDeduct.Label(),
			Value: int(constant.BusinessTypePrepaidDeduct),
			Code:  constant.BusinessTypePrepaidDeduct.Code(),
		},
		{
			Label: constant.BusinessTypeOrderDeduct.Label(),
			Value: int(constant.BusinessTypeOrderDeduct),
			Code:  constant.BusinessTypeOrderDeduct.Code(),
		},
		{
			Label: constant.BusinessTypeRepay.Label(),
			Value: int(constant.BusinessTypeRepay),
			Code:  constant.BusinessTypeRepay.Code(),
		},
	}
}
