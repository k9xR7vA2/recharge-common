package adapters

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/dict/types"
	"strconv"
)

type SupplierAccountStatusDict struct{}

func (d *SupplierAccountStatusDict) GetKey() string {
	return "supplier_account_status"
}

func (d *SupplierAccountStatusDict) GetName() string {
	return "账号池状态"
}

func (d *SupplierAccountStatusDict) GetOptions() []types.DictOption {
	statuses := constant.GetAllSupplierAccountStatuses()
	options := make([]types.DictOption, 0, len(statuses))
	for _, s := range statuses {
		options = append(options, types.DictOption{
			Label:   s.Label,
			Value:   strconv.Itoa(int(s.Value)),
			Code:    strconv.Itoa(int(s.Value)),
			TagType: s.TagType, // DictOption 需补充该字段，见下方说明
		})
	}
	return options
}
