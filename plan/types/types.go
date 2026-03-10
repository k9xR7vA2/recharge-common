package types

import "fmt"

// CarrierCode 运营商+渠道编码
type PackSource string

const (
	SourceJio    PackSource = "jio"
	SourceAirtel PackSource = "airtel"
	SourceVI     PackSource = "vi"
)

// UnifiedPack 统一套餐结构
type UnifiedPack struct {
	PackID          string           `json:"packId"`
	Amount          string           `json:"amount"`
	Description     string           `json:"description"`
	Validity        string           `json:"validity"`
	Category        string           `json:"category"`
	SubCategory     string           `json:"subCategory,omitempty"`
	Type            string           `json:"type"`
	Data            string           `json:"data"`
	CarrierCode     PackSource       `json:"carrierCode"`
	Key             string           `json:"key,omitempty"`
	AirtelOrderInfo *AirtelOrderInfo `json:"airtelOrderInfo,omitempty"`
	RawData         interface{}      `json:"rawData,omitempty"`
}

// AirtelOrderInfo Airtel 下单所需信息
type AirtelOrderInfo struct {
	PackID     string   `json:"packId"`
	Type       string   `json:"type"`       // International Roaming, Data, Unlimited Packs 等
	SubHeading []string `json:"subHeading"` // ["365 days validity", "5 GB data"]
	GstEnabled bool     `json:"gstEnabled"`
}

// BuildAirtelPaymentPayload 构建 Airtel 下单 payload
func (p *UnifiedPack) BuildAirtelPaymentPayload(phoneNumber string) map[string]interface{} {
	if p.AirtelOrderInfo == nil {
		return nil
	}

	amount := p.Amount
	primaryHeading := p.AirtelOrderInfo.Type
	if primaryHeading == "" {
		primaryHeading = "Prepaid"
	}

	return map[string]interface{}{
		"order": map[string]interface{}{
			"lobName":         "PREPAID",
			"paymentAmount":   amount,
			"benefitAmount":   amount,
			"serviceInstance": phoneNumber,
		},
		"billSummary": map[string]interface{}{
			"transactionText": "Prepaid | " + phoneNumber,
			"primaryHeading":  primaryHeading,
			"subHeading":      p.AirtelOrderInfo.SubHeading,
			"benefits":        []interface{}{},
			"gstEnabled":      p.AirtelOrderInfo.GstEnabled,
		},
		"callBack": map[string]interface{}{
			"redirectionUrl": "https://www.airtel.in/pay/summary?omv=v2",
			"cancelUrl":      "https://www.airtel.in/recharge/prepaid/packs",
		},
		"flow": map[string]interface{}{
			"trackingObj": map[string]interface{}{
				"utm_source":   "",
				"utm_campaign": "",
				"utm_medium":   "",
			},
		},
		"result": map[string]interface{}{
			"resultSummary": []map[string]interface{}{
				{"heading": "Prepaid", "value": amount},
			},
			"amountDetails": []map[string]interface{}{
				{"key": "Pack Base Price", "value": fmt.Sprintf("₹%.0f", amount)},
			},
		},
	}
}

// AmountCheckResult 金额检查结果
type AmountCheckResult struct {
	Exists       bool          `json:"exists"`
	MatchedPacks []UnifiedPack `json:"matchedPacks,omitempty"`
	PhoneNumber  string        `json:"phoneNumber"`
	Amount       string        `json:"amount"`
	CarrierCode  PackSource    `json:"carrierCode"`
}
