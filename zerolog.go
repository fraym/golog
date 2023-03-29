package golog

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type zerologLogger struct {
	logger *zerolog.Logger
}

func NewZerologLogger() Logger {
	logger := zerolog.New(os.Stderr).Level(zerolog.InfoLevel).With().Timestamp().Logger()
	zerolog.TimeFieldFormat = time.RFC3339

	return &zerologLogger{
		logger: &logger,
	}
}

func (l *zerologLogger) SetOutput(output io.Writer) {
	newLogger := l.logger.Output(output)
	l.logger = &newLogger
}

func (l *zerologLogger) SetLevel(level Level) {
	switch level {
	case DebugLevel:
		newLogger := l.logger.Level(zerolog.DebugLevel)
		l.logger = &newLogger
	case WarnLevel:
		newLogger := l.logger.Level(zerolog.WarnLevel)
		l.logger = &newLogger
	case ErrorLevel:
		newLogger := l.logger.Level(zerolog.ErrorLevel)
		l.logger = &newLogger
	case FatalLevel:
		newLogger := l.logger.Level(zerolog.FatalLevel)
		l.logger = &newLogger
	default:
		newLogger := l.logger.Level(zerolog.InfoLevel)
		l.logger = &newLogger
	}
}

func (l *zerologLogger) Trace() LoggerInstance {
	return newZerologInstance(l.logger.Trace())
}

func (l *zerologLogger) WriteTrace(args ...any) {
	l.logger.Trace().Msg(fmt.Sprint(args...))
}

func (l *zerologLogger) WriteTracef(format string, args ...any) {
	l.logger.Trace().Msgf(format, args...)
}

func (l *zerologLogger) Debug() LoggerInstance {
	return newZerologInstance(l.logger.Debug())
}

func (l *zerologLogger) WriteDebug(args ...any) {
	l.logger.Debug().Msg(fmt.Sprint(args...))
}

func (l *zerologLogger) WriteDebugf(format string, args ...any) {
	l.logger.Debug().Msgf(format, args...)
}

func (l *zerologLogger) Info() LoggerInstance {
	return newZerologInstance(l.logger.Info())
}

func (l *zerologLogger) WriteInfo(args ...any) {
	l.logger.Info().Msg(fmt.Sprint(args...))
}

func (l *zerologLogger) WriteInfof(format string, args ...any) {
	l.logger.Info().Msgf(format, args...)
}

func (l *zerologLogger) Warn() LoggerInstance {
	return newZerologInstance(l.logger.Warn())
}

func (l *zerologLogger) WriteWarn(args ...any) {
	l.logger.Warn().Msg(fmt.Sprint(args...))
}

func (l *zerologLogger) WriteWarnf(format string, args ...any) {
	l.logger.Warn().Msgf(format, args...)
}

func (l *zerologLogger) Error() LoggerInstance {
	return newZerologInstance(l.logger.Error())
}

func (l *zerologLogger) WriteError(args ...any) {
	l.logger.Error().Msg(fmt.Sprint(args...))
}

func (l *zerologLogger) WriteErrorf(format string, args ...any) {
	l.logger.Error().Msgf(format, args...)
}

func (l *zerologLogger) Fatal() LoggerInstance {
	return newZerologInstance(l.logger.Fatal())
}

func (l *zerologLogger) WriteFatal(args ...any) {
	l.logger.Fatal().Msg(fmt.Sprint(args...))
}

func (l *zerologLogger) WriteFatalf(format string, args ...any) {
	l.logger.Fatal().Msgf(format, args...)
}

type zerologInstance struct {
	logger *zerolog.Event
}

func newZerologInstance(logger *zerolog.Event) LoggerInstance {
	return &zerologInstance{
		logger: logger,
	}
}

func (l *zerologInstance) WithError(err error) LoggerInstance {
	l.logger.Err(err)
	return l
}

func (l *zerologInstance) WithField(key string, value any) LoggerInstance {
	l.logger.Any(key, value)
	return l
}

func (l *zerologInstance) WithFields(fields map[string]any) LoggerInstance {
	if !l.logger.Enabled() {
		return l
	}

	for key, value := range fields {
		l.logger.Any(key, value)
	}
	return l
}

func (l *zerologInstance) Write(args ...any) {
	l.logger.Msg(fmt.Sprint(args...))
}

func (l *zerologInstance) Writef(format string, args ...any) {
	l.logger.Msgf(format, args...)
}
