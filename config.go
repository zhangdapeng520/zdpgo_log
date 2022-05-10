package zdpgo_log

import "path"

// Config zap日志配置核心对象
type Config struct {
	Debug         bool   `env:"debug" yaml:"debug" json:"debug"`                               // 是否为debug模式
	LogLevel      string `env:"log_level" yaml:"log_level" json:"log_level"`                   // 日志级别
	IsWriteDebug  bool   `env:"is_write_debug" yaml:"is_write_debug" json:"is_write_debug"`    // 是否将debug日志写入文件
	IsShowConsole bool   `env:"is_show_console" yaml:"is_show_console" json:"is_show_console"` // 是否在控制台展示
	OpenJsonLog   bool   `env:"open_json_log" yaml:"open_json_log" json:"open_json_log"`       // 是否开启json日志格式
	OpenFileName  bool   `env:"open_file_name" yaml:"open_file_name" json:"open_file_name"`    // 是否输出文件名和行号
	LogFilePath   string `env:"log_file_path" yaml:"log_file_path" json:"log_file_path"`       // 日志路径
	MaxSize       uint   `env:"max_size" yaml:"max_size" json:"max_size"`                      // 日志最大保存多少M
	MaxBackups    uint   `env:"max_backups" yaml:"max_backups" json:"max_backups"`             // 日志保留多少个备份
	MaxAge        uint   `env:"max_age" yaml:"max_age" json:"max_age"`                         // 最多保留多少天日志
	Compress      bool   `env:"compress" yaml:"compress" json:"compress"`                      // 是否压缩
}

// 获取默认的配置
func getDefaultConfig(c Config) Config {
	// 日志路径
	if c.LogFilePath == "" {
		// 创建日志文件夹
		err := createMultiDir("logs/zdpgo")
		if err != nil {
			return c
		}
		c.LogFilePath = "logs/zdpgo/zdpgo_log.log"
	} else {
		// 提取目录名
		dirName := path.Dir(c.LogFilePath)

		// 创建日志文件夹
		err := createMultiDir(dirName)
		if err != nil {
			return c
		}
	}

	// 日志级别
	if c.LogLevel == "" {
		c.LogLevel = "DEBUG"
	}

	// 日志文件大小：默认33M
	if c.MaxSize == 0 {
		c.MaxSize = 33
	}

	// 日志文件个数：默认33个
	if c.MaxBackups == 0 {
		c.MaxBackups = 33
	}

	// 日志文件存放天数：默认33天
	if c.MaxAge == 0 {
		c.MaxAge = 33
	}

	// 返回初始化以后的配置
	return c
}
