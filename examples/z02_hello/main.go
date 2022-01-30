package main

import "github.com/zhangdapeng520/zdpgo_zap"

func main() {
	l := zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:        true,
		OpenGlobal:   true,
		OpenFileName: false,
	})
	l.Debug("debug日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Info("info日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Warning("warning日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Error("error日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
}
