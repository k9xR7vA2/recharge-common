package retry

import "time"

// Strategy 重试策略定义
type Strategy struct {
	MaxRetry   int
	DelayFunc  func(count int) time.Duration
	OnlyErrors []error
}

// 预置延迟函数
var (
	FixedDelay = func(seconds int) func(count int) time.Duration {
		return func(count int) time.Duration {
			return time.Duration(seconds) * time.Second
		}
	}

	LinearDelay = func(baseSeconds int) func(count int) time.Duration {
		return func(count int) time.Duration {
			return time.Duration(count*baseSeconds) * time.Second
		}
	}

	ExponentialDelay = func(baseSeconds int) func(count int) time.Duration {
		return func(count int) time.Duration {
			return time.Duration(1<<uint(count)) * time.Second * time.Duration(baseSeconds)
		}
	}
)

// 预置常用策略
var (
	DefaultStrategy = Strategy{MaxRetry: 3, DelayFunc: LinearDelay(2)}
	HighPriority    = Strategy{MaxRetry: 5, DelayFunc: ExponentialDelay(2)}
	NoRetry         = Strategy{MaxRetry: 0, DelayFunc: FixedDelay(0)}
)
