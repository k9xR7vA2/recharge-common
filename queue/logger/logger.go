package logger

import "github.com/k9xR7vA2/recharge-common/logger"

// AsynqLogger 适配公共库 Logger 接口到 asynq Logger 接口
type AsynqLogger struct {
	log logger.Logger
}

func NewAsynqLogger(log logger.Logger) *AsynqLogger {
	return &AsynqLogger{log: log}
}

func (l *AsynqLogger) Debug(args ...interface{}) { l.log.Debug(args...) }
func (l *AsynqLogger) Info(args ...interface{})  { l.log.Info(args...) }
func (l *AsynqLogger) Warn(args ...interface{})  { l.log.Warn(args...) }
func (l *AsynqLogger) Error(args ...interface{}) { l.log.Error(args...) }
func (l *AsynqLogger) Fatal(args ...interface{}) { l.log.Fatal(args...) }
