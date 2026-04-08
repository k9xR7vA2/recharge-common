package cookie

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LoginRecord 登录记录表，跟踪每次登录行为
// 集合名: login_records
type LoginRecord struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChannelCode string             `bson:"channel_code" json:"channel_code"`
	Phone       string             `bson:"phone"         json:"phone"`

	// 关联成功后的 Cookie（失败为空）
	CookieID *primitive.ObjectID `bson:"cookie_id,omitempty" json:"cookieId,omitempty"`

	// 号商 / 接码渠道
	Channel        string `bson:"channel"          json:"channel"`        // 渠道标识
	ChannelOrderSn string `bson:"channel_order_sn" json:"channelOrderSn"` // 渠道订单号，用于对账

	// 成本
	Cost float64 `bson:"cost" json:"cost"` // 本次登录成本

	// 登录请求现场
	LoginFailCount int    `bson:"login_fail_count" json:"login_fail_count"`
	LoginRequest   string `bson:"login_request"  json:"loginRequest"`
	LoginHeaders   string `bson:"login_headers"  json:"loginHeaders"`
	LoginResponse  string `bson:"login_response" json:"loginResponse"`
	IPAddress      string `bson:"ip_address"     json:"ipAddress"`
	UserAgent      string `bson:"user_agent"     json:"userAgent"`

	// 结果
	Status     int    `bson:"status"     json:"status"` // 1成功 2失败 3超时
	FailReason string `bson:"fail_reason" json:"failReason"`

	CreatedAt int64 `bson:"created_at" json:"createdAt"`
}
