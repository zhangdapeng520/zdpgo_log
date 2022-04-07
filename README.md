# zdpgo_zap

基于zap二次封装的日志库

## 版本历史

- 版本0.1.0：2022年1月30日
- 版本0.1.1：2022年1月30日
- 版本0.2.0：2022年2月10日 新增日志备份功能
- 版本0.2.1：2022年2月10日 新增自动创建日志目录功能
- 版本0.2.2：2022年2月11日 将日志对象开放给外部使用
- 版本0.2.3：2022年2月11日 bug修复
- 版本1.0.0：2022年3月8日 优化日志方法
- 版本1.0.1：2022年4月5日 优化代码结构
- 版本1.0.2：2022年4月7日 移除没必要的第三方依赖

## 使用案例

### 基本使用
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_log"
)

func main() {
	l := zdpgo_log.New(zdpgo_log.Config{
		Debug:        true,
		OpenGlobal:   true,
		OpenFileName: false,
	})
	l.Debug("debug日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Info("info日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Warning("warning日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")
	l.Error("error日志", "a", 111, "b", 22.22, "c", true, "d", "bbb")

	// 全局日志
	zdpgo_log.S().Debug("全局的debug日志全局的debug日志全局的debug日志全局的debug日志全局的debug日志全局的debug日志全局的debug日志全局的debug日志全局的debug日志。。。。")
	zdpgo_log.S().Info("全局的info日志。。。。")
	zdpgo_log.S().Warn("全局的warning日志。。。")
	zdpgo_log.S().Error("全局的error日志。。。。")
}
```
