package tests

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/util/random"
	"github.com/lmittmann/tint"
	"gitlab.com/greyxor/slogor"
)

func TestSlog_1(t *testing.T) {
	//arg1 := slog.String("version", "v1.0.0")
	//arg2 := slog.Int("version", 1234)
	for i := 0; i < 20; i++ {
		slog.Default().Info("[duration] " + random.String(100))
	}
}

func TestErrors_1(t *testing.T) {
	//slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, &slogor.Config{
	//	TimeFormat: time.Stamp,
	//	Level:      slog.LevelDebug,
	//	ShowSource: false,
	//})))

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

func TestColorLog_1(t *testing.T) {
	w := os.Stderr
	//thdlr := tint.NewHandler(os.Stderr, nil)
	// set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	slog.Error("Error with args", slogor.Err(errors.New("i'm an error")))
	slog.Warn("Warn with args", slog.Int("the_answer", 42))
	slog.Debug("Debug with args", slog.String("a_string", "🐛"))
}

func TestCtx_1(t *testing.T) {
	type ctxKey struct{}
	ctx := context.WithValue(context.Background(), ctxKey{}, time.Now())

	if t, ok := ctx.Value(ctxKey{}).(time.Time); ok {
		slog.Info("time.....  ", slog.Time("t", t))
	}
}

func TestXsqlLog(t *testing.T) {
	opt := &xsql.Config{
		Logger: func(format string, args ...any) {
			slog.Default().Info(fmt.Sprintf(format, args...))
		},
	}

	opt.Logger("hello, %s", "world")
}
