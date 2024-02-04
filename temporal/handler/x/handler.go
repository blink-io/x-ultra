package x

import (
	"context"

	"go.temporal.io/sdk/client"
)

type ServiceRegistrar interface {
	client.Client
	Client() client.Client
}

type RegistrarFunc[S any] func(ServiceRegistrar, S)

type CtxRegistrarFunc[S any] func(context.Context, ServiceRegistrar, S)

type RegistrarErrFunc[S any] func(ServiceRegistrar, S) error

type CtxRegistrarErrFunc[S any] func(context.Context, ServiceRegistrar, S) error

type Handler interface {
	HandleTemporal(context.Context, ServiceRegistrar) error
}

type handler[S any] struct {
	s S
	f CtxRegistrarErrFunc[S]
}

var _ Handler = (*handler[any])(nil)

func (h handler[S]) HandleTemporal(ctx context.Context, r ServiceRegistrar) error {
	return h.f(ctx, r, h.s)
}

func NewHandler[S any](s S, f RegistrarFunc[S]) Handler {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxErrHandler(s, cf)
}

func NewCtxHandler[S any](s S, f CtxRegistrarFunc[S]) Handler {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(ctx, r, s)
		return nil
	}
	return NewCtxErrHandler(s, cf)
}

func NewErrHandler[S any](s S, f RegistrarFunc[S]) Handler {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxErrHandler(s, cf)
}

func NewCtxErrHandler[S any](s S, f CtxRegistrarErrFunc[S]) Handler {
	h := &handler[S]{
		s: s,
		f: f,
	}
	return h
}
