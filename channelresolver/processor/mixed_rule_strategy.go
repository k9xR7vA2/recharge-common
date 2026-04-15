package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/constant"
	"math/rand"
	"sort"
	"strconv"
)

// MixedRuleConfig 混合权重模式配置
// {"10":{"1":0,"2":0},"20":{"1":0,"2":0},"30":{"1":100,"2":0}}
type MixedRuleConfig struct {
	AmountWeights constant.MixedRuleConfig // key1: amount, key2: channelID, value: weight
}

func (c *MixedRuleConfig) UnmarshalJSON(data []byte) error {
	// 直接解析为嵌套的 map
	return json.Unmarshal(data, &c.AmountWeights)
}

func (c *MixedRuleConfig) Validate() error {
	if len(c.AmountWeights) == 0 {
		return errors.New("未配置金额区间权重")
	}
	for amountStr, weights := range c.AmountWeights {
		// 将字符串金额转换为数字
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			return fmt.Errorf("金额格式无效: %s", amountStr)
		}
		if amount <= 0 {
			return errors.New("金额配置无效")
		}
		// 检查每个金额区间的权重总和
		totalWeight := 0
		for _, weight := range weights {
			if weight < 0 {
				return errors.New("权重值不能小于0")
			}
			totalWeight += weight
		}

		// 如果配置了权重，总和必须为100
		if totalWeight > 0 && totalWeight != 100 {
			return fmt.Errorf("金额 %s 的权重总和必须为100", amount)
		}
	}
	return nil
}

func (c *MixedRuleConfig) ToNumeric() map[int]map[int]int {
	numericWeights := make(map[int]map[int]int)
	for amount, weights := range c.AmountWeights {
		amountNum, _ := strconv.Atoi(amount)
		numericWeights[amountNum] = make(map[int]int)
		for channelID, weight := range weights {
			channelNum, _ := strconv.Atoi(channelID)
			numericWeights[amountNum][channelNum] = weight
		}
	}
	return numericWeights
}

// ExtractChannelID 先根据金额找到对应的权重配置，再根据权重选择通道
func (c *MixedRuleConfig) ExtractChannelID(amount string) (uint, error) {
	// 获取该金额对应的权重配置
	weightConfig, exists := c.AmountWeights[amount]
	if !exists {
		return 0, errors.New("no weight configuration for the specified amount")
	}

	if len(weightConfig) == 0 {
		return 0, errors.New("empty weight configuration for the specified amount")
	}

	// 计算总权重
	totalWeight := 0
	for _, weight := range weightConfig {
		totalWeight += weight
	}

	if totalWeight <= 0 {
		return 0, errors.New("total weight must be positive for the specified amount")
	}

	// 随机数在总权重范围内
	randomNum := rand.Intn(totalWeight)

	// 根据权重区间选择通道
	currentWeight := 0
	for channelID, weight := range weightConfig {
		currentWeight += weight
		if randomNum < currentWeight {
			cid, _ := strconv.Atoi(channelID)
			return uint(cid), nil
		}
	}

	// 理论上不应该到这里
	return 0, errors.New("failed to select channel based on mixed rule")
}

// GetSupportedAmounts 获取策略支持的所有金额
func (c *MixedRuleConfig) GetSupportedAmounts() []int {
	// 从混合配置中提取所有金额
	uniqueAmounts := make(map[int]bool)
	for amountStr := range c.AmountWeights {
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			continue // 忽略不能转换为整数的金额
		}
		uniqueAmounts[amount] = true
	}

	// 将去重后的金额转换为切片
	result := make([]int, 0, len(uniqueAmounts))
	for amount := range uniqueAmounts {
		result = append(result, amount)
	}

	sort.Ints(result) // 排序
	return result
}

// GetAllChannelIDs 提取所有通道ID
func (c *MixedRuleConfig) GetAllChannelIDs() []int {
	// 从混合配置中提取所有通道ID并去重
	uniqueChannelIDs := make(map[string]bool)

	for _, weightConfig := range c.AmountWeights {
		for channelID := range weightConfig {
			uniqueChannelIDs[channelID] = true
		}
	}

	// 将去重后的通道ID转换为切片
	channelIDs := make([]int, 0, len(uniqueChannelIDs))
	for channelID := range uniqueChannelIDs {
		cid, _ := strconv.Atoi(channelID)
		channelIDs = append(channelIDs, cid)
	}
	sort.Ints(channelIDs) // 排序以便结果一致
	return channelIDs
}

func (c *MixedRuleConfig) ExtractAllChannelIDsSortedByWeight() ([]uint, error) {
	return nil, errors.New("no valid channel with positive weight")
}
