package zdpgo_zap

import (
	"fmt"
	"testing"
)

func prepareZap() *Zap {
	z := New(ZapConfig{
		Debug:        true,
		OpenGlobal:   true,
		OpenFileName: true,
		LogFilePath:  "zdpgo_zap.log",
	})
	return z
}

func TestZap_New(t *testing.T) {
	z := prepareZap()
	fmt.Println(z)
}
