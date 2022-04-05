package zapcore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"
)

var (
	testEntry = Entry{
		LoggerName: "main",
		Level:      InfoLevel,
		Message:    `hello`,
		Time:       _epoch,
		Stack:      "fake-stack",
		Caller:     EntryCaller{Defined: true, File: "foo.go", Line: 42, Function: "foo.Foo"},
	}
)

func TestConsoleSeparator(t *testing.T) {
	tests := []struct {
		desc        string
		separator   string
		wantConsole string
	}{
		{
			desc:        "space console separator",
			separator:   " ",
			wantConsole: "0 info main foo.go:42 foo.Foo hello\nfake-stack\n",
		},
		{
			desc:        "default console separator",
			separator:   "",
			wantConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc:        "tag console separator",
			separator:   "\t",
			wantConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc:        "dash console separator",
			separator:   "--",
			wantConsole: "0--info--main--foo.go:42--foo.Foo--hello\nfake-stack\n",
		},
	}

	for _, tt := range tests {
		console := NewConsoleEncoder(encoderTestEncoderConfig(tt.separator))
		t.Run(tt.desc, func(t *testing.T) {
			entry := testEntry
			consoleOut, err := console.EncodeEntry(entry, nil)
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(
				t,
				tt.wantConsole,
				consoleOut.String(),
				"Unexpected console output",
			)
		})

	}
}

func encoderTestEncoderConfig(separator string) EncoderConfig {
	testEncoder := testEncoderConfig()
	testEncoder.ConsoleSeparator = separator
	return testEncoder
}
