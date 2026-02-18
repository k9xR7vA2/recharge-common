package options

import (
	"errors"
	"fmt"
	"github.com/small-cat1/recharge-common/orderpool/entities"
)

type FetchMobileHandlerOptionsBuilder struct {
	options FetchMobileHandlerOptions
	errors  []error
}

func NewFetchMobileOrderPoolOptions() *FetchMobileHandlerOptionsBuilder {
	return &FetchMobileHandlerOptionsBuilder{
		options: FetchMobileHandlerOptions{},
		errors:  []error{},
	}
}

func (b *FetchMobileHandlerOptionsBuilder) WithTenantInfo(info TenantInfo) *FetchMobileHandlerOptionsBuilder {
	b.options.tenantID = info.TenantID
	b.options.roleType = info.RoleType
	b.options.businessType = info.BusinessType
	return b
}

func (b *FetchMobileHandlerOptionsBuilder) WithFetchRedisKeys(keys FetchRedisKeys) *FetchMobileHandlerOptionsBuilder {
	b.options.statsKey = keys.StatsKey
	b.options.highPriorityPoolKey = keys.HighPriorityPoolKey
	b.options.normalPriorityPoolKey = keys.NormalPriorityPoolKey
	return b
}

// WithPoolArgs 和 MobileHandlerOptionsBuilder 共用同一个分组结构体
func (b *FetchMobileHandlerOptionsBuilder) WithPoolArgs(args entities.MobilePoolArgs) *FetchMobileHandlerOptionsBuilder {
	b.options.poolArgs = args
	return b
}

// ---------- Build ----------

func (b *FetchMobileHandlerOptionsBuilder) Build() (IFetchMobileOrderOptions, error) {
	if b.options.tenantID == 0 {
		b.errors = append(b.errors, errors.New("tenantID is required"))
	}
	if b.options.roleType == "" {
		b.errors = append(b.errors, errors.New("roleType is required"))
	}
	if b.options.businessType == "" {
		b.errors = append(b.errors, errors.New("businessType is required"))
	}
	if b.options.statsKey == "" {
		b.errors = append(b.errors, errors.New("statsKey is required"))
	}
	if b.options.highPriorityPoolKey == "" {
		b.errors = append(b.errors, errors.New("highPriorityPoolKey is required"))
	}
	if b.options.normalPriorityPoolKey == "" {
		b.errors = append(b.errors, errors.New("normalPriorityPoolKey is required"))
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
