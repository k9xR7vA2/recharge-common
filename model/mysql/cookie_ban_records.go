package mysql

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"time"
)

// CookieBanRecord 封控历史事件，运营复盘/永久废弃判断，存 MySQL
// 表名: cookie_ban_records
type CookieBanRecord struct {
	ID          uint             `gorm:"primarykey"                json:"id"`
	TenantID    string           `gorm:"column:tenant_id;not null" json:"tenant_id"`
	CookieID    string           `gorm:"column:cookie_id;not null" json:"cookie_id"` // MongoDB ObjectId 字符串
	ChannelCode string           `gorm:"column:channel_code;type:varchar(50);comment:通道编码" json:"channel_code" `
	BanType     constant.BanType `gorm:"column:ban_type;not null"  json:"ban_type"`

	TriggerAt     *time.Time `gorm:"column:trigger_at"     json:"trigger_at"`
	DetectMsg     string     `gorm:"column:detect_msg;type:text" json:"detect_msg"` // 触发时的响应信息
	CooldownUntil *time.Time `gorm:"column:cooldown_until" json:"cooldown_until"`
	ResolvedAt    *time.Time `gorm:"column:resolved_at"    json:"resolved_at"` // NULL=未恢复

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (CookieBanRecord) TableName() string {
	return "cookie_ban_records"
}
