package zdpgo_zap

// Info 记录info类型的日志
func (z *Zap) Info(msg string, kvs ...interface{}) {
	z.sugar.Infow(msg, kvs...)
}

// Debug 记录debug类型的日志
func (z *Zap) Debug(msg string, kvs ...interface{}) {
	z.sugar.Debugw(msg, kvs...)
}

// Warning 记录warning类型的日志
func (z *Zap) Warning(msg string, kvs ...interface{}) {
	z.sugar.Warnw(msg, kvs...)
}

// Error 记录error类型的日志
func (z *Zap) Error(msg string, kvs ...interface{}) {
	z.sugar.Errorw(msg, kvs...)
}

// Panic 记录panic类型的日志
func (z *Zap) Panic(msg string, kvs ...interface{}) {
	z.sugar.Panicw(msg, kvs...)
}

// Fatal 记录fatal类型的日志
func (z *Zap) Fatal(msg string, kvs ...interface{}) {
	z.sugar.Fatalw(msg, kvs...)
}
