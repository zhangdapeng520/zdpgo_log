package main

import (
	"github.com/zhangdapeng520/zdpgo_log"
)

/*
@Time : 2022/6/24 17:55
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description: 彩色日志输出
*/

func main() {
	var log *zdpgo_log.Log

	// 开发环境的DEBUG日志不会写入日志文件
	log = zdpgo_log.GetDevLog()

	log.Debug("Debug日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Info("Info日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Warning("Warning日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Error("Error日志", "a", 1, "b", 2.2, "c", "ccc")

	// 生成环境的日志不会展示在控制台，是结构化的json日志
	log = zdpgo_log.GetProductLog("log.log")
	log.Debug("Debug日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Info("Info日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Warning("Warning日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Error("Error日志", "a", 1, "b", 2.2, "c", "ccc")
}
