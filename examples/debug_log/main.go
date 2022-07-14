package main

import "github.com/zhangdapeng520/zdpgo_log"

func main() {
	// 默认不写入debug日志
	l := zdpgo_log.Tmp
	l.Debug("debug 111111日志。。。")

	// 设置写入Debug日志
	l = zdpgo_log.NewWithConfig(&zdpgo_log.LogConfig{
		Debug:         true,
		IsWriteDebug:  true,
		OpenJsonLog:   false,
		IsShowConsole: true,
	})
	l.Debug("debug 222222日志。。。")

	// 设置写入Debug日志，但是不显示在控制台
	l = zdpgo_log.NewWithConfig(&zdpgo_log.LogConfig{
		Debug:         true,
		IsWriteDebug:  true,
		OpenJsonLog:   false,
		IsShowConsole: false,
	})
	l.Debug("debug 3333333日志。。。")

	// 设置写入Debug日志，但是显示在控制台，但是日志级别为Warning
	l = zdpgo_log.NewWithConfig(&zdpgo_log.LogConfig{
		Debug:         true,
		IsWriteDebug:  true,
		OpenJsonLog:   false,
		IsShowConsole: true,
		LogLevel:      "warning",
	})
	l.Debug("debug 4444444444日志。。。")
}
