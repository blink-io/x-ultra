package i18n

import (
	"log/slog"
	"net"
	"testing"

	gslog "github.com/blink-io/x/grpc/slog"
	"github.com/blink-io/x/internal/testutil"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/grpclog"
)

func init() {
	grpclog.SetLoggerV2(gslog.NewLogger(slog.Default()))
}

func TestGRPC_Server_1(t *testing.T) {
	gsrv := testutil.CreateGRPCServer(true)

	zhHansJSON := `{"name":"广州", "language":"简体中文"}`
	enUSJSON := `{"name":"gz", "language":"American English"}`

	entries := map[string]*Entry{
		"zh-Hans": {
			Path:     "zh-Hans.json",
			Language: "zh-Hans",
			Valid:    true,
			Payload:  []byte(zhHansJSON),
		},
		"en-US": {
			Path:     "en-US.json",
			Language: "en-US",
			Valid:    true,
			Payload:  []byte(enUSJSON),
		},
		"en-UK": {
			Path:     "en-UK.json",
			Language: "en-UK",
			Valid:    false,
			Payload:  []byte(""),
		},
	}

	var ff = Entries(entries)

	RegisterEntryHandler(gsrv, ff)

	ln, err := net.Listen("tcp", ":9999")
	require.NoError(t, err)

	require.NoError(t, gsrv.Serve(ln))
}

func TestNewGRPCLoader_1(t *testing.T) {
	cc := testutil.CreateGRPCClient(":9999", true)
	ld := NewGRPCLoader(cc, []string{"zh-Hans"})
	err := ld.Load(bb)

	require.NoError(t, err)
}
