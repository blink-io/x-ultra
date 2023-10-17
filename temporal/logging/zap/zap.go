package zap

import (
	"fmt"

	"github.com/blink-io/x/temporal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	i *zap.Logger
}

var _ temporal.Logger = (*logger)(nil)

func NewLogger(zlog *zap.Logger) temporal.Logger {
	return &logger{zlog}
}

func (l *logger) Debug(msg string, keyvals ...interface{}) {
	l.Log(zapcore.DebugLevel, msg, keyvals...)
}

func (l *logger) Info(msg string, keyvals ...interface{}) {
	l.Log(zapcore.InfoLevel, msg, keyvals...)
}

func (l *logger) Warn(msg string, keyvals ...interface{}) {
	l.Log(zapcore.WarnLevel, msg, keyvals...)
}

func (l *logger) Error(msg string, keyvals ...interface{}) {
	l.Log(zapcore.ErrorLevel, msg, keyvals...)
}

func (l *logger) Log(level zapcore.Level, msg string, keyvals ...interface{}) {
	keylen := len(keyvals)
	if keylen > 0 && keylen%2 != 0 {
		keylen -= 1
	}
	data := make([]zap.Field, 0, (keylen/2)+1)
	for i := 0; i < keylen; i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	switch level {
	case zapcore.DebugLevel:
		l.i.Debug(msg, data...)
	case zapcore.InfoLevel:
		l.i.Info(msg, data...)
	case zapcore.WarnLevel:
		l.i.Warn(msg, data...)
	case zapcore.ErrorLevel:
		l.i.Error(msg, data...)
	case zapcore.FatalLevel:
		l.i.Fatal(msg, data...)
	}
}

func (l *logger) Sync() error {
	return l.i.Sync()
}

func (l *logger) Close() error {
	return l.Sync()
}
