package constant

// ChannelMatchRule 通道匹配规则
type ChannelMatchRule string

const (
	// ChannelMatchRuleWeight 子通道权重模式
	ChannelMatchRuleWeight ChannelMatchRule = "WEIGHT"
	// ChannelMatchRuleAmountMapping 金额子通道映射
	ChannelMatchRuleAmountMapping ChannelMatchRule = "AMOUNT_MAPPING"
	// ChannelMatchRuleMixed 混合权重模式
	ChannelMatchRuleMixed ChannelMatchRule = "MIXED"
)

// String 实现匹配规则的字符串表示
func (r ChannelMatchRule) String() string {
	switch r {
	case ChannelMatchRuleWeight:
		return "子通道权重模式"
	case ChannelMatchRuleAmountMapping:
		return "金额子通道映射"
	case ChannelMatchRuleMixed:
		return "混合权重模式"
	default:
		return "未知规则"
	}
}

// IsValid 验证匹配规则是否有效
func (r ChannelMatchRule) IsValid() bool {
	switch r {
	case ChannelMatchRuleWeight, ChannelMatchRuleAmountMapping, ChannelMatchRuleMixed:
		return true
	default:
		return false
	}
}

// GetAllMatchRules 获取所有匹配规则
func GetAllMatchRules() []struct {
	Label string           `json:"label"`
	Value ChannelMatchRule `json:"value"`
} {
	return []struct {
		Label string           `json:"label"`
		Value ChannelMatchRule `json:"value"`
	}{
		{Label: ChannelMatchRuleWeight.String(), Value: ChannelMatchRuleWeight},
		{Label: ChannelMatchRuleAmountMapping.String(), Value: ChannelMatchRuleAmountMapping},
		{Label: ChannelMatchRuleMixed.String(), Value: ChannelMatchRuleMixed},
	}
}

// WeightRuleConfig 1. 权重模式配置
type WeightRuleConfig map[string]int // channelID -> weight

// AmountMappingRuleConfig 2. 金额映射配置
type AmountMappingRuleConfig map[string]uint // amount -> channelID

// MixedRuleConfig 3. 混合权重模式配置
type MixedRuleConfig map[string]WeightRuleConfig // amount -> (channelID -> weight)
