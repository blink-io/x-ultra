package bun

import (
	"context"
	"database/sql"
)

type (
	rtx = RawTx

	GenericTx[M ModelType, I IDType] interface {
		RawIDB

		GenericsBase[M, I]

		// Commit the transaction
		Commit() error

		// Rollback the transaction
		Rollback() error
	}

	genericTx[M ModelType, I IDType] struct {
		rtx
	}
)

var _ GenericTx[ModelType, int] = (*genericTx[ModelType, int])(nil)

func NewGenericTx[M ModelType, I IDType](rtx RawTx) GenericTx[M, I] {
	return &genericTx[M, I]{rtx}
}

func NewGenericTxWithDB[M ModelType, I IDType](ctx context.Context, db IDB, opts *sql.TxOptions) (GenericTx[M, I], error) {
	itx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return NewGenericTx[M, I](itx), nil
}

func (x *genericTx[M, I]) Insert(ctx context.Context, m *M, ops ...DoInsertOption) error {
	return DoInsert[M](ctx, x.rtx, m, ops...)
}

func (x *genericTx[M, I]) BulkInsert(ctx context.Context, ms ModelSlice[M], ops ...DoInsertOption) error {
	return DoBulkInsert[M](ctx, x.rtx, ms, ops...)
}

func (x *genericTx[M, I]) Update(ctx context.Context, m *M, ops ...DoUpdateOption) error {
	return DoUpdate[M](ctx, x.rtx, m, ops...)
}

func (x *genericTx[M, I]) Delete(ctx context.Context, ID I, ops ...DoDeleteOption) error {
	return DoDelete[M](ctx, x.rtx, ID, IDField, ops...)
}

func (x *genericTx[M, I]) BulkDelete(ctx context.Context, IDs IDSlice[I], ops ...DoDeleteOption) error {
	return DoBulkDelete[M, I](ctx, x.rtx, IDs, IDField, ops...)
}

func (x *genericTx[M, I]) Get(ctx context.Context, ID I, ops ...DoSelectOption) (*M, error) {
	return DoGet[M, I](ctx, x.rtx, ID, IDField, ops...)
}

func (x *genericTx[M, I]) One(ctx context.Context, ops ...DoSelectOption) (*M, error) {
	return DoOne[M](ctx, x.rtx, ops...)
}

func (x *genericTx[M, I]) All(ctx context.Context, ops ...DoSelectOption) (ModelSlice[M], error) {
	return DoAll[M](ctx, x.rtx, ops...)
}

func (x *genericTx[M, I]) Count(ctx context.Context, ops ...DoSelectOption) (int, error) {
	return DoCount[M](ctx, x.rtx, ops...)
}

func (x *genericTx[M, I]) Exists(ctx context.Context, ops ...DoSelectOption) (bool, error) {
	return DoExists[M](ctx, x.rtx, ops...)
}
