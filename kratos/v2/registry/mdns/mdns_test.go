package mdns

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/require"
)

func TestDNS_Client_1(t *testing.T) {
	cli := &dns.Client{
		Net: "udp",
	}
	msg1 := &dns.Msg{}

	rm, rtt, err := cli.Exchange(msg1, ":9998")
	require.NoError(t, err)
	require.NotNil(t, rm)
	fmt.Println("rtt: ", rtt)
}

func TestDNS_Server_1(t *testing.T) {
	srv := &dns.Server{
		Addr: ":9998",
		Net:  "udp",
	}
	slog.Info("DNS Server is listening on", "add", srv.Addr)
	err := srv.ListenAndServe()
	require.NoError(t, err)
}
