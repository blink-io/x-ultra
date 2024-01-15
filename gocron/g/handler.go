package cron

import (
	"context"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

type Service interface {
	NewJob(gocron.JobDefinition, gocron.Task, ...gocron.JobOption) (gocron.Job, error)

	Update(uuid.UUID, gocron.JobDefinition, gocron.Task, ...gocron.JobOption) (gocron.Job, error)
}

type RegistrarFunc[S any] func(Service, S)

type CtxRegistrarFunc[S any] func(context.Context, Service, S)

type RegistrarErrFunc[S any] func(Service, S) error

type CtxRegistrarErrFunc[S any] func(context.Context, Service, S) error

type Handler interface {
	HandleCron(context.Context, Service) error
}

var _ Handler = (*handler[any])(nil)

func NewHandler[S any](s S, f RegistrarFunc[S]) Handler {
	cf := func(ctx context.Context, r Service, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxErrHandler(s, cf)
}

func NewCtxHandler[S any](s S, f CtxRegistrarFunc[S]) Handler {
	cf := func(ctx context.Context, r Service, s S) error {
		f(ctx, r, s)
		return nil
	}
	return NewCtxErrHandler(s, cf)
}

func NewErrHandler[S any](s S, f RegistrarErrFunc[S]) Handler {
	cf := func(ctx context.Context, r Service, s S) error {
		return f(r, s)
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

type handler[S any] struct {
	s S
	f CtxRegistrarErrFunc[S]
}

func (h handler[S]) HandleCron(ctx context.Context, r Service) error {
	return h.f(ctx, r, h.s)
}
