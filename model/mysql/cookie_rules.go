package mysql

import "time"

// CookieRule cookie 规则表
type CookieRule struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	ChannelCode string `gorm:"column:channel_code;type:varchar(50);comment:通道编码" json:"channel_code"`

	// 额度规则
	DailyOrderMax   int `gorm:"column:daily_order_max;default:3"      json:"daily_order_max"`
	MonthlyOrderMax int `gorm:"column:monthly_order_max;default:30"   json:"monthly_order_max"`
	MinIntervalSec  int `gorm:"column:min_interval_sec;default:1800"  json:"min_interval_sec"`
	RandDelayMinSec int `gorm:"column:rand_delay_min_sec;default:300" json:"rand_delay_min_sec"`
	RandDelayMaxSec int `gorm:"column:rand_delay_max_sec;default:900" json:"rand_delay_max_sec"`

	// 封控判断规则
	MaxFailCount int `gorm:"column:max_fail_count;default:3"   json:"max_fail_count"`
	CooldownSec  int `gorm:"column:cooldown_sec;default:86400" json:"cooldown_sec"`
	SuspendSec   int `gorm:"column:suspend_sec;default:604800" json:"suspend_sec"`

	// 探测规则
	ProbeEnabled     bool `gorm:"column:probe_enabled;default:true"      json:"probe_enabled"`
	ProbeIntervalSec int  `gorm:"column:probe_interval_sec;default:3600" json:"probe_interval_sec"`

	// 刷新规则
	RefreshEnabled     bool `gorm:"column:refresh_enabled;default:true" json:"refresh_enabled"`
	RefreshIntervalSec int  `gorm:"column:refresh_interval_sec;default:72000"  json:"refresh_interval_sec"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (CookieRule) TableName() string {
	return "as_cookie_rules"
}
