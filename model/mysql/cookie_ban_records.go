package mysql

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"time"
)

// CookieBanRecord 封控历史事件，存 MySQL
// 表名: cookie_ban_records
type CookieBanRecord struct {
	ID       uint             `gorm:"primarykey"                json:"id"`
	TenantID string           `gorm:"column:tenant_id;not null" json:"tenantId"`
	CookieID string           `gorm:"column:cookie_id;not null" json:"cookieId"` // MongoDB ObjectId 字符串
	Platform string           `gorm:"column:platform;not null"  json:"platform"`
	BanType  constant.BanType `gorm:"column:ban_type;not null"  json:"banType"`

	TriggerAt     *time.Time `gorm:"column:trigger_at"     json:"triggerAt"`
	DetectMsg     string     `gorm:"column:detect_msg;type:text" json:"detectMsg"` // 触发时的响应信息
	CooldownUntil *time.Time `gorm:"column:cooldown_until" json:"cooldownUntil"`
	ResolvedAt    *time.Time `gorm:"column:resolved_at"    json:"resolvedAt"` // NULL=未恢复

	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (CookieBanRecord) TableName() string {
	return "cookie_ban_records"
}
