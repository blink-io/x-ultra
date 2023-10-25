package zap

import (
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestZap_1(t *testing.T) {
	logger := zap.NewExample()
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
