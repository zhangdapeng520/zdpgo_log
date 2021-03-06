package zdpgo_log

import (
	"testing"
)

func TestLog_New(t *testing.T) {
	var log *Log
	log = NewWithDebug(true, "c:/tmp/log.log")
	log.Debug("ok")

	Tmp.Debug("it is log by tmp log")
}

// 测试debug环境日志
func TestLog_Debug(t *testing.T) {
	// 默认不写入debug日志
	l := Tmp
	l.Debug("debug 111111日志。。。")

	// 设置写入Debug日志
	l = NewWithConfig(&LogConfig{
		Debug:         true,
		IsWriteDebug:  true,
		OpenJsonLog:   false,
		IsShowConsole: true,
	})
	l.Debug("debug 222222日志。。。")

	// 设置写入Debug日志，但是不显示在控制台
	l = NewWithConfig(&LogConfig{
		Debug:         true,
		IsWriteDebug:  true,
		OpenJsonLog:   false,
		IsShowConsole: false,
	})
	l.Debug("debug 3333333日志。。。")

	// 设置写入Debug日志，但是显示在控制台，但是日志级别为Warning
	l = NewWithConfig(&LogConfig{
		Debug:         true,
		IsWriteDebug:  true,
		OpenJsonLog:   false,
		IsShowConsole: true,
		LogLevel:      "warning",
	})
	l.Debug("debug 4444444444日志。。。")
}

// 测试生产环境日志
func TestLog_Product(t *testing.T) {
	l := NewWithConfig(&LogConfig{
		Debug:         true,
		IsWriteDebug:  false,
		IsShowConsole: true,
		OpenJsonLog:   false,
		LogFilePath:   "c:/tmp/log.log",
	})
	l.Debug("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)
	l.Info("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)
	l.Warning("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)
	l.Error("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)
}
