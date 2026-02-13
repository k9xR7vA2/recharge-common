package mysql

import (
	"time"
)

// TenantAuthority  租户角色表
type TenantAuthority struct {
	CreatedAt       time.Time          // 创建时间
	UpdatedAt       time.Time          // 更新时间
	DeletedAt       *time.Time         `sql:"index"`
	AuthorityId     uint               `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	TenantID        uint               `json:"tenant_id" gorm:"column:tenant_id;type:int unsigned;not null" `       //租户ID
	AuthorityName   string             `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	ParentId        *uint              `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	DefaultRouter   string             `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`                 // 默认菜单(默认dashboard)
	Children        []TenantAuthority  `json:"children" gorm:"-"`
	TenantBaseMenus []TenantBaseMenu   `json:"menus" gorm:"many2many:tenant_authority_menus;foreignKey:authority_id;joinForeignKey:tenant_authority_authority_id;References:id;joinReferences:tenant_base_menu_id"`
	Users           []TenantUser       `json:"-" gorm:"many2many:tenant_user_authority;foreignKey:authority_id;joinForeignKey:tenant_authority_authority_id;References:id;joinReferences:tenant_user_id"`
	DataAuthorityId []*TenantAuthority `json:"dataAuthorityId" gorm:"many2many:tenant_data_authority_id;foreignKey:AuthorityId;joinForeignKey:tenant_authority_authority_id;References:AuthorityId;joinReferences:data_authority_id_authority_id"`
}

func (TenantAuthority) TableName() string {
	return "tenant_authorities"
}

type TenantDataAuthorityId struct {
	TenantID                   uint `json:"tenant_id" gorm:"column:tenant_id;type:int unsigned;not null" ` //租户ID
	TenantAuthorityAuthorityId uint `json:"tenant_authority_authority_id" gorm:"column:tenant_authority_authority_id;type:int unsigned;not null" `
	DataAuthorityIdAuthorityId uint `json:"data_authority_id_authority_id" gorm:"column:data_authority_id_authority_id;type:int unsigned;not null" `
}

func (TenantDataAuthorityId) TableName() string {
	return "tenant_data_authority_id"
}
