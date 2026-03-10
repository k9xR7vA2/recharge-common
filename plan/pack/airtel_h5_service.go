package pack

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/plan/client"
	"github.com/k9xR7vA2/recharge-common/plan/types"
	"strings"
	"time"
)

const airtelPacksAPIURL = "https://digi-api.airtel.in/airtel-selfcare/rest/services/common/recharge/v2/packs"

// AirtelPacksResponse Airtel 套餐响应
type AirtelPacksResponse struct {
	SiNumber string               `json:"siNumber"`
	Lob      string               `json:"lob"`
	IsAirtel bool                 `json:"isAirtel"`
	Packs    []AirtelPackCategory `json:"packs"`
}

// AirtelPackCategory 套餐分类
type AirtelPackCategory struct {
	Category      string       `json:"category"`
	NewCategory   bool         `json:"newCategory"`
	CategoryPacks []AirtelPack `json:"categoryPacks"`
}

// AirtelPack 单个套餐
type AirtelPack struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	OriginalPrice float64 `json:"originalPrice"`
	Price         float64 `json:"price"`
	Type          string  `json:"type"`
	Validity      string  `json:"validity"`
	PackID        string  `json:"packId"`
	CalloutText   string  `json:"calloutText,omitempty"`
	Hidden        bool    `json:"hidden"`
	Benefits      struct {
		ThanksBenefits []struct {
			Name        string `json:"name"`
			HeroBenefit bool   `json:"heroBenefit"`
		} `json:"thanksBenefits"`
		CoreBenefits []struct {
			Type         string `json:"type"`
			FinalBenefit struct {
				Quota    string `json:"quota"`
				Unit     string `json:"unit,omitempty"`
				Duration string `json:"duration,omitempty"`
			} `json:"finalBenefit"`
		} `json:"coreBenefits"`
		BenefitDescription struct {
			ShowText bool   `json:"showText"`
			Text     string `json:"text"`
		} `json:"benefitDescription,omitempty"`
	} `json:"benefits"`
	OrderDetailReq struct {
		SiNumber string `json:"siNumber"`
		Lob      string `json:"lob"`
		PackID   string `json:"packId"`
	} `json:"orderDetailReq"`
	OptimisedCheckoutReq struct {
		SiNumber           string      `json:"siNumber"`
		PaymentLob         string      `json:"paymentLob"`
		TotalAmount        string      `json:"totalAmount"`
		UndiscountedAmount string      `json:"undiscountedAmount"`
		Subtitle           []string    `json:"subtitle"`
		ProductCategory    interface{} `json:"productCategory"`
		Number             string      `json:"number"`
		RequestorName      string      `json:"requestorName"`
		AllowedPayModes    interface{} `json:"allowedPayModes"`
		Order              struct {
			Id     string `json:"id"`
			Url    string `json:"url"`
			Header struct {
			} `json:"header"`
			Request struct {
				Items []struct {
					Autopay        interface{} `json:"autopay"`
					OrderDetailReq struct {
						SiNumber    string `json:"siNumber"`
						Lob         string `json:"lob"`
						PackId      string `json:"packId"`
						WinBackFlow bool   `json:"winBackFlow"`
						ApiKeyMap   struct {
							Biller interface{} `json:"biller"`
						} `json:"apiKeyMap"`
						DirectPay bool `json:"directPay"`
						Fwa       bool `json:"fwa"`
					} `json:"orderDetailReq"`
				} `json:"items"`
				JourneyType string      `json:"journeyType"`
				Rtn         interface{} `json:"rtn"`
				Fwa         bool        `json:"fwa"`
			} `json:"request"`
		} `json:"order"`
	} `json:"optimisedCheckoutReq"`
}

// AirtelH5PackService Airtel H5 套餐服务
type AirtelH5PackService struct {
	*BasePackService
	client *client.AirtelClient
}

// NewAirtelH5PackService 创建服务
func NewAirtelH5PackService() *AirtelH5PackService {
	return &AirtelH5PackService{
		BasePackService: NewBasePackService(types.SourceAirtel),
		client:          client.NewAirtelClient(30 * time.Second),
	}
}

func (s *AirtelH5PackService) GetPacks(ctx context.Context, phoneNumber string) ([]types.UnifiedPack, error) {

	// 请求 API
	params := map[string]string{
		"siNumber": phoneNumber,
		"lob":      "PREPAID",
		"jk10":     "false",
	}

	data, err := s.client.GetDecrypted(ctx, airtelPacksAPIURL, params)
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

func (s *AirtelH5PackService) CheckAmountExists(ctx context.Context, phoneNumber, amount string) (bool, []types.UnifiedPack, error) {
	packs, err := s.GetPacks(ctx, phoneNumber)
	if err != nil {
		return false, nil, err
	}
	exists, matched := s.CheckAmountExistsFromPacks(packs, amount)
	return exists, matched, nil
}

func (s *AirtelH5PackService) GetPackByAmount(ctx context.Context, phoneNumber, amount string) (*types.UnifiedPack, error) {
	packs, err := s.GetPacks(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}
	return s.GetPackByAmountFromPacks(packs, amount), nil
}

func (s *AirtelH5PackService) parsePacks(data []byte) ([]types.UnifiedPack, error) {
	var resp AirtelPacksResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("parse failed: %w", err)
	}

	var packs []types.UnifiedPack
	// 增加去重容器
	seenPackIDs := make(map[string]bool)
	for _, category := range resp.Packs {
		for _, pack := range category.CategoryPacks {
			if pack.Hidden {
				continue
			}
			// 如果 PackID 已经处理过，直接跳过
			if pack.PackID != "" {
				if seenPackIDs[pack.PackID] {
					continue
				}
				seenPackIDs[pack.PackID] = true
			}

			dataStr := s.extractDataBenefit(pack)
			subHeading := s.buildSubHeading(pack)
			amount := pack.OptimisedCheckoutReq.TotalAmount
			if amount == "" {
				amount = fmt.Sprint(pack.Price) // 兜底：279 -> "279"
			}
			orderPackID := pack.OrderDetailReq.PackID
			if orderPackID == "" {
				orderPackID = pack.ID
			}
			packs = append(packs, types.UnifiedPack{
				PackID:      pack.PackID,
				Amount:      amount,
				Description: pack.Name,
				Validity:    pack.Validity,
				Category:    category.Category,
				Type:        pack.Type,
				Data:        dataStr,
				CarrierCode: s.GetPackSource(),
				// Airtel 下单信息
				AirtelOrderInfo: &types.AirtelOrderInfo{
					PackID:     orderPackID,
					Type:       pack.Type,
					SubHeading: subHeading,
					GstEnabled: false,
				},

				RawData: pack,
			})
		}
	}

	return packs, nil
}

// buildSubHeading 构建 subHeading
func (s *AirtelH5PackService) buildSubHeading(pack AirtelPack) []string {
	var subHeading []string
	// 从 name 解析，格式: "Data : 1GB | Validity : 1 day"
	if pack.Name != "" {
		parts := strings.Split(pack.Name, "|")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			lower := strings.ToLower(part)
			kv := strings.Split(part, ":")
			if len(kv) < 2 {
				continue // 如果没有冒号，跳过当前片段，防止崩溃
			}
			val := strings.TrimSpace(kv[1])
			if strings.Contains(lower, "validity") {
				// "Validity : 28 days" -> "28 days validity"
				subHeading = append(subHeading, val+" validity")
			} else if strings.Contains(lower, "data") {
				// "Data : 1GB" -> "1GB data"
				subHeading = append(subHeading, val+" data")
			}
		}
	}

	// 备用：从 benefits 构建
	if len(subHeading) == 0 {
		if pack.Validity != "" {
			subHeading = append(subHeading, pack.Validity+" validity")
		}
		dataStr := s.extractDataBenefit(pack)
		if dataStr != "" {
			subHeading = append(subHeading, dataStr+" data")
		}
	}

	return subHeading
}

// extractDataBenefit 提取数据流量信息
func (s *AirtelH5PackService) extractDataBenefit(pack AirtelPack) string {
	for _, benefit := range pack.Benefits.CoreBenefits {
		if benefit.Type == "DATA" {
			quota := benefit.FinalBenefit.Quota
			unit := benefit.FinalBenefit.Unit
			duration := benefit.FinalBenefit.Duration

			if unit != "" {
				if duration == "PER_DAY" {
					return fmt.Sprintf("%s%s/day", quota, unit)
				}
				return fmt.Sprintf("%s%s", quota, unit)
			}
			return quota // "Unlimited" 等情况
		}
	}
	return ""
}
