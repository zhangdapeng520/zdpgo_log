package main

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_log/libs/zap"
)

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

	// 全局日志
	zap.S().Debug("全局的debug日志。。。。")
	zap.S().Info("全局的info日志。。。。")
	zap.S().Warn("全局的warning日志。。。。")
	zap.S().Error("全局的error日志。。。。")

	// 全局日志
	zdpgo_log.S().Debug("全局的debug日志全局的debug日志全局的debug日志全局的debug日志全局的debug日志全局的debug日志全局的debug日志全局的debug日志全局的debug日志。。。。")
	zdpgo_log.S().Info("全局的info日志。。。。")
	zdpgo_log.S().Warn("全局的warning日志。。。")
	zdpgo_log.S().Error("全局的error日志。。。。")
}
