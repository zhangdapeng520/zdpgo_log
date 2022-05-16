package zapcore

// Core 是一个最小的、快速的记录器接口。它是为库的作者设计的，以包装一个更友好的API。
type Core interface {
	LevelEnabler

	// With 向核心添加结构化上下文。
	With([]Field) Core

	// Check 决定是否应该记录所提供的条目(使用嵌入的LevelEnabler和一些额外的逻辑)。
	// 如果该条目应该被记录，Core将自己添加到CheckedEntry并返回结果。
	// 调用者必须在调用Write之前使用Check。
	Check(Entry, *CheckedEntry) *CheckedEntry

	// Write 序列化日志站点上提供的Entry和任何Fields，并将它们写到它们的目的地。
	// 如果被调用，Write应该总是记录Entry和Fields;它不应该复制Check的逻辑。
	Write(Entry, []Field) error

	// Sync 刷新缓冲的日志(如果有的话)。
	Sync() error
}

type nopCore struct{}

// NewNopCore returns a no-op Core.
func NewNopCore() Core                                        { return nopCore{} }
func (nopCore) Enabled(Level) bool                            { return false }
func (n nopCore) With([]Field) Core                           { return n }
func (nopCore) Check(_ Entry, ce *CheckedEntry) *CheckedEntry { return ce }
func (nopCore) Write(Entry, []Field) error                    { return nil }
func (nopCore) Sync() error                                   { return nil }

// NewCore 创建一个写入日志到WriteSyncer的Core。
func NewCore(enc Encoder, ws WriteSyncer, enab LevelEnabler) Core {
	return &ioCore{
		LevelEnabler: enab,
		enc:          enc,
		out:          ws,
	}
}

type ioCore struct {
	LevelEnabler
	enc Encoder
	out WriteSyncer
}

func (c *ioCore) With(fields []Field) Core {
	clone := c.clone()
	addFields(clone.enc, fields)
	return clone
}

func (c *ioCore) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *ioCore) Write(ent Entry, fields []Field) error {
	buf, err := c.enc.EncodeEntry(ent, fields)
	if err != nil {
		return err
	}
	_, err = c.out.Write(buf.Bytes())
	buf.Free()
	if err != nil {
		return err
	}
	if ent.Level > ErrorLevel {
		// Since we may be crashing the program, sync the output. Ignore Sync
		// errors, pending a clean solution to issue #370.
		c.Sync()
	}
	return nil
}

func (c *ioCore) Sync() error {
	return c.out.Sync()
}

func (c *ioCore) clone() *ioCore {
	return &ioCore{
		LevelEnabler: c.LevelEnabler,
		enc:          c.enc.Clone(),
		out:          c.out,
	}
}
