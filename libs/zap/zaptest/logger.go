package zaptest

import (
	"bytes"

	"github.com/zhangdapeng520/zdpgo_log/libs/zap"
	"github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"
)

// LoggerOption configures the test logger built by NewLogger.
type LoggerOption interface {
	applyLoggerOption(*loggerOptions)
}

type loggerOptions struct {
	Level      zapcore.LevelEnabler
	zapOptions []zap.Option
}

type loggerOptionFunc func(*loggerOptions)

func (f loggerOptionFunc) applyLoggerOption(opts *loggerOptions) {
	f(opts)
}

// Level controls which messages are logged by a test Logger built by
// NewLogger.
func Level(enab zapcore.LevelEnabler) LoggerOption {
	return loggerOptionFunc(func(opts *loggerOptions) {
		opts.Level = enab
	})
}

// WrapOptions adds zap.Option's to a test Logger built by NewLogger.
func WrapOptions(zapOpts ...zap.Option) LoggerOption {
	return loggerOptionFunc(func(opts *loggerOptions) {
		opts.zapOptions = zapOpts
	})
}

// NewLogger builds a new Logger that logs all messages to the given
// testing.TB.
//
//   logger := zaptest.NewLogger(t)
//
// Use this with a *testing.T or *testing.B to get logs which get printed only
// if a test fails or if you ran go test -v.
//
// The returned logger defaults to logging debug level messages and above.
// This may be changed by passing a zaptest.Level during construction.
//
//   logger := zaptest.NewLogger(t, zaptest.Level(zap.WarnLevel))
//
// You may also pass zap.Option's to customize test logger.
//
//   logger := zaptest.NewLogger(t, zaptest.WrapOptions(zap.AddCaller()))
func NewLogger(t TestingT, opts ...LoggerOption) *zap.Logger {
	cfg := loggerOptions{
		Level: zapcore.DebugLevel,
	}
	for _, o := range opts {
		o.applyLoggerOption(&cfg)
	}

	writer := newTestingWriter(t)
	zapOptions := []zap.Option{
		// Send zap errors to the same writer and mark the test as failed if
		// that happens.
		zap.ErrorOutput(writer.WithMarkFailed(true)),
	}
	zapOptions = append(zapOptions, cfg.zapOptions...)

	return zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			writer,
			cfg.Level,
		),
		zapOptions...,
	)
}

// testingWriter is a WriteSyncer that writes to the given testing.TB.
type testingWriter struct {
	t TestingT

	// If true, the test will be marked as failed if this testingWriter is
	// ever used.
	markFailed bool
}

func newTestingWriter(t TestingT) testingWriter {
	return testingWriter{t: t}
}

// WithMarkFailed returns a copy of this testingWriter with markFailed set to
// the provided value.
func (w testingWriter) WithMarkFailed(v bool) testingWriter {
	w.markFailed = v
	return w
}

func (w testingWriter) Write(p []byte) (n int, err error) {
	n = len(p)

	// Strip trailing newline because t.Log always adds one.
	p = bytes.TrimRight(p, "\n")

	// Note: t.Log is safe for concurrent use.
	w.t.Logf("%s", p)
	if w.markFailed {
		w.t.Fail()
	}

	return n, nil
}

func (w testingWriter) Sync() error {
	return nil
}
