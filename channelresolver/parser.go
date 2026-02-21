package channelresolver

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/small-cat1/recharge-common/channelresolver/processor"
	"github.com/small-cat1/recharge-common/channelresolver/types"
	"github.com/small-cat1/recharge-common/constant"
	"gorm.io/datatypes"
)

// RuleConfig 规则配置接口

// RuleConfigParser 规则配置解析器
type RuleConfigParser struct{}

func NewRuleConfigParser() *RuleConfigParser {
	return &RuleConfigParser{}
}

func (p *RuleConfigParser) ParseConfig(rule constant.ChannelMatchRule, configStr datatypes.JSON) (types.RuleConfig, error) {
	if configStr == nil {
		return nil, errors.New("配置内容为空")
	}
	var config types.RuleConfig
	switch rule {
	case constant.ChannelMatchRuleWeight:
		config = &processor.WeightRuleConfig{}
	case constant.ChannelMatchRuleAmountMapping:
		config = &processor.AmountMappingRuleConfig{}
	case constant.ChannelMatchRuleMixed:
		config = &processor.MixedRuleConfig{}
	default:
		return nil, errors.New("不支持的规则类型")
	}

	// 解析JSON
	if err := json.Unmarshal(configStr, config); err != nil {
		return nil, fmt.Errorf("配置解析失败: %v", err)
	}
	// 验证配置
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}

// ParseRuleConfig 定义解析器
func ParseRuleConfig(RuleType constant.ChannelMatchRule, RuleConfig datatypes.JSON) (types.RuleConfig, error) {
	parser := &RuleConfigParser{}
	config, err := parser.ParseConfig(RuleType, RuleConfig)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", processor.ErrConfigParse, err)
	}
	return config, nil
}
