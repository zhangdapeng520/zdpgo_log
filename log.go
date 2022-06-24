package zdpgo_log

import (
	"github.com/zhangdapeng520/zdpgo_log/colorable"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/zhangdapeng520/zdpgo_log/lumberjack"
	"github.com/zhangdapeng520/zdpgo_log/zap"
	"github.com/zhangdapeng520/zdpgo_log/zap/zapcore"
)

// Log 日志核心对象
type Log struct {
	Config *Config // 配置对象

	// 日志方法
	Debug   func(msg string, args ...interface{})
	Info    func(msg string, args ...interface{})
	Warning func(msg string, args ...interface{})
	Error   func(msg string, args ...interface{})
	Panic   func(msg string, args ...interface{})
	Fatal   func(msg string, args ...interface{})
}

var (
	Tmp *Log
)

func init() {
	switch runtime.GOOS {
	case "windows":
		Tmp = NewWithDebug(true, "c:/tmp/log.log")
	default:
		Tmp = NewWithDebug(true, "/tmp/log.log")
	}
}

// New 创建zap实例
func New() *Log {
	return NewWithConfig(&Config{Debug: true, OpenJsonLog: true})
}

// GetDevLog 获取开发环境的日志
func GetDevLog() *Log {
	return NewWithConfig(&Config{
		Debug:         true,
		LogLevel:      "DEBUG",
		IsWriteDebug:  false,
		IsShowConsole: true,
		OpenJsonLog:   false,
		OpenFileName:  false,
		LogFilePath:   "",
	})
}

// GetProductLog 获取开发环境的日志
func GetProductLog(logFilePath string) *Log {
	return NewWithConfig(&Config{
		Debug:         false,
		LogLevel:      "INFO",
		IsWriteDebug:  true,
		IsShowConsole: false,
		OpenJsonLog:   true,
		OpenFileName:  false,
		LogFilePath:   logFilePath,
	})
}

// NewWithConfig 创建zap实例
func NewWithConfig(config *Config) *Log {
	// 创建日志对象
	z := &Log{}

	// 日志路径
	if config.LogFilePath == "" {
		switch runtime.GOOS {
		case "windows":
			config.LogFilePath = "c:\\tmp\\log.log"
		default:
			config.LogFilePath = "/tmp/log.log"
		}
	}

	// 提取目录名
	dirName := path.Dir(config.LogFilePath)

	// 创建日志文件夹
	_ = createMultiDir(dirName)

	// 日志级别
	if config.LogLevel == "" {
		if config.Debug {
			config.LogLevel = "DEBUG"
		} else {
			config.LogLevel = "INFO"
		}
	}

	// 日志文件大小：默认33M
	if config.MaxSize == 0 {
		config.MaxSize = 33
	}

	// 日志文件个数：默认33个
	if config.MaxBackups == 0 {
		config.MaxBackups = 33
	}

	// 日志文件存放天数：默认33天
	if config.MaxAge == 0 {
		config.MaxAge = 33
	}
	z.Config = config

	// 日志级别
	var logLevel zapcore.Level
	switch strings.ToUpper(config.LogLevel) {
	case "DEBUG":
		logLevel = zap.DebugLevel
	case "INFO":
		logLevel = zap.InfoLevel
	case "WARNING":
		logLevel = zap.WarnLevel
	case "ERROR":
		logLevel = zap.ErrorLevel
	case "PANIC":
		logLevel = zap.PanicLevel
	default:
		logLevel = zap.DebugLevel
	}

	// 创建日志
	writeSyncer := getLogWriter(*config)
	encoder := getEncoder(*config)
	var core zapcore.Core

	// DEBUG日志不要写入文件
	var (
		logger           *zap.Logger
		sugarLogger      *zap.SugaredLogger
		debugSugarLogger *zap.SugaredLogger
	)

	// 创建在控制台显示debug日志，但是不写入到文件中
	if config.Debug && !config.IsWriteDebug {
		core = zapcore.NewCore(encoder, zapcore.AddSync(colorable.NewColorableStdout()), zapcore.DebugLevel)
		debugSugarLogger = zap.New(core, zap.AddCaller()).Sugar()
		z.Debug = debugSugarLogger.Debugw
	}

	// 是否在控制台展示日志
	if config.IsShowConsole {
		writerObj := zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(colorable.NewColorableStdout()))
		core = zapcore.NewCore(encoder, writerObj, logLevel)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, logLevel)
	}

	// 创建日志对象
	logger = zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
	defer logger.Sync()
	defer sugarLogger.Sync()

	if z.Debug == nil {
		z.Debug = sugarLogger.Debugw
	}

	// 输出文件名和行号
	if config.OpenFileName {
		logger.WithOptions(zap.AddCaller())
	}

	// 初始化日志方法
	z.Info = sugarLogger.Infow
	z.Warning = sugarLogger.Warnw
	z.Error = sugarLogger.Errorw
	z.Panic = sugarLogger.Panicw
	z.Fatal = sugarLogger.Fatalw

	return z
}

// NewWithDebug 根据debug值和日志路径创建日志对象
func NewWithDebug(debug bool, logFilePath string) *Log {
	logConfig := &Config{
		Debug:       debug,
		OpenJsonLog: false,
		LogFilePath: logFilePath,
	}
	if debug {
		logConfig.IsShowConsole = true
	}
	log := NewWithConfig(logConfig)
	return log
}

// 调用os.MkdirAll递归创建文件夹
func createMultiDir(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	return err == nil
}

// 获取日志编码器
func getEncoder(config Config) zapcore.Encoder {
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
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 返回
	return encoder
}

// 获取日志写入对象
func getLogWriter(config Config) zapcore.WriteSyncer {
	// 处理配置
	config = getDefaultConfig(config)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.LogFilePath,     // 日志输出文件
		MaxSize:    int(config.MaxSize),    // 日志最大保存1M
		MaxBackups: int(config.MaxBackups), // 就日志保留5个备份
		MaxAge:     int(config.MaxAge),     // 最多保留30个日志 和MaxBackups参数配置1个就可以
		Compress:   config.Compress,        // 自动打 gzip包 默认false
	}
	return zapcore.AddSync(lumberJackLogger)
}
