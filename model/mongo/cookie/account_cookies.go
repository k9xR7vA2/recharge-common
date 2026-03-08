package cookie

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// AccountCookie 精简后的Cookie表，只存储业务数据
type AccountCookie struct {
	ID       primitive.ObjectID    `bson:"_id,omitempty" json:"id"`
	Platform constant.PlatformType `bson:"platform"      json:"platform"`
	Phone    string                `bson:"phone"         json:"phone"`

	// 溯源（可追回登录现场）
	LoginRecordID primitive.ObjectID `bson:"login_record_id" json:"loginRecordId"`
	// Cookie 数据
	LoginCookies  string `bson:"login_cookies"  json:"loginCookies"` // AES加密
	CookieVersion int64  `bson:"cookie_version" json:"cookieVersion"`

	// 账号辅助信息
	AccountAlias string `bson:"account_alias" json:"accountAlias"`
	ProxyIP      string `bson:"proxy_ip"      json:"proxyIp"`
	UserAgent    string `bson:"user_agent"    json:"userAgent"`

	// 使用计数
	TodayUsed       int    `bson:"today_used"        json:"today_used"`
	MonthUsed       int    `bson:"month_used"        json:"month_used"`
	TotalUsed       int    `bson:"total_used"        json:"total_used"`
	CountResetDate  string `bson:"count_reset_date"  json:"count_reset_date"`
	CountResetMonth string `bson:"count_reset_month" json:"count_reset_month"`

	// 风控
	HealthScore    int        `bson:"health_score"     json:"health_score"`
	LoginFailCount int        `bson:"login_fail_count" json:"login_fail_count"`
	RiskTags       []string   `bson:"risk_tags"        json:"risk_tags"`
	CooldownUntil  *time.Time `bson:"cooldown_until"   json:"cooldown_until"`

	// 状态
	Status constant.CookieStatus `bson:"status"    json:"status"`
	Remark string                `bson:"remark"    json:"remark"`

	// 时间
	LastUsedAt  int64 `bson:"last_used_at"  json:"last_used_at"`
	LastLoginAt int64 `bson:"last_login_at" json:"last_login_at"`
	ExpiredAt   int64 `bson:"expired_at"    json:"expired_at"`
	CreatedAt   int64 `bson:"created_at"    json:"created_at"`
	UpdatedAt   int64 `bson:"updated_at"    json:"updated_at"`
	IsDeleted   bool  `bson:"is_deleted"    json:"is_deleted"`
}
