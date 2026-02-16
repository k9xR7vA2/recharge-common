package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/small-cat1/recharge-common/constant"
	"github.com/small-cat1/recharge-common/utils"
	"gorm.io/datatypes"
	"strconv"
)

// MobileAttributes 话费属性
type MobileAttributes struct {
	Carrier      constant.CarrierType `json:"carrier"`
	ChargeSpeed  constant.ChargeSpeed `json:"charge_speed"`
	AreaCode     constant.AreaScope   `json:"area_code"`
	ProvinceCode []int                `json:"province_code"` // 支持的省份列表 [200,210,220]
	IsCheckIsp   int                  `json:"is_check_isp"`
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
	if attr.AreaCode == constant.Province {
		var ProvinceCodeSlice []string
		for _, v := range constant.ProvinceList {
			code := strconv.Itoa(v.Value)
			ProvinceCodeSlice = append(ProvinceCodeSlice, code)
		}
		for _, v := range attr.ProvinceCode {
			if !utils.IsInSlice(v, ProvinceCodeSlice) {
				return nil, fmt.Errorf("省份编码不正确，%d", v)
			}
		}
	}
	if attr.AreaCode == constant.Province {
		err = utils.ValidateProvince(attr.ProvinceCode)
		if err != nil {
			return nil, err
		}
	}

	return &attr, err
}
