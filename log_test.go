package zdpgo_log

import (
	"fmt"
	"testing"
	"time"
)

func prepareLog() *Log {
	z := New()
	return z
}

func TestLog_New(t *testing.T) {
	z := prepareLog()
	fmt.Println(z)
}

// 测试debug环境日志
func TestLog_Debug(t *testing.T) {
	// 默认不写入debug日志
	l := New()
	l.Debug("debug 111111日志。。。")

	// 设置写入Debug日志
	l = NewWithConfig(Config{
		Debug:         true,
		IsWriteDebug:  true,
		OpenJsonLog:   true,
		IsShowConsole: true,
	})
	l.Debug("debug 222222日志。。。")

	// 设置写入Debug日志，但是不显示在控制台
	l = NewWithConfig(Config{
		Debug:         true,
		IsWriteDebug:  true,
		OpenJsonLog:   true,
		IsShowConsole: false,
	})
	l.Debug("debug 3333333日志。。。")

	// 设置写入Debug日志，但是显示在控制台，但是日志级别为Warning
	l = NewWithConfig(Config{
		Debug:         true,
		IsWriteDebug:  true,
		OpenJsonLog:   true,
		IsShowConsole: true,
		LogLevel:      "warning",
	})
	l.Debug("debug 4444444444日志。。。")
}

// 测试生产环境日志
func TestLog_Product(t *testing.T) {
	l := prepareLog()
	l.Debug("日志。。。")
	l.Info("日志。。。")
	l.Warning("日志。。。")
	l.Error("日志。。。")
}

// 测试json日志
func TestLog_Json(t *testing.T) {
	l := prepareLog()
	l.Debug("日志。。。")
	l.Info("日志。。。")
	l.Warning("日志。。。")
	l.Error("日志。。。")
}

// 测试备份
func TestLog_Backup(t *testing.T) {
	l := NewWithConfig(Config{
		Debug:        false,
		OpenJsonLog:  true,
		OpenFileName: false,
		MaxSize:      1,
		MaxBackups:   3,
		MaxAge:       3,
		Compress:     false,
	})

	s := "测试日志备份 最多3个日志，最大1M"
	for i := 0; i < 10; i++ {
		s += s
	}

	for i := 0; i < 1000; i++ {
		l.Info(s)
		time.Sleep(time.Millisecond * 100)
	}
}
