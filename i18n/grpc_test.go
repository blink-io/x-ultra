package i18n

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/blink-io/x/testdata"
	"github.com/stretchr/testify/require"
)

func TestGRPC_Server_1(t *testing.T) {
	gsrv := testdata.CreateGRPCServer(true)

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
	}

	var ff = EntryHandlerFunc(func(ctx context.Context, languages []string) map[string]*Entry {
		return entries
	})

	RegisterEntryHandler(gsrv, ff)

	ln, err := net.Listen("tcp", ":9999")
	require.NoError(t, err)

	require.NoError(t, gsrv.Serve(ln))
}

func TestNewGRPCLoader_1(t *testing.T) {
	cc := testdata.CreateGRPCClient(":9999", true)
	ld := NewGRPCLoader(cc, []string{"zh-Hans"})
	err := ld.Load(bb)
	require.NoError(t, err)
}

func TestPB_Desc(t *testing.T) {
	pbdesc := file_i18n_proto_rawDescGZIP()
	fmt.Println(string(pbdesc))

	err := os.WriteFile("./i18n.pb.bin", pbdesc, 0544)
	require.NoError(t, err)
}
