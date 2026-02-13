package mysql

type TenantBaseMenuBtn struct {
	BaseModel
	Name             string `json:"name" gorm:"comment:按钮关键key"`
	Desc             string `json:"desc" gorm:"按钮备注"`
	TenantBaseMenuID uint   `json:"tenant_base_menu_id" gorm:"comment:菜单ID"`
}

func (TenantBaseMenuBtn) TableName() string {
	return "tenant_base_menu_btns"
}
