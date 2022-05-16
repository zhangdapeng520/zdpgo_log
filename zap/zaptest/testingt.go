package zaptest

// TestingT is a subset of the API provided by all *testing.T and *testing.B
// objects.
type TestingT interface {
	// Logs the given message without failing the test.
	Logf(string, ...interface{})

	// Logs the given message and marks the test as failed.
	Errorf(string, ...interface{})

	// Marks the test as failed.
	Fail()

	// Returns true if the test has been marked as failed.
	Failed() bool

	// Returns the name of the test.
	Name() string

	// Marks the test as failed and stops execution of that test.
	FailNow()
}

// Note: We currently only rely on Logf. We are including Errorf and FailNow
// in the interface in anticipation of future need since we can't extend the
// interface without a breaking change.
