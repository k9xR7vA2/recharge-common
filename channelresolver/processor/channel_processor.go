package processor

import (
	"errors"
	"github.com/small-cat1/recharge-common/channelresolver/types"
	"sort"
)

// 定义常用的错误类型
var (
	ErrConfigParse = errors.New("规则配置解析失败")
)

// WeightProcessor 权重模式处理器
type WeightProcessor struct{}

// ProcessAmount  处理器
func (p *WeightProcessor) ProcessAmount(config types.RuleConfig, baseChannelMap map[uint]types.ChannelAmount) []types.ChannelAmount {
	wc := config.(*WeightRuleConfig)
	numericWeights := wc.ToNumeric()
	var result []types.ChannelAmount
	for channelID := range numericWeights {
		if options, exists := baseChannelMap[uint(channelID)]; exists {
			result = append(result, options)
		}
	}
	return result
}

// AmountMappingProcessor 金额映射处理器
type AmountMappingProcessor struct{}

func (p *AmountMappingProcessor) ProcessAmount(config types.RuleConfig, baseChannelMap map[uint]types.ChannelAmount) []types.ChannelAmount {
	wc := config.(*AmountMappingRuleConfig)
	numericMapping := wc.ToNumeric()
	channelAmountMap := make(map[uint][]int)
	var result []types.ChannelAmount

	// 按通道ID分组金额
	for amount, channelID := range numericMapping {
		channelAmountMap[channelID] = append(channelAmountMap[channelID], amount)
	}

	// 处理每个通道的金额
	for channelID, amounts := range channelAmountMap {
		if baseChannel, exists := baseChannelMap[channelID]; exists {
			options := make([]types.AmountOptionResp, len(amounts))
			for i, amount := range amounts {
				options[i] = types.AmountOptionResp{
					Value: amount,
					Label: amount,
				}
			}
			result = append(result, types.ChannelAmount{
				ChannelName: baseChannel.ChannelName,
				Options:     options,
			})
		}
	}
	return result
}

// MixedProcessor 混合权重处理器
type MixedProcessor struct{}

func (p *MixedProcessor) ProcessAmount(config types.RuleConfig, baseChannelMap map[uint]types.ChannelAmount) []types.ChannelAmount {
	wc := config.(*MixedRuleConfig)
	numericWeights := wc.ToNumeric()
	channelWeightMap := make(map[uint][]types.AmountOptionResp)
	var result []types.ChannelAmount
	// 处理金额和权重
	for amount, weightMap := range numericWeights {
		hasValidWeight := false
		for _, weight := range weightMap {
			if weight > 0 {
				hasValidWeight = true
				break
			}
		}
		if !hasValidWeight {
			continue
		}
		for channelID, weight := range weightMap {
			if weight > 0 {
				option := types.AmountOptionResp{
					Value: amount,
					Label: amount,
				}
				channelWeightMap[uint(channelID)] = append(channelWeightMap[uint(channelID)], option)
			}
		}
	}
	// 生成结果
	for channelID, options := range channelWeightMap {
		if baseChannel, exists := baseChannelMap[channelID]; exists {
			sortOptionsByAmount(options)
			result = append(result, types.ChannelAmount{
				ChannelName: baseChannel.ChannelName,
				Options:     options,
			})
		}
	}
	return result
}

func sortOptionsByAmount(options []types.AmountOptionResp) {
	sort.Slice(options, func(i, j int) bool {
		return options[i].Value < options[j].Value
	})
}
