package cookie

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AccountCookie Cookie 主体 + 风控状态，权威来源
type AccountCookie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChannelCode string             `bson:"channel_code" json:"channel_code"`
	Phone       string             `bson:"phone"         json:"phone"`

	// 溯源（可追回登录现场）
	LoginRecordID primitive.ObjectID `bson:"login_record_id" json:"loginRecordId"`
	LoginCookies  string             `bson:"login_cookies"  json:"loginCookies"` // AES加密
	CookieVersion int64              `bson:"cookie_version" json:"cookieVersion"`

	// 账号辅助信息
	AccountAlias string `bson:"account_alias" json:"accountAlias"`
	ProxyIP      string `bson:"proxy_ip"      json:"proxyIp"`
	UserAgent    string `bson:"user_agent"    json:"userAgent"`

	// 风控
	HealthScore          int      `bson:"health_score"     json:"health_score"`
	RiskTags             []string `bson:"risk_tags"        json:"risk_tags"`
	FailCount            int      `bson:"fail_count"    json:"fail_count"` // 使用失败次数（触发风控用）
	CooldownUntil        int64    `bson:"cooldown_until"   json:"cooldown_until"`
	SuspendUntil         int64    `bson:"suspend_until" json:"suspend_until"`                  // 封控到期时间
	NextProbeAt          int64    `bson:"next_probe_at" json:"next_probe_at"`                  // 下次探测时间
	BillQueryBannedUntil int64    `bson:"bill_query_banned_until" json:"billQueryBannedUntil"` // 账单查询封禁到期时间（到期前不用于查询）

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

func (AccountCookie) CollName() string {
	return "cookies"
}
