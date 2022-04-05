package zdpgo_log

// Config zap日志配置核心对象
type Config struct {
	Debug        bool   `yaml:"debug" json:"debug"`                   // 是否为debug模式
	OpenJsonLog  bool   `yaml:"open_json_log" json:"open_json_log"`   // 是否开启json日志格式
	OpenGlobal   bool   `yaml:"open_global" json:"open_global"`       // 是否开启全局日志
	OpenFileName bool   `yaml:"open_file_name" json:"open_file_name"` // 是否输出文件名和行号
	LogFilePath  string `yaml:"log_file_path" json:"log_file_path"`   // 日志路径
	MaxSize      uint   `yaml:"max_size" json:"max_size"`             // 日志最大保存多少M
	MaxBackups   uint   `yaml:"max_backups" json:"max_backups"`       // 日志保留多少个备份
	MaxAge       uint   `yaml:"max_age" json:"max_age"`               // 最多保留多少个日志
	Compress     bool   `yaml:"compress" json:"compress"`             // 是否压缩
}
