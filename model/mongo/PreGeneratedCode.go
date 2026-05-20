package mongo

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PreGeneratedCode struct {
	ID                   primitive.ObjectID  `bson:"_id"`
	ChannelCode          string              `bson:"channel_code"`
	ProductCode          string              `bson:"product_code"`
	Amount               string              `bson:"amount"`
	Payment              string              `bson:"payment"`
	PaymentMethod        string              `bson:"payment_method"`
	AccountID            primitive.ObjectID  `bson:"account_id"`             // ：cookie 账号 ID
	Account              string              `bson:"account"`                // ✅ 新增：账号（手机号）
	CodeValue            string              `bson:"code_value"`             // 实际的码
	PayUrl               string              `bson:"pay_url"`                // UPI 链接
	Status               constant.CodeStatus `bson:"status"`                 // available/used/expired
	RefNo                string              `bson:"ref_no"`                 // 三方流水号
	OfficialSerialNumber string              `bson:"official_serial_number"` //官方流水号
	TransactionID        string              `bson:"transaction_id"`         //官方交易ID
	ExpiredAt            int64               `bson:"expired_at"`             // 过期时间
	UsedAt               int64               `bson:"used_at"`
	CreatedAt            int64               `bson:"created_at"`
	OrderSn              string              `bson:"order_sn"` // 绑定merchantOrder
	// ↓ 新增
	TradeNo     string                 `bson:"trade_no,omitempty"`     // 三方请求交易号，用于追溯
	ErrorMsg    string                 `bson:"error_msg,omitempty"`    // 失败原因
	ApiRequest  map[string]interface{} `bson:"api_request,omitempty"`  // 三方请求体
	ApiResponse map[string]interface{} `bson:"api_response,omitempty"` // 三方响应体
}

func (PreGeneratedCode) CollName() string {
	return "pre_generated_codes"
}
