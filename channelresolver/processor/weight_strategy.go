package processor

import (
	"encoding/json"
	"errors"
	"github.com/small-cat1/recharge-common/constant"
	"math/rand"
	"sort"
	"strconv"
)

// WeightRuleConfig 权重模式配置
// {"1": 50, "2": 50}
type WeightRuleConfig struct {
	ChannelWeights constant.WeightRuleConfig // 直接映射原始数据结构
	// 权重模式不依赖于金额，但在后端可能会配置允许的金额列表
	AllowedAmounts []int
}

func (c *WeightRuleConfig) UnmarshalJSON(data []byte) error {
	// 直接解析为 map[string]int
	return json.Unmarshal(data, &c.ChannelWeights)
}

func (c *WeightRuleConfig) ToNumeric() map[int]int {
	numericWeights := make(map[int]int)
	for k, v := range c.ChannelWeights {
		channelID, _ := strconv.Atoi(k)
		numericWeights[channelID] = v
	}
	return numericWeights
}

func (c *WeightRuleConfig) Validate() error {
	if len(c.ChannelWeights) == 0 {
		return errors.New("未配置通道权重")
	}
	totalWeight := 0
	for _, weight := range c.ChannelWeights {
		if weight < 0 {
			return errors.New("权重值不能小于0")
		}
		totalWeight += weight
	}
	if totalWeight != 100 {
		return errors.New("权重总和必须为100")
	}
	return nil
}

// 根据权重随机选择通道
func (c *WeightRuleConfig) ExtractChannelID(amount string) (uint, error) {
	// 检查金额是否在允许列表中
	if len(c.AllowedAmounts) > 0 {
		amountInt, err := strconv.Atoi(amount)
		if err != nil {
			return 0, errors.New("invalid amount format")
		}

		found := false
		for _, allowed := range c.AllowedAmounts {
			if allowed == amountInt {
				found = true
				break
			}
		}

		if !found {
			return 0, errors.New("amount not allowed in this strategy")
		}
	}

	if len(c.ChannelWeights) == 0 {
		return 0, errors.New("empty weight rule config")
	}

	// 计算总权重
	totalWeight := 0
	for _, weight := range c.ChannelWeights {
		totalWeight += weight
	}
	if totalWeight <= 0 {
		return 0, errors.New("total weight must be positive")
	}
	// 随机数在总权重范围内
	randomNum := rand.Intn(totalWeight)
	// 根据权重区间选择通道
	currentWeight := 0
	for channelID, weight := range c.ChannelWeights {
		currentWeight += weight
		if randomNum < currentWeight {
			cid, _ := strconv.Atoi(channelID)
			return uint(cid), nil
		}
	}
	// 理论上不应该到这里
	return 0, errors.New("failed to select channel based on weight")
}

// GetSupportedAmounts 获取策略支持的所有金额
func (c *WeightRuleConfig) GetSupportedAmounts() []int {
	result := make([]int, len(c.AllowedAmounts))
	copy(result, c.AllowedAmounts)
	sort.Ints(result) // 排序
	return result
}

// GetAllChannelIDs 提取所有通道ID
func (c *WeightRuleConfig) GetAllChannelIDs() []int {
	channelIDs := make([]int, 0, len(c.ChannelWeights))
	for channelID := range c.ChannelWeights {
		cid, _ := strconv.Atoi(channelID)
		channelIDs = append(channelIDs, cid)
	}
	sort.Ints(channelIDs) // 排序以便结果一致
	return channelIDs
}
