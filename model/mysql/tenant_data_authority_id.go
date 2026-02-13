package mysql

type TenantDataAuthorityID struct {
	TenantAuthorityAuthorityID uint `gorm:"primaryKey;column:tenant_authority_authority_id"`  // 角色ID
	DataAuthorityIDAuthorityID uint `gorm:"primaryKey;column:data_authority_id_authority_id"` //角色对应的角色ID
	TenantID                   uint `gorm:"primaryKey;column:tenant_id"`                      //租户ID
}

func (TenantDataAuthorityID) TableName() string {
	return "tenant_data_authority_id"
}
