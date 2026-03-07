package mysql

import "time"

// PlatformRule 平台使用规则，存在 MySQL 公共库，tenant_id 隔离
// 表名: platform_rules
type PlatformRule struct {
	ID       uint   `gorm:"primarykey"                    json:"id"`
	Platform string `gorm:"column:platform;not null"      json:"platform"`

	// 额度规则
	DailyOrderMax   int `gorm:"column:daily_order_max;default:3"      json:"dailyOrderMax"`
	MonthlyOrderMax int `gorm:"column:monthly_order_max;default:30"   json:"monthlyOrderMax"`
	MinIntervalSec  int `gorm:"column:min_interval_sec;default:1800"  json:"minIntervalSec"`  // 两次使用最小间隔
	RandDelayMinSec int `gorm:"column:rand_delay_min_sec;default:300" json:"randDelayMinSec"` // 随机延迟下限
	RandDelayMaxSec int `gorm:"column:rand_delay_max_sec;default:900" json:"randDelayMaxSec"` // 随机延迟上限

	// 封控判断规则
	MaxFailCount int `gorm:"column:max_fail_count;default:3"          json:"maxFailCount"` // 连续失败进风控阈值
	CooldownSec  int `gorm:"column:cooldown_sec;default:86400"         json:"cooldownSec"` // 软风控冷却时长（秒）
	SuspendSec   int `gorm:"column:suspend_sec;default:604800"         json:"suspendSec"`  // 封控沉淀时长（秒）

	// 探测规则
	ProbeEnabled     bool `gorm:"column:probe_enabled;default:true"         json:"probeEnabled"` // 沉淀期满是否探测
	ProbeIntervalSec int  `gorm:"column:probe_interval_sec;default:3600"    json:"probeIntervalSec"`

	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (PlatformRule) TableName() string {
	return "as_platform_rules"
}
