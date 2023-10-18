package tests

import (
	"log/slog"
	"testing"
)

func TestSlog_1(t *testing.T) {
	arg1 := slog.String("version", "v1.0.0")
	arg2 := slog.Int("version", 1234)
	slog.Info("maa", arg1, arg2)
}

func TestErrors_1(t *testing.T) {
}
