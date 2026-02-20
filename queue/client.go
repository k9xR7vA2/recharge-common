package queue

import (
	"context"
	"github.com/hibiken/asynq"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// TraceInjectable 业务 payload 实现此接口即可自动注入 trace
type TraceInjectable interface {
	SetTraceCtx(carrier map[string]string)
}

type AsynqClient struct {
	client     *asynq.Client
	propagator propagation.TextMapPropagator
}

func NewAsynqClient(cfg AsynqConfig) *AsynqClient {
	redisOpt := asynq.RedisClientOpt{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	}
	return &AsynqClient{
		client:     asynq.NewClient(redisOpt),
		propagator: otel.GetTextMapPropagator(),
	}
}

func (c *AsynqClient) GetClient() *asynq.Client {
	return c.client
}

// InjectTraceContext 业务项目 enqueue 前调用
func (c *AsynqClient) InjectTraceContext(ctx context.Context, payload TraceInjectable) {
	carrier := make(map[string]string)
	c.propagator.Inject(ctx, propagation.MapCarrier(carrier))
	payload.SetTraceCtx(carrier)
}

func (c *AsynqClient) Close() {
	_ = c.client.Close()
}
