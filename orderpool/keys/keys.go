package keys

import (
	"fmt"
	"strings"
)

// KeyBuilder Redis key 构建器
type KeyBuilder struct {
	components []string
}

// NewKeyBuilder 创建新的 KeyBuilder，只添加租户前缀
func NewKeyBuilder(tenantId uint) *KeyBuilder {
	return &KeyBuilder{
		components: []string{fmt.Sprintf("tenant_%d", tenantId)},
	}
}

// Add 添加 key 组件
func (kb *KeyBuilder) Add(component interface{}) *KeyBuilder {
	if component != nil {
		kb.components = append(kb.components, fmt.Sprint(component))
	}
	return kb
}

// Build 构建最终的 key
func (kb *KeyBuilder) Build() string {
	return strings.Join(kb.components, KeySeparator)
}
