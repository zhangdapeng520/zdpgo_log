package zaptest

import "github.com/zhangdapeng520/zdpgo_log/zap/internal/ztest"

type (
	// A Syncer is a spy for the Sync portion of zapcore.WriteSyncer.
	Syncer = ztest.Syncer

	// A Discarder sends all writes to ioutil.Discard.
	Discarder = ztest.Discarder

	// FailWriter is a WriteSyncer that always returns an error on writes.
	FailWriter = ztest.FailWriter

	// ShortWriter is a WriteSyncer whose write method never returns an error,
	// but always reports that it wrote one byte less than the input slice's
	// length (thus, a "short write").
	ShortWriter = ztest.ShortWriter

	// Buffer is an implementation of zapcore.WriteSyncer that sends all writes to
	// a bytes.Buffer. It has convenience methods to split the accumulated buffer
	// on newlines.
	Buffer = ztest.Buffer
)
