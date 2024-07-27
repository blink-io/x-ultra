package registrar

import (
	"context"

	"go.temporal.io/sdk/client"
)

type ServiceRegistrar interface {
	client.Client
	Client() client.Client
}

type RegistrarFunc[S any] func(ServiceRegistrar, S)

type RegistrarFuncWithErr[S any] func(ServiceRegistrar, S) error

type CtxRegistrarFunc[S any] func(context.Context, ServiceRegistrar, S)

type CtxRegistrarFuncWithErr[S any] func(context.Context, ServiceRegistrar, S) error

type Registrar interface {
	RegisterToTemporal(context.Context, ServiceRegistrar) error
}

type registrar[S any] struct {
	s S
	f CtxRegistrarFuncWithErr[S]
}

var _ Registrar = (*registrar[any])(nil)

func (h *registrar[S]) RegisterToTemporal(ctx context.Context, r ServiceRegistrar) error {
	return h.f(ctx, r, h.s)
}

func New[S any](s S, f RegistrarFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxWithErr(s, cf)
}

func NewCtx[S any](s S, f CtxRegistrarFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(ctx, r, s)
		return nil
	}
	return NewCtxWithErr(s, cf)
}

func NewWithErr[S any](s S, f RegistrarFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxWithErr(s, cf)
}

func NewCtxWithErr[S any](s S, f CtxRegistrarFuncWithErr[S]) Registrar {
	h := &registrar[S]{
		s: s,
		f: f,
	}
	return h
}
