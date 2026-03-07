package mongo

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AccountCookie 存储在租户 MongoDB 分库中
// 集合名: account_cookies
// 无 tenant_id，库名本身即租户隔离
type AccountCookie struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"    json:"id"`
	Platform string             `bson:"platform"         json:"platform"` // amazon_us / ebay_us 等

	// 账号身份
	AccountEmail string `bson:"account_email"    json:"accountEmail"`
	AccountAlias string `bson:"account_alias"    json:"accountAlias"` // 内部备注
	ProxyIP      string `bson:"proxy_ip"         json:"proxyIp"`      // 1:1 绑定代理
	UserAgent    string `bson:"user_agent"       json:"userAgent"`

	// Cookie 本体
	CookieRaw  string                 `bson:"cookie_raw"       json:"-"`          // AES 加密存储，不对外暴露
	CookieJSON map[string]interface{} `bson:"cookie_json"      json:"cookieJson"` // 结构化 KV

	// 生命周期
	ImportedAt     *time.Time `bson:"imported_at"      json:"importedAt"`
	ExpiresAt      *time.Time `bson:"expires_at"       json:"expiresAt"` // TTL 索引字段
	LastVerifiedAt *time.Time `bson:"last_verified_at" json:"lastVerifiedAt"`
	LastUsedAt     *time.Time `bson:"last_used_at"     json:"lastUsedAt"`

	// 使用计数（纯统计，规则判断在 platform_rules）
	TodayUsed       int    `bson:"today_used"       json:"todayUsed"`
	MonthUsed       int    `bson:"month_used"       json:"monthUsed"`
	TotalUsed       int    `bson:"total_used"       json:"totalUsed"`
	CountResetDate  string `bson:"count_reset_date" json:"countResetDate"`   // "2025-03-06"
	CountResetMonth string `bson:"count_reset_month" json:"countResetMonth"` // "2025-03"

	// 状态
	Status        constant.CookieStatus `bson:"status"           json:"status"`
	CooldownUntil *time.Time            `bson:"cooldown_until"   json:"cooldownUntil"`

	ExpiredAt int64 `bson:"expired_at"`
	CreatedAt int64 `bson:"created_at"`
	UpdatedAt int64 `bson:"updated_at"`
	IsDeleted bool  `bson:"is_deleted"`
}
