package grpc

import (
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
)

type Doer interface {
	DoGRPC(*kgrpc.Server) error
}
