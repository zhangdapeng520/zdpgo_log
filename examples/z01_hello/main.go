package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
)

// 使用简单
func method1() {
	// 创建日志
	logger, _ := zap.NewProduction()

	// 写入日志
	defer logger.Sync() // flushes buffer, if any

	// 记录日志
	url := "http://www.baidu.com"
	sugar := logger.Sugar()

	// 方式1
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)

	// 方式2
	sugar.Infof("Failed to fetch URL: %s", url)
}

// 性能最高
func method2() {
	// 创建日志
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// 记录日志
	url := "http://www.baidu.com"
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

// 全局日志
func method3() {
	zap.L().Info("global Logger before")
	zap.S().Info("global SugaredLogger before")

	logger := zap.NewExample()
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	zap.L().Info("global Logger after")
	zap.S().Info("global SugaredLogger after")
}

// 预设字段
func method4() {
	logger := zap.NewExample(zap.Fields(
		zap.Int("serverId", 90),
		zap.String("serverName", "awesome web"),
	))

	logger.Info("hello world")
}

// 配合标准日志使用
func method5() {
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.RedirectStdLog(logger)
	log.Print("redirected standard library")
	log.Print("redirected standard library111")
	undo() // 取消重定向

	log.Print("restored standard library")
}

// 输出文件名和行号
func method6() {
	logger, _ := zap.NewProduction(zap.AddCaller())
	defer logger.Sync()

	logger.Info("hello world")
}

// 输出到文件
func method7() {
	rawJSON := []byte(`{
    "level":"debug",
    "encoding":"json",
    "outputPaths": ["stdout", "server.log"],
    "errorOutputPaths": ["stderr"],
    "initialFields":{"name":"dj"},
    "encoderConfig": {
      "messageKey": "message",
      "levelKey": "level",
      "levelEncoder": "lowercase"
    }
  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("server start work successfully!")
}

// 输出到文件 另一种方式
func method8() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{ // 输出路径
		"os.stdout",
		"zdpgo_zap.log",
	}
	logger, err := config.Build()
	fmt.Println(logger, err)
}

func main() {
	method1()
	method2()
	method3()
	method4()
	method5()
	method6()
	method7()
}
