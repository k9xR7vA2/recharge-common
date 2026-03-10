package plan

import (
	"SaasApi/pkg/plan/types"
	"sort"
	"strconv"
	"strings"
)

type FinalOutput struct {
	Amount string   `json:"amount"`
	Packs  []string `json:"packs"`
}

func GetOutputPacks(packs []types.UnifiedPack) []FinalOutput {
	resultMap := make(map[string][]string)
	var amountKeys []float64

	// 记录已经添加过的金额，防止 amountKeys 重复
	seen := make(map[float64]bool)

	for _, item := range packs {
		amt, _ := strconv.ParseFloat(item.Amount, 64)
		// 重要：统一格式化 Key，确保存和取的字符串完全一致
		// 'f', -1 会把 589.0 变成 "589"，把 589.5 变成 "589.5"
		cleanAmtStr := strconv.FormatFloat(amt, 'f', -1, 64)
		if !seen[amt] {
			amountKeys = append(amountKeys, amt)
			seen[amt] = true
		}
		resultMap[cleanAmtStr] = append(resultMap[cleanAmtStr], item.PackID)
		//// 根据 source 归类
		//var val string
		//if packSource == types.SourceJioH5 {
		//	val = item.Key
		//} else if packSource == types.SourceAirtelH5 {
		//	val = item.PackID
		//}
		//if val != "" {
		//	// 使用统一后的 cleanAmtStr 作为 Key
		//}
	}
	// 2. 排序
	sort.Float64s(amountKeys)

	// 3. 构造输出
	var sortedResponse []FinalOutput
	for _, amt := range amountKeys {
		amtStr := strconv.FormatFloat(amt, 'f', -1, 64)
		sortedResponse = append(sortedResponse, FinalOutput{
			Amount: amtStr,
			Packs:  resultMap[amtStr], // 现在这里一定能取到值了
		})
	}
	return sortedResponse
}

type SimplePlan struct {
	PlanID   string `json:"plan_id"`
	Amount   string `json:"amount"`
	Desc     string `json:"desc"`
	Category string `json:"category"`
}

func GetSimplePlans(packs []types.UnifiedPack) []SimplePlan {
	var result []SimplePlan
	for _, p := range packs {

		// 优先用 AirtelOrderInfo 的 SubHeading 拼描述，否则用 Validity + Data
		desc := p.Description
		if p.AirtelOrderInfo != nil && len(p.AirtelOrderInfo.SubHeading) > 0 {
			desc = strings.Join(p.AirtelOrderInfo.SubHeading, ", ")
		} else if desc == "" {
			desc = p.Validity + ", " + p.Data
		}

		result = append(result, SimplePlan{
			PlanID: p.PackID,
			Amount: p.Amount,
			Desc:   desc,
		})
	}

	// 按金额排序
	//sort.Slice(result, func(i, j int) bool {
	//	return result[i].Amount < result[j].Amount
	//})

	return result
}

type PlanGroup struct {
	Category string       `json:"category"`
	Plans    []SimplePlan `json:"plans"`
}

func buildDesc(p types.UnifiedPack) string {
	// 优先用 AirtelOrderInfo 的 SubHeading
	if p.AirtelOrderInfo != nil && len(p.AirtelOrderInfo.SubHeading) > 0 {
		return strings.Join(p.AirtelOrderInfo.SubHeading, ", ")
	}
	// 其次拼 Data + Validity
	var parts []string
	if p.Data != "" {
		parts = append(parts, p.Data+" data")
	}
	if p.Validity != "" {
		parts = append(parts, p.Validity+" validity")
	}
	if len(parts) > 0 {
		return strings.Join(parts, ", ")
	}
	// 兜底用原始 Description（截断避免太长）
	desc := p.Description
	if len(desc) > 100 {
		desc = desc[:100] + "..."
	}
	return desc
}
func GetGroupedPlans(packs []types.UnifiedPack) []PlanGroup {
	groupMap := make(map[string][]SimplePlan)
	var categoryOrder []string
	seen := make(map[string]bool)

	for _, p := range packs {
		sp := SimplePlan{
			PlanID:   p.PackID,
			Amount:   p.Amount,
			Desc:     buildDesc(p),
			Category: p.Category,
		}
		if !seen[p.Category] {
			categoryOrder = append(categoryOrder, p.Category)
			seen[p.Category] = true
		}
		groupMap[p.Category] = append(groupMap[p.Category], sp)
	}

	var result []PlanGroup
	for _, cat := range categoryOrder {
		result = append(result, PlanGroup{
			Category: cat,
			Plans:    groupMap[cat],
		})
	}
	return result
}
