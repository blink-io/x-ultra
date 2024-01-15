package g

import (
	"context"

	kgrpc "github.com/blink-io/x/kratos/v2/transport/grpc"
)

type RegistrarFunc[S any] func(kgrpc.ServiceRegistrar, S)

type CtxRegistrarFunc[S any] func(context.Context, kgrpc.ServiceRegistrar, S)

type RegistrarFuncWithErr[S any] func(kgrpc.ServiceRegistrar, S) error

type CtxRegistrarFuncWithErr[S any] func(context.Context, kgrpc.ServiceRegistrar, S) error

type Handler = kgrpc.Handler

type handler[S any] struct {
	s S
	f CtxRegistrarFuncWithErr[S]
}

var _ Handler = (*handler[any])(nil)

func NewHandler[S any](s S, f RegistrarFunc[S]) Handler {
	cf := func(ctx context.Context, r kgrpc.ServiceRegistrar, s S) error {
		f(r, s)
		return nil
	}
	h := &handler[S]{
		s: s,
		f: cf,
	}
	return h
}

func NewCtxHandler[S any](s S, f CtxRegistrarFunc[S]) Handler {
	cf := func(ctx context.Context, r kgrpc.ServiceRegistrar, s S) error {
		f(ctx, r, s)
		return nil
	}
	h := &handler[S]{
		s: s,
		f: cf,
	}
	return h
}

func NewHandlerErr[S any](s S, f RegistrarFuncWithErr[S]) Handler {
	cf := func(ctx context.Context, r kgrpc.ServiceRegistrar, s S) error {
		return f(r, s)
	}
	h := &handler[S]{
		s: s,
		f: cf,
	}
	return h
}

func NewCtxErrHandler[S any](s S, f CtxRegistrarFuncWithErr[S]) Handler {
	h := &handler[S]{
		s: s,
		f: f,
	}
	return h
}

func (h handler[S]) HandleGRPC(ctx context.Context, r kgrpc.ServiceRegistrar) error {
	return h.f(ctx, r, h.s)
}
