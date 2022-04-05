package zapcore_test

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"
)

func BenchmarkJSONLogMarshalerFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		enc := NewJSONEncoder(testEncoderConfig())
		enc.AddObject("nested", ObjectMarshalerFunc(func(enc ObjectEncoder) error {
			enc.AddInt64("i", int64(i))
			return nil
		}))
	}
}

func BenchmarkZapJSONFloat32AndComplex64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			enc := NewJSONEncoder(testEncoderConfig())
			enc.AddFloat32("float32", 3.14)
			enc.AddComplex64("complex64", 2.71+3.14i)
		}
	})
}

func BenchmarkZapJSON(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			enc := NewJSONEncoder(testEncoderConfig())
			enc.AddString("str", "foo")
			enc.AddInt64("int64-1", 1)
			enc.AddInt64("int64-2", 2)
			enc.AddFloat64("float64", 1.0)
			enc.AddString("string1", "\n")
			enc.AddString("string2", "💩")
			enc.AddString("string3", "🤔")
			enc.AddString("string4", "🙊")
			enc.AddBool("bool", true)
			buf, _ := enc.EncodeEntry(Entry{
				Message: "fake",
				Level:   DebugLevel,
			}, nil)
			buf.Free()
		}
	})
}

func BenchmarkStandardJSON(b *testing.B) {
	record := struct {
		Level   string                 `json:"level"`
		Message string                 `json:"msg"`
		Time    time.Time              `json:"ts"`
		Fields  map[string]interface{} `json:"fields"`
	}{
		Level:   "debug",
		Message: "fake",
		Time:    time.Unix(0, 0),
		Fields: map[string]interface{}{
			"str":     "foo",
			"int64-1": int64(1),
			"int64-2": int64(1),
			"float64": float64(1.0),
			"string1": "\n",
			"string2": "💩",
			"string3": "🤔",
			"string4": "🙊",
			"bool":    true,
		},
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			json.Marshal(record)
		}
	})
}
