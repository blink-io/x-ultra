package grpc_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/blink-io/x/session"
	sessgrpc "github.com/blink-io/x/session/grpc"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type commonService struct {
	UnimplementedCommonServer
}

func (s *commonService) Health(ctx context.Context, req *HealthRequest) (*HealthResponse, error) {
	from := req.GetFrom()
	log.Printf("[Health] from: %s", from)

	res := new(HealthResponse)
	res.Code = 200
	res.Message = "success"

	sm, ok := session.FromContext(ctx)
	if ok {
		sm.Put(ctx, "from", from)
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
		fmt.Println("[Version] Value stored in session:  ", v)
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

func TestGRPC_Server_1(t *testing.T) {
	sm := session.NewManager()
	sh := sessgrpc.NewSessionHandler(
		sessgrpc.WithHeader(sessgrpc.DefaultHeader),
		sessgrpc.WithSessionManager(sm))
	gsrv := grpc.NewServer(
		grpc.UnaryInterceptor(sh.UnaryServerInterceptor),
	)
	svc := &commonService{}

	RegisterCommonServer(gsrv, svc)

	ln, err := net.Listen("tcp", ":9999")
	require.NoError(t, err)

	log.Printf("GRPC is listening on %s", ln.Addr().String())
	require.NoError(t, gsrv.Serve(ln))
}
