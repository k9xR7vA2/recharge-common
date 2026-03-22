package constant

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/datatypes"
	"strconv"
	"strings"
)

// MobileAttributes 话费属性
type MobileAttributes struct {
	Carrier      CarrierType `json:"carrier"`
	ChargeSpeed  ChargeSpeed `json:"charge_speed"`
	AreaCode     AreaScope   `json:"area_code"`
	ProvinceCode []int       `json:"province_code"` // 支持的省份列表 [200,210,220]
	IsCheckIsp   int         `json:"is_check_isp"`
	ValidTime    int         `json:"valid_time"`
}

// MobileAttributes 实现接口
func (a *MobileAttributes) GetCarrier() int             { return int(a.Carrier) }
func (a *MobileAttributes) GetChargeSpeed() ChargeSpeed { return a.ChargeSpeed }
func (a *MobileAttributes) GetIsCheckIsp() int          { return a.IsCheckIsp }
func (a *MobileAttributes) GetValidTime() int           { return a.ValidTime }

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
	if attr.ValidTime <= 0 {
		return nil, errors.New("订单有效期参数不正确")
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
	HasSku      int              `json:"has_sku"`
	ValidTime   int              `json:"valid_time"`
}

// IndiaMobileAttributes 实现接口
func (a *IndiaMobileAttributes) GetCarrier() int             { return int(a.Carrier) }
func (a *IndiaMobileAttributes) GetChargeSpeed() ChargeSpeed { return a.ChargeSpeed }
func (a *IndiaMobileAttributes) GetIsCheckIsp() int          { return a.IsCheckIsp }
func (a *IndiaMobileAttributes) GetValidTime() int           { return a.ValidTime }

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
	skuFlag := []int{1, 2}
	flg := false
	for _, v := range skuFlag {
		if v == attr.HasSku {
			flg = true
			break
		}
	}
	if flg == false {
		return nil, errors.New("sku参数错误")
	}
	if attr.ValidTime <= 0 {
		return nil, errors.New("订单有效期参数不正确")
	}
	return &attr, err
}
