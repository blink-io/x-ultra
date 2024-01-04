package g

import (
	"context"

	khttp "github.com/blink-io/x/kratos/v2/transport/http"
)

type RegistrarFunc[S any] func(khttp.ServerRouter, S)

type Handler[S any] interface {
	HandleHTTP(context.Context, khttp.ServerRouter)
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

func (h handler[S]) HandleHTTP(ctx context.Context, r khttp.ServerRouter) {
	h.f(r, h.s)
}

type CtxRegistrarFunc[S any] func(context.Context, khttp.ServerRouter, S)

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

func (h ctxHandler[S]) HandleHTTP(ctx context.Context, r khttp.ServerRouter) {
	h.f(ctx, r, h.s)
}
