package zdpgo_log

import (
	"fmt"
	"testing"
	"time"
)

func prepareLog() *Log {
	z := New(Config{
		Debug:        true,
		OpenGlobal:   true,
		OpenFileName: true,
	})
	return z
}

func TestLog_New(t *testing.T) {
	z := prepareLog()
	fmt.Println(z)
}

// 测试debug环境日志
func TestLog_Debug(t *testing.T) {
	l := NewDebug()
	l.Debug("日志。。。")
	l.Info("日志。。。")
	l.Warning("日志。。。")
	l.Error("日志。。。")
}

// 测试全局日志
func TestLog_GlobalInfo(t *testing.T) {
	prepareLog()
	Info("这是一条info类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
	Debug("这是一条debug类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
	Warning("这是一条warning类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
	Error("这是一条error类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
}

// 测试生产环境日志
func TestLog_Product(t *testing.T) {
	l := NewProduct()
	l.Debug("日志。。。")
	l.Info("日志。。。")
	l.Warning("日志。。。")
	l.Error("日志。。。")
}

// 测试json日志
func TestLog_Json(t *testing.T) {
	l := New(Config{
		Debug:       true,
		OpenJsonLog: true,
	})
	l.Debug("日志。。。")
	l.Info("日志。。。")
	l.Warning("日志。。。")
	l.Error("日志。。。")
}

// 测试备份
func TestLog_Backup(t *testing.T) {
	l := New(Config{
		Debug:       true,
		OpenJsonLog: false,
		MaxSize:     1,
		MaxBackups:  3,
	})

	s := "测试日志备份 最多3个日志，最大1M"
	for i := 0; i < 10; i++ {
		s += s
	}

	for i := 0; i < 100000; i++ {
		l.Info(s)
		time.Sleep(time.Millisecond * 10)
	}
}
