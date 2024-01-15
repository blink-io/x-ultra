package g

import (
	"context"

	khttp "github.com/blink-io/x/kratos/v2/transport/http"
)

type RegistrarFunc[S any] func(khttp.ServerRouter, S)

type CtxRegistrarFunc[S any] func(context.Context, khttp.ServerRouter, S)

type RegistrarFuncWithErr[S any] func(khttp.ServerRouter, S) error

type CtxRegistrarFuncWithErr[S any] func(context.Context, khttp.ServerRouter, S) error

type Handler = khttp.Handler

type handler[S any] struct {
	s S
	f CtxRegistrarFuncWithErr[S]
}

var _ Handler = (*handler[any])(nil)

func NewHandler[S any](s S, f RegistrarFunc[S]) Handler {
	cf := func(ctx context.Context, r khttp.ServerRouter, s S) error {
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
	cf := func(ctx context.Context, r khttp.ServerRouter, s S) error {
		f(ctx, r, s)
		return nil
	}
	h := &handler[S]{
		s: s,
		f: cf,
	}
	return h
}

// NewErrHandler creates a handler with returning error.
func NewErrHandler[S any](s S, f RegistrarFuncWithErr[S]) Handler {
	cf := func(ctx context.Context, r khttp.ServerRouter, s S) error {
		return f(r, s)
	}
	h := &handler[S]{
		s: s,
		f: cf,
	}
	return h
}

// NewCtxErrHandler creates a handler with a context parameter and returning error.
func NewCtxErrHandler[S any](s S, f CtxRegistrarFuncWithErr[S]) Handler {
	h := &handler[S]{
		s: s,
		f: f,
	}
	return h
}

func (h handler[S]) HandleHTTP(ctx context.Context, r khttp.ServerRouter) error {
	return h.f(ctx, r, h.s)
}
