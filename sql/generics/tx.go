package generics

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type (
	itx = bun.Tx

	Tx[M Model, I ID] interface {
		base[M, I]
		// Commit the transaction
		Commit() error
		// Rollback the transaction
		Rollback() error
		// RunInTx runs the function in a transaction. If the function returns an error,
		// the transaction is rolled back. Otherwise, the transaction is committed.
		RunInTx(context.Context, *sql.TxOptions, func(context.Context, bun.Tx) error) error
	}

	tx[M Model, I ID] struct {
		itx
	}
)

var _ Tx[Model, ID] = (*tx[Model, ID])(nil)

func NewTx[M Model, I ID](itx bun.Tx) Tx[M, I] {
	return &tx[M, I]{itx}
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
	return Delete[M](ctx, g.itx, ID, ops...)
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
