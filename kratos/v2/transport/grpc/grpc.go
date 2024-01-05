package grpc

import (
	"context"

	"github.com/blink-io/x/kratos/v2/transport"
	"github.com/blink-io/x/kratos/v2/util"
	"github.com/go-kratos/kratos/v2/middleware"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"

	"google.golang.org/grpc"
)

type Validator = util.Validator

type ServiceRegistrar = grpc.ServiceRegistrar

type Server interface {
	ServiceRegistrar

	transport.Server

	transport.Endpointer

	Validator

	Use(selector string, m ...middleware.Middleware)

	GetServiceInfo() map[string]grpc.ServiceInfo

	Raw() *grpc.Server
}

var _ Server = (*server)(nil)

type isrv = kgrpc.Server

type server struct {
	*isrv
}

func (s *server) Validate(ctx context.Context) error {
	return nil
}

func NewServer(opts ...ServerOption) Server {
	isrv := kgrpc.NewServer(opts...)
	s := &server{
		isrv: isrv,
	}
	return s
}

// Raw exposes the raw grpc.Server.
func (s *server) Raw() *grpc.Server {
	return s.Server
}
