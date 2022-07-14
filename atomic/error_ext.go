package atomic

// atomic.Value panics on nil inputs, or if the underlying type changes.
// Stabilize by always storing a custom struct that we control.

//go:generate bin/gen-atomicwrapper -name=Error -type=error -wrapped=Value -pack=packError -unpack=unpackError -file=error.go

type packedError struct{ Value error }

func packError(v error) interface{} {
	return packedError{v}
}

func unpackError(v interface{}) error {
	if err, ok := v.(packedError); ok {
		return err.Value
	}
	return nil
}
