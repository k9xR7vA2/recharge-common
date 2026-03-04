package base

// Logger 公共库日志接口，业务项目注入具体实现（zap、logrus 等均可）
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
