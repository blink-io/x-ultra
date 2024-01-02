package grpc

import (
	"google.golang.org/grpc"
)

type Handler interface {
	HandleGRPC(grpc.ServiceRegistrar)
}
