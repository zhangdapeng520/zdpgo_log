package atomic

import (
	"encoding/json"
	"time"
)

// Duration is an atomic type-safe wrapper for time.Duration values.
type Duration struct {
	_ nocmp // disallow non-atomic comparison

	v Int64
}

var _zeroDuration time.Duration

// NewDuration creates a new Duration.
func NewDuration(v time.Duration) *Duration {
	x := &Duration{}
	if v != _zeroDuration {
		x.Store(v)
	}
	return x
}

// Load atomically loads the wrapped time.Duration.
func (x *Duration) Load() time.Duration {
	return time.Duration(x.v.Load())
}

// Store atomically stores the passed time.Duration.
func (x *Duration) Store(v time.Duration) {
	x.v.Store(int64(v))
}

// CAS is an atomic compare-and-swap for time.Duration values.
func (x *Duration) CAS(o, n time.Duration) bool {
	return x.v.CAS(int64(o), int64(n))
}

// Swap atomically stores the given time.Duration and returns the old
// value.
func (x *Duration) Swap(o time.Duration) time.Duration {
	return time.Duration(x.v.Swap(int64(o)))
}

// MarshalJSON encodes the wrapped time.Duration into JSON.
func (x *Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Load())
}

// UnmarshalJSON decodes a time.Duration from JSON.
func (x *Duration) UnmarshalJSON(b []byte) error {
	var v time.Duration
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	x.Store(v)
	return nil
}
