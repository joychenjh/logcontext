package logcontext

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

var _traceIDContextKey = "_traceIDContexKey"
var _traceSTContextKey = "_traceSTContexKey"

type ContextLog interface {
	ContextInfo(ctx context.Context) []zap.Field
	Key() string
}

type nilContext struct {
}

func NewNilContext() ContextLog {
	return &nilContext{}
}

func (nilContext) ContextInfo(ctx context.Context) []zap.Field {
	return []zap.Field{}
}

func (nilContext) Key() string {
	return ""
}

type TraceContext struct {
	KInfo string
}

func NewTraceContext() ContextLog {
	return &TraceContext{KInfo: "_traceid"}
}
func (tContext *TraceContext) ContextInfo(ctx context.Context) []zap.Field {

	_ustr := ""
	if _v, ok := ctx.Value(_traceSTContextKey).(time.Time); ok {
		_ustr = time.Since(_v).String()
	}
	return []zap.Field{zap.String("_traceid", fmt.Sprint(ctx.Value(_traceIDContextKey))), zap.String("_ut", _ustr)}
}

func (tContext *TraceContext) Key() string {
	return tContext.KInfo
}
