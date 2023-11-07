package grpc_test

import (
	"context"
	"log/slog"
	"net"
	"os"
	"testing"
	"time"

	"github.com/blink-io/x/session"
	sessgrpc "github.com/blink-io/x/session/grpc"
	"github.com/blink-io/x/session/store/goredis"
	"github.com/blink-io/x/testdata"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"gitlab.com/greyxor/slogor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

func init() {
	w := os.Stderr
	//slog.SetDefault(slog.New(
	//	tint.NewHandler(w, &tint.Options{
	//		Level:      slog.LevelDebug,
	//		TimeFormat: time.Kitchen,
	//	}),
	//))
	slog.SetDefault(slog.New(slogor.NewHandler(w, &slogor.Options{
		TimeFormat: time.Kitchen,
		Level:      slog.LevelDebug,
		ShowSource: false,
	})))
}

type commonService struct {
	UnimplementedCommonServer
}

func (s *commonService) Health(ctx context.Context, req *HealthRequest) (*HealthResponse, error) {
	from := req.GetFrom()
	slog.Info("[Health] Invoke", "from", from)

	res := new(HealthResponse)
	res.Code = 200
	res.Message = "success"

	sm, ok := session.FromContext(ctx)
	if ok {
		sm.Put(ctx, "from", from)
		sm.Put(ctx, "is_admin", true)
		sm.Put(ctx, "version", &VersionRequest{
			From: "From Version",
		})
	}

	return res, nil
}

func (s *commonService) Version(ctx context.Context, req *VersionRequest) (*VersionResponse, error) {
	res := new(VersionResponse)
	res.Code = 200
	res.Message = "Very OK"
	res.Data = &VersionResponse_Data{
		Version: "v1.0.0",
		BuiltAt: "build-v1.0.0",
	}

	sm, ok := session.FromContext(ctx)
	if ok {
		v := sm.GetString(ctx, "from")
		slog.Info("[Version] Invoke, stored in session", "value", v)
	}

	return res, nil
}

func (s *commonService) Testing(ctx context.Context, req *TestingRequest) (*TestingResponse, error) {
	action := req.Action
	if action == "error" {
		//errMD := map[string]string{
		//	"Action":    action,
		//	"Operation": "OperationCommonTesting",
		//}
		return nil, status.Error(codes.InvalidArgument, "You input an error, I return it")
	}
	res := &TestingResponse{
		Code:    200,
		Message: "Action received: " + action,
		Data: &TestingResponse_Data{
			Action: action,
		},
	}
	return res, nil
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

func TestGRPC_Server_1(t *testing.T) {
	creds := credentials.NewTLS(testdata.GetTLSConfig())

	rc := redis.NewUniversalClient(&redis.UniversalOptions{})
	rs := goredis.New(rc)
	sm := session.NewManager(session.Store(rs))
	sh := sessgrpc.NewSessionHandler(
		sessgrpc.WithExposeExpiry(),
		sessgrpc.WithHeader(sessgrpc.DefaultHeader),
		sessgrpc.WithSessionManager(sm))

	gsrv := grpc.NewServer(
		grpc.Creds(creds),
		grpc.StatsHandler(&statsHandler{}),
		grpc.ChainUnaryInterceptor(sh.UnaryServerInterceptor),
		grpc.ChainStreamInterceptor(sh.StreamServerInterceptor),
	)

	svc := &commonService{}

	RegisterCommonServer(gsrv, svc)
	service.RegisterChannelzServiceToServer(gsrv)

	ln, err := net.Listen("tcp", ":9999")
	require.NoError(t, err)

	slog.Info("GRPC is listening on", "addr", ln.Addr().String())
	require.NoError(t, gsrv.Serve(ln))
}
