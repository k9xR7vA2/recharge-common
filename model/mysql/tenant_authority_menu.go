package mysql

type TenantMenu struct {
	TenantBaseMenu
	MenuId      uint                      `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId uint                      `json:"-" gorm:"comment:角色ID"`
	Children    []TenantMenu              `json:"children" gorm:"-"`
	Parameters  []TenantBaseMenuParameter `json:"parameters" gorm:"foreignKey:TenantBaseMenuID;references:MenuId"`
	Btns        map[string]uint           `json:"btns" gorm:"-"`
}

type TenantAuthorityMenu struct {
	MenuId      uint `json:"menuId" gorm:"comment:菜单ID;column:tenant_base_menu_id"`      //菜单ID
	AuthorityId uint `json:"-" gorm:"comment:角色ID;column:tenant_authority_authority_id"` //角色ID
	TenantID    uint `json:"tenant_id" gorm:"column:tenant_id;"`                         //租户ID
}

func (s TenantAuthorityMenu) TableName() string {
	return "tenant_authority_menus"
}
