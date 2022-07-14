package goleak

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_log/goleak/internal/stack"
)

// TestingT is the minimal subset of testing.TB that we use.
type TestingT interface {
	Error(...interface{})
}

// filterStacks will filter any stacks excluded by the given opts.
// filterStacks modifies the passed in stacks slice.
func filterStacks(stacks []stack.Stack, skipID int, opts *opts) []stack.Stack {
	filtered := stacks[:0]
	for _, stack := range stacks {
		// Always skip the running goroutine.
		if stack.ID() == skipID {
			continue
		}
		// Run any default or user-specified filters.
		if opts.filter(stack) {
			continue
		}
		filtered = append(filtered, stack)
	}
	return filtered
}

// Find looks for extra goroutines, and returns a descriptive error if
// any are found.
func Find(options ...Option) error {
	cur := stack.Current().ID()

	opts := buildOpts(options...)
	var stacks []stack.Stack
	retry := true
	for i := 0; retry; i++ {
		stacks = filterStacks(stack.All(), cur, opts)

		if len(stacks) == 0 {
			return nil
		}
		retry = opts.retry(i)
	}

	return fmt.Errorf("found unexpected goroutines:\n%s", stacks)
}

// VerifyNone marks the given TestingT as failed if any extra goroutines are
// found by Find. This is a helper method to make it easier to integrate in
// tests by doing:
// 	defer VerifyNone(t)
func VerifyNone(t TestingT, options ...Option) {
	if err := Find(options...); err != nil {
		t.Error(err)
	}
}
