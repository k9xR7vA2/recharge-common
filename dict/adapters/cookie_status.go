package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

// CookieStatusDict Cookie 状态字典
type CookieStatusDict struct{}

func (d *CookieStatusDict) GetKey() string {
	return "cookie_status"
}

func (d *CookieStatusDict) GetName() string {
	return "Cookie状态"
}

func (d *CookieStatusDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{Label: constant.CookieStatusNormal.String(), Value: int(constant.CookieStatusNormal), Code: constant.CookieStatusNormal.Code()},
		{Label: constant.CookieStatusCooldown.String(), Value: int(constant.CookieStatusCooldown), Code: constant.CookieStatusCooldown.Code()},
		{Label: constant.CookieStatusSuspected.String(), Value: int(constant.CookieStatusSuspected), Code: constant.CookieStatusSuspected.Code()},
		{Label: constant.CookieStatusSuspended.String(), Value: int(constant.CookieStatusSuspended), Code: constant.CookieStatusSuspended.Code()},
		{Label: constant.CookieStatusBanned.String(), Value: int(constant.CookieStatusBanned), Code: constant.CookieStatusBanned.Code()},
		{Label: constant.CookieStatusExpired.String(), Value: int(constant.CookieStatusExpired), Code: constant.CookieStatusExpired.Code()},
	}
}
