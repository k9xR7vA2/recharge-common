package pack

import (
	"context"
	"github.com/k9xR7vA2/recharge-common/plan/types"
)

// PackService 套餐服务接口（只负责获取数据）
type PackService interface {
	// GetPackSource 获取服务编码
	GetPackSource() types.PackSource

	// GetPacks 获取套餐列表（从第三方 API）
	GetPacks(ctx context.Context, phoneNumber string) ([]types.UnifiedPack, error)

	// CheckAmountExists 检查金额是否存在
	CheckAmountExists(ctx context.Context, phoneNumber, amount string) (bool, []types.UnifiedPack, error)

	// GetPackByAmount 根据金额获取套餐
	GetPackByAmount(ctx context.Context, phoneNumber, amount string) (*types.UnifiedPack, error)
}
