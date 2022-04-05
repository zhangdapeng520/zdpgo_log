package observer

import (
	"testing"

	"github.com/zhangdapeng520/zdpgo_log/libs/zap"
	"github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"

	"github.com/stretchr/testify/assert"
)

func TestLoggedEntryContextMap(t *testing.T) {
	tests := []struct {
		msg    string
		fields []zapcore.Field
		want   map[string]interface{}
	}{
		{
			msg:    "no fields",
			fields: nil,
			want:   map[string]interface{}{},
		},
		{
			msg: "simple",
			fields: []zapcore.Field{
				zap.String("k1", "v"),
				zap.Int64("k2", 10),
			},
			want: map[string]interface{}{
				"k1": "v",
				"k2": int64(10),
			},
		},
		{
			msg: "overwrite",
			fields: []zapcore.Field{
				zap.String("k1", "v1"),
				zap.String("k1", "v2"),
			},
			want: map[string]interface{}{
				"k1": "v2",
			},
		},
		{
			msg: "nested",
			fields: []zapcore.Field{
				zap.String("k1", "v1"),
				zap.Namespace("nested"),
				zap.String("k2", "v2"),
			},
			want: map[string]interface{}{
				"k1": "v1",
				"nested": map[string]interface{}{
					"k2": "v2",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			entry := LoggedEntry{
				Context: tt.fields,
			}
			assert.Equal(t, tt.want, entry.ContextMap())
		})
	}
}
