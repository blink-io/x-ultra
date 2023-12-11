package zap

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestZap_1(t *testing.T) {
	zlog, err := zap.NewDevelopment()
	require.NoError(t, err)

	log := NewLogger(zlog)
	log.Info("INFO odd", "ver", "v1", "score", 99.0, "kkk")
	log.Info("INFO even", "host", "127.0.0.1", "port", 3306)
	log.Info("INFO empty")
}
