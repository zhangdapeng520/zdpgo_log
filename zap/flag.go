package zap

import (
	"flag"

	"github.com/zhangdapeng520/zdpgo_log/zap/zapcore"
)

// LevelFlag uses the standard library's flag.Var to declare a global flag
// with the specified name, default, and usage guidance. The returned value is
// a pointer to the value of the flag.
//
// If you don't want to use the flag package's global state, you can use any
// non-nil *Level as a flag.Value with your own *flag.FlagSet.
func LevelFlag(name string, defaultLevel zapcore.Level, usage string) *zapcore.Level {
	lvl := defaultLevel
	flag.Var(&lvl, name, usage)
	return &lvl
}
