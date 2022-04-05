# zdpgo_zap

基于zap二次封装的日志库

## 功能

- 持久化日志
- json日志
- 区分测试环境和生产环境
- 全局日志
- 日志备份

## 版本日志

- 版本0.1.0：2022年1月30日
- 版本0.1.1：2022年1月30日
- 版本0.2.0：2022年2月10日 新增日志备份功能
- 版本0.2.1：2022年2月10日 新增自动创建日志目录功能
- 版本0.2.2：2022年2月11日 将日志对象开放给外部使用
- 版本0.2.3：2022年2月11日 bug修复
- 版本1.0.0：2022年3月8日 优化日志方法
- 版本1.0.1：2022年4月5日 优化代码结构

## 常用方法

- Info 记录info类型的日志
- Debug 记录debug类型的日志
- Warning 记录warning类型的日志
- Error 记录error类型的日志
- Panic 记录panic类型的日志
- Fatal 记录fatal类型的日志

## 快速入门

```go
package main

import "github.com/zhangdapeng520/zdpgo_zap"

func main() {
	l := zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:        true,
		OpenGlobal:   true,
		OpenFileName: false,
	})
	l.Debug("debug日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Info("info日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Warning("warning日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Error("error日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
}
```

## 创建日志的便捷方式

```go
func TestZap_Debug(t *testing.T) {
l := NewDebug()
l.Debug("日志。。。")
l.Info("日志。。。")
l.Warning("日志。。。")
l.Error("日志。。。")
}

func TestZap_Product(t *testing.T) {
l := NewProduct()
l.Debug("日志。。。")
l.Info("日志。。。")
l.Warning("日志。。。")
l.Error("日志。。。")
}
```