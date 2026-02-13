package mysql

// TenantUserAuthority 是 AsUser 和 AsAuthority 的连接表
type TenantUserAuthority struct {
	TenantUserId               uint `gorm:"column:tenant_user_id"`
	TenantAuthorityAuthorityId uint `gorm:"column:tenant_authority_authority_id"`
	TenantID                   uint `gorm:"primaryKey;column:tenant_id"` //租户ID
}

func (s *TenantUserAuthority) TableName() string {
	return "tenant_user_authority"
}
