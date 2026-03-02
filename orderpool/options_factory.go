package orderpool

import (
	"context"
	"github.com/k9xR7vA2/recharge-common/constant"
	"github.com/k9xR7vA2/recharge-common/orderpool/entities"
	"github.com/k9xR7vA2/recharge-common/orderpool/keys"
	"github.com/k9xR7vA2/recharge-common/orderpool/options"
)

type IOrderPoolOptionsFactory interface {
	BuildMobileOrderOptions(ctx context.Context,
		businessType string,
		mobilePoolArgs entities.AddMobilePoolArgs) (options.IMobileHandlerOptions, error)
	BuildMobileOrderFetchOptions(ctx context.Context, businessType string, mobilePoolArgs entities.MobileMatchmakingArgs) (options.IFetchMobileOrderOptions, error)
}

type OrderPoolOptionsFactory struct {
}

func NewOrderPoolOptionsFactory() *OrderPoolOptionsFactory {
	return &OrderPoolOptionsFactory{}
}

// BuildMobileOrderOptions  为供应商订单入池构建选项
func (f *OrderPoolOptionsFactory) BuildMobileOrderOptions(ctx context.Context,
	businessType string, mobilePoolArgs entities.AddMobilePoolArgs) (options.IMobileHandlerOptions, error) {
	keysGen := keys.NewRedisKeysGenerate(mobilePoolArgs.TenantID, keys.RoleSupplier, businessType)
	orderKey := keysGen.OrderKey(mobilePoolArgs.SystemOrderSn)
	poolKey := keysGen.GenerateMobilePoolKey(mobilePoolArgs.Priority, mobilePoolArgs.MobilePoolArgs)
	return options.NewMobileOrderPoolOptions().
		WithRedisKeys(options.RedisKeys{
			OrderKey: orderKey,
			PoolKey:  poolKey,
		}).
		WithOrderInfo(options.OrderInfo{
			OrderSn:         mobilePoolArgs.SystemOrderSn,
			SupplierOrderSN: mobilePoolArgs.SupplierOrderSn,
			Priority:        mobilePoolArgs.Priority,
		}).
		WithOrderData(mobilePoolArgs.SupplierOrderStr).
		WithTimeArgs(options.TimeArgs{
			ValidTime: mobilePoolArgs.ValidTime,
			ExpiredAt: mobilePoolArgs.ExpiredAt,
		}).
		WithPoolArgs(mobilePoolArgs.MobilePoolArgs).
		Build()
}

func (f *OrderPoolOptionsFactory) BuildMobileOrderFetchOptions(ctx context.Context,
	businessType string,
	mobilePoolArgs entities.MobileMatchmakingArgs) (options.IFetchMobileOrderOptions, error) {
	keysGen := keys.NewRedisKeysGenerate(mobilePoolArgs.TenantID, keys.RoleSupplier, businessType)
	// 生成两个优先级的 poolKey
	highPriorityPoolKey := keysGen.GenerateMobilePoolKey(constant.HighPriority.String(), mobilePoolArgs.MobilePoolArgs)
	normalPriorityPoolKey := keysGen.GenerateMobilePoolKey(constant.NormalPriority.String(), mobilePoolArgs.MobilePoolArgs)
	return options.NewFetchMobileOrderPoolOptions().
		WithTenantInfo(options.TenantInfo{
			TenantID:     mobilePoolArgs.TenantID,
			RoleType:     keys.RoleSupplier,
			BusinessType: businessType,
		}).
		WithFetchRedisKeys(options.FetchRedisKeys{
			HighPriorityPoolKey:   highPriorityPoolKey,
			NormalPriorityPoolKey: normalPriorityPoolKey,
		}).
		WithPoolArgs(mobilePoolArgs.MobilePoolArgs).
		Build()
}
