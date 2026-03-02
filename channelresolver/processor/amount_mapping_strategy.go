package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/constant"
	"sort"
	"strconv"
)

// AmountMappingRuleConfig 金额映射配置
// {"30": 1, "50": 2, "100": 1, "200": 2, "300": 1, "500": 2}
type AmountMappingRuleConfig struct {
	AmountChannelMap constant.AmountMappingRuleConfig // key: amount, value: channelID
}

func (c *AmountMappingRuleConfig) UnmarshalJSON(data []byte) error {
	// 直接解析为 map[string]uint
	return json.Unmarshal(data, &c.AmountChannelMap)
}

func (c *AmountMappingRuleConfig) Validate() error {
	if len(c.AmountChannelMap) == 0 {
		return errors.New("未配置金额映射")
	}
	for amountStr, channelID := range c.AmountChannelMap {
		// 将字符串金额转换为数字
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			return fmt.Errorf("金额格式无效: %s", amountStr)
		}
		if amount <= 0 {
			return errors.New("金额配置无效")
		}
		if channelID == 0 {
			return errors.New("通道ID不能为空")
		}
	}
	return nil
}

func (c *AmountMappingRuleConfig) ToNumeric() map[int]uint {
	numericMapping := make(map[int]uint)
	for k, v := range c.AmountChannelMap {
		amount, _ := strconv.Atoi(k)
		numericMapping[amount] = v
	}
	return numericMapping
}

// ExtractChannelID 根据金额直接映射到通道
func (c *AmountMappingRuleConfig) ExtractChannelID(amount string) (uint, error) {
	channelID, exists := c.AmountChannelMap[amount]
	if !exists {
		return 0, errors.New("no channel mapping for the specified amount")
	}
	return channelID, nil
}

// GetSupportedAmounts 获取策略支持的所有金额
func (c *AmountMappingRuleConfig) GetSupportedAmounts() []int {
	// 从映射配置中提取所有金额
	uniqueAmounts := make(map[int]bool)
	for amountStr := range c.AmountChannelMap {
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
func (c *AmountMappingRuleConfig) GetAllChannelIDs() []int {
	// 从映射配置中提取所有通道ID并去重
	uniqueChannelIDs := make(map[uint]bool)
	for _, channelID := range c.AmountChannelMap {
		uniqueChannelIDs[channelID] = true
	}

	// 将去重后的通道ID转换为切片
	channelIDs := make([]int, 0, len(uniqueChannelIDs))
	for channelID := range uniqueChannelIDs {

		channelIDs = append(channelIDs, int(channelID))
	}
	sort.Ints(channelIDs) // 排序以便结果一致
	return channelIDs
}
