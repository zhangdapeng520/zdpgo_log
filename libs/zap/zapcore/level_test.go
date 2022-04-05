package zapcore

import (
	"bytes"
	"flag"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试日志级别字符串
func TestLevelString(t *testing.T) {
	// 表格驱动测试：参数、期望值
	tests := map[Level]string{
		DebugLevel:  "debug",
		InfoLevel:   "info",
		WarnLevel:   "warn",
		ErrorLevel:  "error",
		DPanicLevel: "dpanic",
		PanicLevel:  "panic",
		FatalLevel:  "fatal",
		Level(-42):  "Level(-42)",
	}

	// 遍历参数和期望值
	for lvl, stringLevel := range tests {
		// 字符串表示符合期望
		assert.Equal(t, stringLevel, lvl.String(), "意外的小写级别字符串。")

		// 字符串大写表示符合期望
		assert.Equal(t, strings.ToUpper(stringLevel), lvl.CapitalString(), "意外的全大写级别字符串。")
	}
}

// 测试日志级别文本
func TestLevelText(t *testing.T) {
	tests := []struct {
		text  string
		level Level
	}{
		{"debug", DebugLevel},
		{"info", InfoLevel},
		{"", InfoLevel}, // 默认的日志级别
		{"warn", WarnLevel},
		{"error", ErrorLevel},
		{"dpanic", DPanicLevel},
		{"panic", PanicLevel},
		{"fatal", FatalLevel},
	}
	// 日志文本，日志级别
	for _, tt := range tests {
		if tt.text != "" {
			// 日志级别
			lvl := tt.level

			// 序列化文本
			marshaled, err := lvl.MarshalText()

			// 期望不报错
			assert.NoError(t, err, "解析日志级别%v错误", &lvl)

			// 期望相等
			assert.Equal(t, tt.text, string(marshaled), "解析日志级别%v错误", &lvl)
		}

		var unmarshaled Level
		err := unmarshaled.UnmarshalText([]byte(tt.text))
		assert.NoError(t, err, `反向解析字符串%q错误`, tt.text)
		assert.Equal(t, tt.level, unmarshaled, `反向解析字符串%q错误`, tt.text)
	}
}

// 测试解析全大写级别
func TestCapitalLevelsParse(t *testing.T) {
	tests := []struct {
		text  string
		level Level
	}{
		{"DEBUG", DebugLevel},
		{"INFO", InfoLevel},
		{"WARN", WarnLevel},
		{"ERROR", ErrorLevel},
		{"DPANIC", DPanicLevel},
		{"PANIC", PanicLevel},
		{"FATAL", FatalLevel},
	}
	for _, tt := range tests {
		var unmarshaled Level
		err := unmarshaled.UnmarshalText([]byte(tt.text))
		assert.NoError(t, err, `反向解析字符串%q错误`, tt.text)
		assert.Equal(t, tt.level, unmarshaled, `反向解析字符串%q错误`, tt.text)
	}
}

func TestWeirdLevelsParse(t *testing.T) {
	tests := []struct {
		text  string
		level Level
	}{
		// I guess...
		{"Debug", DebugLevel},
		{"Info", InfoLevel},
		{"Warn", WarnLevel},
		{"Error", ErrorLevel},
		{"Dpanic", DPanicLevel},
		{"Panic", PanicLevel},
		{"Fatal", FatalLevel},

		// What even is...
		{"DeBuG", DebugLevel},
		{"InFo", InfoLevel},
		{"WaRn", WarnLevel},
		{"ErRor", ErrorLevel},
		{"DpAnIc", DPanicLevel},
		{"PaNiC", PanicLevel},
		{"FaTaL", FatalLevel},
	}
	for _, tt := range tests {
		var unmarshaled Level
		err := unmarshaled.UnmarshalText([]byte(tt.text))
		assert.NoError(t, err, `Unexpected error unmarshaling text %q to level.`, tt.text)
		assert.Equal(t, tt.level, unmarshaled, `Text %q unmarshaled to an unexpected level.`, tt.text)
	}
}

func TestLevelNils(t *testing.T) {
	var l *Level

	// The String() method will not handle nil level properly.
	assert.Panics(t, func() {
		assert.Equal(t, "Level(nil)", l.String(), "Unexpected result stringifying nil *Level.")
	}, "Level(nil).String() should panic")

	assert.Panics(t, func() {
		l.MarshalText()
	}, "Expected to panic when marshalling a nil level.")

	err := l.UnmarshalText([]byte("debug"))
	assert.Equal(t, errUnmarshalNilLevel, err, "Expected to error unmarshalling into a nil Level.")
}

func TestLevelUnmarshalUnknownText(t *testing.T) {
	var l Level
	err := l.UnmarshalText([]byte("foo"))
	assert.Contains(t, err.Error(), "unrecognized level", "Expected unmarshaling arbitrary text to fail.")
}

func TestLevelAsFlagValue(t *testing.T) {
	var (
		buf bytes.Buffer
		lvl Level
	)
	fs := flag.NewFlagSet("levelTest", flag.ContinueOnError)
	fs.SetOutput(&buf)
	fs.Var(&lvl, "level", "log level")

	for _, expected := range []Level{DebugLevel, InfoLevel, WarnLevel, ErrorLevel, DPanicLevel, PanicLevel, FatalLevel} {
		assert.NoError(t, fs.Parse([]string{"-level", expected.String()}))
		assert.Equal(t, expected, lvl, "Unexpected level after parsing flag.")
		assert.Equal(t, expected, lvl.Get(), "Unexpected output using flag.Getter API.")
		assert.Empty(t, buf.String(), "Unexpected error output parsing level flag.")
		buf.Reset()
	}

	assert.Error(t, fs.Parse([]string{"-level", "nope"}))
	assert.Equal(
		t,
		`invalid value "nope" for flag -level: unrecognized level: "nope"`,
		strings.Split(buf.String(), "\n")[0], // second line is help message
		"Unexpected error output from invalid flag input.",
	)
}
