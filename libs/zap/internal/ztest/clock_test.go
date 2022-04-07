package ztest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhangdapeng520/zdpgo_log/libs/atomic"
)

func TestMockClock_NewTicker(t *testing.T) {
	var n atomic.Int32
	clock := NewMockClock()

	done := make(chan struct{})
	defer func() { <-done }() // wait for end

	quit := make(chan struct{})
	// Create a channel to increment every microsecond.
	go func(ticker *time.Ticker) {
		defer close(done)
		for {
			select {
			case <-quit:
				ticker.Stop()
				return
			case <-ticker.C:
				n.Inc()
			}
		}
	}(clock.NewTicker(time.Microsecond))

	// Move clock forward.
	clock.Add(2 * time.Microsecond)
	assert.Equal(t, int32(2), n.Load())
	close(quit)
}
