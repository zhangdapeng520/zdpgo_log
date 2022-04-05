package zap

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"
)

const (
	_stdLogDefaultDepth      = 1
	_loggerWriterDepth       = 2
	_programmerErrorTemplate = "您发现了一个Bug，请到https://github.com/zhangdapeng520/zdpgo_log/issues/new提交该错误: %v"
)

var (
	_globalMu sync.RWMutex       // 全局的互斥锁
	_globalL  = NewNop()         // 全局的Logger日志对象
	_globalS  = _globalL.Sugar() // 全局的Sugar日志对象
)

// L 返回使用ReplaceGlobals替换的全局Logger对象，是线程安全的
func L() *Logger {
	_globalMu.RLock()
	l := _globalL
	_globalMu.RUnlock()
	return l
}

// S 返回使用ReplaceGlobals替换的SugaredLogger对象，这是线程安全的
func S() *SugaredLogger {
	_globalMu.RLock()
	s := _globalS
	_globalMu.RUnlock()
	return s
}

// ReplaceGlobals 替换全局的Logger和SugaredLogger对象，返回一个函数，执行该函数会恢复之前的全局对象。这是线程安全的方法。
func ReplaceGlobals(logger *Logger) func() {
	_globalMu.Lock()
	prev := _globalL
	_globalL = logger
	_globalS = logger.Sugar()
	_globalMu.Unlock()
	return func() { ReplaceGlobals(prev) }
}

// NewStdLog 返回一个*log.Logger对象，使用info级别的Logger
func NewStdLog(l *Logger) *log.Logger {
	logger := l.WithOptions(AddCallerSkip(_stdLogDefaultDepth + _loggerWriterDepth))
	f := logger.Info
	return log.New(&loggerWriter{f}, "" /* prefix */, 0 /* flags */)
}

// NewStdLogAt 使用指定的日志界别创建日志对象
func NewStdLogAt(l *Logger, level zapcore.Level) (*log.Logger, error) {
	logger := l.WithOptions(AddCallerSkip(_stdLogDefaultDepth + _loggerWriterDepth))
	logFunc, err := levelToFunc(logger, level)
	if err != nil {
		return nil, err
	}
	return log.New(&loggerWriter{logFunc}, "" /* prefix */, 0 /* flags */), nil
}

// RedirectStdLog 将标准库的包全局记录器的输出重定向到info level提供的记录器。
// 因为zap已经处理了调用者注解、时间戳等，所以它会自动禁用标准库的注解和前缀。
// 它返回一个函数来恢复原来的前缀和标志，并将标准库的输出重置为os.Stderr。
func RedirectStdLog(l *Logger) func() {
	f, err := redirectStdLogAt(l, InfoLevel)
	if err != nil {
		// Can't get here, since passing InfoLevel to redirectStdLogAt always
		// works.
		panic(fmt.Sprintf(_programmerErrorTemplate, err))
	}
	return f
}

// RedirectStdLogAt 将输出从标准库的包全局记录器重定向到指定级别上提供的记录器。
// 因为zap已经处理了调用者注解、时间戳等，所以它会自动禁用标准库的注解和前缀。
// 它返回一个函数来恢复原来的前缀和标志，并将标准库的输出重置为os.Stderr。
func RedirectStdLogAt(l *Logger, level zapcore.Level) (func(), error) {
	return redirectStdLogAt(l, level)
}

func redirectStdLogAt(l *Logger, level zapcore.Level) (func(), error) {
	flags := log.Flags()
	prefix := log.Prefix()
	log.SetFlags(0)
	log.SetPrefix("")
	logger := l.WithOptions(AddCallerSkip(_stdLogDefaultDepth + _loggerWriterDepth))
	logFunc, err := levelToFunc(logger, level)
	if err != nil {
		return nil, err
	}
	log.SetOutput(&loggerWriter{logFunc})
	return func() {
		log.SetFlags(flags)
		log.SetPrefix(prefix)
		log.SetOutput(os.Stderr)
	}, nil
}

func levelToFunc(logger *Logger, lvl zapcore.Level) (func(string, ...Field), error) {
	switch lvl {
	case DebugLevel:
		return logger.Debug, nil
	case InfoLevel:
		return logger.Info, nil
	case WarnLevel:
		return logger.Warn, nil
	case ErrorLevel:
		return logger.Error, nil
	case DPanicLevel:
		return logger.DPanic, nil
	case PanicLevel:
		return logger.Panic, nil
	case FatalLevel:
		return logger.Fatal, nil
	}
	return nil, fmt.Errorf("unrecognized level: %q", lvl)
}

type loggerWriter struct {
	logFunc func(msg string, fields ...Field)
}

func (l *loggerWriter) Write(p []byte) (int, error) {
	p = bytes.TrimSpace(p)
	l.logFunc(string(p))
	return len(p), nil
}
