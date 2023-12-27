package grpc

import (
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
)

type Handler interface {
	HandleGRPC(*kgrpc.Server)
}
