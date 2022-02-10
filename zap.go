package zdpgo_zap

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// Zap zap日志核心对象
type Zap struct {
	log    *zap.Logger        // 日志对象
	sugar  *zap.SugaredLogger // sugar日志对象
	config *ZapConfig         // 配置对象
}

// New 创建zap实例
func New(config ZapConfig) *Zap {
	z := Zap{}

	// 日志路径
	if config.LogFilePath == "" {
		// 创建日志文件夹
		err := createMultiDir("logs/zdpgo")
		if err != nil {
			return nil
		}
		config.LogFilePath = "logs/zdpgo/zdpgo_zap.log"
	}

	// 创建日志
	writeSyncer := getLogWriter(config)
	encoder := getEncoder(config)
	var core zapcore.Core
	if config.Debug {
		writerObj := zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))
		core = zapcore.NewCore(encoder, writerObj, zapcore.DebugLevel)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	}
	logger := zap.New(core, zap.AddCaller())
	sugarLogger := logger.Sugar()
	defer sugarLogger.Sync()

	// sugar日志
	z.log = logger
	z.sugar = sugarLogger

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

// 调用os.MkdirAll递归创建文件夹
func createMultiDir(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败,error info:", err)
			return err
		}
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 获取日志编码器
func getEncoder(config ZapConfig) zapcore.Encoder {
	// 配置对象
	var encoderConfig zapcore.EncoderConfig
	if config.Debug {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	// 指定时间格式
	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeTime = customTimeEncoder

	//显示完整文件路径
	if !config.Debug {
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	}

	// 编码器
	var encoder zapcore.Encoder
	if config.OpenJsonLog {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return encoder
}

// 获取日志写入对象
func getLogWriter(config ZapConfig) zapcore.WriteSyncer {
	// 处理配置
	if config.MaxSize == 0 {
		config.MaxSize = 33
	}
	if config.MaxBackups == 0 {
		config.MaxBackups = 33
	}
	if config.MaxAge == 0 {
		config.MaxAge = 33
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.LogFilePath,     // 日志输出文件
		MaxSize:    int(config.MaxSize),    // 日志最大保存1M
		MaxBackups: int(config.MaxBackups), // 就日志保留5个备份
		MaxAge:     int(config.MaxAge),     // 最多保留30个日志 和MaxBackups参数配置1个就可以
		Compress:   config.Compress,        // 自动打 gzip包 默认false
	}
	return zapcore.AddSync(lumberJackLogger)
}

// NewDebug 创建debug环境下的日志
func NewDebug() *Zap {
	return New(ZapConfig{
		Debug:        true,
		OpenGlobal:   true,
		OpenFileName: false,
	})
}

// NewProduct 创建生产环境下的日志
func NewProduct() *Zap {
	return New(ZapConfig{
		Debug:        false,
		OpenGlobal:   true,
		OpenFileName: true,
		OpenJsonLog:  true,
	})
}
