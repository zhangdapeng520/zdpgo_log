package zaptest

import "testing"

// Just a compile-time test to ensure that TestingT matches the testing.TB
// interface. We could do this in testingt.go but that would put a dependency
// on the "testing" package from zaptest.

var _ TestingT = (testing.TB)(nil)
