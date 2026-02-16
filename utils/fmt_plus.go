package utils

import (
	"encoding/json"
	"fmt"
	"github.com/small-cat1/recharge-common/constant"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

func ToJsonString(v interface{}) string {
	if v == nil {
		return "nil"
	}
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("json marshal error: %v, raw data: %+v", err, v)
	}
	return string(jsonBytes)
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

// StringMapToString 将 map[string]string 转换为字符串表示
func StringMapToString(data map[string]string) string {
	var parts []string
	for k, v := range data {
		parts = append(parts, fmt.Sprintf("%s=%s", k, v))
	}
	return "{" + strings.Join(parts, ", ") + "}"
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// MaheHump 将字符串转换为驼峰命名
func MaheHump(s string) string {
	words := strings.Split(s, "-")

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, "")
}

// 随机字符串
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[RandomInt(0, len(letters))]
	}
	return string(b)
}

// generateRandomDigits 生成指定长度的随机数字字符串
func GenerateRandomDigits(length int) string {
	const digits = "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = digits[time.Now().UnixNano()%10]
		time.Sleep(time.Nanosecond)
	}
	return string(result)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func GetProvinceLabel(code int) (string, bool) {
	for _, province := range constant.ProvinceList {
		if province.Value == code {
			return province.Label, true
		}
	}
	return "", false
}

func IsProvinceMatch(area, provinceLabel string) bool {
	return strings.Contains(area, provinceLabel) || strings.Contains(provinceLabel, area)
}
