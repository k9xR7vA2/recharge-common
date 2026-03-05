package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

type TimezoneDict struct{}

func (d *TimezoneDict) GetKey() string {
	return "timezone"
}

func (d *TimezoneDict) GetName() string {
	return "时区"
}

func (d *TimezoneDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{Label: constant.TimezoneShanghai.Label(), Value: int(constant.TimezoneShanghai), Code: constant.TimezoneShanghai.ToIANA()},
		{Label: constant.TimezoneKolkata.Label(), Value: int(constant.TimezoneKolkata), Code: constant.TimezoneKolkata.ToIANA()},
		{Label: constant.TimezoneNewYork.Label(), Value: int(constant.TimezoneNewYork), Code: constant.TimezoneNewYork.ToIANA()},
		{Label: constant.TimezoneLosAngeles.Label(), Value: int(constant.TimezoneLosAngeles), Code: constant.TimezoneLosAngeles.ToIANA()},
		{Label: constant.TimezoneLondon.Label(), Value: int(constant.TimezoneLondon), Code: constant.TimezoneLondon.ToIANA()},
		{Label: constant.TimezoneDubai.Label(), Value: int(constant.TimezoneDubai), Code: constant.TimezoneDubai.ToIANA()},
		{Label: constant.TimezoneManila.Label(), Value: int(constant.TimezoneManila), Code: constant.TimezoneManila.ToIANA()},
	}
}
