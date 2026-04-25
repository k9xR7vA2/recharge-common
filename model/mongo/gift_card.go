package mongo

import (
	"github.com/k9xR7vA2/recharge-common/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GiftCard struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ProductCode string             `bson:"product_code"` // 产品编码，如 amazonGiftCard
	ChannelCode string             `bson:"channel_code"` // 通道编码

	// 卡密信息
	CardNo    string `bson:"card_no"`    // 卡号
	CardPin   string `bson:"card_pin"`   // 卡密（AES加密存储）
	FaceValue string `bson:"face_value"` // 面值
	Currency  string `bson:"currency"`   // 货币
	ExpiredAt int64  `bson:"expired_at"` // 卡密过期时间

	// 库存状态
	Status constant.GiftCardStatus `bson:"status"` // 1未使用 2已分配 3已核销 4已失效

	// 分配信息（已分配/已核销时填充）
	AssignedOrderSn string `bson:"assigned_order_sn"` // 分配给哪个订单
	AssignedAt      int64  `bson:"assigned_at"`
	VerifiedAt      int64  `bson:"verified_at"` // 核销时间

	// 来源追踪
	SourceOrderSn string `bson:"source_order_sn"` // 平台采购时的上游订单号
	PurchasedAt   int64  `bson:"purchased_at"`    // 采购时间

	CreatedAt int64 `bson:"created_at"`
	UpdatedAt int64 `bson:"updated_at"`
	IsDeleted bool  `bson:"is_deleted"`
}
