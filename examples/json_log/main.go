package main

import "github.com/zhangdapeng520/zdpgo_log"

func main() {
	// 普通日志
	l := zdpgo_log.NewWithConfig(&zdpgo_log.LogConfig{
		Debug:         true,
		IsWriteDebug:  false,
		IsShowConsole: true,
		OpenJsonLog:   false,
		LogFilePath:   "log.log",
	})
	l.Debug("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)
	l.Info("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)
	l.Warning("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)
	l.Error("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)

	// json日志
	l = zdpgo_log.NewWithConfig(&zdpgo_log.LogConfig{
		Debug:         true,
		IsWriteDebug:  false,
		IsShowConsole: true,
		OpenJsonLog:   true,
		LogFilePath:   "log.log",
	})
	l.Debug("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)
	l.Info("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)
	l.Warning("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)
	l.Error("日志。。。", "a", 1, "b", 2.2, "c", "333", "d", true)
}
