package mysql

import "strconv"

type TenantCasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"` // 主键ID
	Ptype string `gorm:"size:100" json:"ptype"`              // 策略类型：p 或 g
	V0    string `gorm:"size:100" json:"v0"`                 // 主体（如用户、角色）
	V1    string `gorm:"size:100" json:"v1"`                 // 资源（如URL、对象）
	V2    string `gorm:"size:100" json:"v2"`                 // 操作（如read、write）
	V3    string `gorm:"size:100" json:"tenant_id"`          // 租户字段
	V4    string `gorm:"size:100" json:"v4,omitempty"`       // 扩展字段2
	V5    string `gorm:"size:100" json:"v5,omitempty"`       // 扩展字段3
}

func (TenantCasbinRule) TableName() string {
	return "tenant_casbin_rules"
}

func GetTenantCasbinRules(AuthorityId uint, TenantID string, tenantApi []TenantApi) []TenantCasbinRule {
	v0 := strconv.Itoa(int(AuthorityId))
	var rules []TenantCasbinRule
	for _, v := range tenantApi {
		rules = append(rules, TenantCasbinRule{
			Ptype: "p",
			V0:    v0,
			V1:    v.Path,
			V2:    v.Method,
			V3:    TenantID,
		})
	}
	return rules
}
