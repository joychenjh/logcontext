package logcontext

import (
	"context"

	"go.uber.org/zap"
)

type LogLevel int

const (
	// FatalLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel LogLevel = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

var defLog LogContext

func init() {
	defLog = NewDefStdLog()
}

type LogContext interface {
	SetContextLog(c ContextLog) LogContext
	SetLogLevel(level LogLevel) LogContext
	SetCallerSkip(calldepth int) LogContext

	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Debugc(ctx context.Context, v ...interface{})
	Debugcf(ctx context.Context, format string, v ...interface{})
	Debugw(msg string, f ...zap.Field)
	Debugwc(ctx context.Context, msg string, f ...zap.Field)

	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infoc(ctx context.Context, v ...interface{})
	Infocf(ctx context.Context, format string, v ...interface{})
	Infow(msg string, f ...zap.Field)
	Infowc(ctx context.Context, msg string, f ...zap.Field)

	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Warnc(ctx context.Context, v ...interface{})
	Warncf(ctx context.Context, format string, v ...interface{})
	Warnw(msg string, f ...zap.Field)
	Warnwc(ctx context.Context, msg string, f ...zap.Field)

	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorc(ctx context.Context, v ...interface{})
	Errorcf(ctx context.Context, format string, v ...interface{})
	Errorw(msg string, f ...zap.Field)
	Errorwc(ctx context.Context, msg string, f ...zap.Field)

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalc(ctx context.Context, v ...interface{})
	Fatalcf(ctx context.Context, format string, v ...interface{})
	Fatalw(msg string, f ...zap.Field)
	Fatalwc(ctx context.Context, msg string, f ...zap.Field)

	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicc(ctx context.Context, v ...interface{})
	Paniccf(ctx context.Context, format string, v ...interface{})
	Panicw(msg string, f ...zap.Field)
	Panicwc(ctx context.Context, msg string, f ...zap.Field)
}

func SetLog(log LogContext) LogContext {
	defLog = log
	return defLog
}
func SetContextLog(c ContextLog) LogContext {
	return defLog.SetContextLog(c)
}
func SetLogLevel(level LogLevel) LogContext {
	return defLog.SetLogLevel(level)
}
func SetCallerSkip(calldepth int) LogContext {
	return defLog.SetCallerSkip(calldepth)
}

func Debug(v ...interface{}) {
	defLog.Debug(v...)
}
func Debugf(format string, v ...interface{}) {
	defLog.Debugf(format, v...)
}

func Debugc(ctx context.Context, v ...interface{}) {
	defLog.Debugc(ctx, v...)
}
func Debugcf(ctx context.Context, format string, v ...interface{}) {
	defLog.Debugcf(ctx, format, v...)
}
func Debugw(msg string, f ...zap.Field) {
	defLog.Debugw(msg, f...)
}
func Debugwc(ctx context.Context, msg string, f ...zap.Field) {
	defLog.Debugwc(ctx, msg, f...)
}
