package zapcore_test

import (
	"testing"

	. "github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"
)

func BenchmarkZapConsole(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			enc := NewConsoleEncoder(humanEncoderConfig())
			enc.AddString("str", "foo")
			enc.AddInt64("int64-1", 1)
			enc.AddInt64("int64-2", 2)
			enc.AddFloat64("float64", 1.0)
			enc.AddString("string1", "\n")
			enc.AddString("string2", "💩")
			enc.AddString("string3", "🤔")
			enc.AddString("string4", "🙊")
			enc.AddBool("bool", true)
			buf, _ := enc.EncodeEntry(Entry{
				Message: "fake",
				Level:   DebugLevel,
			}, nil)
			buf.Free()
		}
	})
}
