package logger

// LoggerStub is a stub of the readtack/logger.Logger interface
type LoggerStub struct {
	OnDebug func(string, ...interface{})
	OnInfo  func(string, ...interface{})
	OnPrint func(string, ...interface{})
	OnWarn  func(string, ...interface{})
	OnError func(string, ...interface{})
	OnFatal func(string, ...interface{})
	OnPanic func(string, ...interface{})
}

func (l *LoggerStub) Debug(format string, args ...interface{}) {
	l.OnDebug(format, args...)
}

func (l *LoggerStub) Info(format string, args ...interface{}) {
	l.OnInfo(format, args...)
}

func (l *LoggerStub) Print(format string, args ...interface{}) {
	l.OnPrint(format, args...)
}

func (l *LoggerStub) Warn(format string, args ...interface{}) {
	l.OnWarn(format, args...)
}

func (l *LoggerStub) Error(format string, args ...interface{}) {
	l.OnError(format, args...)
}

func (l *LoggerStub) Fatal(format string, args ...interface{}) {
	l.OnFatal(format, args...)
}

func (l *LoggerStub) Panic(format string, args ...interface{}) {
	l.OnPanic(format, args...)
}
