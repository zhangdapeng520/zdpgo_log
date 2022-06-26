# zdpgo_zap

基于zap二次封装的日志库

## 版本历史

- v0.1.0 2022/01/30
- v0.1.1 2022/01/30
- v0.2.0 2022/02/10 新增日志备份功能
- v0.2.1 2022/02/10 新增自动创建日志目录功能
- v0.2.2 2022/02/11 将日志对象开放给外部使用
- v0.2.3 2022/02/11 bug修复
- v1.0.0 2022/03/08 优化日志方法
- v1.0.1 2022/04/05 优化代码结构
- v1.0.2 2022/04/07 移除没必要的第三方依赖
- v1.3.1 2022/04/26 优化：移除全局日志，优化日志创建方式
- v1.3.2 2022/05/10 新增：Debug日志可以只展示在控制台但是不写入日志文件
- v1.3.3 2022/05/11 BUG修复：解决某些情况下Debug方法不能用
- v1.3.4 2022/05/16 新增：根据debug值和日志路径创建日志对象
- v1.3.5 2022/06/17 新增：Tmp临时日志
- v1.3.6 2022/06/24 新增：支持彩色日志
- v1.3.7 2022/06/26 优化：移除第三方依赖

## 使用案例

### 基本使用

```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_log"
)

/*
@Time : 2022/6/24 17:55
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description: 彩色日志输出
*/

func main() {
	var log *zdpgo_log.Log

	// 开发环境的DEBUG日志不会写入日志文件
	log = zdpgo_log.GetDevLog()

	log.Debug("Debug日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Info("Info日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Warning("Warning日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Error("Error日志", "a", 1, "b", 2.2, "c", "ccc")

	// 生成环境的日志不会展示在控制台，是结构化的json日志
	log = zdpgo_log.GetProductLog("log.log")
	log.Debug("Debug日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Info("Info日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Warning("Warning日志", "a", 1, "b", 2.2, "c", "ccc")
	log.Error("Error日志", "a", 1, "b", 2.2, "c", "ccc")
}
```
