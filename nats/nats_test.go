package nats

import (
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/require"
)

func TestNATS_1(t *testing.T) {
	cfg := nats.GetDefaultOptions()
	cc, err := cfg.Connect()
	require.NoError(t, err)

	ecc, err := nats.NewEncodedConn(cc, "json")
	require.NoError(t, err)
	require.NotNil(t, ecc)
}
