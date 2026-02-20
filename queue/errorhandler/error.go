package errorhandler

import (
	"context"
	"github.com/hibiken/asynq"
	"time"
)

// TaskErrorRecord 错误记录结构，业务项目可按需扩展存储
type TaskErrorRecord struct {
	TaskID     string    `bson:"task_id"`
	TaskType   string    `bson:"task_type"`
	Payload    string    `bson:"payload"`
	Error      string    `bson:"error"`
	ErrorType  string    `json:"error_type"`
	ErrorStack string    `json:"error_stack"`
	RetryCount int       `bson:"retry_count"`
	MaxRetry   int       `bson:"max_retry"`
	CreatedAt  time.Time `bson:"created_at"`
}

// 公共库
type ErrorStorage interface {
	Save(ctx context.Context, record TaskErrorRecord) error
}

type AsynqErrorHandler struct {
	storage ErrorStorage // 可为 nil
}

func NewAsynqErrorHandler(storage ErrorStorage) *AsynqErrorHandler {
	return &AsynqErrorHandler{storage: storage}
}

func (h *AsynqErrorHandler) HandleError(ctx context.Context, task *asynq.Task, err error) {
	if h.storage == nil {
		return
	}
	retryCount, _ := asynq.GetRetryCount(ctx)
	maxRetry, _ := asynq.GetMaxRetry(ctx)
	_ = h.storage.Save(ctx, TaskErrorRecord{
		TaskType:   task.Type(),
		Payload:    string(task.Payload()),
		Error:      err.Error(),
		RetryCount: retryCount,
		MaxRetry:   maxRetry,
		CreatedAt:  time.Now(),
	})
}
