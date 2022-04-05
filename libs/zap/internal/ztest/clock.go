package ztest

import (
	"time"

	"github.com/benbjohnson/clock"
)

// MockClock provides control over the time.
type MockClock struct{ m *clock.Mock }

// NewMockClock builds a new mock clock that provides control of time.
func NewMockClock() *MockClock {
	return &MockClock{clock.NewMock()}
}

// Now reports the current time.
func (c *MockClock) Now() time.Time {
	return c.m.Now()
}

// NewTicker returns a time.Ticker that ticks at the specified frequency.
func (c *MockClock) NewTicker(d time.Duration) *time.Ticker {
	return &time.Ticker{C: c.m.Ticker(d).C}
}

// Add progresses time by the given duration.
func (c *MockClock) Add(d time.Duration) {
	c.m.Add(d)
}
