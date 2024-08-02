package bun

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/gofrs/uuid/v5"
	guuid "github.com/google/uuid"
)

const (
	IDField = "id"
)

type (
	// IDType defines the generic type for I in repository
	IDType interface {
		~string | uuid.UUID | guuid.UUID |
			~int | ~uint |
			~int8 | ~uint8 |
			~int16 | ~uint16 |
			~int32 | ~uint32 |
			~int64 | ~uint64
	}

	// ModelType defines the generic type for ModelType in repository
	ModelType = any

	GenericBase[M ModelType, I IDType] interface {
		// Insert a new record.
		Insert(context.Context, *M, ...DoInsertOption) error
		// BulkInsert inserts more than one record
		BulkInsert(context.Context, ModelSlice[M], ...DoInsertOption) error
		// Update a record by I
		Update(context.Context, *M, ...DoUpdateOption) error
		// Delete a record by I
		Delete(context.Context, I, ...DoDeleteOption) error
		// BulkDelete deletes by IDs
		BulkDelete(context.Context, IDSlice[I], ...DoDeleteOption) error
		// Get a record by I
		Get(context.Context, I, ...DoSelectOption) (*M, error)
		// One get one record by criteria
		One(context.Context, ...DoSelectOption) (*M, error)
		// All fetch all data from repository
		All(context.Context, ...DoSelectOption) (ModelSlice[M], error)
		// Count rows
		Count(context.Context, ...DoSelectOption) (int, error)
		// Exists check
		Exists(context.Context, ...DoSelectOption) (bool, error)
	}

	Generic[M ModelType, I IDType] interface {
		IDB

		GenericBase[M, I]

		// DB .
		DB() IDB

		// ModelType defines
		ModelType() *M

		// TableType defines
		TableType() *TableType

		// Tx defines
		Tx(context.Context, *sql.TxOptions) (GenericTx[M, I], error)
	}

	idb = IDB

	generic[M ModelType, I IDType] struct {
		idb
		mm *M
		tt *TableType
	}
)

// Do type check
var _ Generic[BaseModel, int] = (*generic[BaseModel, int])(nil)

func NewGenericDB[M ModelType, I IDType](idb IDB) Generic[M, I] {
	mm := (*M)(nil)
	idb.RegisterModel(mm)
	tt := idb.Table(reflect.TypeOf(mm))
	return &generic[M, I]{idb: idb, mm: mm, tt: tt}
}

func (g *generic[M, I]) Insert(ctx context.Context, m *M, ops ...DoInsertOption) error {
	return DoInsert[M](ctx, g.idb, m, ops...)
}

func (g *generic[M, I]) BulkInsert(ctx context.Context, ms ModelSlice[M], ops ...DoInsertOption) error {
	return DoBulkInsert[M](ctx, g.idb, ms, ops...)
}

func (g *generic[M, I]) Update(ctx context.Context, m *M, ops ...DoUpdateOption) error {
	return DoUpdate[M](ctx, g.idb, m, ops...)
}

func (g *generic[M, I]) Delete(ctx context.Context, ID I, ops ...DoDeleteOption) error {
	return DoDelete[M](ctx, g.idb, ID, IDField, ops...)
}

func (g *generic[M, I]) BulkDelete(ctx context.Context, IDs IDSlice[I], ops ...DoDeleteOption) error {
	return DoBulkDelete[M, I](ctx, g.idb, IDs, IDField, ops...)
}

func (g *generic[M, I]) Get(ctx context.Context, ID I, ops ...DoSelectOption) (*M, error) {
	return DoGet[M, I](ctx, g.idb, ID, IDField, ops...)
}

func (g *generic[M, I]) One(ctx context.Context, ops ...DoSelectOption) (*M, error) {
	return DoOne[M](ctx, g.idb, ops...)
}

func (g *generic[M, I]) All(ctx context.Context, ops ...DoSelectOption) (ModelSlice[M], error) {
	return DoAll[M](ctx, g.idb, ops...)
}

func (g *generic[M, I]) Count(ctx context.Context, ops ...DoSelectOption) (int, error) {
	return DoCount[M](ctx, g.idb, ops...)
}

func (g *generic[M, I]) Exists(ctx context.Context, ops ...DoSelectOption) (bool, error) {
	return DoExists[M](ctx, g.idb, ops...)
}

func (g *generic[M, I]) Tx(ctx context.Context, opts *sql.TxOptions) (GenericTx[M, I], error) {
	return NewGenericTxWithDB[M, I](ctx, g.idb, opts)
}

func (g *generic[M, I]) DB() IDB {
	return g.idb
}

func (g *generic[M, I]) ModelType() *M {
	return g.mm
}

func (g *generic[M, I]) TableType() *TableType {
	return g.tt
}
