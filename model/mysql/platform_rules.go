package mysql

import "time"

// PlatformRule 平台规则配置，管理后台维护
// 表名: platform_rules
type PlatformRule struct {
	ID          uint   `gorm:"primarykey"                    json:"id"`
	ChannelCode string `gorm:"column:channel_code;type:varchar(50);comment:通道编码" json:"channel_code" `
	// 额度规则
	DailyOrderMax   int `gorm:"column:daily_order_max;default:3"      json:"daily_order_max"`
	MonthlyOrderMax int `gorm:"column:monthly_order_max;default:30"   json:"monthly_order_max"`
	MinIntervalSec  int `gorm:"column:min_interval_sec;default:1800"  json:"min_interval_sec"`   // 两次使用最小间隔
	RandDelayMinSec int `gorm:"column:rand_delay_min_sec;default:300" json:"rand_delay_min_sec"` // 随机延迟下限
	RandDelayMaxSec int `gorm:"column:rand_delay_max_sec;default:900" json:"rand_delay_max_sec"` // 随机延迟上限

	// 封控判断规则
	MaxFailCount int `gorm:"column:max_fail_count;default:3"          json:"max_fail_count"` // 连续失败进风控阈值
	CooldownSec  int `gorm:"column:cooldown_sec;default:86400"         json:"cooldown_sec"`  // 软风控冷却时长（秒）
	SuspendSec   int `gorm:"column:suspend_sec;default:604800"         json:"suspend_sec"`   // 封控沉淀时长（秒）

	// 探测规则
	ProbeEnabled     bool `gorm:"column:probe_enabled;default:true"         json:"probe_enabled"` // 沉淀期满是否探测
	ProbeIntervalSec int  `gorm:"column:probe_interval_sec;default:3600"    json:"probe_interval_sec"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (PlatformRule) TableName() string {
	return "as_platform_rules"
}
