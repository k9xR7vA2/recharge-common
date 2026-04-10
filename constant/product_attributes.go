package constant

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/dict"
	"gorm.io/datatypes"
	"strconv"
	"strings"
	"sync"
)

// MobileAttributes 话费属性
type MobileAttributes struct {
	Carrier      CarrierType `json:"carrier"`
	ChargeSpeed  ChargeSpeed `json:"charge_speed"`
	AreaCode     AreaScope   `json:"area_code"`
	ProvinceCode []int       `json:"province_code"` // 支持的省份列表 [200,210,220]
	IsCheckIsp   int         `json:"is_check_isp"`
}

// MobileAttributes 实现接口
func (a *MobileAttributes) GetCarrier() int             { return int(a.Carrier) }
func (a *MobileAttributes) GetChargeSpeed() ChargeSpeed { return a.ChargeSpeed }
func (a *MobileAttributes) GetIsCheckIsp() int          { return a.IsCheckIsp }

func ParseMobileProductAttrs(data datatypes.JSON) (*MobileAttributes, error) {
	var attr MobileAttributes
	err := json.Unmarshal(data, &attr)
	if err != nil {
		return nil, err
	}
	if !attr.Carrier.IsValid() {
		return nil, errors.New("运营商类型不正确")
	}
	if !attr.ChargeSpeed.IsValid() {
		return nil, errors.New("充值速度不正确！")
	}
	if !attr.AreaCode.IsValid() {
		return nil, errors.New("区域参数不正确")
	}
	if attr.AreaCode == Province {
		err = ValidateProvince(attr.ProvinceCode)
		if err != nil {
			return nil, err
		}
	}
	return &attr, err
}

func ValidateProvince(ProvinceCode []int) error {
	var ProvinceCodeSlice []string
	for _, v := range ProvinceList {
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

type IndiaMobileAttributes struct {
	IsCheckIsp  int              `json:"is_check_isp"`
	ChargeSpeed ChargeSpeed      `json:"charge_speed"`
	Carrier     IndiaCarrierType `json:"carrier"`
}

// IndiaMobileAttributes 实现接口
func (a *IndiaMobileAttributes) GetCarrier() int             { return int(a.Carrier) }
func (a *IndiaMobileAttributes) GetChargeSpeed() ChargeSpeed { return a.ChargeSpeed }
func (a *IndiaMobileAttributes) GetIsCheckIsp() int          { return a.IsCheckIsp }

func ParseIndiaMobileProductAttrs(data datatypes.JSON) (*IndiaMobileAttributes, error) {
	var attr IndiaMobileAttributes
	err := json.Unmarshal(data, &attr)
	if err != nil {
		return nil, err
	}
	if !attr.Carrier.IsValid() {
		return nil, errors.New("运营商类型不正确")
	}
	if !attr.ChargeSpeed.IsValid() {
		return nil, errors.New("充值速度不正确！")
	}
	return &attr, err
}

// IndiaDthAttributes IndiaDTH卫星电视属性
type IndiaDthAttributes struct {
	Operator []IndiaDthOperatorType `json:"operator"` // 改为切片，支持多运营商
}

func ParseDthProductAttrs(data datatypes.JSON) (*IndiaDthAttributes, error) {
	var attr IndiaDthAttributes
	if err := json.Unmarshal(data, &attr); err != nil {
		return nil, err
	}
	if len(attr.Operator) == 0 {
		return nil, errors.New("至少选择一个DTH运营商")
	}
	for _, op := range attr.Operator {
		if !op.IsValid() {
			return nil, fmt.Errorf("DTH运营商类型不正确: %d", op)
		}
	}
	return &attr, nil
}

var indiaElectricOperatorIds map[int]bool
var indiaElectricOnce sync.Once

func getIndiaElectricOperatorIds() map[int]bool {
	indiaElectricOnce.Do(func() {
		indiaElectricOperatorIds = make(map[int]bool)
		d := dict.GetDict("india_electric_operator")
		for _, item := range d.Options {
			switch v := item.Value.(type) {
			case int:
				indiaElectricOperatorIds[v] = true
			case float64:
				indiaElectricOperatorIds[int(v)] = true
			case int64:
				indiaElectricOperatorIds[int(v)] = true
			}
		}
	})
	return indiaElectricOperatorIds
}

// IndiaElectricAttributes 印度电费属性
type IndiaElectricAttributes struct {
	OperatorIds []int `json:"operator_ids"`
}

func ParseIndiaElectricProductAttrs(data datatypes.JSON) (*IndiaElectricAttributes, error) {
	var attr IndiaElectricAttributes
	if err := json.Unmarshal(data, &attr); err != nil {
		return nil, err
	}
	if len(attr.OperatorIds) == 0 {
		return nil, errors.New("至少选择一个电费运营商")
	}
	validIds := getIndiaElectricOperatorIds()
	for _, id := range attr.OperatorIds {
		if !validIds[id] {
			return nil, fmt.Errorf("运营商ID不合法: %d", id)
		}
	}
	return &attr, nil
}
