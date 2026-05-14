package mongo

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PreGeneratedCode struct {
	ID            primitive.ObjectID  `bson:"_id"`
	ChannelCode   string              `bson:"channel_code"`
	ProductCode   string              `bson:"product_code"`
	Amount        string              `bson:"amount"`
	Payment       string              `bson:"payment"`
	PaymentMethod string              `bson:"payment_method"`
	AccountID     primitive.ObjectID  `bson:"account_id"` // ：cookie 账号 ID
	Account       string              `bson:"account"`    // ✅ 新增：账号（手机号）
	CodeValue     string              `bson:"code_value"` // 实际的码
	PayUrl        string              `bson:"pay_url"`    // UPI 链接
	Status        constant.CodeStatus `bson:"status"`     // available/used/expired
	RefNo         string              `bson:"ref_no"`     // 三方流水号
	ExpiredAt     int64               `bson:"expired_at"` // 过期时间
	UsedAt        int64               `bson:"used_at"`
	CreatedAt     int64               `bson:"created_at"`
	OrderSn       string              `bson:"order_sn"` // 使用时写入
}

func (PreGeneratedCode) CollName() string {
	return "pre_generated_codes"
}
