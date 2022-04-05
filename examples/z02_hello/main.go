package main

import "github.com/zhangdapeng520/zdpgo_log"

func main() {
	l := zdpgo_log.New(zdpgo_log.Config{
		Debug:        true,
		OpenGlobal:   true,
		OpenFileName: false,
	})
	l.Debug("debug日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Info("info日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Warning("warning日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Error("error日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
}
