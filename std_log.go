package logcontext

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type stdLogger struct {
	stdLog    *log.Logger
	ctxOpt    ContextLog
	atom      zap.AtomicLevel
	fencode   zapcore.Encoder
	calldepth int
}

const (
	_stdcalldepth = 3
)

var (
	_gatomLogLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
)

func NewDefStdLog() LogContext {
	return &stdLogger{
		stdLog:    log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.LstdFlags|log.Lshortfile),
		ctxOpt:    nilContext{},
		atom:      _gatomLogLevel,
		calldepth: 1,
		fencode:   zapcore.NewJSONEncoder(zapcore.EncoderConfig{}),
	}
}

func (l *stdLogger) SetContextLog(c ContextLog) LogContext {
	l.ctxOpt = c
	return l
}

func (l *stdLogger) SetCallerSkip(calldepth int) LogContext {
	l.calldepth = calldepth
	return l
}

func (l *stdLogger) SetLogLevel(level LogLevel) LogContext {
	if zapcore.Level(level) <= zapcore.FatalLevel && zapcore.Level(level) >= zapcore.DebugLevel {
		l.atom.SetLevel(zapcore.Level(level))
	}
	return l
}

//Debug

func (l *stdLogger) Debug(v ...interface{}) {
	l.logMsgData(nil, DebugLevel, "", v...)
}

func (l *stdLogger) Debugf(format string, v ...interface{}) {
	l.logMsgData(nil, DebugLevel, format, v...)
}

func (l *stdLogger) Debugc(ctx context.Context, v ...interface{}) {
	l.logMsgData(ctx, DebugLevel, "", v...)
}

func (l *stdLogger) Debugcf(ctx context.Context, format string, v ...interface{}) {
	l.logMsgData(ctx, DebugLevel, format, v...)
}

func (l *stdLogger) Debugw(msg string, f ...zap.Field) {
	l.wlogFieldData(nil, DebugLevel, msg, f...)
}

func (l *stdLogger) Debugwc(ctx context.Context, msg string, f ...zap.Field) {
	l.wlogFieldData(ctx, DebugLevel, msg, f...)
}

//Info

func (l *stdLogger) Info(v ...interface{}) {
	l.logMsgData(nil, InfoLevel, "", v...)
}

func (l *stdLogger) Infof(format string, v ...interface{}) {
	l.logMsgData(nil, InfoLevel, format, v...)
}

func (l *stdLogger) Infoc(ctx context.Context, v ...interface{}) {
	l.logMsgData(ctx, InfoLevel, "", v...)
}

func (l *stdLogger) Infocf(ctx context.Context, format string, v ...interface{}) {
	l.logMsgData(ctx, InfoLevel, format, v...)
}

func (l *stdLogger) Infow(msg string, f ...zap.Field) {
	l.wlogFieldData(nil, InfoLevel, msg, f...)
}

func (l *stdLogger) Infowc(ctx context.Context, msg string, f ...zap.Field) {
	l.wlogFieldData(ctx, InfoLevel, msg, f...)
}

//Warn

func (l *stdLogger) Warn(v ...interface{}) {
	l.logMsgData(nil, WarnLevel, "", v...)
}
func (l *stdLogger) Warnf(format string, v ...interface{}) {
	l.logMsgData(nil, WarnLevel, format, v...)
}

func (l *stdLogger) Warnc(ctx context.Context, v ...interface{}) {
	l.logMsgData(ctx, WarnLevel, "", v...)
}

func (l *stdLogger) Warncf(ctx context.Context, format string, v ...interface{}) {
	l.logMsgData(ctx, WarnLevel, format, v...)
}

func (l *stdLogger) Warnw(msg string, f ...zap.Field) {
	l.wlogFieldData(nil, WarnLevel, msg, f...)
}

func (l *stdLogger) Warnwc(ctx context.Context, msg string, f ...zap.Field) {
	l.wlogFieldData(ctx, WarnLevel, msg, f...)
}

//Error
func (l *stdLogger) Error(v ...interface{}) {
	l.logMsgData(nil, ErrorLevel, "", v...)
}
func (l *stdLogger) Errorf(format string, v ...interface{}) {
	l.logMsgData(nil, ErrorLevel, format, v...)
}

func (l *stdLogger) Errorc(ctx context.Context, v ...interface{}) {
	l.logMsgData(ctx, ErrorLevel, "", v...)
}

func (l *stdLogger) Errorcf(ctx context.Context, format string, v ...interface{}) {
	l.logMsgData(ctx, ErrorLevel, format, v...)
}

func (l *stdLogger) Errorw(msg string, f ...zap.Field) {
	l.wlogFieldData(nil, ErrorLevel, msg, f...)
}

func (l *stdLogger) Errorwc(ctx context.Context, msg string, f ...zap.Field) {
	l.wlogFieldData(ctx, ErrorLevel, msg, f...)
}

//Fatal
func (l *stdLogger) Fatal(v ...interface{}) {
	l.logMsgData(nil, FatalLevel, "", v...)
}
func (l *stdLogger) Fatalf(format string, v ...interface{}) {
	l.logMsgData(nil, FatalLevel, format, v...)
}

func (l *stdLogger) Fatalc(ctx context.Context, v ...interface{}) {
	l.logMsgData(ctx, FatalLevel, "", v...)
}

func (l *stdLogger) Fatalcf(ctx context.Context, format string, v ...interface{}) {
	l.logMsgData(ctx, FatalLevel, format, v...)
}

func (l *stdLogger) Fatalw(msg string, f ...zap.Field) {
	l.wlogFieldData(nil, FatalLevel, msg, f...)
}

func (l *stdLogger) Fatalwc(ctx context.Context, msg string, f ...zap.Field) {
	l.wlogFieldData(ctx, FatalLevel, msg, f...)
}

//Panic
func (l *stdLogger) Panic(v ...interface{}) {
	l.logMsgData(nil, PanicLevel, "", v...)
}
func (l *stdLogger) Panicf(format string, v ...interface{}) {
	l.logMsgData(nil, PanicLevel, format, v...)
}

func (l *stdLogger) Panicc(ctx context.Context, v ...interface{}) {
	l.logMsgData(ctx, PanicLevel, "", v...)
}

func (l *stdLogger) Paniccf(ctx context.Context, format string, v ...interface{}) {
	l.logMsgData(ctx, PanicLevel, format, v...)
}

func (l *stdLogger) Panicw(msg string, f ...zap.Field) {
	l.wlogFieldData(nil, PanicLevel, msg, f...)
}

func (l *stdLogger) Panicwc(ctx context.Context, msg string, f ...zap.Field) {
	l.wlogFieldData(ctx, PanicLevel, msg, f...)
}

//Print
func (l *stdLogger) Print(v ...interface{}) {
	l.logMsgData(nil, InfoLevel, "", v...)
}

//Printf
func (l *stdLogger) Printf(format string, v ...interface{}) {
	l.logMsgData(nil, InfoLevel, format, v...)
}

//Printf
func (l *stdLogger) Println(v ...interface{}) {
	l.logMsgData(nil, InfoLevel, "", v...)
}

func (log *stdLogger) logMsgData(ctx context.Context, level LogLevel, format string, args ...interface{}) {

	_zlevel := zapcore.Level(level)
	if !log.atom.Enabled(_zlevel) {
		return
	}

	//TODO msg可以优化.不使用+
	_msg := ""
	switch _zlevel {
	case zapcore.DebugLevel:
		_msg = Blue.Add(_zlevel.CapitalString())
	case zapcore.InfoLevel:
		_msg = Green.Add(_zlevel.CapitalString())
	case zapcore.WarnLevel:
		_msg = Yellow.Add(_zlevel.CapitalString())
	case zapcore.ErrorLevel:
		_msg = Red.Add(_zlevel.CapitalString())
	case zapcore.DPanicLevel:
		_msg = Magenta.Add(_zlevel.CapitalString())
	case zapcore.PanicLevel:
		_msg = Magenta.Add(_zlevel.CapitalString())
	case zapcore.FatalLevel:
		_msg = Magenta.Add(_zlevel.CapitalString())
	}
	_msg += "\t"
	if ctx != nil {
		_buf, err := log.fencode.EncodeEntry(zapcore.Entry{}, log.ctxOpt.ContextInfo(ctx))
		if err == nil {
			_buf.TrimNewline()
			_msg += _buf.String() + " "
			_buf.Free()
		}

		if len(format) == 0 {
			_msg += fmt.Sprint(args...)
		} else {
			_msg += fmt.Sprintf(format, args...)
		}

	} else {
		if len(format) == 0 {
			_msg += fmt.Sprint(args...)
		} else {
			_msg += fmt.Sprintf(format, args...)
		}
	}

	log.stdLog.Output(log.calldepth+_stdcalldepth, _msg)
}

func (log *stdLogger) wlogFieldData(ctx context.Context, level LogLevel, msg string, fields ...zap.Field) {
	_zlevel := zapcore.Level(level)
	if !log.atom.Enabled(_zlevel) {
		return
	}

	//TODO msg可以优化.不使用+
	_msg := ""
	switch _zlevel {
	case zapcore.DebugLevel:
		_msg = Blue.Add(_zlevel.CapitalString())
	case zapcore.InfoLevel:
		_msg = Green.Add(_zlevel.CapitalString())
	case zapcore.WarnLevel:
		_msg = Yellow.Add(_zlevel.CapitalString())
	case zapcore.DPanicLevel:
		_msg = Magenta.Add(_zlevel.CapitalString())
	case zapcore.PanicLevel:
		_msg = Magenta.Add(_zlevel.CapitalString())
	case zapcore.FatalLevel:
		_msg = Magenta.Add(_zlevel.CapitalString())
	}
	_msg += "\t"
	if ctx != nil {
		_buf, err := log.fencode.EncodeEntry(zapcore.Entry{}, append(fields, log.ctxOpt.ContextInfo(ctx)...))
		if err == nil {
			_buf.TrimNewline()
			_msg += " " + _buf.String()
			_buf.Free()
		}
	} else {
		_buf, err := log.fencode.EncodeEntry(zapcore.Entry{}, fields)
		if err == nil {
			_buf.TrimNewline()
			_msg += " " + _buf.String()
			_buf.Free()
		}
	}
	_msg += " " + msg

	log.stdLog.Output(log.calldepth+_stdcalldepth, _msg)
}
