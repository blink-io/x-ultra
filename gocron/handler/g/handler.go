package cron

import (
	"context"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

type ServiceRegistrar interface {
	// NewJob creates a new job in the Scheduler. The job is scheduled per the provided
	// definition when the Scheduler is started. If the Scheduler is already running
	// the job will be scheduled when the Scheduler is started.
	NewJob(gocron.JobDefinition, gocron.Task, ...gocron.JobOption) (gocron.Job, error)
	// RemoveByTags removes all jobs that have at least one of the provided tags.
	RemoveByTags(...string)
	// RemoveJob removes the job with the provided id.
	RemoveJob(uuid.UUID) error
	// Update replaces the existing Job's JobDefinition with the provided
	// JobDefinition. The Job's Job.ID() remains the same.
	Update(uuid.UUID, gocron.JobDefinition, gocron.Task, ...gocron.JobOption) (gocron.Job, error)
}

type RegistrarFunc[S any] func(ServiceRegistrar, S)

type CtxRegistrarFunc[S any] func(context.Context, ServiceRegistrar, S)

type RegistrarErrFunc[S any] func(ServiceRegistrar, S) error

type CtxRegistrarErrFunc[S any] func(context.Context, ServiceRegistrar, S) error

type Handler interface {
	HandleCron(context.Context, ServiceRegistrar) error
}

var _ Handler = (*handler[any])(nil)

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

func NewErrHandler[S any](s S, f RegistrarErrFunc[S]) Handler {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
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

func (h handler[S]) HandleCron(ctx context.Context, r ServiceRegistrar) error {
	return h.f(ctx, r, h.s)
}
