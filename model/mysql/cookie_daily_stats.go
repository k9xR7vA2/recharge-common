package mysql

import "time"

// CookieDailyStats 运营日报快照，定时任务生成，存 MySQL
// 表名: cookie_daily_stats
type CookieDailyStats struct {
	ID       uint   `gorm:"primarykey"                json:"id"`
	TenantID string `gorm:"column:tenant_id;not null" json:"tenant_id"`
	Platform string `gorm:"column:platform;not null"  json:"platform"`
	StatDate string `gorm:"column:stat_date;not null" json:"stat_date"` // "2025-03-06"

	// 存量与流量
	CarryOver int `gorm:"column:carry_over;default:0" json:"carry_over"` // 昨日结转可用数
	Imported  int `gorm:"column:imported;default:0"   json:"imported"`   // 当日新导入
	Used      int `gorm:"column:used;default:0"       json:"used"`       // 当日使用次数
	Banned    int `gorm:"column:banned;default:0"     json:"banned"`     // 当日废弃数
	Available int `gorm:"column:available;default:0"  json:"available"`  // 当日结束可用 = carryOver + imported - banned

	// 成本（选填）
	UnitCost  float64 `gorm:"column:unit_cost;default:0"  json:"unit_cost"`  // 单个 Cookie 成本
	CostTotal float64 `gorm:"column:cost_total;default:0" json:"cost_total"` // 当日总成本

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (CookieDailyStats) TableName() string {
	return "cookie_daily_stats"
}
