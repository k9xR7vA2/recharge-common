package planlog

import "github.com/k9xR7vA2/recharge-common/logger"

// Logger 直接复用公共库已有接口，不重复定义
type Logger = logger.Logger

// noopLogger 默认空实现
type noopLogger struct{}

func (n *noopLogger) Debug(args ...interface{}) {}
func (n *noopLogger) Info(args ...interface{})  {}
func (n *noopLogger) Warn(args ...interface{})  {}
func (n *noopLogger) Error(args ...interface{}) {}
func (n *noopLogger) Fatal(args ...interface{}) {}

// 全局 logger，默认 noop
var globalLogger Logger = &noopLogger{}

func SetLogger(l Logger) {
	if l != nil {
		globalLogger = l
	}
}

func GetLogger() Logger {
	return globalLogger
}
