package slog

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/stretchr/testify/require"
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
	logger := slog.Default()
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}

	si1 := logging.UnaryServerInterceptor(InterceptorLogger(logger), opts...)
	require.NotNil(t, si1)

	si2 := logging.StreamServerInterceptor(InterceptorLogger(logger), opts...)
	require.NotNil(t, si2)

	si3 := logging.UnaryClientInterceptor(InterceptorLogger(logger), opts...)
	require.NotNil(t, si3)

	si4 := logging.StreamClientInterceptor(InterceptorLogger(logger), opts...)
	require.NotNil(t, si4)
}
