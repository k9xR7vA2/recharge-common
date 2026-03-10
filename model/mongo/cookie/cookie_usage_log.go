package cookie

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CookieUsageLog 不可变操作流水，排查/分析用
type CookieUsageLog struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CookieID     primitive.ObjectID `json:"cookie_id" bson:"cookie_id"` // 关联Cookie
	Platform     string             `json:"platform" bson:"platform"`
	Action       string             `json:"action" bson:"action"` // get,release,refresh,fail
	Result       int8               `json:"result" bson:"result"` // 1成功,2失败
	ErrorCode    string             `json:"error_code,omitempty" bson:"error_code,omitempty"`
	ErrorMessage string             `json:"error_message,omitempty" bson:"error_message,omitempty"`
	ResponseTime int                `json:"response_time" bson:"response_time"` // 响应时间ms
	IPAddress    string             `json:"ip_address,omitempty" bson:"ip_address,omitempty"`
	CreatedAt    int64              `json:"created_at" bson:"created_at"`
}
