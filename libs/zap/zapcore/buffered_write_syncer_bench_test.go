package zapcore

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkBufferedWriteSyncer(b *testing.B) {
	b.Run("write file with buffer", func(b *testing.B) {
		file, err := ioutil.TempFile("", "log")
		require.NoError(b, err)

		defer func() {
			assert.NoError(b, file.Close())
			assert.NoError(b, os.Remove(file.Name()))
		}()

		w := &BufferedWriteSyncer{
			WS: AddSync(file),
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
