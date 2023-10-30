package quic

import (
	"testing"

	"github.com/quic-go/quic-go/http3"
	"github.com/stretchr/testify/require"
)

func TestQuic_1(t *testing.T) {
	err := http3.ListenAndServe("", "", "", nil)
	require.NoError(t, err)
}
