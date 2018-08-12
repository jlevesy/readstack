package logger

import (
	"bytes"
	"log"
)

const (
	debugPrefix = "[DEBUG] "
	infoPrefix  = "[INFO] "
	warnPrefix  = "[WARN] "
	errorPrefix = "[ERROR] "
	fatalPrefix = "[FATAL] "
	panicPrefix = "[PANIC] "
)

type stdLogger struct {
	*log.Logger
}

func NewStdLogger(logger *log.Logger) Logger {
	return &stdLogger{logger}
}

func (s *stdLogger) Debug(format string, args ...interface{}) {
	buf := bytes.NewBufferString(debugPrefix)
	buf.WriteString(format)
	s.Logger.Printf(buf.String(), args...)
}

func (s *stdLogger) Info(format string, args ...interface{}) {
	buf := bytes.NewBufferString(infoPrefix)
	buf.WriteString(format)
	s.Logger.Printf(buf.String(), args...)
}

func (s *stdLogger) Print(format string, args ...interface{}) {
	s.Logger.Printf(format, args...)
}

func (s *stdLogger) Warn(format string, args ...interface{}) {
	buf := bytes.NewBufferString(warnPrefix)
	buf.WriteString(format)
	s.Logger.Printf(buf.String(), args...)
}

func (s *stdLogger) Error(format string, args ...interface{}) {
	buf := bytes.NewBufferString(errorPrefix)
	buf.WriteString(format)
	s.Logger.Printf(buf.String(), args...)
}

func (s *stdLogger) Fatal(format string, args ...interface{}) {
	buf := bytes.NewBufferString(fatalPrefix)
	buf.WriteString(format)
	s.Logger.Fatalf(buf.String(), args...)
}

func (s *stdLogger) Panic(format string, args ...interface{}) {
	buf := bytes.NewBufferString(panicPrefix)
	buf.WriteString(format)
	s.Logger.Panicf(buf.String(), args...)
}
