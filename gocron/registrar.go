package gocron

import (
	"context"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

type (
	Task          = gocron.Task
	Job           = gocron.Job
	JobOption     = gocron.JobOption
	JobDefinition = gocron.JobDefinition
)

type RegisterToGocronFunc func(context.Context, ServiceRegistrar) error

type WithRegistrar interface {
	GocronRegistrar(context.Context) RegisterToGocronFunc
}

type ServiceRegistrar interface {
	// NewJob creates a new job in the Scheduler. The job is scheduled per the provided
	// definition when the Scheduler is started. If the Scheduler is already running
	// the job will be scheduled when the Scheduler is started.
	NewJob(JobDefinition, Task, ...JobOption) (Job, error)
	// RemoveByTags removes all jobs that have at least one of the provided tags.
	RemoveByTags(...string)
	// RemoveJob removes the job with the provided id.
	RemoveJob(uuid.UUID) error
	// Update replaces the existing Job's JobDefinition with the provided
	// JobDefinition. The Job's Job.ID() remains the same.
	Update(uuid.UUID, JobDefinition, Task, ...JobOption) (Job, error)
}

type RegistrarFunc[S any] func(ServiceRegistrar, S)

type RegistrarErrFunc[S any] func(ServiceRegistrar, S) error

type CtxRegistrarFunc[S any] func(context.Context, ServiceRegistrar, S)

type CtxRegistrarErrFunc[S any] func(context.Context, ServiceRegistrar, S) error

type Registrar interface {
	RegisterToGocron(context.Context, ServiceRegistrar) error
}

var _ Registrar = (*registrar[any])(nil)

func NewRegistrar[S any](s S, f RegistrarFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(r, s)
		return nil
	}
	return NewCtxRegistrarWithErr(s, cf)
}

func NewRegistrarWithErr[S any](s S, f RegistrarErrFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		return f(r, s)
	}
	return NewCtxRegistrarWithErr(s, cf)
}

func NewCtxRegistrar[S any](s S, f CtxRegistrarFunc[S]) Registrar {
	cf := func(ctx context.Context, r ServiceRegistrar, s S) error {
		f(ctx, r, s)
		return nil
	}
	return NewCtxRegistrarWithErr(s, cf)
}

func NewCtxRegistrarWithErr[S any](s S, f CtxRegistrarErrFunc[S]) Registrar {
	h := &registrar[S]{
		s: s,
		f: f,
	}
	return h
}

type registrar[S any] struct {
	s S
	f CtxRegistrarErrFunc[S]
}

func (h *registrar[S]) RegisterToGocron(ctx context.Context, r ServiceRegistrar) error {
	return h.f(ctx, r, h.s)
}
