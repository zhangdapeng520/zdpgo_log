//go:build go1.13
// +build go1.13

package multierr

import "errors"

// As attempts to find the first error in the error list that matches the type
// of the value that target points to.
//
// This function allows errors.As to traverse the values stored on the
// multierr error.
func (merr *multiError) As(target interface{}) bool {
	for _, err := range merr.Errors() {
		if errors.As(err, target) {
			return true
		}
	}
	return false
}

// Is attempts to match the provided error against errors in the error list.
//
// This function allows errors.Is to traverse the values stored on the
// multierr error.
func (merr *multiError) Is(target error) bool {
	for _, err := range merr.Errors() {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}
