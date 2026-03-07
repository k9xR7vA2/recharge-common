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
	TodayUsed       int    `bson:"today_used"        json:"todayUsed"`
	MonthUsed       int    `bson:"month_used"        json:"monthUsed"`
	TotalUsed       int    `bson:"total_used"        json:"totalUsed"`
	CountResetDate  string `bson:"count_reset_date"  json:"countResetDate"`
	CountResetMonth string `bson:"count_reset_month" json:"countResetMonth"`

	// 风控
	HealthScore    int        `bson:"health_score"     json:"healthScore"`
	LoginFailCount int        `bson:"login_fail_count" json:"loginFailCount"`
	RiskTags       []string   `bson:"risk_tags"        json:"riskTags"`
	CooldownUntil  *time.Time `bson:"cooldown_until"   json:"cooldownUntil"`

	// 状态
	Status constant.CookieStatus `bson:"status"    json:"status"`
	Remark string                `bson:"remark"    json:"remark"`

	// 时间
	LastUsedAt  int64 `bson:"last_used_at"  json:"lastUsedAt"`
	LastLoginAt int64 `bson:"last_login_at" json:"lastLoginAt"`
	ExpiredAt   int64 `bson:"expired_at"    json:"expiredAt"`
	CreatedAt   int64 `bson:"created_at"    json:"createdAt"`
	UpdatedAt   int64 `bson:"updated_at"    json:"updatedAt"`
	IsDeleted   bool  `bson:"is_deleted"    json:"isDeleted"`
}
