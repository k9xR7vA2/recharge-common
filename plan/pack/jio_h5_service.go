package pack

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/plan/client"
	"github.com/k9xR7vA2/recharge-common/plan/types"
	"time"
)

// JioPacksResponse Jio 套餐响应
type JioPacksResponse struct {
	PlanCategories []JioPlanCategory `json:"planCategories"`
	ProductType    string            `json:"productType"`
	PrimeMember    bool              `json:"primeMember"`
	Type           string            `json:"type"`
}

// JioPlanCategory 套餐分类
type JioPlanCategory struct {
	Type          string           `json:"type"`
	HeaderType    string           `json:"headerType"`
	SubCategories []JioSubCategory `json:"subCategories"`
}

// JioSubCategory 子分类
type JioSubCategory struct {
	Type  string    `json:"type"`
	Plans []JioPlan `json:"plans"`
}

// JioPlan 单个套餐
type JioPlan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Amount      string `json:"amount"` // 字符串，需转换
	Description string `json:"description"`
	VoucherID   string `json:"voucherId"`
	PlanName    string `json:"planName"`
	Key         string `json:"key"`
	BillingType string `json:"billingType"`

	PrimeData struct {
		Plan           string `json:"plan"`
		OfferBenefits1 string `json:"offerBenefits1"` // 数据量
		OfferBenefits2 string `json:"offerBenefits2"` // 数据单位 (GB/Day)
		OfferBenefits3 string `json:"offerBenefits3"` // 有效期
		OfferBenefits4 string `json:"offerBenefits4"` // 有效期单位 (Days)
		Subscriptions  string `json:"subscriptions"`
	} `json:"primeData"`

	FilterKeys struct {
		Data          string `json:"data"`
		Validity      string `json:"validity"`
		ValidityRange string `json:"validityRange"`
		Price         string `json:"price"`
		PriceRange    string `json:"priceRange"`
		OTT           string `json:"ott"`
	} `json:"filterKeys"`
}

// JioH5PackService Jio H5 套餐服务
type JioH5PackService struct {
	*BasePackService
}

// NewJioH5PackService 创建服务
func NewJioH5PackService() *JioH5PackService {
	return &JioH5PackService{
		BasePackService: NewBasePackService(types.SourceJio),
	}
}

func (s *JioH5PackService) GetPacks(ctx context.Context, phoneNumber string) ([]types.UnifiedPack, error) {
	// 创建客户端
	jioClient := client.NewJioClient(60 * time.Second)

	// Step 1: 初始化 session
	if err := jioClient.InitSession(ctx, phoneNumber); err != nil {
		return nil, fmt.Errorf("init session failed: %w", err)
	}

	// Step 2: 获取套餐
	url := fmt.Sprintf("https://www.jio.com/api/jio-recharge-service/recharge/plans/serviceId/%s", phoneNumber)
	referer := fmt.Sprintf("https://www.jio.com/selfcare/recharge/mobility/plans/?serviceType=mobility&serviceId=%s&next=PREPAID&billingType=PREPAID&entrysource=Widget", phoneNumber)

	data, err := jioClient.GetWithSession(ctx, url, referer)
	if err != nil {
		return nil, fmt.Errorf("get packs failed: %w", err)
	}

	// 解析
	packs, err := s.parsePacks(data)
	if err != nil {
		return nil, err
	}

	return packs, nil
}

func (s *JioH5PackService) CheckAmountExists(ctx context.Context, phoneNumber, amount string) (bool, []types.UnifiedPack, error) {
	packs, err := s.GetPacks(ctx, phoneNumber)
	if err != nil {
		return false, nil, err
	}
	exists, matched := s.CheckAmountExistsFromPacks(packs, amount)
	return exists, matched, nil
}

func (s *JioH5PackService) GetPackByAmount(ctx context.Context, phoneNumber, amount string) (*types.UnifiedPack, error) {
	packs, err := s.GetPacks(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}
	return s.GetPackByAmountFromPacks(packs, amount), nil
}

func (s *JioH5PackService) parsePacks(data []byte) ([]types.UnifiedPack, error) {
	var resp JioPacksResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("parse failed: %w", err)
	}

	packs := make([]types.UnifiedPack, 0, 256)

	for _, category := range resp.PlanCategories {
		for _, subCategory := range category.SubCategories {
			for _, subPlan := range subCategory.Plans {

				packID := subPlan.VoucherID
				if packID == "" {
					packID = subPlan.ID
				}

				packs = append(packs, types.UnifiedPack{
					PackID:      subPlan.Key,
					Amount:      subPlan.Amount, // 👈 原样 string
					Description: subPlan.Description,
					Validity:    s.extractValidity(subPlan),
					Category:    category.Type,
					SubCategory: subCategory.Type,
					Type:        subPlan.BillingType,
					Data:        s.extractData(subPlan),
					CarrierCode: s.GetPackSource(),

					Key:     subPlan.Key,
					RawData: subPlan,
				})
			}
		}
	}

	return packs, nil
}

// extractData 提取数据流量信息
func (s *JioH5PackService) extractData(plan JioPlan) string {
	benefits1 := plan.PrimeData.OfferBenefits1
	benefits2 := plan.PrimeData.OfferBenefits2

	if benefits1 != "" && benefits2 != "" {
		return benefits1 + benefits2 // 如 "2GB/Day"
	}

	// 备用：从 filterKeys 获取
	if plan.FilterKeys.Data != "" {
		return plan.FilterKeys.Data + "GB/Day"
	}

	return ""
}

// extractValidity 提取有效期
func (s *JioH5PackService) extractValidity(plan JioPlan) string {
	benefits3 := plan.PrimeData.OfferBenefits3
	benefits4 := plan.PrimeData.OfferBenefits4

	if benefits3 != "" && benefits4 != "" {
		return benefits3 + " " + benefits4 // 如 "28 Days"
	}

	// 备用：从 filterKeys 获取
	if plan.FilterKeys.Validity != "" {
		return plan.FilterKeys.Validity + " Days"
	}

	return ""
}
