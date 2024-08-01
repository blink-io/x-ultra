package thrift

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/blink-io/x/i18n/grpc"

	i18nthrift "github.com/blink-io/x/i18n/thrift"
)

func TestServer(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := context.Background()

	zhHansJSON := `{"name":"广州", "language":"简体中文", "from":"测试模式"}`
	enUSJSON := `{"name":"gz", "language":"American English", "from":"TestMode"}`

	entries := map[string]*grpc.Entry{
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

	th := i18nthrift.NewHandler(grpc.Entries(entries))
	srv := NewServer(
		WithAddress(":7700"),
		WithProcessor(i18nthrift.NewI18NProcessor(th)),
	)

	if err := srv.Start(ctx); err != nil {
		panic(err)
	}

	defer func() {
		if err := srv.Stop(ctx); err != nil {
			t.Errorf("expected nil got %v", err)
		}
	}()

	<-interrupt
}
