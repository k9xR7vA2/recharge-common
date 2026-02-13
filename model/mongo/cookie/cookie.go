package cookie

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Cookie 精简后的Cookie表，只存储业务数据
type Cookie struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Phone    string             `json:"phone" bson:"phone"`
	Platform string             `json:"platform" bson:"platform"`
	Operator int                `json:"operator" bson:"operator"` // 运营商:1联通,2电信,3移动,4广电
	Status   int                `json:"status" bson:"status"`     // 状态:1正常,2暂停,3失效,4封禁

	// 关联平台规则
	PlatformRuleID primitive.ObjectID `json:"platform_rule_id,omitempty" bson:"platform_rule_id,omitempty"`
	PlatformCode   string             `json:"platform_code" bson:"platform_code"`

	// 基础属性 - Cookie固有数据
	HealthScore    int     `json:"health_score" bson:"health_score"`                       // 健康度0-100
	UseCount       int     `json:"use_count" bson:"use_count"`                             // 历史使用次数
	LastUsedAt     int64   `json:"last_used_at,omitempty" bson:"last_used_at,omitempty"`   // 最后使用时间
	LastLoginAt    int64   `json:"last_login_at,omitempty" bson:"last_login_at,omitempty"` // 最后登录时间
	Channel        string  `json:"channel" bson:"channel"`                                 // 接码渠道
	Cost           float64 `json:"cost" bson:"cost"`                                       // 成本费用
	LoginFailCount int     `json:"login_fail_count" bson:"login_fail_count"`               // 登录失败次数

	// Cookie核心数据
	LoginRequest  string `json:"login_request" bson:"login_request"`   // 登录请求参数
	LoginHeaders  string `json:"login_headers" bson:"login_headers"`   // 登录请求头
	LoginCookies  string `json:"login_cookies" bson:"login_cookies"`   // Cookie数据
	LoginResponse string `json:"login_response" bson:"login_response"` // 登录响应
	HasCookie     bool   `json:"has_cookie" bson:"has_cookie"`         // 是否有Cookie
	CookieVersion int64  `json:"cookie_version" bson:"cookie_version"` // Cookie版本

	// 风控相关
	IPAddress string   `json:"ip_address,omitempty" bson:"ip_address,omitempty"`
	UserAgent string   `json:"user_agent,omitempty" bson:"user_agent,omitempty"`
	RiskTags  []string `json:"risk_tags,omitempty" bson:"risk_tags,omitempty"`
	Remark    string   `json:"remark,omitempty" bson:"remark,omitempty"`

	// 时间字段
	ExpiredAt int64 `json:"expired_at,omitempty" bson:"expired_at,omitempty"`
	CreatedAt int64 `json:"created_at" bson:"created_at"`
	UpdatedAt int64 `json:"updated_at" bson:"updated_at"`
	IsDeleted bool  `json:"is_deleted" bson:"is_deleted"`
}
