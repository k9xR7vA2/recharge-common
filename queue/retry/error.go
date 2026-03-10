package retry

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/k9xR7vA2/recharge-common/logger"
)

type ErrorHandler struct {
	logger   logger.Logger // 公共库自己的 Logger 接口
	registry *PolicyRegistry
}

func NewErrorHandler(log logger.Logger, registry *PolicyRegistry) asynq.ErrorHandlerFunc {
	h := &ErrorHandler{logger: log, registry: registry}
	return h.HandleError
}

func (h *ErrorHandler) HandleError(ctx context.Context, task *asynq.Task, err error) {
	retryCount, _ := asynq.GetRetryCount(ctx)
	maxRetry, _ := asynq.GetMaxRetry(ctx)

	if retryCount >= maxRetry {
		// 达到最大重试次数
		h.logger.Error("task reached max retry",
			"taskType", task.Type(),
			"retryCount", retryCount,
			"maxRetry", maxRetry,
			"error", err.Error(),
		)
		return
	}

	h.logger.Warn("task failed, will retry",
		"taskType", task.Type(),
		"retryCount", retryCount,
		"maxRetry", maxRetry,
		"error", err.Error(),
	)
}
