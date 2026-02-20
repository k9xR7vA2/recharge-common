package logger

// Logger 公共库日志接口，业务项目实现
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

// AsynqLogger 适配公共库 Logger 接口到 asynq Logger 接口
type AsynqLogger struct {
	log Logger
}

func NewAsynqLogger(log Logger) *AsynqLogger {
	return &AsynqLogger{log: log}
}

func (l *AsynqLogger) Debug(args ...interface{}) { l.log.Debug(args...) }
func (l *AsynqLogger) Info(args ...interface{})  { l.log.Info(args...) }
func (l *AsynqLogger) Warn(args ...interface{})  { l.log.Warn(args...) }
func (l *AsynqLogger) Error(args ...interface{}) { l.log.Error(args...) }
func (l *AsynqLogger) Fatal(args ...interface{}) { l.log.Fatal(args...) }
