package slog

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/blink-io/x/log"
	"gitlab.com/greyxor/slogor"
)

func init() {
	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, &slogor.Options{
		TimeFormat: time.Stamp,
		Level:      slog.LevelDebug,
		ShowSource: false,
	})))
}

func TestSlog_1(t *testing.T) {
	sl := NewLogger(slog.Default())

	sl.Log(log.LevelInfo, "name", "heison", "reason", "internal error")
	sl.Log(log.LevelInfo, "name", "heison", "reason", "internal error", "loc")
}
