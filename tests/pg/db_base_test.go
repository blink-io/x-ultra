package pg

import (
	"context"
	"fmt"
	"log/slog"
	"testing"

	"github.com/blink-io/x/sql/hooks"
	"github.com/blink-io/x/sql/hooks/timing"
)

var ctx = context.Background()

func init() {
	infoEnabled := slog.Default().Enabled(ctx, slog.LevelInfo)
	fmt.Println("Info enabled: ", infoEnabled)
	//crlog := slog.New(color.NewHandler(
	//	os.Stdout,
	//	&color.Options{
	//		Level: slog.LevelInfo,
	//	}))
	//slog.SetDefault(crlog)
	//defLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	//	Level: slog.LevelInfo,
	//}))
	//slog.SetDefault(defLogger)
}

func newDriverHooks() []hooks.Hooks {
	drvHooks := make([]hooks.Hooks, 0)
	drvHooks = append(drvHooks, timing.New(timing.Logf(func(format string, args ...any) {
		msg := fmt.Sprintf(format, args...)
		slog.Default().Info(msg)
	})))

	return drvHooks
}

func TestSlogDef(t *testing.T) {
	slog.Default().Warn("Slog Warn Level")
	slog.Default().Info("Slog Info Level")
}
