package dict

import "recharge-common/dict/types"

// DictItem 字典项接口，所有字典类型都需要实现
type DictItem interface {
	Value() interface{} // 返回值（用于前端 value 字段）
	Label() string      // 返回标签（用于前端 label 字段）
	Code() string       // 返回编码（可选，用于某些场景）
}

// DictProvider 字典提供者接口
type DictProvider interface {
	GetKey() string                 // 获取字典键名
	GetName() string                // 获取字典名称
	GetOptions() []types.DictOption // 获取所有选项
}

// Manager 字典管理器
type Manager struct {
	providers map[string]DictProvider
}

// NewManager 创建字典管理器
func NewManager() *Manager {
	return &Manager{
		providers: make(map[string]DictProvider),
	}
}

// Register 注册字典提供者
func (m *Manager) Register(provider DictProvider) {
	m.providers[provider.GetKey()] = provider
}

// Get 获取单个字典
func (m *Manager) Get(key string) *types.DictResponse {
	provider, exists := m.providers[key]
	if !exists {
		return nil
	}
	return &types.DictResponse{
		Key:     provider.GetKey(),
		Name:    provider.GetName(),
		Options: provider.GetOptions(),
	}
}

// GetAll 获取所有字典
func (m *Manager) GetAll() map[string]*types.DictResponse {
	result := make(map[string]*types.DictResponse)
	for key, provider := range m.providers {
		result[key] = &types.DictResponse{
			Key:     provider.GetKey(),
			Name:    provider.GetName(),
			Options: provider.GetOptions(),
		}
	}
	return result
}

// GetMultiple 批量获取字典
func (m *Manager) GetMultiple(keys []string) map[string]*types.DictResponse {
	result := make(map[string]*types.DictResponse)
	for _, key := range keys {
		if dict := m.Get(key); dict != nil {
			result[key] = dict
		}
	}

	return result
}

// GetKeys 获取所有字典键名
func (m *Manager) GetKeys() []string {
	keys := make([]string, 0, len(m.providers))
	for key := range m.providers {
		keys = append(keys, key)
	}
	return keys
}
