package zdpgo_zap

import "testing"

func TestZap_GlobalInfo(t *testing.T) {
	prepareZap()
	Info("这是一条info类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
	Debug("这是一条debug类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
	Warning("这是一条warning类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
	Error("这是一条error类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
}
