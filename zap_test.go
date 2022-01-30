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
		LogFilePath:  "zdpgo_zap.log",
	})
	return z
}

func TestZap_New(t *testing.T) {
	z := prepareZap()
	fmt.Println(z)
}

func TestZap_Debug(t *testing.T) {
	l := NewDebug()
	l.Debug("日志。。。")
	l.Info("日志。。。")
	l.Warning("日志。。。")
	l.Error("日志。。。")
}

func TestZap_Product(t *testing.T) {
	l := NewProduct()
	l.Debug("日志。。。")
	l.Info("日志。。。")
	l.Warning("日志。。。")
	l.Error("日志。。。")
}
