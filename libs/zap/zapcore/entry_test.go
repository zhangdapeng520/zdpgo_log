package zapcore

import (
	"sync"
	"testing"

	"github.com/zhangdapeng520/zdpgo_log/libs/zap/internal/exit"

	"github.com/stretchr/testify/assert"
)

func assertGoexit(t *testing.T, f func()) {
	var finished bool
	recovered := make(chan interface{})
	go func() {
		defer func() {
			recovered <- recover()
		}()

		f()
		finished = true
	}()

	assert.Nil(t, <-recovered, "Goexit should cause recover to return nil")
	assert.False(t, finished, "Goroutine should not finish after Goexit")
}

func TestPutNilEntry(t *testing.T) {
	// Pooling nil entries defeats the purpose.
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			putCheckedEntry(nil)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			ce := getCheckedEntry()
			assert.NotNil(t, ce, "Expected only non-nil CheckedEntries in pool.")
			assert.False(t, ce.dirty, "Unexpected dirty bit set.")
			assert.Nil(t, ce.ErrorOutput, "Non-nil ErrorOutput.")
			assert.Equal(t, WriteThenNoop, ce.should, "Unexpected terminal behavior.")
			assert.Equal(t, 0, len(ce.cores), "Expected empty slice of cores.")
			assert.True(t, cap(ce.cores) > 0, "Expected pooled CheckedEntries to pre-allocate slice of Cores.")
		}
	}()

	wg.Wait()
}

func TestEntryCaller(t *testing.T) {
	tests := []struct {
		caller EntryCaller
		full   string
		short  string
	}{
		{
			caller: NewEntryCaller(100, "/path/to/foo.go", 42, false),
			full:   "undefined",
			short:  "undefined",
		},
		{
			caller: NewEntryCaller(100, "/path/to/foo.go", 42, true),
			full:   "/path/to/foo.go:42",
			short:  "to/foo.go:42",
		},
		{
			caller: NewEntryCaller(100, "to/foo.go", 42, true),
			full:   "to/foo.go:42",
			short:  "to/foo.go:42",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.full, tt.caller.String(), "Unexpected string from EntryCaller.")
		assert.Equal(t, tt.full, tt.caller.FullPath(), "Unexpected FullPath from EntryCaller.")
		assert.Equal(t, tt.short, tt.caller.TrimmedPath(), "Unexpected TrimmedPath from EntryCaller.")
	}
}

func TestCheckedEntryWrite(t *testing.T) {
	t.Run("nil is safe", func(t *testing.T) {
		var ce *CheckedEntry
		assert.NotPanics(t, func() { ce.Write() }, "Unexpected panic writing nil CheckedEntry.")
	})

	t.Run("WriteThenPanic", func(t *testing.T) {
		var ce *CheckedEntry
		ce = ce.Should(Entry{}, WriteThenPanic)
		assert.Panics(t, func() { ce.Write() }, "Expected to panic when WriteThenPanic is set.")
	})

	t.Run("WriteThenGoexit", func(t *testing.T) {
		var ce *CheckedEntry
		ce = ce.Should(Entry{}, WriteThenGoexit)
		assertGoexit(t, func() { ce.Write() })
	})

	t.Run("WriteThenFatal", func(t *testing.T) {
		var ce *CheckedEntry
		ce = ce.Should(Entry{}, WriteThenFatal)
		stub := exit.WithStub(func() {
			ce.Write()
		})
		assert.True(t, stub.Exited, "Expected to exit when WriteThenFatal is set.")
	})
}
