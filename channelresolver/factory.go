package channelresolver

import (
	"fmt"
	"github.com/small-cat1/recharge-common/channelresolver/processor"
	"github.com/small-cat1/recharge-common/channelresolver/types"
	"github.com/small-cat1/recharge-common/constant"
)

type ChannelProcessorFactory struct{}

func (cp *ChannelProcessorFactory) CreateProcessor(ruleType constant.ChannelMatchRule) (types.ChannelProcessor, error) {
	switch ruleType {
	case constant.ChannelMatchRuleWeight:
		return &processor.WeightProcessor{}, nil
	case constant.ChannelMatchRuleAmountMapping:
		return &processor.AmountMappingProcessor{}, nil
	case constant.ChannelMatchRuleMixed:
		return &processor.MixedProcessor{}, nil
	default:
		return nil, fmt.Errorf("unsupported rule type: %s", ruleType)
	}
}
