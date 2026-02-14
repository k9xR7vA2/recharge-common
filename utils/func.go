package utils

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func HandleDBResult(result *gorm.DB, notFoundErr error) error {
	if result.Error == nil {
		return nil
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return notFoundErr // 传入自定义的"未找到"错误
	}

	return result.Error // 返回原始数据库错误
}

// IsZeroValue  检查一个值是否为其类型的"零值"或空值
func IsZeroValue(v interface{}) bool {
	if v == nil {
		return true
	}

	val := reflect.ValueOf(v)

	// 处理指针
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return true
		}
		// 解引用指针并检查其值
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.String:
		return val.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.Slice, reflect.Map, reflect.Array:
		return val.Len() == 0
	case reflect.Struct:
		// 对于时间，检查是否为零时间
		if val.Type() == reflect.TypeOf(time.Time{}) {
			return val.Interface().(time.Time).IsZero()
		}
		// 对于其他结构体，可以根据需要定义什么是"空"
		// 例如，可以检查所有字段是否都为零值
		return reflect.DeepEqual(val.Interface(), reflect.New(val.Type()).Elem().Interface())
	}

	return false
}

func IsInSlice(ele interface{}, slice []string) bool {
	// 将ele转为字符串
	var eleStr string
	switch val := ele.(type) {
	case string:
		eleStr = val
	case int:
		eleStr = strconv.Itoa(val)
	case int64:
		eleStr = strconv.FormatInt(val, 10)
	case uint:
		eleStr = strconv.FormatUint(uint64(val), 10)
	case uint64:
		eleStr = strconv.FormatUint(val, 10)
	case float64:
		// 如果是整数值的float,转为整数字符串
		if float64(int(val)) == val {
			eleStr = strconv.Itoa(int(val))
		} else {
			eleStr = strconv.FormatFloat(val, 'f', -1, 64)
		}
	default:
		eleStr = fmt.Sprintf("%v", val)
	}

	// 对每个slice中的值尝试数字比较
	if eleNum, err := strconv.ParseInt(eleStr, 10, 64); err == nil {
		for _, v := range slice {
			if vNum, err := strconv.ParseInt(v, 10, 64); err == nil {
				if eleNum == vNum {
					return true
				}
			}
		}
	}

	// 再尝试字符串比较
	for _, v := range slice {
		if eleStr == strings.TrimSpace(v) {
			return true
		}
	}

	return false
}

func SliceEqualUsingMap(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}

	countMap := make(map[uint]uint)

	// 统计第一个切片中元素出现次数
	for _, v := range a {
		countMap[v]++
	}

	// 减去第二个切片中元素出现的次数
	for _, v := range b {
		countMap[v]--
		if countMap[v] < 0 {
			return false
		}
	}

	// 检查所有计数是否为0
	for _, count := range countMap {
		if count != 0 {
			return false
		}
	}

	return true
}
