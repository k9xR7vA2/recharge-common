package options

import (
	"errors"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/orderpool/entities"
)

type MobileHandlerOptionsBuilder struct {
	options MobileHandlerOptions
	errors  []error // 收集校验错误
}

func NewMobileOrderPoolOptions() *MobileHandlerOptionsBuilder {
	return &MobileHandlerOptionsBuilder{
		options: MobileHandlerOptions{},
		errors:  []error{},
	}
}

// ---------- 链式设置方法 ----------
func (b *MobileHandlerOptionsBuilder) WithRedisKeys(keys RedisKeys) *MobileHandlerOptionsBuilder {
	b.options.orderKey = keys.OrderKey
	b.options.poolKey = keys.PoolKey
	return b
}

func (b *MobileHandlerOptionsBuilder) WithOrderInfo(info OrderInfo) *MobileHandlerOptionsBuilder {
	b.options.orderSn = info.OrderSn
	b.options.supplierOrderSN = info.SupplierOrderSN
	b.options.priority = info.Priority
	return b
}

func (b *MobileHandlerOptionsBuilder) WithOrderData(data string) *MobileHandlerOptionsBuilder {
	b.options.orderData = data
	return b
}

func (b *MobileHandlerOptionsBuilder) WithTimeArgs(args TimeArgs) *MobileHandlerOptionsBuilder {
	b.options.validTime = args.ValidTime
	b.options.expiredAt = args.ExpiredAt
	return b
}

func (b *MobileHandlerOptionsBuilder) WithPoolArgs(args entities.MobilePoolArgs) *MobileHandlerOptionsBuilder {
	b.options.poolArgs = args
	return b
}

// ---------- Build：统一校验并返回 ----------
func (b *MobileHandlerOptionsBuilder) Build() (IMobileHandlerOptions, error) {
	// 参数校验
	if b.options.orderKey == "" {
		b.errors = append(b.errors, errors.New("orderKey is required"))
	}
	if b.options.poolKey == "" {
		b.errors = append(b.errors, errors.New("poolKey is required"))
	}
	if b.options.orderSn == "" {
		b.errors = append(b.errors, errors.New("systemOrderSn is required"))
	}
	if b.options.orderData == "" {
		b.errors = append(b.errors, errors.New("orderData is required"))
	}
	if b.options.validTime <= 0 {
		b.errors = append(b.errors, errors.New("validTime must be > 0"))
	}
	if b.options.poolArgs.Amount == "" {
		b.errors = append(b.errors, errors.New("poolArgs.Amount is required"))
	}
	if b.options.poolArgs.Carrier == "" {
		b.errors = append(b.errors, errors.New("poolArgs.Carrier is required"))
	}

	if len(b.errors) > 0 {
		return nil, fmt.Errorf("validation failed: %v", b.errors)
	}

	return b.options, nil
}
