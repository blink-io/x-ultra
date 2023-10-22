package tests

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"gitlab.com/greyxor/slogor"
)

func TestSlog_1(t *testing.T) {
	arg1 := slog.String("version", "v1.0.0")
	arg2 := slog.Int("version", 1234)
	slog.Info("maa", arg1, arg2)
}

func TestErrors_1(t *testing.T) {
	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, &slogor.Options{
		TimeFormat: time.Stamp,
		Level:      slog.LevelDebug,
		ShowSource: false,
	})))

	slog.Info("I'm an information message, everything's fine")
	slog.Warn("I'm a warning, that's ok.")
	slog.Error("I'm an error message, this is serious...")
	slog.Debug("Useful debug message.")

	fmt.Println("")
	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, &slogor.Options{
		TimeFormat: time.Kitchen,
		Level:      slog.LevelDebug,
		ShowSource: false,
	})))

	slog.Info("Example with kitchen time.")
	slog.Debug("Example with kitchen time.")
	fmt.Println("")
	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, &slogor.Options{
		TimeFormat: time.RFC3339Nano,
		Level:      slog.LevelDebug,
		ShowSource: true,
	})))
	slog.Info("Example with RFC 3339 time and source path")

	fmt.Println("")
	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, &slogor.Options{
		TimeFormat: time.Stamp,
		Level:      slog.LevelDebug,
		ShowSource: false,
	})))

	slog.Error("Error with args", slogor.Err(errors.New("i'm an error")))
	slog.Warn("Warn with args", slog.Int("the_answer", 42))
	slog.Debug("Debug with args", slog.String("a_string", "🐛"))

}
