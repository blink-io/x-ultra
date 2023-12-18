package testutil

import (
	"context"
	"log"
	"log/slog"

	"github.com/blink-io/x/internal/testdata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/stats"
)

func CreateGRPCClient(target string, secure bool, ops ...grpc.DialOption) *grpc.ClientConn {
	var creds credentials.TransportCredentials
	if secure {
		creds = credentials.NewTLS(testdata.CreateClientTLSConfig())
	} else {
		creds = insecure.NewCredentials()
	}
	var newops = make([]grpc.DialOption, 0)
	newops = append(newops, grpc.WithTransportCredentials(creds))
	if len(ops) > 0 {
		newops = append(newops, ops...)
	}
	c, err := grpc.Dial(target, newops...)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func CreateGRPCServer(secure bool, ops ...grpc.ServerOption) *grpc.Server {
	var creds credentials.TransportCredentials
	if secure {
		creds = credentials.NewTLS(testdata.GetTLSConfig())
	} else {
		creds = insecure.NewCredentials()
	}

	var newops = make([]grpc.ServerOption, 0)
	newops = append(newops, grpc.Creds(creds), grpc.StatsHandler(&statsHandler{}))
	if len(ops) > 0 {
		newops = append(newops, ops...)
	}
	gsrv := grpc.NewServer(newops...)

	return gsrv
}

var _ stats.Handler = (*statsHandler)(nil)

type statsHandler struct {
}

func (s *statsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	slog.Info("Invoke [TagRPC]", "info", info)
	return ctx
}

func (s *statsHandler) HandleRPC(ctx context.Context, rpcStats stats.RPCStats) {
	slog.Info("Invoke [HandleRPC]", "rpcStats", rpcStats)
}

func (s *statsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	slog.Info("Invoke [TagConn]", "remote addr", info.RemoteAddr)
	slog.Info("Invoke [TagConn]", "local addr", info.LocalAddr)
	return ctx
}

func (s *statsHandler) HandleConn(ctx context.Context, connStats stats.ConnStats) {
	slog.Info("Invoke [HandleConn]", "connStats", connStats)
}
