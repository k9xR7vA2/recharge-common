package cookie

import "go.mongodb.org/mongo-driver/bson/primitive"

type CookiePoolStatus struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CookieID primitive.ObjectID `json:"cookie_id" bson:"cookie_id"` // 关联Cookie
	Platform string             `json:"platform" bson:"platform"`   // 冗余字段，便于查询
	// 池子状态
	IsInPool      bool  `json:"is_in_pool" bson:"is_in_pool"`             // 是否在池子中
	PoolPriority  int   `json:"pool_priority" bson:"pool_priority"`       // 池子优先级
	AddedToPoolAt int64 `json:"added_to_pool_at" bson:"added_to_pool_at"` // 加入池子时间
	LastPoolCheck int64 `json:"last_pool_check" bson:"last_pool_check"`   // 最后检查时间
	// 运行时状态
	CurrentUsage    int   `json:"current_usage" bson:"current_usage"`       // 当前使用数
	LastUsageAt     int64 `json:"last_usage_at" bson:"last_usage_at"`       // 最后使用时间
	TotalAcquires   int   `json:"total_acquires" bson:"total_acquires"`     // 总获取次数
	SuccessAcquires int   `json:"success_acquires" bson:"success_acquires"` // 成功次数

	// 统计信息
	AvgResponseTime int     `json:"avg_response_time" bson:"avg_response_time"` // 平均响应时间(ms)
	FailureRate     float64 `json:"failure_rate" bson:"failure_rate"`           // 失败率

	CreatedAt int64 `json:"created_at" bson:"created_at"`
	UpdatedAt int64 `json:"updated_at" bson:"updated_at"`
}
