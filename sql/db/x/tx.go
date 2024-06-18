package x

import (
	"context"

	rdb "github.com/blink-io/x/sql/db"
)

type (
	rtx = rdb.RawTx

	Tx[M Model, I ID] interface {
		rdb.RawIDB

		base[M, I]
		// Commit the transaction
		Commit() error
		// Rollback the transaction
		Rollback() error
	}

	tx[M Model, I ID] struct {
		rtx
	}
)

var _ Tx[Model, IDType] = (*tx[Model, IDType])(nil)

func NewTx[M Model, I ID](itx rdb.RawTx) Tx[M, I] {
	return &tx[M, I]{itx}
}

func NewTxWithDB[M Model, I ID](db rdb.IDB) (Tx[M, I], error) {
	itx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return NewTx[M, I](itx), nil
}

func (x *tx[M, I]) Insert(ctx context.Context, m *M, ops ...InsertOption) error {
	return Insert[M](ctx, x.rtx, m, ops...)
}

func (x *tx[M, I]) BulkInsert(ctx context.Context, ms []*M, ops ...InsertOption) error {
	return BulkInsert[M](ctx, x.rtx, ms, ops...)
}

func (x *tx[M, I]) Update(ctx context.Context, m *M, ops ...UpdateOption) error {
	return Update[M](ctx, x.rtx, m, ops...)
}

func (x *tx[M, I]) Delete(ctx context.Context, ID I, ops ...DeleteOption) error {
	return Delete[M](ctx, x.rtx, ID, IDField, ops...)
}

func (x *tx[M, I]) BulkDelete(ctx context.Context, IDs []I, ops ...DeleteOption) error {
	return BulkDelete[M, I](ctx, x.rtx, IDs, IDField, ops...)
}

func (x *tx[M, I]) Get(ctx context.Context, ID I, ops ...SelectOption) (*M, error) {
	return Get[M, I](ctx, x.rtx, ID, IDField, ops...)
}

func (x *tx[M, I]) One(ctx context.Context, ops ...SelectOption) (*M, error) {
	return One[M](ctx, x.rtx, ops...)
}

func (x *tx[M, I]) All(ctx context.Context, ops ...SelectOption) ([]*M, error) {
	return All[M](ctx, x.rtx, ops...)
}

func (x *tx[M, I]) Count(ctx context.Context, ops ...SelectOption) (int, error) {
	return Count[M](ctx, x.rtx, ops...)
}

func (x *tx[M, I]) Exists(ctx context.Context, ops ...SelectOption) (bool, error) {
	return Exists[M](ctx, x.rtx, ops...)
}
