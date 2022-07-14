package atomic

// Error is an atomic type-safe wrapper for error values.
type Error struct {
	_ nocmp // disallow non-atomic comparison

	v Value
}

var _zeroError error

// NewError creates a new Error.
func NewError(v error) *Error {
	x := &Error{}
	if v != _zeroError {
		x.Store(v)
	}
	return x
}

// Load atomically loads the wrapped error.
func (x *Error) Load() error {
	return unpackError(x.v.Load())
}

// Store atomically stores the passed error.
func (x *Error) Store(v error) {
	x.v.Store(packError(v))
}
