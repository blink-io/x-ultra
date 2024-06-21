package x

import (
	"context"
	"database/sql"
	"reflect"

	xbun "github.com/blink-io/x/bun"
)

const (
	IDField = "id"
)

type (
	// IDType defines the generic type for ID in repository
	IDType interface {
		string |
			~int | ~uint |
			~int8 | ~uint8 |
			~int16 | ~uint16 |
			~int32 | ~uint32 |
			~int64 | ~uint64
	}

	// ModelType defines the generic type for ModelType in repository
	ModelType = any

	DBer interface {
		DB() *xbun.DB
	}

	// Replacer replaces raw DB
	Replacer interface {
		Replace(*xbun.DB)
	}

	base[M ModelType, I IDType] interface {
		// Insert a new record.
		Insert(context.Context, *M, ...InsertOption) error
		// BulkInsert inserts more than one record
		BulkInsert(context.Context, ModelSlice[M], ...InsertOption) error
		// Update a record by ID
		Update(context.Context, *M, ...UpdateOption) error
		// Delete a record by ID
		Delete(context.Context, I, ...DeleteOption) error
		// BulkDelete deletes by IDs
		BulkDelete(context.Context, IDSlice[I], ...DeleteOption) error
		// Get a record by ID
		Get(context.Context, I, ...SelectOption) (*M, error)
		// One get one record by criteria
		One(context.Context, ...SelectOption) (*M, error)
		// All fetch all data from repository
		All(context.Context, ...SelectOption) (ModelSlice[M], error)
		// Count rows
		Count(context.Context, ...SelectOption) (int, error)
		// Exists check
		Exists(context.Context, ...SelectOption) (bool, error)
	}

	DB[M ModelType, I IDType] interface {
		xbun.IDB

		base[M, I]

		// DB .
		DB() xbun.IDB

		// ModelType defines
		ModelType() *M

		// TableType defines
		TableType() *xbun.TableType

		// Tx defines
		Tx(context.Context, *sql.TxOptions) (Tx[M, I], error)
	}

	idb = xbun.IDB

	db[M ModelType, I IDType] struct {
		idb
		mm *M
		tt *xbun.TableType
	}
)

// Do type check
var _ DB[xbun.BaseModel, int] = (*db[xbun.BaseModel, int])(nil)

func NewDB[M ModelType, I IDType](idb xbun.IDB) DB[M, I] {
	mm := (*M)(nil)
	idb.RegisterModel(mm)
	tt := idb.Table(reflect.TypeOf(mm))
	return &db[M, I]{idb: idb, mm: mm, tt: tt}
}

func (x *db[M, I]) Insert(ctx context.Context, m *M, ops ...InsertOption) error {
	return Insert[M](ctx, x.idb, m, ops...)
}

func (x *db[M, I]) BulkInsert(ctx context.Context, ms ModelSlice[M], ops ...InsertOption) error {
	return BulkInsert[M](ctx, x.idb, ms, ops...)
}

func (x *db[M, I]) Update(ctx context.Context, m *M, ops ...UpdateOption) error {
	return Update[M](ctx, x.idb, m, ops...)
}

func (x *db[M, I]) Delete(ctx context.Context, ID I, ops ...DeleteOption) error {
	return Delete[M](ctx, x.idb, ID, IDField, ops...)
}

func (x *db[M, I]) BulkDelete(ctx context.Context, IDs IDSlice[I], ops ...DeleteOption) error {
	return BulkDelete[M, I](ctx, x.idb, IDs, IDField, ops...)
}

func (x *db[M, I]) Get(ctx context.Context, ID I, ops ...SelectOption) (*M, error) {
	return Get[M, I](ctx, x.idb, ID, IDField, ops...)
}

func (x *db[M, I]) One(ctx context.Context, ops ...SelectOption) (*M, error) {
	return One[M](ctx, x.idb, ops...)
}

func (x *db[M, I]) All(ctx context.Context, ops ...SelectOption) (ModelSlice[M], error) {
	return All[M](ctx, x.idb, ops...)
}

func (x *db[M, I]) Count(ctx context.Context, ops ...SelectOption) (int, error) {
	return Count[M](ctx, x.idb, ops...)
}

func (x *db[M, I]) Exists(ctx context.Context, ops ...SelectOption) (bool, error) {
	return Exists[M](ctx, x.idb, ops...)
}

func (x *db[M, I]) Tx(ctx context.Context, opts *sql.TxOptions) (Tx[M, I], error) {
	return NewTxWithDB[M, I](ctx, x.idb, opts)
}

func (x *db[M, I]) DB() xbun.IDB {
	return x.idb
}

func (x *db[M, I]) ModelType() *M {
	return x.mm
}

func (x *db[M, I]) TableType() *xbun.TableType {
	return x.tt
}
