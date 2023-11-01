package quic

import (
	"testing"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

func TestQuic_Types(t *testing.T) {
	var _ http3.QUICEarlyListener = (*quic.EarlyListener)(nil)
}
