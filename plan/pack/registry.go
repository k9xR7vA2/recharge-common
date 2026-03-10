package pack

import (
	"context"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/plan/types"
	"sync"
)

var (
	registry     = make(map[types.PackSource]PackService)
	registryLock sync.RWMutex
)

// Register 注册套餐服务
func Register(service PackService) {
	registryLock.Lock()
	defer registryLock.Unlock()
	registry[service.GetPackSource()] = service
}

// GetService 获取套餐服务
func GetService(code types.PackSource) (PackService, error) {
	registryLock.RLock()
	defer registryLock.RUnlock()

	service, ok := registry[code]
	if !ok {
		return nil, fmt.Errorf("pack service not found: %s", code)
	}
	return service, nil
}

// GetAllServices 获取所有已注册的服务
func GetAllServices() map[types.PackSource]PackService {
	registryLock.RLock()
	defer registryLock.RUnlock()

	result := make(map[types.PackSource]PackService)
	for k, v := range registry {
		result[k] = v
	}
	return result
}

// ===== 便捷方法 =====

// GetPacks 根据编码获取套餐
func GetPacks(ctx context.Context, code types.PackSource, phoneNumber string) ([]types.UnifiedPack, error) {
	service, err := GetService(code)
	if err != nil {
		return nil, err
	}
	return service.GetPacks(ctx, phoneNumber)
}
