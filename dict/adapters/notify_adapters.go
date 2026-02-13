package adapters

import (
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/dict/types"
)

// GlobalNotifyStatusDict 通知状态字典
type GlobalNotifyStatusDict struct{}

func (d *GlobalNotifyStatusDict) GetKey() string {
	return "notify_status"
}

func (d *GlobalNotifyStatusDict) GetName() string {
	return "通知状态"
}

func (d *GlobalNotifyStatusDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.NotifyStatusPreparing.String(),
			Value: int(constant.NotifyStatusPreparing),
			Code:  constant.NotifyStatusPreparing.Code(),
		},
		{
			Label: constant.NotifyStatusProcessing.String(),
			Value: int(constant.NotifyStatusProcessing),
			Code:  constant.NotifyStatusProcessing.Code(),
		},
		{
			Label: constant.NotifyStatusSuccess.String(),
			Value: int(constant.NotifyStatusSuccess),
			Code:  constant.NotifyStatusSuccess.Code(),
		},
		{
			Label: constant.NotifyStatusError.String(),
			Value: int(constant.NotifyStatusError),
			Code:  constant.NotifyStatusError.Code(),
		},
		{
			Label: constant.NotifyStatusTimeout.String(),
			Value: int(constant.NotifyStatusTimeout),
			Code:  constant.NotifyStatusTimeout.Code(),
		},
	}
}
