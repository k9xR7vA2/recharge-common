package filter

import "strings"

// Filter 定义一个泛型过滤器函数类型，T可以是任意类型
type Filter[T any] func(item T) bool

// ApplyFilter 应用单个过滤器到数据集
func ApplyFilter[T any](items []T, filter Filter[T]) []T {
	var result []T
	for _, item := range items {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}

// ApplyFilters 应用多个过滤器到数据集（与逻辑，全部满足）
func ApplyFilters[T any](items []T, filters []Filter[T]) []T {
	result := items
	for _, filter := range filters {
		result = ApplyFilter(result, filter)
	}
	return result
}

// ApplyAnyFilter 应用多个过滤器到数据集（或逻辑，满足任一条件）
func ApplyAnyFilter[T any](items []T, filters []Filter[T]) []T {
	if len(filters) == 0 {
		return items
	}

	var result []T
	for _, item := range items {
		for _, filter := range filters {
			if filter(item) {
				result = append(result, item)
				break // 一旦有一个过滤器满足，就添加并跳出内循环
			}
		}
	}
	return result
}

// CombineFilters 组合多个过滤器为一个（与逻辑）
func CombineFilters[T any](filters ...Filter[T]) Filter[T] {
	return func(item T) bool {
		for _, filter := range filters {
			if !filter(item) {
				return false
			}
		}
		return true
	}
}

// CombineAnyFilters 组合多个过滤器为一个（或逻辑）
func CombineAnyFilters[T any](filters ...Filter[T]) Filter[T] {
	return func(item T) bool {
		for _, filter := range filters {
			if filter(item) {
				return true
			}
		}
		return false
	}
}

// StringContains 创建一个检查string字段是否包含特定子串的过滤器
func StringContains[T any](getValue func(T) string, substr string) Filter[T] {
	return func(item T) bool {
		if substr == "" {
			return true // 如果子串为空，不进行过滤
		}
		value := getValue(item)
		return strings.Contains(value, substr)
	}
}
