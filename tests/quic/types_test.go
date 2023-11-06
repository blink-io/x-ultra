package quic

import (
	"net"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/stretchr/testify/require"
)

func TestQuic_Types(t *testing.T) {
	var _ http3.QUICEarlyListener = (*quic.EarlyListener)(nil)
}

func TestNet_Listen(t *testing.T) {
	c := qt.New(t)
	require.NotNil(t, c)

	ln, err := net.Listen("tcp", ":9999")
	require.NoError(t, err)
	require.NotNil(t, ln)

	uaddr, err := net.ResolveUDPAddr("udp", ":9999")
	require.NoError(t, err)
	require.NotNil(t, uaddr)

	ucc, err := net.ListenUDP("udp", uaddr)
	require.NotNil(t, err)
	require.NotNil(t, ucc)
}
