package golog

import "io"

type Level uint32

const (
	FatalLevel Level = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

type Logger interface {
	SetLevel(level Level)
	SetOutput(output io.Writer)
	Trace() LoggerInstance
	WriteTrace(args ...any)
	WriteTracef(format string, args ...any)
	Debug() LoggerInstance
	WriteDebug(args ...any)
	WriteDebugf(format string, args ...any)
	Info() LoggerInstance
	WriteInfo(args ...any)
	WriteInfof(format string, args ...any)
	Warn() LoggerInstance
	WriteWarn(args ...any)
	WriteWarnf(format string, args ...any)
	Error() LoggerInstance
	WriteError(args ...any)
	WriteErrorf(format string, args ...any)
	Fatal() LoggerInstance
	WriteFatal(args ...any)
	WriteFatalf(format string, args ...any)
}

type LoggerInstance interface {
	WithError(err error) LoggerInstance
	WithField(key string, value any) LoggerInstance
	WithFields(fields map[string]any) LoggerInstance
	Write(args ...any)
	Writef(format string, args ...any)
}
