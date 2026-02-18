package keys

import (
	"fmt"
	"github.com/small-cat1/recharge-common/constant"
	"strconv"
	"strings"
)

// KeyParser Redis key 解析器
type KeyParser struct {
	components []string
}

// NewKeyParser 创建新的 KeyParser
func NewKeyParser(key string) *KeyParser {
	return &KeyParser{
		components: strings.Split(key, KeySeparator),
	}
}

// ParsedKey 解析后的 key 结构
type ParsedKey struct {
	TenantID     uint
	Role         string
	BusinessType string
	Type         string    // pool/order/stats/payment
	Details      KeyDetail // 不同类型的详细信息
}

// KeyDetail 接口定义不同类型的 key 详情
type KeyDetail interface {
	KeyType() string
}

// PoolKeyDetail pool类型的 key 详情
type PoolKeyDetail struct {
	Amount     string
	Carrier    string
	ChargeType string
	Region     string
	Province   string
}

func (p PoolKeyDetail) KeyType() string {
	return TypePool
}

// OrderKeyDetail order类型的 key 详情
type OrderKeyDetail struct {
	OrderSn string
}

func (o OrderKeyDetail) KeyType() string {
	return TypeOrder
}

// Parse 解析 key
func (kp *KeyParser) Parse() (*ParsedKey, error) {
	if len(kp.components) < 3 {
		return nil, fmt.Errorf("invalid key format: insufficient components")
	}

	// 解析租户ID
	tenantIDStr := strings.TrimPrefix(kp.components[0], "tenant_")
	tenantID, err := strconv.ParseUint(tenantIDStr, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid tenant id: %s", tenantIDStr)
	}

	result := &ParsedKey{
		TenantID:     uint(tenantID),
		Role:         kp.components[1],
		BusinessType: kp.components[2],
	}

	// 至少要有类型标识
	if len(kp.components) < 4 {
		return nil, fmt.Errorf("invalid key format: missing type")
	}

	// 根据类型解析详情
	switch kp.components[3] {
	case TypePool:
		detail, err := kp.parsePoolDetail()
		if err != nil {
			return nil, err
		}
		result.Type = TypePool
		result.Details = detail

	case TypeOrder:
		if len(kp.components) != 5 {
			return nil, fmt.Errorf("invalid order key format")
		}
		result.Type = TypeOrder
		result.Details = &OrderKeyDetail{
			OrderSn: kp.components[4],
		}

	case TypeStats:
		result.Type = TypeStats

	case TypePayment:
		if len(kp.components) != 5 {
			return nil, fmt.Errorf("invalid payment key format")
		}
		result.Type = TypePayment
		result.Details = &OrderKeyDetail{
			OrderSn: kp.components[4],
		}
	}

	return result, nil
}

// parsePoolDetail 解析订单池的详细信息
func (kp *KeyParser) parsePoolDetail() (*PoolKeyDetail, error) {
	// pool key 至少需要 8 个组件
	if len(kp.components) < 8 {
		return nil, fmt.Errorf("invalid pool key format")
	}

	detail := &PoolKeyDetail{
		Amount:     kp.components[4],
		Carrier:    kp.components[5],
		ChargeType: kp.components[6],
		Region:     kp.components[7],
	}

	// 如果是省份级别，还需要包含省份信息
	if detail.Region == constant.Province.Code() {
		if len(kp.components) != 9 {
			return nil, fmt.Errorf("invalid province pool key format")
		}
		detail.Province = kp.components[8]
	}

	return detail, nil
}
