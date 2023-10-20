package generics

import (
	"context"

	"github.com/blink-io/x/sql"
	"github.com/uptrace/bun"
)

type (
	itx = bun.Tx

	Tx[M Model, I ID] interface {
		bun.IDB
		base[M, I]
		// Commit the transaction
		Commit() error
		// Rollback the transaction
		Rollback() error
	}

	tx[M Model, I comparable] struct {
		itx
	}
)

var _ Tx[Model, IDType] = (*tx[Model, IDType])(nil)

func NewTx[M Model, I ID](itx bun.Tx) Tx[M, I] {
	return &tx[M, I]{itx}
}

func NewTxWithDB[M Model, I ID](db *sql.DB) (Tx[M, I], error) {
	itx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return NewTx[M, I](itx), nil
}

func (g *tx[M, I]) Insert(ctx context.Context, m *M, ops ...InsertOption) error {
	return Insert[M](ctx, g.itx, m, ops...)
}

func (g *tx[M, I]) BulkInsert(ctx context.Context, ms []*M, ops ...InsertOption) error {
	return BulkInsert[M](ctx, g.itx, ms, ops...)
}

func (g *tx[M, I]) Update(ctx context.Context, m *M, ops ...UpdateOption) error {
	return Update[M](ctx, g.itx, m, ops...)
}

func (g *tx[M, I]) Delete(ctx context.Context, ID I, ops ...DeleteOption) error {
	return Delete[M](ctx, g.itx, ID, IDField, ops...)
}

func (g *tx[M, I]) BulkDelete(ctx context.Context, IDs []I, ops ...DeleteOption) error {
	return BulkDelete[M, I](ctx, g.itx, IDs, IDField, ops...)
}

func (g *tx[M, I]) Get(ctx context.Context, ID I, ops ...SelectOption) (*M, error) {
	return Get[M, I](ctx, g.itx, ID, IDField, ops...)
}

func (g *tx[M, I]) One(ctx context.Context, ops ...SelectOption) (*M, error) {
	return One[M](ctx, g.itx, ops...)
}

func (g *tx[M, I]) All(ctx context.Context, ops ...SelectOption) ([]*M, error) {
	return All[M](ctx, g.itx, ops...)
}

func (g *tx[M, I]) Count(ctx context.Context, ops ...SelectOption) (int, error) {
	return Count[M](ctx, g.itx, ops...)
}

func (g *tx[M, I]) Exists(ctx context.Context, ops ...SelectOption) (bool, error) {
	return Exists[M](ctx, g.itx, ops...)
}
