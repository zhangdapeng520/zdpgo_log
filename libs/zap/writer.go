package zap

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"

	"go.uber.org/multierr"
)

// Open is a high-level wrapper that takes a variadic number of URLs, opens or
// creates each of the specified resources, and combines them into a locked
// WriteSyncer. It also returns any error encountered and a function to close
// any opened files.
//
// Passing no URLs returns a no-op WriteSyncer. Zap handles URLs without a
// scheme and URLs with the "file" scheme. Third-party code may register
// factories for other schemes using RegisterSink.
//
// URLs with the "file" scheme must use absolute paths on the local
// filesystem. No user, password, port, fragments, or query parameters are
// allowed, and the hostname must be empty or "localhost".
//
// Since it's common to write logs to the local filesystem, URLs without a
// scheme (e.g., "/var/log/foo.log") are treated as local file paths. Without
// a scheme, the special paths "stdout" and "stderr" are interpreted as
// os.Stdout and os.Stderr. When specified without a scheme, relative file
// paths also work.
func Open(paths ...string) (zapcore.WriteSyncer, func(), error) {
	writers, close, err := open(paths)
	if err != nil {
		return nil, nil, err
	}

	writer := CombineWriteSyncers(writers...)
	return writer, close, nil
}

func open(paths []string) ([]zapcore.WriteSyncer, func(), error) {
	writers := make([]zapcore.WriteSyncer, 0, len(paths))
	closers := make([]io.Closer, 0, len(paths))
	close := func() {
		for _, c := range closers {
			c.Close()
		}
	}

	var openErr error
	for _, path := range paths {
		sink, err := newSink(path)
		if err != nil {
			openErr = multierr.Append(openErr, fmt.Errorf("couldn't open sink %q: %v", path, err))
			continue
		}
		writers = append(writers, sink)
		closers = append(closers, sink)
	}
	if openErr != nil {
		close()
		return writers, nil, openErr
	}

	return writers, close, nil
}

// CombineWriteSyncers is a utility that combines multiple WriteSyncers into a
// single, locked WriteSyncer. If no inputs are supplied, it returns a no-op
// WriteSyncer.
//
// It's provided purely as a convenience; the result is no different from
// using zapcore.NewMultiWriteSyncer and zapcore.Lock individually.
func CombineWriteSyncers(writers ...zapcore.WriteSyncer) zapcore.WriteSyncer {
	if len(writers) == 0 {
		return zapcore.AddSync(ioutil.Discard)
	}
	return zapcore.Lock(zapcore.NewMultiWriteSyncer(writers...))
}
