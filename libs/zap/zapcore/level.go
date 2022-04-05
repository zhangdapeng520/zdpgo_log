package zapcore

import (
	"bytes"
	"errors"
	"fmt"
)

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")

// Level 是日志优先级。级别越高越重要。
type Level int8

const (
	DebugLevel  Level = iota - 1 //日志通常是大量的，并且通常在生产中禁用
	InfoLevel                    //是默认的日志优先级
	WarnLevel                    //日志比Info更重要，但不需要单独的人工检查
	ErrorLevel                   //日志是高优先级。如果应用程序运行顺利，它不应该生成任何错误级别的日志
	DPanicLevel                  //日志是特别重要的错误。在开发过程中，日志记录器在写完消息后会感到恐慌
	PanicLevel                   //记录消息，然后是panic抛出异常
	FatalLevel                   //记录一条消息，然后调用os.Exit(1)。

	_minLevel = DebugLevel // 最小日志等级
	_maxLevel = FatalLevel // 最大日志等级
)

// String 返回日志级别的小写ASCII表示。
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case DPanicLevel:
		return "dpanic"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

// CapitalString 返回日志级别的全大写ASCII表示。
func (l Level) CapitalString() string {
	// 全大写打印级别非常常见，因此我们应该导出此功能。
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case DPanicLevel:
		return "DPANIC"
	case PanicLevel:
		return "PANIC"
	case FatalLevel:
		return "FATAL"
	default:
		return fmt.Sprintf("LEVEL(%d)", l)
	}
}

// MarshalText 将Level封送为文本。注意，文本表示去掉了-Level后缀(参见示例)。
func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

// UnmarshalText 解组文本到一个水平。像MarshalText一样，UnmarshalText希望Level的文本表示去掉-Level后缀(参见示例)。
// 特别是，这使得使用YAML、TOML或JSON文件配置日志级别变得很容易。
func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLevel
	}
	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("未知的日志等级: %q", text)
	}
	return nil
}

func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "debug", "DEBUG":
		*l = DebugLevel
	case "info", "INFO", "": // make the zero value useful
		*l = InfoLevel
	case "warn", "WARN":
		*l = WarnLevel
	case "error", "ERROR":
		*l = ErrorLevel
	case "dpanic", "DPANIC":
		*l = DPanicLevel
	case "panic", "PANIC":
		*l = PanicLevel
	case "fatal", "FATAL":
		*l = FatalLevel
	default:
		return false
	}
	return true
}

// Set 设置标志的级别，flag.Value的接口
func (l *Level) Set(s string) error {
	return l.UnmarshalText([]byte(s))
}

// Get flag.Getter 的接口
func (l *Level) Get() interface{} {
	return *l
}

// Enabled 如果给定级别等于或高于此级别，则返回true。
func (l Level) Enabled(lvl Level) bool {
	return lvl >= l
}

// LevelEnabler 决定在记录消息时是否启用给定的日志级别。
// 使能器是用来实现确定性过滤器的;
// 像采样这样的问题作为Core来实现会更好。
// 每个具体的Level值实现了一个静态的LevelEnabler，它为自己和所有更高的日志级别返回true。
// 例如，WarnLevel. enabled()对于WarnLevel、ErrorLevel、DPanicLevel、PanicLevel和FatalLevel将返回true，而对于infollevel和DebugLevel将返回false。
type LevelEnabler interface {
	Enabled(Level) bool
}
