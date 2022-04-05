package zapcore

import (
	"testing"
	"time"

	"github.com/zhangdapeng520/zdpgo_log/libs/zap/internal/ztest"
)

// Verify that the mock clock satisfies the Clock interface.
var _ Clock = (*ztest.MockClock)(nil)

func TestSystemClock_NewTicker(t *testing.T) {
	want := 3

	var n int
	timer := DefaultClock.NewTicker(time.Millisecond)
	for range timer.C {
		n++
		if n == want {
			return
		}
	}
}
