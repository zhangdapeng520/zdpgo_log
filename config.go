package zdpgo_zap

// ZapConfig zap日志配置核心对象
type ZapConfig struct {
	Debug        bool   // 是否为debug模式
	OpenJsonLog  bool   // 是否开启json日志格式
	OpenGlobal   bool   // 是否开启全局日志
	OpenFileName bool   // 是否输出文件名和行号
	LogFilePath  string // 日志路径
	MaxSize      uint   // 日志最大保存多少M
	MaxBackups   uint   // 日志保留多少个备份
	MaxAge       uint   // 最多保留多少个日志
	Compress     bool   // 是否压缩
}
