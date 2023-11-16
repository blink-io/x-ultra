package slog

import (
	"errors"
	"log/slog"
	"testing"
)

func TestSlog_1(t *testing.T) {
	l := New(slog.Default())
	l.Print()
	l.Print("abc", errors.New("kkk"), true, 123)
	l.Print("abc", errors.New("kkk"), true, 123, 3.1415)
}
