package retry

import (
	"github.com/hibiken/asynq"
	"time"
)

type Handler struct {
	defaultStrategy Strategy
	registry        *PolicyRegistry // ← 原来是直接用 PolicyMap
}

func NewHandler(defaultStrategy Strategy, registry *PolicyRegistry) *Handler {
	return &Handler{
		defaultStrategy: defaultStrategy,
		registry:        registry,
	}
}

func (h *Handler) GetDelayFunc() func(count int, err error, task *asynq.Task) time.Duration {
	return func(count int, err error, task *asynq.Task) time.Duration {
		strategy, exists := h.registry.Get(task.Type()) // ← 原来是 PolicyMap[task.Type()]
		if !exists {
			strategy = h.defaultStrategy
		}
		if strategy.MaxRetry == 0 || count > strategy.MaxRetry {
			return 0
		}
		if len(strategy.OnlyErrors) > 0 {
			shouldRetry := false
			for _, retryErr := range strategy.OnlyErrors {
				if err.Error() == retryErr.Error() {
					shouldRetry = true
					break
				}
			}
			if !shouldRetry {
				return 0
			}
		}
		return strategy.DelayFunc(count)
	}
}
