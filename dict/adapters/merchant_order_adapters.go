package adapters

import (
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/dict/types"
)

// MerOrderMainStatDict 商户订单主状态字典
type MerOrderMainStatDict struct{}

func (d *MerOrderMainStatDict) GetKey() string {
	return "mer_order_main_stat"
}

func (d *MerOrderMainStatDict) GetName() string {
	return "商户订单主状态"
}

func (d *MerOrderMainStatDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.MerOrderMainStatCreated.String(),
			Value: int(constant.MerOrderMainStatCreated),
			Code:  constant.MerOrderMainStatCreated.Code(),
		},
		{
			Label: constant.MerOrderMainStatProcessing.String(),
			Value: int(constant.MerOrderMainStatProcessing),
			Code:  constant.MerOrderMainStatProcessing.Code(),
		},
		{
			Label: constant.MerOrderMainStatSuccess.String(),
			Value: int(constant.MerOrderMainStatSuccess),
			Code:  constant.MerOrderMainStatSuccess.Code(),
		},
		{
			Label: constant.MerOrderMainStatFailed.String(),
			Value: int(constant.MerOrderMainStatFailed),
			Code:  constant.MerOrderMainStatFailed.Code(),
		},
		{
			Label: constant.MerOrderMainStatTimeout.String(),
			Value: int(constant.MerOrderMainStatTimeout),
			Code:  constant.MerOrderMainStatTimeout.Code(),
		},
		{
			Label: constant.MerOrderMainStatException.String(),
			Value: int(constant.MerOrderMainStatException),
			Code:  constant.MerOrderMainStatException.Code(),
		},
	}
}

// MerOrderSubStatDict 商户订单子状态字典
type MerOrderSubStatDict struct{}

func (d *MerOrderSubStatDict) GetKey() string {
	return "mer_order_sub_stat"
}

func (d *MerOrderSubStatDict) GetName() string {
	return "商户订单子状态"
}

func (d *MerOrderSubStatDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		// 下单成功(10)的子状态
		{
			Label: constant.MerOrderSubStatInitiated.String(),
			Value: int(constant.MerOrderSubStatInitiated),
			Code:  constant.MerOrderSubStatInitiated.Code(),
		},
		{
			Label: constant.MerOrderSubStatMatching.String(),
			Value: int(constant.MerOrderSubStatMatching),
			Code:  constant.MerOrderSubStatMatching.Code(),
		},
		{
			Label: constant.MerOrderSubStatMatchFailed.String(),
			Value: int(constant.MerOrderSubStatMatchFailed),
			Code:  constant.MerOrderSubStatMatchFailed.Code(),
		},
		{
			Label: constant.MerOrderSubStatMatchSuccess.String(),
			Value: int(constant.MerOrderSubStatMatchSuccess),
			Code:  constant.MerOrderSubStatMatchSuccess.Code(),
		},
		// 支付中(20)的子状态
		{
			Label: constant.MerOrderSubStatAuthorizing.String(),
			Value: int(constant.MerOrderSubStatAuthorizing),
			Code:  constant.MerOrderSubStatAuthorizing.Code(),
		},
		{
			Label: constant.MerOrderSubStatAuthSuccess.String(),
			Value: int(constant.MerOrderSubStatAuthSuccess),
			Code:  constant.MerOrderSubStatAuthSuccess.Code(),
		},
		{
			Label: constant.MerOrderSubStatPayMatching.String(),
			Value: int(constant.MerOrderSubStatPayMatching),
			Code:  constant.MerOrderSubStatPayMatching.Code(),
		},
		{
			Label: constant.MerOrderSubStatPayMatchFailed.String(),
			Value: int(constant.MerOrderSubStatPayMatchFailed),
			Code:  constant.MerOrderSubStatPayMatchFailed.Code(),
		},
		{
			Label: constant.MerOrderSubStatPayMatchSuccess.String(),
			Value: int(constant.MerOrderSubStatPayMatchSuccess),
			Code:  constant.MerOrderSubStatPayMatchSuccess.Code(),
		},
		{
			Label: constant.MerOrderSubStatCodeGenSuccess.String(),
			Value: int(constant.MerOrderSubStatCodeGenSuccess),
			Code:  constant.MerOrderSubStatCodeGenSuccess.Code(),
		},
		{
			Label: constant.MerOrderSubStatCodeGenFailed.String(),
			Value: int(constant.MerOrderSubStatCodeGenFailed),
			Code:  constant.MerOrderSubStatCodeGenFailed.Code(),
		},
		{
			Label: constant.MerOrderSubStatPaying.String(),
			Value: int(constant.MerOrderSubStatPaying),
			Code:  constant.MerOrderSubStatPaying.Code(),
		},
		{
			Label: constant.MerOrderSubStatQuerying.String(),
			Value: int(constant.MerOrderSubStatQuerying),
			Code:  constant.MerOrderSubStatQuerying.Code(),
		},
		// 支付成功(30)的子状态
		{
			Label: constant.MerOrderSubStatCompleted.String(),
			Value: int(constant.MerOrderSubStatCompleted),
			Code:  constant.MerOrderSubStatCompleted.Code(),
		},
		// 支付失败(40)的子状态
		{
			Label: constant.MerOrderSubStatUnpaid.String(),
			Value: int(constant.MerOrderSubStatUnpaid),
			Code:  constant.MerOrderSubStatUnpaid.Code(),
		},
		{
			Label: constant.MerOrderSubStatRefunded.String(),
			Value: int(constant.MerOrderSubStatRefunded),
			Code:  constant.MerOrderSubStatRefunded.Code(),
		},
		// 超时支付(50)的子状态
		{
			Label: constant.MerOrderSubStatPayTimeout.String(),
			Value: int(constant.MerOrderSubStatPayTimeout),
			Code:  constant.MerOrderSubStatPayTimeout.Code(),
		},
		{
			Label: constant.MerOrderSubStatOrderTimeout.String(),
			Value: int(constant.MerOrderSubStatOrderTimeout),
			Code:  constant.MerOrderSubStatOrderTimeout.Code(),
		},
		{
			Label: constant.MerOrderSubStatAuthTimeout.String(),
			Value: int(constant.MerOrderSubStatAuthTimeout),
			Code:  constant.MerOrderSubStatAuthTimeout.Code(),
		},
		// 订单异常(60)的子状态
		{
			Label: constant.MerOrderSubStatBlacklist.String(),
			Value: int(constant.MerOrderSubStatBlacklist),
			Code:  constant.MerOrderSubStatBlacklist.Code(),
		},
		{
			Label: constant.MerOrderSubStatRiskControl.String(),
			Value: int(constant.MerOrderSubStatRiskControl),
			Code:  constant.MerOrderSubStatRiskControl.Code(),
		},
		{
			Label: constant.MerOrderSubStatSysError.String(),
			Value: int(constant.MerOrderSubStatSysError),
			Code:  constant.MerOrderSubStatSysError.Code(),
		},
	}
}
