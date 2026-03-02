package utils

import (
	"fmt"
	"github.com/k9xR7vA2/recharge-common/constant"
	"regexp"
	"strconv"
)

func ValidateAmountArray(amounts []int) error {
	//var amounts []int
	//if err := json.Unmarshal(amountJSON, &amounts); err != nil {
	//	return fmt.Errorf("金额格式不正确: %w", err)
	//}
	// 检查是否为空数组
	if len(amounts) == 0 {
		return fmt.Errorf("金额列表不能为空")
	}
	// 可选：验证每个金额是否为正数
	for _, amount := range amounts {
		if amount <= 0 {
			return fmt.Errorf("金额必须为正数: %d", amount)
		}
	}
	return nil
}

func ValidateProvince(ProvinceCode []int) error {
	var ProvinceCodeSlice []string
	for _, v := range constant.ProvinceList {
		code := strconv.Itoa(v.Value)
		ProvinceCodeSlice = append(ProvinceCodeSlice, code)
	}
	for _, v := range ProvinceCode {
		if !IsInSlice(v, ProvinceCodeSlice) {
			return fmt.Errorf("省份编码不正确，%d", v)
		}
	}
	return nil
}

func ValidateAPIPath(path string) bool {
	regex := regexp.MustCompile(`^/[a-zA-Z0-9_/.:][a-zA-Z0-9_/.:.-]*[a-zA-Z0-9_.:.-]$`)
	return regex.MatchString(path)
}
