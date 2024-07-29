package repository

import (
	"context"
	bunx "github.com/blink-io/x/bun"
)

type CrudRepository[M bunx.ModelType, ID bunx.IDType] interface {
	Create(ctx context.Context, m M) error
	Update(ctx context.Context, m M) error
	Delete(ctx context.Context, m M) error
	One(ctx context.Context, m M) (*M, error)
	All(ctx context.Context, m M) (*[]M, error)
}
