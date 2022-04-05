package zapcore_test

import (
	"testing"

	"github.com/zhangdapeng520/zdpgo_log/libs/zap/internal/ztest"
	. "github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"
)

func withBenchedTee(b *testing.B, f func(Core)) {
	fac := NewTee(
		NewCore(NewJSONEncoder(testEncoderConfig()), &ztest.Discarder{}, DebugLevel),
		NewCore(NewJSONEncoder(testEncoderConfig()), &ztest.Discarder{}, InfoLevel),
	)
	b.ResetTimer()
	f(fac)
}

func BenchmarkTeeCheck(b *testing.B) {
	cases := []struct {
		lvl Level
		msg string
	}{
		{DebugLevel, "foo"},
		{InfoLevel, "bar"},
		{WarnLevel, "baz"},
		{ErrorLevel, "babble"},
	}
	withBenchedTee(b, func(core Core) {
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				tt := cases[i]
				entry := Entry{Level: tt.lvl, Message: tt.msg}
				if cm := core.Check(entry, nil); cm != nil {
					cm.Write(Field{Key: "i", Integer: int64(i), Type: Int64Type})
				}
				i = (i + 1) % len(cases)
			}
		})
	})
}
