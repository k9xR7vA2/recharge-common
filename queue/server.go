package queue

import (
	"github.com/hibiken/asynq"
	"github.com/k9xR7vA2/recharge-common/queue/logger"
)

type IHandlerRegister interface {
	RegisterHandlers(mux *asynq.ServeMux)
}

type AsynqServer struct {
	server *asynq.Server
	mux    *asynq.ServeMux
	log    logger.Logger
}

func NewAsynqServer(cfg AsynqConfig, serverCfg ServerConfig) *AsynqServer {
	redisOpt := asynq.RedisClientOpt{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	}
	srv := asynq.NewServer(redisOpt, NewServerConfig(serverCfg))
	return &AsynqServer{
		server: srv,
		mux:    asynq.NewServeMux(),
		log:    serverCfg.Log,
	}
}

func (s *AsynqServer) Start(register IHandlerRegister) error {
	register.RegisterHandlers(s.mux)
	s.log.Info("starting asynq server...")
	return s.server.Run(s.mux)
}

func (s *AsynqServer) Close() {
	if s.server != nil {
		s.server.Shutdown()
	}
}
