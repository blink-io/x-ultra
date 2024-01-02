package slog

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
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

	sl.Log(log.LevelDebug, "name", "heison", "reason", "this is a debug info")

	sl.Log(log.LevelInfo,
		"enabled", true,
		"name", "fighter",
		"reason", "internal error",
		"loc", "12,34",
		"level", 10,
		"score", 101.1234,
	)

	sl.Log(log.LevelWarn, "warn info", "WARN is warning")

	sl.Log(log.LevelError, "reason", "system fault")
}
