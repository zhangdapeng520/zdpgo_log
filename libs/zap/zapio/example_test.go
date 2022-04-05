package zapio_test

import (
	"io"
	"log"

	"github.com/zhangdapeng520/zdpgo_log/libs/zap"
	"github.com/zhangdapeng520/zdpgo_log/libs/zap/zapio"
)

func ExampleWriter() {
	logger := zap.NewExample()
	w := &zapio.Writer{Log: logger}

	io.WriteString(w, "starting up\n")
	io.WriteString(w, "running\n")
	io.WriteString(w, "shutting down\n")

	if err := w.Close(); err != nil {
		log.Fatal(err)
	}

	// Output:
	// {"level":"info","msg":"starting up"}
	// {"level":"info","msg":"running"}
	// {"level":"info","msg":"shutting down"}
}
