package zdpgo_log

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_log/core"
	"github.com/zhangdapeng520/zdpgo_log/multierr"
)

const (
	_oddNumberErrMsg    = "Ignored key without a value."
	_nonStringKeyErrMsg = "Ignored key-value pairs with non-string keys."
)

// SugaredLogger 包含了基础的Logger，但是API更轻量级。任何Logger都能够转换为SugaredLogger的方法。
type SugaredLogger struct {
	base *Logger
}

// Desugar 取消包裹SugaredLogger，暴露一个原始的Logger
func (s *SugaredLogger) Desugar() *Logger {
	base := s.base.clone()
	base.callerSkip -= 2
	return base
}

// Named 修改Logger的名称
func (s *SugaredLogger) Named(name string) *SugaredLogger {
	return &SugaredLogger{base: s.base.Named(name)}
}

// With 添加日志上下文
func (s *SugaredLogger) With(args ...interface{}) *SugaredLogger {
	return &SugaredLogger{base: s.base.With(s.sweetenFields(args)...)}
}

// Debug 使用 fmt.Sprint 结构输出日志消息
func (s *SugaredLogger) Debug(args ...interface{}) {
	s.log(DebugLevel, "", args, nil)
}

// Info 使用 fmt.Sprint 结构化输出日志消息
func (s *SugaredLogger) Info(args ...interface{}) {
	s.log(InfoLevel, "", args, nil)
}

// Warn 使用 fmt.Sprint 结构化输出日志消息
func (s *SugaredLogger) Warn(args ...interface{}) {
	s.log(WarnLevel, "", args, nil)
}

// Error 使用 fmt.Sprint 结构化输出日志消息
func (s *SugaredLogger) Error(args ...interface{}) {
	s.log(ErrorLevel, "", args, nil)
}

// DPanic 使用 fmt.Sprint 结构化输出日志消息
func (s *SugaredLogger) DPanic(args ...interface{}) {
	s.log(DPanicLevel, "", args, nil)
}

// Panic 使用 fmt.Sprint 结构化输出日志消息，然后调用panic
func (s *SugaredLogger) Panic(args ...interface{}) {
	s.log(PanicLevel, "", args, nil)
}

// Fatal 使用 fmt.Sprint 结构化输出日志消息，然后调用 os.Exit.
func (s *SugaredLogger) Fatal(args ...interface{}) {
	s.log(FatalLevel, "", args, nil)
}

// Debugf 使用 fmt.Sprintf 记录模板消息
func (s *SugaredLogger) Debugf(template string, args ...interface{}) {
	s.log(DebugLevel, template, args, nil)
}

// Infof 使用 fmt.Sprintf 记录模板消息
func (s *SugaredLogger) Infof(template string, args ...interface{}) {
	s.log(InfoLevel, template, args, nil)
}

// Warnf 使用 fmt.Sprintf 记录模板消息
func (s *SugaredLogger) Warnf(template string, args ...interface{}) {
	s.log(WarnLevel, template, args, nil)
}

// Errorf 使用 fmt.Sprintf 记录模板消息
func (s *SugaredLogger) Errorf(template string, args ...interface{}) {
	s.log(ErrorLevel, template, args, nil)
}

// DPanicf 使用 fmt.Sprintf 记录模板消息
func (s *SugaredLogger) DPanicf(template string, args ...interface{}) {
	s.log(DPanicLevel, template, args, nil)
}

// Panicf 使用 fmt.Sprintf 记录模板消息
func (s *SugaredLogger) Panicf(template string, args ...interface{}) {
	s.log(PanicLevel, template, args, nil)
}

// Fatalf 使用 fmt.Sprintf 记录模板消息, 然后调用 os.Exit.
func (s *SugaredLogger) Fatalf(template string, args ...interface{}) {
	s.log(FatalLevel, template, args, nil)
}

// Debugw 使用 fmt.Sprintf 记录模板消息
func (s *SugaredLogger) Debugw(msg string, keysAndValues ...interface{}) {
	s.log(DebugLevel, msg, nil, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) Infow(msg string, keysAndValues ...interface{}) {
	s.log(InfoLevel, msg, nil, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) Warnw(msg string, keysAndValues ...interface{}) {
	s.log(WarnLevel, msg, nil, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) Errorw(msg string, keysAndValues ...interface{}) {
	s.log(ErrorLevel, msg, nil, keysAndValues)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) DPanicw(msg string, keysAndValues ...interface{}) {
	s.log(DPanicLevel, msg, nil, keysAndValues)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (s *SugaredLogger) Panicw(msg string, keysAndValues ...interface{}) {
	s.log(PanicLevel, msg, nil, keysAndValues)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (s *SugaredLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	s.log(FatalLevel, msg, nil, keysAndValues)
}

// Sync flushes any buffered log entries.
func (s *SugaredLogger) Sync() error {
	return s.base.Sync()
}

func (s *SugaredLogger) log(lvl core.Level, template string, fmtArgs []interface{}, context []interface{}) {
	// If logging at this level is completely disabled, skip the overhead of
	// string formatting.
	if lvl < DPanicLevel && !s.base.Core().Enabled(lvl) {
		return
	}

	msg := getMessage(template, fmtArgs)
	if ce := s.base.Check(lvl, msg); ce != nil {
		ce.Write(s.sweetenFields(context)...)
	}
}

// getMessage format with Sprint, Sprintf, or neither.
func getMessage(template string, fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}

func (s *SugaredLogger) sweetenFields(args []interface{}) []Field {
	if len(args) == 0 {
		return nil
	}

	// Allocate enough space for the worst case; if users pass only structured
	// fields, we shouldn't penalize them with extra allocations.
	fields := make([]Field, 0, len(args))
	var invalid invalidPairs

	for i := 0; i < len(args); {
		// This is a strongly-typed field. Consume it and move on.
		if f, ok := args[i].(Field); ok {
			fields = append(fields, f)
			i++
			continue
		}

		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			s.base.Error(_oddNumberErrMsg, Any("ignored", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			// Subsequent errors are likely, so allocate once up front.
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(args)/2)
			}
			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			fields = append(fields, Any(keyStr, val))
		}
		i += 2
	}

	// If we encountered any invalid key-value pairs, log an error.
	if len(invalid) > 0 {
		s.base.Error(_nonStringKeyErrMsg, Array("invalid", invalid))
	}
	return fields
}

type invalidPair struct {
	position   int
	key, value interface{}
}

func (p invalidPair) MarshalLogObject(enc core.ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	Any("key", p.key).AddTo(enc)
	Any("value", p.value).AddTo(enc)
	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc core.ArrayEncoder) error {
	var err error
	for i := range ps {
		err = multierr.Append(err, enc.AppendObject(ps[i]))
	}
	return err
}
