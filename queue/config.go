package queue

import (
	"github.com/hibiken/asynq"
	logger2 "github.com/k9xR7vA2/recharge-common/logger"
	"github.com/k9xR7vA2/recharge-common/queue/errorhandler"
	"github.com/k9xR7vA2/recharge-common/queue/logger"
	"github.com/k9xR7vA2/recharge-common/queue/retry"
)

type AsynqConfig struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	Concurrency   int
	RetryLimit    int
	Queues        map[string]int // 由业务项目传入
}

type ServerConfig struct {
	Concurrency    int
	Queues         map[string]int
	RetryDelayFunc *retry.Handler
	Log            logger2.Logger
	ErrorStorage   errorhandler.ErrorStorage
}

func NewServerConfig(cfg ServerConfig) asynq.Config {
	return asynq.Config{
		Concurrency:    cfg.Concurrency,
		Queues:         cfg.Queues,
		Logger:         logger.NewAsynqLogger(cfg.Log),
		ErrorHandler:   errorhandler.NewAsynqErrorHandler(cfg.ErrorStorage),
		RetryDelayFunc: cfg.RetryDelayFunc.GetDelayFunc(),
	}
}
