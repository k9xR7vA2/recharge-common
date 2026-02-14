package mysql

import (
	"github.com/gofrs/uuid/v5"
	"github.com/small-cat1/recharge-common/constant"
)

type Login interface {
	GetUsername() string
	GetNickname() string
	GetUUID() uuid.UUID
	GetUserId() uint
	GetAuthorityId() uint
	GetTenantId() uint
	GetUserInfo() any
	IsTenantAdmin() bool
}

var _ Login = new(TenantUser)

type TenantUser struct {
	BaseModel
	UUID             uuid.UUID                    `json:"uuid" gorm:"index;comment:用户UUID"`                             // 用户UUID
	TenantID         uint                         `gorm:"column:tenant_id;type:int unsigned;not null" json:"tenant_id"` //租户ID
	Username         string                       `json:"userName" gorm:"index;comment:用户登录名"`                          // 用户登录名
	Password         string                       `json:"-"  gorm:"comment:用户登录密码"`                                     // 用户登录密码
	NickName         string                       `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                    // 用户昵称
	AuthorityId      uint                         `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                // 用户角色ID
	Authority        TenantAuthority              `json:"authority" gorm:"foreignKey:AuthorityId;References:AuthorityId;comment:用户角色"`
	Enable           constant.GlobalAccountStatus `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	GoogleAuthKey    string                       `json:"-" gorm:"column:google_auth_key;comment:google_auth_key;"`
	GoogleAuthStatus bool                         `json:"googleAuthStatus" gorm:"default:0;comment:用户是否绑定谷歌验证器 0未绑定 1绑定"`                                                                                                     //  用户是否绑定谷歌验证器 0未绑定 1绑定
	IsTenant         int                          `json:"is_tenant" gorm:"column:is_tenant;type:tinyint;not null;default:1;comment:'是否是租户'" `                                                                                 // 是否是租户1不是，2是
	Authorities      []TenantAuthority            `json:"authorities" gorm:"many2many:tenant_user_authority;foreignKey:ID;joinForeignKey:tenant_user_id;References:AuthorityId;joinReferences:tenant_authority_authority_id"` // 多角色
}

func (TenantUser) TableName() string {
	return "tenant_users"
}

func (s *TenantUser) GetUsername() string {
	return s.Username
}

func (s *TenantUser) GetNickname() string {
	return s.NickName
}

func (s *TenantUser) GetUUID() uuid.UUID {
	return s.UUID
}

func (s *TenantUser) GetUserId() uint {
	return s.ID
}

func (s *TenantUser) GetAuthorityId() uint {
	return s.AuthorityId
}

func (s *TenantUser) GetTenantId() uint {
	return s.TenantID
}

func (s *TenantUser) GetUserInfo() any {
	return *s
}

func (s *TenantUser) IsTenantAdmin() bool {
	return s.IsTenant == 2
}

type TenantView struct {
	TenantID      uint   `json:"tenant_id,omitempty" gorm:"autoIncrement:true;primaryKey;column:tenant_id;type:int unsigned;not null" `   // 租户ID
	TenantName    string `json:"tenant_name,omitempty" gorm:"column:tenant_name;type:varchar(255);not null;comment:'租户名称'" `              // 租户名称
	TenantAccount string `json:"tenant_account,omitempty" gorm:"column:tenant_account;unique;type:varchar(255);not null;comment:'租户账号'" ` // 租户账号
}

// TableName get sql table name.获取数据库表名
func (TenantView) TableName() string {
	return "as_tenants"
}
