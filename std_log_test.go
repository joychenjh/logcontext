package logcontext

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"go.uber.org/zap/zapcore"
)

func Test_StdLog(t *testing.T) {

	sLog := stdLogger{
		stdLog:  log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.LstdFlags|log.Lshortfile),
		ctxOpt:  NewTraceContext(),
		atom:    _gatomLogLevel,
		fencode: zapcore.NewConsoleEncoder(zapcore.EncoderConfig{}),
	}
	ctx := context.WithValue(context.Background(), _traceIDContextKey, time.Now().UnixNano())
	ctx = context.WithValue(ctx, _traceSTContextKey, time.Now())
	sLog.Debug("DEBUGMGS")

	sLog.Debugc(ctx, "hihi ctx")
	sLog.Debugcf(ctx, "hihi %s", "我是")
	sLog.Println("Prinln")
	sLog.Info("INFO")
	sLog.Warn("WARN")
	sLog.Error("ERROR")
	sLog.Fatal("FATAL")
}
