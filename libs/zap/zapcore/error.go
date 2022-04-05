package zapcore

import (
	"fmt"
	"reflect"
	"sync"
)

// Encodes the given error into fields of an object. A field with the given
// name is added for the error message.
//
// If the error implements fmt.Formatter, a field with the name ${key}Verbose
// is also added with the full verbose error message.
//
// Finally, if the error implements errorGroup (from go.uber.org/multierr) or
// causer (from github.com/pkg/errors), a ${key}Causes field is added with an
// array of objects containing the errors this error was comprised of.
//
//  {
//    "error": err.Error(),
//    "errorVerbose": fmt.Sprintf("%+v", err),
//    "errorCauses": [
//      ...
//    ],
//  }
func encodeError(key string, err error, enc ObjectEncoder) (retErr error) {
	// Try to capture panics (from nil references or otherwise) when calling
	// the Error() method
	defer func() {
		if rerr := recover(); rerr != nil {
			// If it's a nil pointer, just say "<nil>". The likeliest causes are a
			// error that fails to guard against nil or a nil pointer for a
			// value receiver, and in either case, "<nil>" is a nice result.
			if v := reflect.ValueOf(err); v.Kind() == reflect.Ptr && v.IsNil() {
				enc.AddString(key, "<nil>")
				return
			}

			retErr = fmt.Errorf("PANIC=%v", rerr)
		}
	}()

	basic := err.Error()
	enc.AddString(key, basic)

	switch e := err.(type) {
	case errorGroup:
		return enc.AddArray(key+"Causes", errArray(e.Errors()))
	case fmt.Formatter:
		verbose := fmt.Sprintf("%+v", e)
		if verbose != basic {
			// This is a rich error type, like those produced by
			// github.com/pkg/errors.
			enc.AddString(key+"Verbose", verbose)
		}
	}
	return nil
}

type errorGroup interface {
	// Provides read-only access to the underlying list of errors, preferably
	// without causing any allocs.
	Errors() []error
}

// Note that errArray and errArrayElem are very similar to the version
// implemented in the top-level error.go file. We can't re-use this because
// that would require exporting errArray as part of the zapcore API.

// Encodes a list of errors using the standard error encoding logic.
type errArray []error

func (errs errArray) MarshalLogArray(arr ArrayEncoder) error {
	for i := range errs {
		if errs[i] == nil {
			continue
		}

		el := newErrArrayElem(errs[i])
		arr.AppendObject(el)
		el.Free()
	}
	return nil
}

var _errArrayElemPool = sync.Pool{New: func() interface{} {
	return &errArrayElem{}
}}

// Encodes any error into a {"error": ...} re-using the same errors logic.
//
// May be passed in place of an array to build a single-element array.
type errArrayElem struct{ err error }

func newErrArrayElem(err error) *errArrayElem {
	e := _errArrayElemPool.Get().(*errArrayElem)
	e.err = err
	return e
}

func (e *errArrayElem) MarshalLogArray(arr ArrayEncoder) error {
	return arr.AppendObject(e)
}

func (e *errArrayElem) MarshalLogObject(enc ObjectEncoder) error {
	return encodeError("error", e.err, enc)
}

func (e *errArrayElem) Free() {
	e.err = nil
	_errArrayElemPool.Put(e)
}
