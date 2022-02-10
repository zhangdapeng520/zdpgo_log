package zdpgo_zap

import (
	"fmt"
	"testing"
)

func prepareZap() *Zap {
	z := New(ZapConfig{
		Debug:        true,
		OpenGlobal:   true,
		OpenFileName: true,
	})
	return z
}

func TestZap_New(t *testing.T) {
	z := prepareZap()
	fmt.Println(z)
}

// 测试debug环境日志
func TestZap_Debug(t *testing.T) {
	l := NewDebug()
	l.Debug("日志。。。")
	l.Info("日志。。。")
	l.Warning("日志。。。")
	l.Error("日志。。。")
}

// 测试生产环境日志
func TestZap_Product(t *testing.T) {
	l := NewProduct()
	l.Debug("日志。。。")
	l.Info("日志。。。")
	l.Warning("日志。。。")
	l.Error("日志。。。")
}

// 测试json日志
func TestZap_Json(t *testing.T) {
	l := New(ZapConfig{
		Debug:       true,
		OpenJsonLog: true,
	})
	l.Debug("日志。。。")
	l.Info("日志。。。")
	l.Warning("日志。。。")
	l.Error("日志。。。")
}

// 测试备份
func TestZap_Backup(t *testing.T) {
	l := New(ZapConfig{
		Debug:       true,
		OpenJsonLog: false,
		MaxSize:     1,
		MaxBackups:  3,
	})
	for i := 0; i < 1000000; i++ {
		l.Info("测试日志备份 最多3个日志，最大1M")
	}
}
