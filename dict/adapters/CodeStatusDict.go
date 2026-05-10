package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

type CodeStatusDict struct{}

func (d *CodeStatusDict) GetKey() string {
	return "code_status"
}

func (d *CodeStatusDict) GetName() string {
	return "码状态"
}

func (d *CodeStatusDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.CodeStatusAvailable.Label(),
			Value: constant.CodeStatusAvailable.Val(),
		},
		{
			Label: constant.CodeStatusUsed.Label(),
			Value: constant.CodeStatusUsed.Val(),
		},
		{
			Label: constant.CodeStatusExpired.Label(),
			Value: constant.CodeStatusExpired.Val(),
		},
	}
}
