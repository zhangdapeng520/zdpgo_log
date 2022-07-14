package atomic

import (
	"encoding/json"
)

// Bool is an atomic type-safe wrapper for bool values.
type Bool struct {
	_ nocmp // disallow non-atomic comparison

	v Uint32
}

var _zeroBool bool

// NewBool creates a new Bool.
func NewBool(v bool) *Bool {
	x := &Bool{}
	if v != _zeroBool {
		x.Store(v)
	}
	return x
}

// Load atomically loads the wrapped bool.
func (x *Bool) Load() bool {
	return truthy(x.v.Load())
}

// Store atomically stores the passed bool.
func (x *Bool) Store(v bool) {
	x.v.Store(boolToInt(v))
}

// CAS is an atomic compare-and-swap for bool values.
func (x *Bool) CAS(o, n bool) bool {
	return x.v.CAS(boolToInt(o), boolToInt(n))
}

// Swap atomically stores the given bool and returns the old
// value.
func (x *Bool) Swap(o bool) bool {
	return truthy(x.v.Swap(boolToInt(o)))
}

// MarshalJSON encodes the wrapped bool into JSON.
func (x *Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Load())
}

// UnmarshalJSON decodes a bool from JSON.
func (x *Bool) UnmarshalJSON(b []byte) error {
	var v bool
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	x.Store(v)
	return nil
}
