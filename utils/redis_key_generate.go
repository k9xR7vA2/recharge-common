package utils

import (
	"fmt"
	"github.com/k9xR7vA2/recharge-common/constant"
	"strings"
	"time"
)

// StatsEntity 统计实体配置
type StatsEntity struct {
	Type         string // 实体类型
	ID           uint   // 实体ID
	ChannelType  int    // 通道类型（如果有）
	BusinessType int    // 业务类型（如果有）
}

// RedisKeyGenerator Redis键生成器
type RedisKeyGenerator struct {
	prefix string
}

// NewRedisKeyGenerator 创建新的Redis键生成器
func NewRedisKeyGenerator() *RedisKeyGenerator {
	return &RedisKeyGenerator{
		prefix: "stats", // 统一前缀
	}
}

// FormatDate 格式化日期为YYYY-MM-DD
func (g *RedisKeyGenerator) FormatDate(date time.Time) string {
	return date.Format("20060102")
}

// GenerateBaseKey 生成基础统计键
// entityType: 实体类型 (merchant, supplier, channel, product, all_merchants等)
// dateStr: 日期字符串
// tenantID: 租户ID
// entityID: 实体ID (可选，全局统计为0)
func (g *RedisKeyGenerator) GenerateBaseKey(entityType, dateStr string, tenantID uint, entityID uint) string {
	if entityType == "" {
		return ""
	}

	parts := []string{g.prefix, entityType, dateStr, fmt.Sprintf("%d", tenantID)}

	if entityID > 0 {
		parts = append(parts, fmt.Sprintf("%d", entityID))
	}

	return strings.Join(parts, ":")
}

// GenerateChannelTypeKey 生成带通道类型的统计键
// baseKey: 基础键
// channelType: 通道类型(1:基础通道, 2:混合通道)
func (g *RedisKeyGenerator) GenerateChannelTypeKey(baseKey string, channelType constant.ChannelType) string {
	if baseKey == "" || channelType <= 0 {
		return baseKey
	}

	return fmt.Sprintf("%s:%d", baseKey, channelType)
}

// GenerateBusinessTypeKey 生成业务类型统计键
// entityType: 实体类型 (channel, product, merchant_channel, supplier_product等)
// dateStr: 日期字符串
// tenantID: 租户ID
// entityID: 实体ID
// businessType: 业务类型ID
// relatedID: 关联实体ID（可选，适用于merchant_channel或supplier_product）
func (g *RedisKeyGenerator) GenerateBusinessTypeKey(entityType, dateStr string, tenantID, entityID uint, businessType constant.BusinessType, relatedID ...uint) string {
	// 添加_business后缀标识业务类型统计
	businessEntityType := fmt.Sprintf("%s_business", entityType)
	parts := []string{g.prefix, businessEntityType, dateStr, fmt.Sprintf("%d", tenantID)}
	// 对于关联类型，先添加主实体ID，再添加次级实体ID
	if len(relatedID) > 0 && relatedID[0] > 0 {
		parts = append(parts, fmt.Sprintf("%d", relatedID[0]), fmt.Sprintf("%d", entityID))
	} else {
		parts = append(parts, fmt.Sprintf("%d", entityID))
	}

	// 添加业务类型
	parts = append(parts, fmt.Sprintf("%s", businessType))

	return strings.Join(parts, ":")
}

// GenerateRelationKey 生成关联实体统计键
// relationType: 关联类型 (merchant_channel, supplier_product等)
// dateStr: 日期字符串
// tenantID: 租户ID
// primaryID: 主实体ID (商户ID或供货商ID)
// secondaryID: 次级实体ID (通道ID或产品ID)
// channelType: 通道类型（可选，适用于merchant_channel）
func (g *RedisKeyGenerator) GenerateRelationKey(relationType, dateStr string, tenantID, primaryID, secondaryID uint, channelType ...constant.ChannelType) string {
	if relationType == "" || primaryID <= 0 || secondaryID <= 0 {
		return ""
	}

	parts := []string{g.prefix, relationType, dateStr, fmt.Sprintf("%d", tenantID),
		fmt.Sprintf("%d", primaryID), fmt.Sprintf("%d", secondaryID)}

	// 如果指定了通道类型
	if len(channelType) > 0 && channelType[0] > 0 {
		parts = append(parts, fmt.Sprintf("%d", channelType[0]))
	}

	return strings.Join(parts, ":")
}

// GenerateRelationBusinessKey 生成关联实体业务类型统计键
// relationType: 关联类型 (merchant_channel, supplier_product等)
// dateStr: 日期字符串
// tenantID: 租户ID
// primaryID: 主实体ID (商户ID或供货商ID)
// secondaryID: 次级实体ID (通道ID或产品ID)
// businessType: 业务类型ID
func (g *RedisKeyGenerator) GenerateRelationBusinessKey(relationType, dateStr string, tenantID, primaryID, secondaryID uint, businessType constant.BusinessType) string {
	if relationType == "" || primaryID <= 0 || secondaryID <= 0 || !businessType.IsValid() {
		return ""
	}

	// 添加_business后缀标识业务类型统计
	businessRelationType := fmt.Sprintf("%s_business", relationType)

	parts := []string{g.prefix, businessRelationType, dateStr, fmt.Sprintf("%d", tenantID),
		fmt.Sprintf("%d", primaryID), fmt.Sprintf("%d", secondaryID),
		fmt.Sprintf("%s", businessType)}

	return strings.Join(parts, ":")
}

// GenerateAmountKey 生成金额维度统计键
func (g *RedisKeyGenerator) GenerateAmountKey(
	baseKey string,
	amount uint,
) string {
	if baseKey == "" || amount == 0 {
		return ""
	}
	return baseKey + ":" + "amount" + ":" + fmt.Sprintf("%d", amount)
}

// GenerateEntityAmountKey 生成实体金额维度统计键
func (g *RedisKeyGenerator) GenerateEntityAmountKey(
	entityType string,
	dateStr string,
	tenantID uint,
	entityID uint,
	amount uint,
) string {
	baseKey := g.GenerateBaseKey(entityType, dateStr, tenantID, entityID)
	return g.GenerateAmountKey(baseKey, amount)
}

// GenerateChannelTypeAmountKey 生成通道类型金额维度统计键
func (g *RedisKeyGenerator) GenerateChannelTypeAmountKey(
	baseKey string,
	channelType constant.ChannelType,
	amount uint,
) string {
	if baseKey == "" || channelType <= 0 || amount == 0 {
		return ""
	}
	typeKey := g.GenerateChannelTypeKey(baseKey, channelType)
	return g.GenerateAmountKey(typeKey, amount)
}

// GenerateBusinessTypeAmountKey 生成业务类型金额维度统计键
func (g *RedisKeyGenerator) GenerateBusinessTypeAmountKey(
	entityType string,
	dateStr string,
	tenantID uint,
	entityID uint,
	businessType constant.BusinessType,
	amount uint,
) string {
	if !businessType.IsValid() || amount == 0 {
		return ""
	}
	businessKey := g.GenerateBusinessTypeKey(entityType, dateStr, tenantID, entityID, businessType)
	return g.GenerateAmountKey(businessKey, amount)
}

// GenerateRelationAmountKey 生成关联实体金额维度统计键
func (g *RedisKeyGenerator) GenerateRelationAmountKey(
	relationType string,
	dateStr string,
	tenantID uint,
	primaryID uint,
	secondaryID uint,
	amount uint,
) string {
	if amount == 0 {
		return ""
	}
	relationKey := g.GenerateRelationKey(relationType, dateStr, tenantID, primaryID, secondaryID)
	return g.GenerateAmountKey(relationKey, amount)
}

// GenerateRelationTypeAmountKey 生成关联实体类型金额维度统计键
func (g *RedisKeyGenerator) GenerateRelationTypeAmountKey(
	relationType string,
	dateStr string,
	tenantID uint,
	primaryID uint,
	secondaryID uint,
	channelType constant.ChannelType,
	amount uint,
) string {
	if channelType <= 0 || amount == 0 {
		return ""
	}
	typeKey := g.GenerateRelationKey(relationType, dateStr, tenantID, primaryID, secondaryID, channelType)
	return g.GenerateAmountKey(typeKey, amount)
}

// GenerateRelationBusinessAmountKey 生成关联实体业务类型金额维度统计键
func (g *RedisKeyGenerator) GenerateRelationBusinessAmountKey(
	relationType string,
	dateStr string,
	tenantID uint,
	primaryID uint,
	secondaryID uint,
	businessType constant.BusinessType,
	amount uint,
) string {
	if !businessType.IsValid() || amount == 0 {
		return ""
	}
	businessKey := g.GenerateRelationBusinessKey(relationType, dateStr, tenantID, primaryID, secondaryID, businessType)
	return g.GenerateAmountKey(businessKey, amount)
}

// GetHashFields 获取统计数据的Hash字段名
// 返回总订单数、成功订单数和成功金额的字段名
func (g *RedisKeyGenerator) GetHashFields() (string, string, string) {
	return "total_orders", "success_orders", "success_amount"
}
