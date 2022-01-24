package zdpgo_zap

import "go.uber.org/zap"

// Info 记录info类型的日志
func Info(msg string, kvs ...interface{}) {
	zap.S().Infow(msg, kvs...)
}

// Debug 记录debug类型的日志
func Debug(msg string, kvs ...interface{}) {
	zap.S().Debugw(msg, kvs...)
}

// Warning 记录warning类型的日志
func Warning(msg string, kvs ...interface{}) {
	zap.S().Warnw(msg, kvs...)
}

// Error 记录error类型的日志
func Error(msg string, kvs ...interface{}) {
	zap.S().Errorw(msg, kvs...)
}

// Panic 记录panic类型的日志
func Panic(msg string, kvs ...interface{}) {
	zap.S().Panicw(msg, kvs...)
}

// Fatal 记录fatal类型的日志
func Fatal(msg string, kvs ...interface{}) {
	zap.S().Fatalw(msg, kvs...)
}
