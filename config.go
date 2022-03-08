package zdpgo_zap

// ZapConfig zap日志配置核心对象
type ZapConfig struct {
	Debug        bool   `mapstructure:"debug" json:"debug"`                   // 是否为debug模式
	OpenJsonLog  bool   `mapstructure:"open_json_log" json:"open_json_log"`   // 是否开启json日志格式
	OpenGlobal   bool   `mapstructure:"open_global" json:"open_global"`       // 是否开启全局日志
	OpenFileName bool   `mapstructure:"open_file_name" json:"open_file_name"` // 是否输出文件名和行号
	LogFilePath  string `mapstructure:"log_file_path" json:"log_file_path"`   // 日志路径
	MaxSize      uint   `mapstructure:"max_size" json:"max_size"`             // 日志最大保存多少M
	MaxBackups   uint   `mapstructure:"max_backups" json:"max_backups"`       // 日志保留多少个备份
	MaxAge       uint   `mapstructure:"max_age" json:"max_age"`               // 最多保留多少个日志
	Compress     bool   `mapstructure:"compress" json:"compress"`             // 是否压缩
}
