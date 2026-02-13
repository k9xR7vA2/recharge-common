package utils

import (
	"reflect"
	"time"
)

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
