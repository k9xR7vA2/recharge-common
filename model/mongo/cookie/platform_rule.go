package cookie

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlatformRule 平台规则
type PlatformRule struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PlatformName string             `json:"platform_name" bson:"platform_name"`
	PlatformCode string             `json:"platform_code" bson:"platform_code"`

	// 核心限制
	MonthlyLimit    int `json:"monthly_limit" bson:"monthly_limit"`
	DailyLimit      int `json:"daily_limit,omitempty" bson:"daily_limit"`
	UseInterval     int `json:"use_interval" bson:"use_interval"`
	ConcurrentLimit int `json:"concurrent_limit" bson:"concurrent_limit"`

	// 时间窗口配置
	TimeWindow TimeWindowConfig `json:"time_window" bson:"time_window"`

	//// 池子管理配置
	//PoolConfig PoolConfig `json:"pool_config" bson:"pool_config"`
	//
	//// 健康检测配置
	//HealthConfig HealthConfig `json:"health_config" bson:"health_config"`

	IsActive  bool  `json:"is_active" bson:"is_active"`
	CreatedAt int64 `json:"created_at" bson:"created_at"`
	UpdatedAt int64 `json:"updated_at" bson:"updated_at"`
}

// TimeWindowConfig 其他配置结构体（简化版）
type TimeWindowConfig struct {
	WindowType  string      `json:"window_type" bson:"window_type"`
	ResetConfig ResetConfig `json:"reset_config" bson:"reset_config"`
}

type ResetConfig struct {
	ResetPeriod string `json:"reset_period" bson:"reset_period"`
	ResetDay    int    `json:"reset_day,omitempty" bson:"reset_day,omitempty"`
	ResetHour   int    `json:"reset_hour,omitempty" bson:"reset_hour,omitempty"`
	RollingDays int    `json:"rolling_days,omitempty" bson:"rolling_days,omitempty"`
}

type PoolConfig struct {
	MinPoolSize       int             `json:"min_pool_size" bson:"min_pool_size"`
	MaxPoolSize       int             `json:"max_pool_size" bson:"max_pool_size"`
	AutoRefillEnabled bool            `json:"auto_refill_enabled" bson:"auto_refill_enabled"`
	RefillThreshold   float64         `json:"refill_threshold" bson:"refill_threshold"`
	PriorityWeights   PriorityWeights `json:"priority_weights" bson:"priority_weights"`
}

type PriorityWeights struct {
	HealthScore  float64 `json:"health_score" bson:"health_score"`
	RemainTime   float64 `json:"remain_time" bson:"remain_time"`
	UseFrequency float64 `json:"use_frequency" bson:"use_frequency"`
	Concurrent   float64 `json:"concurrent" bson:"concurrent"`
}

type HealthConfig struct {
	FailCountThreshold   int  `json:"fail_count_threshold" bson:"fail_count_threshold"`
	UnusedDaysThreshold  int  `json:"unused_days_threshold" bson:"unused_days_threshold"`
	MinHealthScore       int  `json:"min_health_score" bson:"min_health_score"`
	CheckIntervalMinutes int  `json:"check_interval_minutes" bson:"check_interval_minutes"`
	AutoRemoveEnabled    bool `json:"auto_remove_enabled" bson:"auto_remove_enabled"`
}

// CookieStats Cookie统计信息
type CookieStats struct {
	Platform  string `json:"platform"`
	Total     int64  `json:"total"`
	Available int64  `json:"available"`
	Expired   int64  `json:"expired"`
	Invalid   int64  `json:"invalid"`
}
