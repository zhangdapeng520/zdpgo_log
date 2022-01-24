package zdpgo_zap

import "testing"

func TestZap_Info(t *testing.T) {
	z := prepareZap()
	z.Info("这是一条info类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
	z.Debug("这是一条debug类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
	z.Warning("这是一条warning类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
	z.Error("这是一条error类型的日志", "a", 1, "b", 2.22, "c", "333", "d", true)
}
