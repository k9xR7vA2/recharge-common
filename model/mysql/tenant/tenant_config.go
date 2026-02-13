package tenant

import "time"

type TenantConfig struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"` // 主键ID
	TenantID  uint      `gorm:"primaryKey;column:tenant_id;type:int unsigned;not null" json:"tenant_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`   // 配置名称
	Key       string    `gorm:"type:varchar(255);not null" json:"key"`    // 配置键
	Value     *string   `gorm:"type:varchar(255)" json:"value,omitempty"` // 配置值（可为空）
	CreatedAt time.Time `json:"created_at"`                               // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                               // 更新时间
}

func (TenantConfig) TableName() string {
	return "as_tenant_config"
}
