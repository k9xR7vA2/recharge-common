package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

// PriorityDict 优先级字典
type PriorityDict struct{}

func (d *PriorityDict) GetKey() string {
	return "priority"
}

func (d *PriorityDict) GetName() string {
	return "优先级"
}

func (d *PriorityDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.HighPriority.Name(),
			Value: string(constant.HighPriority),
			Code:  constant.HighPriority.String(),
		},
		{
			Label: constant.NormalPriority.Name(),
			Value: string(constant.NormalPriority),
			Code:  constant.NormalPriority.String(),
		},
	}
}

// TradeTypeDict 交易类型字典
type TradeTypeDict struct{}

func (d *TradeTypeDict) GetKey() string {
	return "trade_type"
}

func (d *TradeTypeDict) GetName() string {
	return "交易类型"
}

func (d *TradeTypeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.BackgroundOperations.Text(),
			Value: constant.BackgroundOperations.Value(),
			Code:  constant.BackgroundOperations.Code(),
		},
		{
			Label: constant.OrderSettlement.Text(),
			Value: constant.OrderSettlement.Value(),
			Code:  constant.OrderSettlement.Code(),
		},
	}
}
