package quic

import (
	"fmt"
	"testing"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/stretchr/testify/require"
)

func TestQuic_1(t *testing.T) {
	err := http3.ListenAndServe("", "", "", nil)
	require.NoError(t, err)
}

func TestQuic_Config_1(t *testing.T) {
	var qconf *quic.Config //= new(quic.Config)
	initQuicConfig(qconf)
	require.NotNil(t, qconf)
}

func initQuicConfig(qconf *quic.Config) {
	if qconf == nil {
		qconf = &quic.Config{
			MaxIdleTimeout:     15 * time.Second,
			MaxIncomingStreams: 1000,
			Allow0RTT:          true,
		}
	} else {
		qconf.MaxIdleTimeout = 333 * time.Second
		qconf.MaxIncomingStreams = 555
	}
	fmt.Printf("qconf: %#v", qconf)
}
