package channel

import (
	"github.com/small-cat1/recharge-common/model/mongo/order/merchant"
)

// MerchantMobileOrder 商户基础订单
type MerchantMobileOrder struct {
	merchant.MerchantOrder `bson:",inline"`
}
