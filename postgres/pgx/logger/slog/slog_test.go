package slog

import (
	"context"
	"log/slog"
	"testing"

	"github.com/jackc/pgx/v5/tracelog"
)

func TestSlog_1(t *testing.T) {
	ctx := context.Background()
	ll := NewLogger(slog.Default())

	logWithData := func(data map[string]any) {
		ll.Log(ctx, tracelog.LogLevelInfo, "msg1", data)
	}

	d1 := map[string]any{
		"k1": "ok",
		"k2": true,
		"k3": 123,
		"k4": map[string]any{
			"k4_1": "k4_1",
			"k4_2": 12,
			"k4_3": true,
			"k4_4": 3.1415,
			"k4_5": []string{"one", "two"},
		},
		"k5": []string{"one", "two"},
	}
	logWithData(d1)
}
