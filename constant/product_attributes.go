package constant

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/small-cat1/recharge-common/utils"
	"gorm.io/datatypes"
	"strconv"
)

// MobileAttributes 话费属性
type MobileAttributes struct {
	Carrier      CarrierType `json:"carrier"`
	ChargeSpeed  ChargeSpeed `json:"charge_speed"`
	AreaCode     AreaScope   `json:"area_code"`
	ProvinceCode []int       `json:"province_code"` // 支持的省份列表 [200,210,220]
	IsCheckIsp   int         `json:"is_check_isp"`
}

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
		var ProvinceCodeSlice []string
		for _, v := range ProvinceList {
			code := strconv.Itoa(v.Value)
			ProvinceCodeSlice = append(ProvinceCodeSlice, code)
		}
		for _, v := range attr.ProvinceCode {
			if !utils.IsInSlice(v, ProvinceCodeSlice) {
				return nil, fmt.Errorf("省份编码不正确，%d", v)
			}
		}
	}
	if attr.AreaCode == Province {
		err = utils.ValidateProvince(attr.ProvinceCode)
		if err != nil {
			return nil, err
		}
	}

	return &attr, err
}

type IndiaMobileAttributes struct {
	IsCheckIsp  int              `json:"is_check_isp"`
	ChargeSpeed ChargeSpeed      `json:"charge_speed"`
	Carrier     IndiaCarrierType `json:"carrier"`
	HasSku      int              `json:"has_sku"`
}

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
	return &attr, err
}
