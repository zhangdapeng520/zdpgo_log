package observer

import "github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"

// An LoggedEntry is an encoding-agnostic representation of a log message.
// Field availability is context dependant.
type LoggedEntry struct {
	zapcore.Entry
	Context []zapcore.Field
}

// ContextMap returns a map for all fields in Context.
func (e LoggedEntry) ContextMap() map[string]interface{} {
	encoder := zapcore.NewMapObjectEncoder()
	for _, f := range e.Context {
		f.AddTo(encoder)
	}
	return encoder.Fields
}
