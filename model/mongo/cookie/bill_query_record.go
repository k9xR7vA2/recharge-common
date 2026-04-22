package cookie

import "go.mongodb.org/mongo-driver/bson/primitive"

type BillQueryRecord struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"    json:"id"`
	TenantID    uint               `bson:"tenant_id"        json:"tenantId"`
	SupplierID  uint               `bson:"supplier_id"      json:"supplierId"`
	ChannelCode string             `bson:"channel_code"     json:"channelCode"`
	Account     string             `bson:"account"          json:"account"` // 查询的手机号
	Month       string             `bson:"month"            json:"month"`   // 账单月份 2025-01

	// Cookie 相关
	CookieID primitive.ObjectID `bson:"cookie_id"        json:"cookieId"` // 使用的 cookie，缓存命中时为空

	// 结果
	Source      string  `bson:"source"           json:"source"` // cache / third_party
	Success     bool    `bson:"success"          json:"success"`
	ErrorCode   string  `bson:"error_code"       json:"errorCode"` // 失败时记录错误码
	BillFetchId string  `bson:"bill_fetch_id"    json:"billFetchId"`
	BillAmount  float64 `bson:"bill_amount"      json:"billAmount"`

	// 时间（UTC）
	CreatedAt int64 `bson:"created_at"       json:"createdAt"`
}

func (BillQueryRecord) CollName() string {
	return "bill_query_records"
}
