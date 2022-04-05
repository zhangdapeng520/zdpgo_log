package zdpgo_log

import (
	"fmt"
	"os"
	"time"

	"github.com/zhangdapeng520/zdpgo_log/libs/lumberjack"
	"github.com/zhangdapeng520/zdpgo_log/libs/zap"
	"github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"
)

// 全局日志
var (
	Debug   func(msg string, args ...interface{})
	Info    func(msg string, args ...interface{})
	Warning func(msg string, args ...interface{})
	Error   func(msg string, args ...interface{})
	Panic   func(msg string, args ...interface{})
	Fatal   func(msg string, args ...interface{})

	// S 全局的日志对象
	S = zap.S
)

// Log 日志核心对象
type Log struct {
	Log    *zap.Logger        // 日志对象
	Sugar  *zap.SugaredLogger // sugar日志对象
	config *Config            // 配置对象

	// 日志方法
	Debug   func(msg string, args ...interface{})
	Info    func(msg string, args ...interface{})
	Warning func(msg string, args ...interface{})
	Error   func(msg string, args ...interface{})
	Panic   func(msg string, args ...interface{})
	Fatal   func(msg string, args ...interface{})
}

// New 创建zap实例
func New(config Config) *Log {
	// 创建日志对象
	z := Log{}

	// 初始化配置
	config = getDefaultConfig(config)
	z.config = &config

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

	// sugar日志
	z.Log = logger
	z.Sugar = sugarLogger

	// 记录日志
	defer z.Log.Sync()
	defer z.Sugar.Sync()

	// 全局日志
	if config.OpenGlobal {
		zap.ReplaceGlobals(z.Log)
	}

	// 输出文件名和行号
	if config.OpenFileName {
		z.Log.WithOptions(zap.AddCaller())
	}

	// 初始化日志方法
	z.Debug = sugarLogger.Debugw
	z.Info = sugarLogger.Infow
	z.Warning = sugarLogger.Warnw
	z.Error = sugarLogger.Errorw
	z.Panic = sugarLogger.Panicw
	z.Fatal = sugarLogger.Fatalw

	// 初始化全局日志方法
	Debug = sugarLogger.Debugw
	Info = sugarLogger.Infow
	Warning = sugarLogger.Warnw
	Error = sugarLogger.Errorw
	Panic = sugarLogger.Panicw
	Fatal = sugarLogger.Fatalw
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
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
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

// NewDebug 创建debug环境下的日志
func NewDebug() *Log {
	return New(Config{
		Debug:        true,
		OpenGlobal:   true,
		OpenFileName: false,
	})
}

// NewProduct 创建生产环境下的日志
func NewProduct() *Log {
	return New(Config{
		Debug:        false,
		OpenGlobal:   true,
		OpenFileName: true,
		OpenJsonLog:  true,
	})
}
