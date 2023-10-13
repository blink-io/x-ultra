package repository

import (
	"context"

	"github.com/uptrace/bun"
)

type Tx[M Model, I ID] interface {
	// Insert a new record.
	Insert(context.Context, *M, ...InsertOption) error
	// BulkInsert inserts more than one record
	BulkInsert(context.Context, []*M, ...InsertOption) error
	// Update a record by ID
	Update(context.Context, *M, ...UpdateOption) error
	// Delete a record by ID
	Delete(context.Context, I, ...DeleteOption) error
	// BulkDelete deletes by IDs
	BulkDelete(context.Context, []I, ...DeleteOption) error
	// Get a record by ID
	Get(context.Context, I, ...SelectOption) (*M, error)
	// One get one record by criteria
	One(context.Context, ...SelectOption) (*M, error)
	// All fetch all data from repository
	All(context.Context, ...SelectOption) ([]*M, error)
	// Count rows
	Count(context.Context, ...SelectOption) (int, error)
	// Exists has record
	Exists(context.Context, ...SelectOption) (bool, error)
}

var _ Tx[Model, ID] = (*tx[Model, ID])(nil)

type tx[M Model, I ID] struct {
	rx bun.Tx
}

func NewTx[M Model, I ID](rx bun.Tx) Tx[M, I] {
	return &tx[M, I]{rx: rx}
}

func (g *tx[M, I]) Insert(ctx context.Context, m *M, ops ...InsertOption) error {
	return Insert[M](ctx, g.rx, m, ops...)
}

func (g *tx[M, I]) BulkInsert(ctx context.Context, ms []*M, ops ...InsertOption) error {
	return BulkInsert[M](ctx, g.rx, ms, ops...)
}

func (g *tx[M, I]) Update(ctx context.Context, m *M, ops ...UpdateOption) error {
	return Update[M](ctx, g.rx, m, ops...)
}

func (g *tx[M, I]) Delete(ctx context.Context, ID I, ops ...DeleteOption) error {
	return Delete[M](ctx, g.rx, ID, ops...)
}

func (g *tx[M, I]) BulkDelete(ctx context.Context, IDs []I, ops ...DeleteOption) error {
	return BulkDelete[M, I](ctx, g.rx, IDs, ops...)
}

func (g *tx[M, I]) Get(ctx context.Context, ID I, ops ...SelectOption) (*M, error) {
	return Get[M, I](ctx, g.rx, ID, ops...)
}

func (g *tx[M, I]) One(ctx context.Context, ops ...SelectOption) (*M, error) {
	return One[M](ctx, g.rx, ops...)
}

func (g *tx[M, I]) All(ctx context.Context, ops ...SelectOption) ([]*M, error) {
	return All[M](ctx, g.rx, ops...)
}

func (g *tx[M, I]) Count(ctx context.Context, ops ...SelectOption) (int, error) {
	return Count[M](ctx, g.rx, ops...)
}

func (g *tx[M, I]) Exists(ctx context.Context, ops ...SelectOption) (bool, error) {
	return Exists[M](ctx, g.rx, ops...)
}
