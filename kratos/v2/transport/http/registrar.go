package http

import (
	"context"
)

type RegisterToHTTPFunc func(context.Context, ServerRouter) error

type WithRegistrar interface {
	HTTPRegistrar() RegisterToHTTPFunc
}

type RegistrarFunc[S any] func(ServerRouter, S)

type RegistrarFuncWithErr[S any] func(ServerRouter, S) error

type CtxRegistrarFunc[S any] func(context.Context, ServerRouter, S)

type CtxRegistrarFuncWithErr[S any] func(context.Context, ServerRouter, S) error

type Registrar interface {
	RegisterToHTTP(context.Context, ServerRouter) error
}

type registrar[S any] struct {
	s S
	f CtxRegistrarFuncWithErr[S]
}

var _ Registrar = (*registrar[any])(nil)

func NewRegistrar[S any](s S, f RegistrarFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServerRouter, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxRegistrarWithErr(s, cf)
}

func NewCtxRegistrar[S any](s S, f CtxRegistrarFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServerRouter, s S) error {
		f(ctx, r, s)
		return nil
	}
	return NewCtxRegistrarWithErr(s, cf)
}

// NewRegistrarWithErr creates a registrar with returning error.
func NewRegistrarWithErr[S any](s S, f RegistrarFuncWithErr[S]) Registrar {
	cf := func(ctx context.Context, r ServerRouter, s S) error {
		return f(r, s)
	}
	return NewCtxRegistrarWithErr(s, cf)
}

// NewCtxRegistrarWithErr creates a registrar with a context parameter and returning error.
func NewCtxRegistrarWithErr[S any](s S, f CtxRegistrarFuncWithErr[S]) Registrar {
	h := &registrar[S]{
		s: s,
		f: f,
	}
	return h
}

func (h registrar[S]) RegisterToHTTP(ctx context.Context, r ServerRouter) error {
	return h.f(ctx, r, h.s)
}
