package zapcore

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhangdapeng520/zdpgo_log/libs/zap/internal/ztest"
)

func BenchmarkMultiWriteSyncer(b *testing.B) {
	b.Run("2 discarder", func(b *testing.B) {
		w := NewMultiWriteSyncer(
			&ztest.Discarder{},
			&ztest.Discarder{},
		)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				w.Write([]byte("foobarbazbabble"))
			}
		})
	})
	b.Run("4 discarder", func(b *testing.B) {
		w := NewMultiWriteSyncer(
			&ztest.Discarder{},
			&ztest.Discarder{},
			&ztest.Discarder{},
			&ztest.Discarder{},
		)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				w.Write([]byte("foobarbazbabble"))
			}
		})
	})
	b.Run("4 discarder with buffer", func(b *testing.B) {
		w := &BufferedWriteSyncer{
			WS: NewMultiWriteSyncer(
				&ztest.Discarder{},
				&ztest.Discarder{},
				&ztest.Discarder{},
				&ztest.Discarder{},
			),
		}
		defer w.Stop()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				w.Write([]byte("foobarbazbabble"))
			}
		})
	})
}

func BenchmarkWriteSyncer(b *testing.B) {
	b.Run("write file with no buffer", func(b *testing.B) {
		file, err := ioutil.TempFile("", "log")
		assert.NoError(b, err)
		defer file.Close()
		defer os.Remove(file.Name())

		w := AddSync(file)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				w.Write([]byte("foobarbazbabble"))
			}
		})
	})
}
