package zdpgo_log

import "time"

func timeToMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
