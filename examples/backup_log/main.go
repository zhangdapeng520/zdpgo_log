package main

import (
	"time"

	"github.com/zhangdapeng520/zdpgo_log"
)

func main() {
	l := zdpgo_log.NewWithConfig(&zdpgo_log.LogConfig{
		Debug:        false,
		LogFilePath:  "log.log",
		OpenJsonLog:  true,
		OpenFileName: false,
		MaxSize:      1,
		MaxBackups:   3,
		MaxAge:       3,
		Compress:     false,
	})

	s := "it is test logit is test logit is test logit is test logit is test logit is test logit is test logit is test logit is test logit is test log"
	for i := 0; i < 1000; i++ {
		l.Info(s)
		time.Sleep(time.Millisecond * 100)
	}
}
