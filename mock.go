package golog

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

func (l *MockLogger) SetLevel(level Level) {
	l.Called(level)
}

func (l *MockLogger) SetOutput(output io.Writer) {
	l.Called(output)
}

func (l *MockLogger) Writer(level Level) io.Writer {
	args := l.Called(level)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(io.Writer)
}

func (l *MockLogger) Trace() LoggerInstance {
	args := l.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(LoggerInstance)
}

func (l *MockLogger) WriteTrace(args ...any) {
	l.Called(args...)
}

func (l *MockLogger) WriteTracef(format string, args ...any) {
	var allArgs []any
	allArgs = append(allArgs, format)
	allArgs = append(allArgs, args...)
	l.Called(allArgs...)
}

func (l *MockLogger) Debug() LoggerInstance {
	args := l.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(LoggerInstance)
}

func (l *MockLogger) WriteDebug(args ...any) {
	l.Called(args...)
}

func (l *MockLogger) WriteDebugf(format string, args ...any) {
	var allArgs []any
	allArgs = append(allArgs, format)
	allArgs = append(allArgs, args...)
	l.Called(allArgs...)
}

func (l *MockLogger) Info() LoggerInstance {
	args := l.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(LoggerInstance)
}

func (l *MockLogger) WriteInfo(args ...any) {
	l.Called(args...)
}

func (l *MockLogger) WriteInfof(format string, args ...any) {
	var allArgs []any
	allArgs = append(allArgs, format)
	allArgs = append(allArgs, args...)
	l.Called(allArgs...)
}

func (l *MockLogger) Warn() LoggerInstance {
	args := l.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(LoggerInstance)
}

func (l *MockLogger) WriteWarn(args ...any) {
	l.Called(args...)
}

func (l *MockLogger) WriteWarnf(format string, args ...any) {
	var allArgs []any
	allArgs = append(allArgs, format)
	allArgs = append(allArgs, args...)
	l.Called(allArgs...)
}

func (l *MockLogger) Error() LoggerInstance {
	args := l.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(LoggerInstance)
}

func (l *MockLogger) WriteError(args ...any) {
	l.Called(args...)
}

func (l *MockLogger) WriteErrorf(format string, args ...any) {
	var allArgs []any
	allArgs = append(allArgs, format)
	allArgs = append(allArgs, args...)
	l.Called(allArgs...)
}

func (l *MockLogger) Fatal() LoggerInstance {
	args := l.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(LoggerInstance)
}

func (l *MockLogger) WriteFatal(args ...any) {
	l.Called(args...)
}

func (l *MockLogger) WriteFatalf(format string, args ...any) {
	var allArgs []any
	allArgs = append(allArgs, format)
	allArgs = append(allArgs, args...)
	l.Called(allArgs...)
}
