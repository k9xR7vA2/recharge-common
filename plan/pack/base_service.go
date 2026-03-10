package pack

import (
	"github.com/k9xR7vA2/recharge-common/plan/types"
)

// BasePackService 基础套餐服务（不再管理缓存）
type BasePackService struct {
	providerCode types.PackSource
}

// NewBasePackService 创建基础服务
func NewBasePackService(code types.PackSource) *BasePackService {
	return &BasePackService{
		providerCode: code,
	}
}

func (s *BasePackService) GetPackSource() types.PackSource {
	return s.providerCode
}

// CheckAmountExistsFromPacks 通用金额检查（子类可复用）
func (s *BasePackService) CheckAmountExistsFromPacks(packs []types.UnifiedPack, amount string) (bool, []types.UnifiedPack) {
	var matched []types.UnifiedPack
	for _, pack := range packs {
		if pack.Amount == amount {
			matched = append(matched, pack)
		}
	}
	return len(matched) > 0, matched
}

// GetPackByAmountFromPacks 通用根据金额获取（子类可复用）
func (s *BasePackService) GetPackByAmountFromPacks(packs []types.UnifiedPack, amount string) *types.UnifiedPack {
	for _, pack := range packs {
		if pack.Amount == amount {
			return &pack
		}
	}
	return nil
}
