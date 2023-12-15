package ext

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/blink-io/x/slog/handlers/color"
)

func TestNewLogger_1(t *testing.T) {
	crhdlr := color.NewHandler(os.Stdout, &color.Options{
		TimeFormat: time.RFC3339,
		Level:      slog.LevelInfo,
		ShowSource: true,
	})
	crlog := slog.New(crhdlr)

	elog := NewLogger(crlog)

	elog.Info("Info Log", "level", "info")
	elog.Fatal("Fatal Log", "level", "fatal")
}

func TestNewLogger_Default(t *testing.T) {
	elog := NewLogger(slog.Default())

	elog.Info("Info Log", "level", "info")
	elog.Fatal("Fatal Log", "level", "fatal")
}

func TestSlog_LvlVar(t *testing.T) {
	l := slog.Default()

	l.Info("pwd", "pwd", pwdVal("Password"))
}

type pwdVal string

func (p pwdVal) LogValue() slog.Value {
	return slog.StringValue("***")
}

var _ slog.LogValuer = (*pwdVal)(nil)
