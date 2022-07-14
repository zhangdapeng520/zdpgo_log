package bufferpool

import "github.com/zhangdapeng520/zdpgo_log/buffer"

var (
	_pool = buffer.NewPool()
	// Get retrieves a buffer from the pool, creating one if necessary.
	Get = _pool.Get
)
