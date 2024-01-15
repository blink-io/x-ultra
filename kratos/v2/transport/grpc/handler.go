package grpc

import (
	"context"
)

type WithHandler interface {
	GRPCHandler() Handler
}

type Handler interface {
	HandleGRPC(context.Context, ServiceRegistrar) error
}
