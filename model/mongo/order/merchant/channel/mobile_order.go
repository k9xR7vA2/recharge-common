package channel

import (
	"SaasAdmin/internal/infrastructure/model/mongo/order/merchant"
)

// MerchantMobileOrder 商户基础订单
type MerchantMobileOrder struct {
	merchant.MerchantOrder `bson:",inline"`
}
