package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
)

type GiftCardStatusDict struct{}

func (d *GiftCardStatusDict) GetKey() string {
	return "gift_card_status"
}

func (d *GiftCardStatusDict) GetName() string {
	return "卡密状态"
}

func (d *GiftCardStatusDict) GetOptions() []types.DictOption {
	return []types.DictOption{
		{
			Label: constant.GiftCardStatusUnused.Label(),
			Value: constant.GiftCardStatusUnused.Value(),
		},
		{
			Label: constant.GiftCardStatusAssigned.Label(),
			Value: constant.GiftCardStatusAssigned.Value(),
		},
		{
			Label: constant.GiftCardStatusVerified.Label(),
			Value: constant.GiftCardStatusVerified.Value(),
		},
		{
			Label: constant.GiftCardStatusInvalid.Label(),
			Value: constant.GiftCardStatusInvalid.Value(),
		},
	}
}
