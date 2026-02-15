package utils

import (
	"reflect"
	"strconv"
	"strings"
)

// 判断  是否在 allowedAmounts 中
func IsAllowed(amount interface{}, allowedAmounts interface{}) bool {
	allowedSlice := reflect.ValueOf(allowedAmounts)
	for i := 0; i < allowedSlice.Len(); i++ {
		if reflect.DeepEqual(amount, allowedSlice.Index(i).Interface()) {
			return true
		}
	}
	return false
}

// ValidateSliceIsInSlice 验证 amounts 中的每个值是否都在 allowedAmounts 中
func ValidateSliceIsInSlice(amounts interface{}, allowedAmounts interface{}) bool {
	amountsVal := reflect.ValueOf(amounts)
	// 确保 amounts 和 allowedAmounts 都是切片
	if amountsVal.Kind() != reflect.Slice || reflect.ValueOf(allowedAmounts).Kind() != reflect.Slice {
		return false
	}
	// 检查 amounts 中的每个元素是否在 allowedAmounts 中
	for i := 0; i < amountsVal.Len(); i++ {
		if !IsAllowed(amountsVal.Index(i).Interface(), allowedAmounts) {
			return false
		}
	}
	return true
}

func StringToIntSlice(input string) ([]int, error) {
	// 使用 strings.Split 将字符串按逗号分割
	strSlice := strings.Split(input, ",")
	// 创建一个 int 切片，用于存放转换后的结果
	intSlice := make([]int, len(strSlice))

	// 遍历字符串切片，将每个元素转换为整数
	for i, str := range strSlice {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err // 如果转换失败，返回错误
		}
		intSlice[i] = num
	}
	return intSlice, nil
}
