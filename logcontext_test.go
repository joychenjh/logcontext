package logcontext

import (
	"context"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Test_LogContext_std(t *testing.T) {

	Debugwc(context.TODO(), "hi all")
}

func Test_LogContext_zap(t *testing.T) {

	var zencoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
	}
	var atom = zap.NewAtomicLevelAt(zap.DebugLevel)

	var config = zap.Config{
		Level:            atom,                                                // 日志级别
		Development:      true,                                                // 开发模式，堆栈跟踪
		Encoding:         "json",                                              // 输出格式 console 或 json
		EncoderConfig:    zencoderConfig,                                      // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": "spikeProxy"}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout", "/tmp/zaplog.log"},               // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	tzLog, err := config.Build()
	if err != nil {
		t.Error("init zap log err")
		return
	}

	SetLog(NewZapLog(tzLog, atom))
	SetCallerSkip(3)

	Debug("zapdebug")

}
