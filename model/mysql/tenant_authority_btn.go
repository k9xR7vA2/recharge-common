package mysql

type TenantAuthorityBtn struct {
	AuthorityId         uint              `gorm:"comment:角色ID"`
	TenantMenuID        uint              `gorm:"comment:菜单ID"`
	TenantBaseMenuBtnID uint              `gorm:"comment:菜单按钮ID"`
	TenantID            uint              `json:"tenant_id" gorm:"column:tenant_id;"` //租户ID
	TenantBaseMenuBtn   TenantBaseMenuBtn ` gorm:"comment:按钮详情"`
}
