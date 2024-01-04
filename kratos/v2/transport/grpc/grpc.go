package grpc

import (
	"net"
	"net/http"

	"github.com/blink-io/x/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/middleware"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
)

type server = kgrpc.Server

type Server interface {
	grpc.ServiceRegistrar

	transport.Server

	transport.Endpointer

	Use(selector string, m ...middleware.Middleware)

	GetServiceInfo() map[string]grpc.ServiceInfo
	Serve(lis net.Listener) error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	GracefulStop()
}

var _ Server = (*server)(nil)

func NewServer(opts ...ServerOption) Server {
	return kgrpc.NewServer(opts...)
}
