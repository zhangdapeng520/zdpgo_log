package zdpgo_zap

import "go.uber.org/zap"

// Zap zap日志核心对象
type Zap struct {
	log    *zap.Logger        // 日志对象
	sugar  *zap.SugaredLogger // sugar日志对象
	config *ZapConfig         // 配置对象
}

type ZapConfig struct {
	Debug        bool   // 是否为debug模式
	OpenGlobal   bool   // 是否开启全局日志
	OpenFileName bool   // 是否输出文件名和行号
	LogFilePath  string // 日志路径
}

// New 创建zap实例
func New(config ZapConfig) *Zap {
	var err error
	z := Zap{}

	// 日志路径
	if config.LogFilePath == "" {
		config.LogFilePath = "zdpgo_zap.log"
	}

	// 创建日志
	var c zap.Config
	if config.Debug {
		c = zap.NewDevelopmentConfig()
		c.OutputPaths = []string{ // 输出路径
			"stderr",
			config.LogFilePath,
		}
	} else {
		c = zap.NewProductionConfig()
		c.OutputPaths = []string{ // 输出路径
			config.LogFilePath,
		}
	}
	c.Encoding = "json"
	z.log, err = c.Build()
	if err != nil {
		z.log.Panic("创建zap日志失败：", zap.String("err", err.Error()))
	}

	// 将日志写入文件
	defer z.log.Sync()

	// sugar日志
	z.sugar = z.log.Sugar()

	// 全局日志
	if config.OpenGlobal {
		zap.ReplaceGlobals(z.log)
	}

	// 输出文件名和行号
	if config.OpenFileName {
		z.log.WithOptions(zap.AddCaller())
	}

	return &z
}
