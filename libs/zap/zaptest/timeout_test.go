package zaptest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhangdapeng520/zdpgo_log/libs/zap/internal/ztest"
)

func TestTimeout(t *testing.T) {
	defer ztest.Initialize("2")()
	assert.Equal(t, time.Duration(100), Timeout(50), "Expected to scale up timeout.")
}

func TestSleep(t *testing.T) {
	defer ztest.Initialize("2")()
	const sleepFor = 50 * time.Millisecond
	now := time.Now()
	Sleep(sleepFor)
	elapsed := time.Since(now)
	assert.True(t, 2*sleepFor <= elapsed, "Expected to scale up timeout.")
}
