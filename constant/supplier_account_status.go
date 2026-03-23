package constant

// SupplierAccountStatus 账号池状态
type SupplierAccountStatus int

const (
	SupplierAccountStatusPending    SupplierAccountStatus = 1 // 待审核
	SupplierAccountStatusApproved   SupplierAccountStatus = 2 // 审核通过
	SupplierAccountStatusProcessing SupplierAccountStatus = 3 // 充值中
	SupplierAccountStatusDone       SupplierAccountStatus = 4 // 已完成
	SupplierAccountStatusRejected   SupplierAccountStatus = 5 // 已拒绝
	SupplierAccountStatusDisabled   SupplierAccountStatus = 6 // 禁用
)

func (s SupplierAccountStatus) ShowName() string {
	switch s {
	case SupplierAccountStatusPending:
		return "待审核"
	case SupplierAccountStatusApproved:
		return "审核通过"
	case SupplierAccountStatusProcessing:
		return "充值中"
	case SupplierAccountStatusDone:
		return "已完成"
	case SupplierAccountStatusRejected:
		return "已拒绝"
	case SupplierAccountStatusDisabled:
		return "禁用"
	default:
		return "未知状态"
	}
}

func (s SupplierAccountStatus) Code() string {
	switch s {
	case SupplierAccountStatusPending:
		return "pending"
	case SupplierAccountStatusApproved:
		return "approved"
	case SupplierAccountStatusProcessing:
		return "processing"
	case SupplierAccountStatusDone:
		return "done"
	case SupplierAccountStatusRejected:
		return "rejected"
	case SupplierAccountStatusDisabled:
		return "disabled"
	default:
		return "unknown"
	}
}

func (s SupplierAccountStatus) CanDelete() bool {
	return s == SupplierAccountStatusPending ||
		s == SupplierAccountStatusRejected ||
		s == SupplierAccountStatusDisabled
}

// TagType 对应前端 el-tag 的 type
func (s SupplierAccountStatus) TagType() string {
	switch s {
	case SupplierAccountStatusPending:
		return "warning"
	case SupplierAccountStatusApproved:
		return "primary"
	case SupplierAccountStatusProcessing:
		return "info"
	case SupplierAccountStatusDone:
		return "success"
	case SupplierAccountStatusRejected:
		return "danger"
	case SupplierAccountStatusDisabled:
		return "danger"
	default:
		return "info"
	}
}

func (s SupplierAccountStatus) IsValid() bool {
	switch s {
	case SupplierAccountStatusPending,
		SupplierAccountStatusApproved,
		SupplierAccountStatusProcessing,
		SupplierAccountStatusDone,
		SupplierAccountStatusRejected,
		SupplierAccountStatusDisabled:
		return true
	default:
		return false
	}
}

func GetAllSupplierAccountStatuses() []struct {
	Label   string                `json:"label"`
	Value   SupplierAccountStatus `json:"value"`
	Code    string                `json:"code"`
	TagType string                `json:"tag_type"`
} {
	return []struct {
		Label   string                `json:"label"`
		Value   SupplierAccountStatus `json:"value"`
		Code    string                `json:"code"`
		TagType string                `json:"tag_type"`
	}{
		{
			Label:   SupplierAccountStatusPending.ShowName(),
			Code:    SupplierAccountStatusPending.Code(),
			Value:   SupplierAccountStatusPending,
			TagType: SupplierAccountStatusPending.TagType()},
		{Label: SupplierAccountStatusApproved.ShowName(),
			Code: SupplierAccountStatusApproved.Code(),

			Value: SupplierAccountStatusApproved, TagType: SupplierAccountStatusApproved.TagType()},
		{Label: SupplierAccountStatusProcessing.ShowName(),
			Code: SupplierAccountStatusProcessing.Code(),

			Value: SupplierAccountStatusProcessing, TagType: SupplierAccountStatusProcessing.TagType()},
		{Label: SupplierAccountStatusDone.ShowName(),
			Code:  SupplierAccountStatusDone.Code(),
			Value: SupplierAccountStatusDone, TagType: SupplierAccountStatusDone.TagType()},
		{Label: SupplierAccountStatusRejected.ShowName(),
			Code:  SupplierAccountStatusRejected.Code(),
			Value: SupplierAccountStatusRejected, TagType: SupplierAccountStatusRejected.TagType()},
		{Label: SupplierAccountStatusDisabled.ShowName(),
			Code:  SupplierAccountStatusDisabled.Code(),
			Value: SupplierAccountStatusDisabled, TagType: SupplierAccountStatusDisabled.TagType()},
	}
}
