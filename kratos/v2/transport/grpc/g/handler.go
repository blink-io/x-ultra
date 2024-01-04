package g

import (
	"context"

	kgrpc "github.com/blink-io/x/kratos/v2/transport/grpc"
)

type RegistrarFunc[S any] func(kgrpc.ServiceRegistrar, S)

type Handler[S any] interface {
	HandleGRPC(context.Context, kgrpc.ServiceRegistrar)
}

type handler[S any] struct {
	s S
	f RegistrarFunc[S]
}

var _ Handler[any] = (*handler[any])(nil)

func NewHandler[S any](s S, f RegistrarFunc[S]) Handler[S] {
	h := &handler[S]{
		s: s,
		f: f,
	}
	return h
}

func (h handler[S]) HandleGRPC(ctx context.Context, r kgrpc.ServiceRegistrar) {
	h.f(r, h.s)
}

type CtxRegistrarFunc[S any] func(context.Context, kgrpc.ServiceRegistrar, S)

var _ Handler[any] = (*ctxHandler[any])(nil)

func NewCtxHandler[S any](s S, f CtxRegistrarFunc[S]) Handler[S] {
	h := &ctxHandler[S]{
		s: s,
		f: f,
	}
	return h
}

type ctxHandler[S any] struct {
	s S
	f CtxRegistrarFunc[S]
}

func (h ctxHandler[S]) HandleGRPC(ctx context.Context, r kgrpc.ServiceRegistrar) {
	h.f(ctx, r, h.s)
}
