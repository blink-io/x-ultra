package g

import (
	"context"

	kgrpc "github.com/blink-io/x/kratos/v2/transport/grpc"
)

type RegistrarFunc[S any] func(kgrpc.ServiceRegistrar, S)

type CtxRegistrarFunc[S any] func(context.Context, kgrpc.ServiceRegistrar, S)

type Handler = kgrpc.Handler

type handler[S any] struct {
	s S
	f CtxRegistrarFunc[S]
}

var _ Handler = (*handler[any])(nil)

func NewHandler[S any](s S, f RegistrarFunc[S]) Handler {
	cf := func(ctx context.Context, r kgrpc.ServiceRegistrar, s S) {
		f(r, s)
	}
	h := &handler[S]{
		s: s,
		f: cf,
	}
	return h
}

func NewCtxHandler[S any](s S, f CtxRegistrarFunc[S]) Handler {
	h := &handler[S]{
		s: s,
		f: f,
	}
	return h
}

func (h handler[S]) HandleGRPC(ctx context.Context, r kgrpc.ServiceRegistrar) {
	h.f(ctx, r, h.s)
}
