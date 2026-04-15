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

// WeightRuleConfig 权重模式配置
// {"1": 50, "2": 50}
type WeightRuleConfig struct {
	ChannelWeights constant.WeightRuleConfig // 直接映射原始数据结构
	// 权重模式不依赖于金额，但在后端可能会配置允许的金额列表
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
	if len(c.ChannelWeights) == 0 {
		return 0, errors.New("empty weight rule config")
	}

	totalWeight := 0
	for _, weight := range c.ChannelWeights {
		totalWeight += weight
	}
	if totalWeight <= 0 {
		return 0, errors.New("total weight must be positive")
	}

	// 排序保证遍历顺序一致
	keys := make([]string, 0, len(c.ChannelWeights))
	for k := range c.ChannelWeights {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	randomNum := rand.Intn(totalWeight)
	currentWeight := 0
	for _, channelID := range keys {
		currentWeight += c.ChannelWeights[channelID]
		if randomNum < currentWeight {
			cid, err := strconv.Atoi(channelID)
			if err != nil {
				return 0, fmt.Errorf("无效的通道ID: %s", channelID)
			}
			return uint(cid), nil
		}
	}
	return 0, errors.New("failed to select channel based on weight")
}

// ExtractAllChannelIDsSortedByWeight 返回按权重降序排列的所有通道ID
// 权重高的排前面，下单时优先尝试权重大的通道
func (c *WeightRuleConfig) ExtractAllChannelIDsSortedByWeight() ([]uint, error) {
	if len(c.ChannelWeights) == 0 {
		return nil, errors.New("empty weight rule config")
	}

	type weightedChannel struct {
		channelID uint
		weight    int
	}

	candidates := make([]weightedChannel, 0, len(c.ChannelWeights))
	for k, w := range c.ChannelWeights {
		if w <= 0 {
			continue
		}
		cid, err := strconv.Atoi(k)
		if err != nil {
			return nil, fmt.Errorf("无效的通道ID: %s", k)
		}
		candidates = append(candidates, weightedChannel{
			channelID: uint(cid),
			weight:    w,
		})
	}

	if len(candidates) == 0 {
		return nil, errors.New("no valid channel with positive weight")
	}

	// 权重降序，权重相同时 channelID 升序保证稳定
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].weight != candidates[j].weight {
			return candidates[i].weight > candidates[j].weight
		}
		return candidates[i].channelID < candidates[j].channelID
	})

	ids := make([]uint, len(candidates))
	for i, c := range candidates {
		ids[i] = c.channelID
	}
	return ids, nil
}

// GetSupportedAmounts 获取策略支持的所有金额
func (c *WeightRuleConfig) GetSupportedAmounts() []int {
	return nil
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
