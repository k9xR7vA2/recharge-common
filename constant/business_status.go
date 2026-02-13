package constant

// TenantBusinessStatus 租户业务状态（适用于产品、通道等业务）
type TenantBusinessStatus int

const (
	TenantBusinessStatusEnable  TenantBusinessStatus = iota + 1 // TenantBusinessStatusEnable 启用
	TenantBusinessStatusDisable                                 // TenantBusinessStatusDisable 禁用
	TenantBusinessStatusUnbind                                  // TenantBusinessStatusUnbind  解绑
)

// String 实现状态的字符串表示
func (s TenantBusinessStatus) String() string {
	switch s {
	case TenantBusinessStatusEnable:
		return "启用"
	case TenantBusinessStatusDisable:
		return "禁用"
	case TenantBusinessStatusUnbind:
		return "解绑"
	default:
		return "未知状态"
	}
}

// IsValid 验证状态值是否有效
func (s TenantBusinessStatus) IsValid() bool {
	return s >= TenantBusinessStatusEnable && s <= TenantBusinessStatusUnbind
}

// IsBind 验证是否绑定
func (s TenantBusinessStatus) IsBind() bool {
	return s == TenantBusinessStatusEnable || s == TenantBusinessStatusDisable
}
