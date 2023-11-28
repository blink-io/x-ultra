package xlog

import (
	xlog "github.com/blink-io/x/log"
	"go.temporal.io/sdk/log"
)

type logger struct {
	i xlog.Logger
}

var _ log.Logger = (*logger)(nil)

func NewLogger(i xlog.Logger) log.Logger {
	return &logger{i: i}
}

func (l *logger) Debug(msg string, keyvals ...any) {
	l.i.Log(xlog.LevelDebug, msgWithKeyVals(msg, keyvals...)...)
}

func (l *logger) Info(msg string, keyvals ...any) {
	l.i.Log(xlog.LevelInfo, msgWithKeyVals(msg, keyvals...)...)
}

func (l *logger) Warn(msg string, keyvals ...any) {
	l.i.Log(xlog.LevelWarn, msgWithKeyVals(msg, keyvals...)...)
}

func (l *logger) Error(msg string, keyvals ...any) {
	l.i.Log(xlog.LevelError, msgWithKeyVals(msg, keyvals...)...)
}

func msgWithKeyVals(msg string, keyvals ...any) []any {
	keylen := len(keyvals)
	if keylen > 0 && keylen%2 != 0 {
		keylen -= 1
	}
	data := make([]any, 0, (keylen/2)+2)
	data = append(data, "msg", msg)
	for i := 0; i < keylen; i += 2 {
		data = append(data, keyvals[i], keyvals[i+1])
	}
	return data
}
