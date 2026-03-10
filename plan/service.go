package plan

import (
	"context"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/plan/pack"
	"github.com/k9xR7vA2/recharge-common/plan/types"
)

// Service 套餐服务（对外统一接口）
type Service struct {
}

// NewService 创建套餐服务，可选传入 logger
func NewService() *Service {
	pack.Init()
	return &Service{}
}

// GetPacks 获取套餐列表
func (s *Service) GetPacks(ctx context.Context, providerCode types.PackSource, phoneNumber string) ([]types.UnifiedPack, error) {
	// 从 pack service 获取
	packs, err := pack.GetPacks(ctx, providerCode, phoneNumber)
	if err != nil {
		return nil, err
	}
	return packs, nil
}

// CheckAmountExists 检查金额是否存在
func (s *Service) CheckAmountExists(ctx context.Context, providerCode types.PackSource, phoneNumber string, amount string) (*types.AmountCheckResult, error) {
	packs, err := s.GetPacks(ctx, providerCode, phoneNumber)
	if err != nil {
		return nil, err
	}

	var matched []types.UnifiedPack
	for _, p := range packs {
		if p.Amount == amount {
			matched = append(matched, p)
		}
	}

	return &types.AmountCheckResult{
		Exists:       len(matched) > 0,
		MatchedPacks: matched,
		PhoneNumber:  phoneNumber,
		Amount:       amount,
		CarrierCode:  providerCode,
	}, nil
}

// GetPackByAmount 根据金额获取套餐
func (s *Service) GetPackByAmount(ctx context.Context, providerCode types.PackSource, phoneNumber, amount string) (*types.UnifiedPack, error) {
	packs, err := s.GetPacks(ctx, providerCode, phoneNumber)
	if err != nil {
		return nil, err
	}
	for _, p := range packs {
		if p.Amount == amount {
			return &p, nil
		}
	}
	return nil, nil
}

// GetSupportedCarriers 获取支持的运营商列表
func (s *Service) GetSupportedCarriers() []types.PackSource {
	services := pack.GetAllServices()
	providers := make([]types.PackSource, 0, len(services))
	for code := range services {
		providers = append(providers, code)
	}
	return providers
}

// BuildPaymentPayload 构建支付参数（根据运营商类型）
func (s *Service) BuildPaymentPayload(pack *types.UnifiedPack, phoneNumber string) (interface{}, error) {
	switch pack.CarrierCode {
	case types.SourceAirtel:
		return pack.BuildAirtelPaymentPayload(phoneNumber), nil
	case types.SourceJio:
		// Jio 使用 Key 字段
		return map[string]interface{}{
			"phoneNumber": phoneNumber,
			"key":         pack.Key,
			"amount":      pack.Amount,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported carrier: %s", pack.CarrierCode)
	}
}
