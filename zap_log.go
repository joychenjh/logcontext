package logcontext

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLog struct {
	zlog      *zap.Logger
	atom      zap.AtomicLevel
	ctxLog    ContextLog
	calldepth int
}

func NewZapLog(zlog *zap.Logger, atom zap.AtomicLevel) (zl LogContext) {

	_zt := zlog.WithOptions(zap.AddCallerSkip(3))
	zl = &zapLog{
		zlog:   _zt,
		atom:   atom,
		ctxLog: NewNilContext(),
	}

	return zl
}

func (log *zapLog) SetContextLog(c ContextLog) LogContext {
	log.ctxLog = c
	return log
}

func (log *zapLog) SetCallerSkip(calldepth int) LogContext {
	log.calldepth = calldepth
	log.zlog.WithOptions(zap.AddCallerSkip(calldepth))
	return log
}

func (log *zapLog) SetLogLevel(level LogLevel) LogContext {
	log.atom.SetLevel(zapcore.Level(level))
	return log
}

//Debug
func (log *zapLog) Debug(v ...interface{}) {
	log.logMsgData(nil, zapcore.DebugLevel, "", v...)
}

func (log *zapLog) Debugf(format string, v ...interface{}) {
	log.logMsgData(nil, zapcore.DebugLevel, format, v...)
}

func (log *zapLog) Debugc(ctx context.Context, v ...interface{}) {
	log.logMsgData(ctx, zapcore.DebugLevel, "", v...)
}

func (log *zapLog) Debugcf(ctx context.Context, format string, v ...interface{}) {
	log.logMsgData(ctx, zapcore.DebugLevel, format, v...)
}

func (log *zapLog) Debugw(msg string, f ...zap.Field) {
	log.slogFieldData(nil, zapcore.DebugLevel, "", f...)
}

func (log *zapLog) Debugwc(ctx context.Context, msg string, f ...zap.Field) {
	log.slogFieldData(ctx, zapcore.DebugLevel, msg, f...)
}

//Info
func (log *zapLog) Info(v ...interface{}) {
	log.logMsgData(nil, zapcore.InfoLevel, "", v...)
}

func (log *zapLog) Infof(format string, v ...interface{}) {
	log.logMsgData(nil, zapcore.InfoLevel, format, v...)
}

func (log *zapLog) Infoc(ctx context.Context, v ...interface{}) {
	log.logMsgData(ctx, zapcore.InfoLevel, "", v...)
}

func (log *zapLog) Infocf(ctx context.Context, format string, v ...interface{}) {
	log.logMsgData(ctx, zapcore.InfoLevel, format, v...)
}

func (log *zapLog) Infow(msg string, f ...zap.Field) {
	log.slogFieldData(nil, zapcore.InfoLevel, "", f...)
}

func (log *zapLog) Infowc(ctx context.Context, msg string, f ...zap.Field) {
	log.slogFieldData(ctx, zapcore.InfoLevel, msg, f...)
}

//Warn
func (log *zapLog) Warn(v ...interface{}) {
	log.logMsgData(nil, zapcore.WarnLevel, "", v...)
}

func (log *zapLog) Warnf(format string, v ...interface{}) {
	log.logMsgData(nil, zapcore.WarnLevel, format, v...)
}

func (log *zapLog) Warnc(ctx context.Context, v ...interface{}) {
	log.logMsgData(ctx, zapcore.WarnLevel, "", v...)
}

func (log *zapLog) Warncf(ctx context.Context, format string, v ...interface{}) {
	log.logMsgData(ctx, zapcore.WarnLevel, format, v...)
}

func (log *zapLog) Warnw(msg string, f ...zap.Field) {
	log.slogFieldData(nil, zapcore.WarnLevel, "", f...)
}

func (log *zapLog) Warnwc(ctx context.Context, msg string, f ...zap.Field) {
	log.slogFieldData(ctx, zapcore.WarnLevel, msg, f...)
}

//Error
func (log *zapLog) Error(v ...interface{}) {
	log.logMsgData(nil, zapcore.ErrorLevel, "", v...)
}

func (log *zapLog) Errorf(format string, v ...interface{}) {
	log.logMsgData(nil, zapcore.ErrorLevel, format, v...)
}

func (log *zapLog) Errorc(ctx context.Context, v ...interface{}) {
	log.logMsgData(ctx, zapcore.ErrorLevel, "", v...)
}

func (log *zapLog) Errorcf(ctx context.Context, format string, v ...interface{}) {
	log.logMsgData(ctx, zapcore.ErrorLevel, format, v...)
}

func (log *zapLog) Errorw(msg string, f ...zap.Field) {
	log.slogFieldData(nil, zapcore.ErrorLevel, "", f...)
}

func (log *zapLog) Errorwc(ctx context.Context, msg string, f ...zap.Field) {
	log.slogFieldData(ctx, zapcore.ErrorLevel, msg, f...)
}

//Fatal
func (log *zapLog) Fatal(v ...interface{}) {
	log.logMsgData(nil, zapcore.FatalLevel, "", v...)
}

func (log *zapLog) Fatalf(format string, v ...interface{}) {
	log.logMsgData(nil, zapcore.FatalLevel, format, v...)
}

func (log *zapLog) Fatalc(ctx context.Context, v ...interface{}) {
	log.logMsgData(ctx, zapcore.FatalLevel, "", v...)
}

func (log *zapLog) Fatalcf(ctx context.Context, format string, v ...interface{}) {
	log.logMsgData(ctx, zapcore.FatalLevel, format, v...)
}

func (log *zapLog) Fatalw(msg string, f ...zap.Field) {
	log.slogFieldData(nil, zapcore.FatalLevel, "", f...)
}

func (log *zapLog) Fatalwc(ctx context.Context, msg string, f ...zap.Field) {
	log.slogFieldData(ctx, zapcore.FatalLevel, msg, f...)
}

//Panic
func (log *zapLog) Panic(v ...interface{}) {
	log.logMsgData(nil, zapcore.PanicLevel, "", v...)
}

func (log *zapLog) Panicf(format string, v ...interface{}) {
	log.logMsgData(nil, zapcore.PanicLevel, format, v...)
}

func (log *zapLog) Panicc(ctx context.Context, v ...interface{}) {
	log.logMsgData(ctx, zapcore.PanicLevel, "", v...)
}

func (log *zapLog) Paniccf(ctx context.Context, format string, v ...interface{}) {
	log.logMsgData(ctx, zapcore.PanicLevel, format, v...)
}

func (log *zapLog) Panicw(msg string, f ...zap.Field) {
	log.slogFieldData(nil, zapcore.PanicLevel, "", f...)
}

func (log *zapLog) Panicwc(ctx context.Context, msg string, f ...zap.Field) {
	log.slogFieldData(ctx, zapcore.PanicLevel, msg, f...)
}

//Print
func (log *zapLog) Print(v ...interface{}) {
	log.logMsgData(nil, zapcore.InfoLevel, "", v...)
}

//Printf
func (log *zapLog) Printf(format string, v ...interface{}) {
	log.logMsgData(nil, zapcore.InfoLevel, format, v...)
}

//Printf
func (log *zapLog) Println(v ...interface{}) {
	log.logMsgData(nil, zapcore.InfoLevel, "", v...)
}

func (log *zapLog) logMsgData(ctx context.Context, level zapcore.Level, format string, args ...interface{}) {

	if !log.atom.Enabled(level) {
		return
	}
	logfunc := log.zlog.Debug

	switch level {
	case zapcore.DebugLevel:
		logfunc = log.zlog.Debug
	case zapcore.InfoLevel:
		logfunc = log.zlog.Info
	case zapcore.WarnLevel:
		logfunc = log.zlog.Warn
	case zapcore.DPanicLevel:
		logfunc = log.zlog.DPanic
	case zapcore.PanicLevel:
		logfunc = log.zlog.Panic
	case zapcore.FatalLevel:
		logfunc = log.zlog.Fatal
	}

	if ctx != nil {
		if len(format) == 0 {
			logfunc(fmt.Sprint(args...), log.ctxLog.ContextInfo(ctx)...)
		} else {
			logfunc(fmt.Sprintf(format, args...), log.ctxLog.ContextInfo(ctx)...)
		}
	} else {
		if len(format) == 0 {
			logfunc(fmt.Sprint(args...))
		} else {
			logfunc(fmt.Sprintf(format, args...))
		}
	}
}

func (log *zapLog) slogFieldData(ctx context.Context, level zapcore.Level, msg string, fields ...zap.Field) {
	if !log.atom.Enabled(level) {
		return
	}
	logfunc := log.zlog.Debug

	switch level {
	case zapcore.DebugLevel:
		logfunc = log.zlog.Debug
	case zapcore.InfoLevel:
		logfunc = log.zlog.Info
	case zapcore.WarnLevel:
		logfunc = log.zlog.Warn
	case zapcore.DPanicLevel:
		logfunc = log.zlog.DPanic
	case zapcore.PanicLevel:
		logfunc = log.zlog.Panic
	case zapcore.FatalLevel:
		logfunc = log.zlog.Fatal
	}

	if ctx != nil {
		logfunc(msg, append(log.ctxLog.ContextInfo(ctx), fields...)...)

	} else {
		logfunc(msg, append(log.ctxLog.ContextInfo(ctx), fields...)...)
	}
}
