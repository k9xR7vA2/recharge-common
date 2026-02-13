package adapters

import (
	"recharge-common/constant"
	"recharge-common/dict/types"
)

// SupOrderPoolStatDict 供货商订单池状态字典
type SupOrderPoolStatDict struct{}

func (d *SupOrderPoolStatDict) GetKey() string {
	return "sup_order_pool_stat"
}

func (d *SupOrderPoolStatDict) GetName() string {
	return "供货商订单池状态"
}

func (d *SupOrderPoolStatDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.SupOrderPoolWaiting.String(),
			Value: constant.SupOrderPoolWaiting.Val(),
			Code:  constant.SupOrderPoolWaiting.Code(),
		},
		{
			Label: constant.SupOrderPoolProcessing.String(),
			Value: constant.SupOrderPoolProcessing.Val(),
			Code:  constant.SupOrderPoolProcessing.Code(),
		},
		{
			Label: constant.SupOrderPoolCompleted.String(),
			Value: constant.SupOrderPoolCompleted.Val(),
			Code:  constant.SupOrderPoolCompleted.Code(),
		},
		{
			Label: constant.SupOrderPoolAwaitingCallback.String(),
			Value: constant.SupOrderPoolAwaitingCallback.Val(),
			Code:  constant.SupOrderPoolAwaitingCallback.Code(),
		},
		{
			Label: constant.SupOrderPoolFinalized.String(),
			Value: constant.SupOrderPoolFinalized.Val(),
			Code:  constant.SupOrderPoolFinalized.Code(),
		},
		{
			Label: constant.SupOrderPoolExpired.String(),
			Value: constant.SupOrderPoolExpired.Val(),
			Code:  constant.SupOrderPoolExpired.Code(),
		},
		{
			Label: constant.SupOrderPoolCanceled.String(),
			Value: constant.SupOrderPoolCanceled.Val(),
			Code:  constant.SupOrderPoolCanceled.Code(),
		},
	}
}

// SupOrderStatusDict 供货商订单状态字典
type SupOrderStatusDict struct{}

func (d *SupOrderStatusDict) GetKey() string {
	return "sup_order_status"
}

func (d *SupOrderStatusDict) GetName() string {
	return "供货商订单状态"
}

func (d *SupOrderStatusDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.SupOrderStatusWaiting.String(),
			Value: int(constant.SupOrderStatusWaiting),
			Code:  constant.SupOrderStatusWaiting.Code(),
		},
		{
			Label: constant.SupOrderStatusProcessing.String(),
			Value: int(constant.SupOrderStatusProcessing),
			Code:  constant.SupOrderStatusProcessing.Code(),
		},
		{
			Label: constant.SupOrderStatusSuccess.String(),
			Value: int(constant.SupOrderStatusSuccess),
			Code:  constant.SupOrderStatusSuccess.Code(),
		},
		{
			Label: constant.SupOrderStatusFailed.String(),
			Value: int(constant.SupOrderStatusFailed),
			Code:  constant.SupOrderStatusFailed.Code(),
		},
		{
			Label: constant.SupOrderStatusUnused.String(),
			Value: int(constant.SupOrderStatusUnused),
			Code:  constant.SupOrderStatusUnused.Code(),
		},
		{
			Label: constant.SupOrderStatusBlacklisted.String(),
			Value: int(constant.SupOrderStatusBlacklisted),
			Code:  constant.SupOrderStatusBlacklisted.Code(),
		},
		{
			Label: constant.SupOrderStatusCancelled.String(),
			Value: int(constant.SupOrderStatusCancelled),
			Code:  constant.SupOrderStatusCancelled.Code(),
		},
	}
}
