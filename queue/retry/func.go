package retry

import (
	"context"
	"github.com/hibiken/asynq"
)

const GlobalQueueMaxRetry = 3

func GetRetryInfo(ctx context.Context, taskType string, registry *PolicyRegistry) (retryCount, maxRetry int, shouldRetry bool) {
	retryCount, ok := asynq.GetRetryCount(ctx)
	if !ok {
		retryCount = 0
	}
	maxRetry, ok = asynq.GetMaxRetry(ctx)
	if !ok {
		maxRetry = GlobalQueueMaxRetry
		if strategy, exists := registry.Get(taskType); exists {
			maxRetry = strategy.MaxRetry
		}
	}
	shouldRetry = retryCount <= maxRetry
	return
}
