package adapters

import (
	"recharge-common/constant"
	"recharge-common/dict/types"
)

// ChannelMatchRuleDict 通道匹配规则字典
type ChannelMatchRuleDict struct{}

func (d *ChannelMatchRuleDict) GetKey() string {
	return "channel_match_rule"
}

func (d *ChannelMatchRuleDict) GetName() string {
	return "通道匹配规则"
}

func (d *ChannelMatchRuleDict) GetOptions() []types.DictOption {
	rules := constant.GetAllMatchRules()
	options := make([]types.DictOption, 0, len(rules))

	for _, rule := range rules {
		var code string
		switch rule.Value {
		case constant.ChannelMatchRuleWeight:
			code = "weight"
		case constant.ChannelMatchRuleAmountMapping:
			code = "amount_mapping"
		case constant.ChannelMatchRuleMixed:
			code = "mixed"
		}

		options = append(options, types.DictOption{
			Label: rule.Label,
			Value: string(rule.Value),
			Code:  code,
		})
	}

	return options
}

// ChannelTypeDict 通道类型字典
type ChannelTypeDict struct{}

func (d *ChannelTypeDict) GetKey() string {
	return "channel_type"
}

func (d *ChannelTypeDict) GetName() string {
	return "通道类型"
}

func (d *ChannelTypeDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.ChannelTypeBasic.String(),
			Value: constant.ChannelTypeBasic.Val(),
			Code:  constant.ChannelTypeBasic.Code(),
		},
		{
			Label: constant.ChannelTypeComposite.String(),
			Value: constant.ChannelTypeComposite.Val(),
			Code:  constant.ChannelTypeComposite.Code(),
		},
	}
}
