package cookie

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AccountStats 账号使用统计
type AccountStats struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CookieID primitive.ObjectID `json:"cookie_id" bson:"cookie_id"` // 关联Cookie
	Platform string             `json:"platform" bson:"platform"`

	// 当前周期统计
	CurrentPeriod   string `json:"current_period" bson:"current_period"`       // "2024-05" 或 "2024-05-01_2024-05-31"
	PeriodUsed      int    `json:"period_used" bson:"period_used"`             // 当期已使用次数
	PeriodLimit     int    `json:"period_limit" bson:"period_limit"`           // 当期限制次数
	PeriodStartTime int64  `json:"period_start_time" bson:"period_start_time"` // 周期开始时间
	PeriodEndTime   int64  `json:"period_end_time" bson:"period_end_time"`     // 周期结束时间

	// 今日统计
	Today     string `json:"today" bson:"today"`           // "2024-05-24"
	DailyUsed int    `json:"daily_used" bson:"daily_used"` // 今日已使用次数

	// 最后使用时间
	LastUsedAt int64 `json:"last_used_at,omitempty" bson:"last_used_at,omitempty"`

	UpdatedAt int64 `json:"updated_at" bson:"updated_at"`
}
