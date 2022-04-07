package zap

import (
	"testing"

	"github.com/zhangdapeng520/zdpgo_log/libs/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
