package golog

import (
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
	})

	return &logrusLogger{
		logger: logger,
	}
}

func (l *logrusLogger) SetOutput(output io.Writer) {
	l.logger.SetOutput(output)
}

func (l *logrusLogger) SetLevel(level Level) {
	switch level {
	case DebugLevel:
		l.logger.SetLevel(logrus.DebugLevel)
	case WarnLevel:
		l.logger.SetLevel(logrus.WarnLevel)
	case ErrorLevel:
		l.logger.SetLevel(logrus.ErrorLevel)
	case FatalLevel:
		l.logger.SetLevel(logrus.FatalLevel)
	default:
		l.logger.SetLevel(logrus.InfoLevel)
	}
}

func (l *logrusLogger) Trace() LoggerInstance {
	return newLogrusInstance(logrus.TraceLevel, logrus.NewEntry(l.logger))
}

func (l *logrusLogger) WriteTrace(args ...any) {
	l.logger.Trace(args...)
}

func (l *logrusLogger) WriteTracef(format string, args ...any) {
	l.logger.Tracef(format, args...)
}

func (l *logrusLogger) Debug() LoggerInstance {
	return newLogrusInstance(logrus.DebugLevel, logrus.NewEntry(l.logger))
}

func (l *logrusLogger) WriteDebug(args ...any) {
	l.logger.Debug(args...)
}

func (l *logrusLogger) WriteDebugf(format string, args ...any) {
	l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Info() LoggerInstance {
	return newLogrusInstance(logrus.InfoLevel, logrus.NewEntry(l.logger))
}

func (l *logrusLogger) WriteInfo(args ...any) {
	l.logger.Info(args...)
}

func (l *logrusLogger) WriteInfof(format string, args ...any) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warn() LoggerInstance {
	return newLogrusInstance(logrus.WarnLevel, logrus.NewEntry(l.logger))
}

func (l *logrusLogger) WriteWarn(args ...any) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) WriteWarnf(format string, args ...any) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Error() LoggerInstance {
	return newLogrusInstance(logrus.ErrorLevel, logrus.NewEntry(l.logger))
}

func (l *logrusLogger) WriteError(args ...any) {
	l.logger.Error(args...)
}

func (l *logrusLogger) WriteErrorf(format string, args ...any) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatal() LoggerInstance {
	return newLogrusInstance(logrus.FatalLevel, logrus.NewEntry(l.logger))
}

func (l *logrusLogger) WriteFatal(args ...any) {
	l.logger.Fatal(args...)
}

func (l *logrusLogger) WriteFatalf(format string, args ...any) {
	l.logger.Fatalf(format, args...)
}

type logrusInstance struct {
	level  logrus.Level
	logger *logrus.Entry
}

func newLogrusInstance(level logrus.Level, logger *logrus.Entry) LoggerInstance {
	return &logrusInstance{
		level:  level,
		logger: logger,
	}
}

func (l *logrusInstance) WithError(err error) LoggerInstance {
	l.logger = l.logger.WithError(err)
	return l
}

func (l *logrusInstance) WithField(key string, value any) LoggerInstance {
	l.logger = l.logger.WithField(key, value)
	return l
}

func (l *logrusInstance) WithFields(fields map[string]any) LoggerInstance {
	l.logger = l.logger.WithFields(fields)
	return l
}

func (l *logrusInstance) Write(args ...any) {
	l.logger.Log(l.level, args...)
}

func (l *logrusInstance) Writef(format string, args ...any) {
	l.logger.Logf(l.level, format, args...)
}
