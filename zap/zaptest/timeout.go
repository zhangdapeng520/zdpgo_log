package zaptest

import (
	"time"

	"github.com/zhangdapeng520/zdpgo_log/zap/internal/ztest"
)

// Timeout scales the provided duration by $TEST_TIMEOUT_SCALE.
//
// Deprecated: This function is intended for internal testing and shouldn't be
// used outside zap itself. It was introduced before Go supported internal
// packages.
func Timeout(base time.Duration) time.Duration {
	return ztest.Timeout(base)
}

// Sleep scales the sleep duration by $TEST_TIMEOUT_SCALE.
//
// Deprecated: This function is intended for internal testing and shouldn't be
// used outside zap itself. It was introduced before Go supported internal
// packages.
func Sleep(base time.Duration) {
	ztest.Sleep(base)
}
