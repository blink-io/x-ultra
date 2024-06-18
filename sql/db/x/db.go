package x

import (
	"context"
	"reflect"

	rdb "github.com/blink-io/x/sql/db"
)

const (
	IDField = "id"
)

type (
	// ID defines the generic type for ID in repository
	ID any
	// Model defines the generic type for Model in repository
	Model = any

	IDType = int64

	DBer interface {
		DB() *rdb.DB
	}

	// Replacer replaces raw DB
	Replacer interface {
		Replace(*rdb.DB)
	}

	base[M Model, I ID] interface {
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
		// Exists check
		Exists(context.Context, ...SelectOption) (bool, error)
	}

	DB[M Model, I ID] interface {
		rdb.IDB

		base[M, I]

		// DB .
		DB() rdb.IDB

		// ModelType defines
		ModelType() *M

		// TableType defines
		TableType() *rdb.TableType

		// Tx defines
		Tx() (Tx[M, I], error)
	}

	idb = rdb.IDB

	db[M Model, I ID] struct {
		idb
		mm *M
		tt *rdb.TableType
	}
)

// Do type check
var _ DB[Model, IDType] = (*db[Model, IDType])(nil)

func New[M Model, I ID](idb rdb.IDB) DB[M, I] {
	mm := (*M)(nil)
	idb.RegisterModel(mm)
	tt := idb.Table(reflect.TypeOf(mm))
	return &db[M, I]{idb: idb, mm: mm, tt: tt}
}

func (x *db[M, I]) Insert(ctx context.Context, m *M, ops ...InsertOption) error {
	return Insert[M](ctx, x.idb, m, ops...)
}

func (x *db[M, I]) BulkInsert(ctx context.Context, ms []*M, ops ...InsertOption) error {
	return BulkInsert[M](ctx, x.idb, ms, ops...)
}

func (x *db[M, I]) Update(ctx context.Context, m *M, ops ...UpdateOption) error {
	return Update[M](ctx, x.idb, m, ops...)
}

func (x *db[M, I]) Delete(ctx context.Context, ID I, ops ...DeleteOption) error {
	return Delete[M](ctx, x.idb, ID, IDField, ops...)
}

func (x *db[M, I]) BulkDelete(ctx context.Context, IDs []I, ops ...DeleteOption) error {
	return BulkDelete[M, I](ctx, x.idb, IDs, IDField, ops...)
}

func (x *db[M, I]) Get(ctx context.Context, ID I, ops ...SelectOption) (*M, error) {
	return Get[M, I](ctx, x.idb, ID, IDField, ops...)
}

func (x *db[M, I]) One(ctx context.Context, ops ...SelectOption) (*M, error) {
	return One[M](ctx, x.idb, ops...)
}

func (x *db[M, I]) All(ctx context.Context, ops ...SelectOption) ([]*M, error) {
	return All[M](ctx, x.idb, ops...)
}

func (x *db[M, I]) Count(ctx context.Context, ops ...SelectOption) (int, error) {
	return Count[M](ctx, x.idb, ops...)
}

func (x *db[M, I]) Exists(ctx context.Context, ops ...SelectOption) (bool, error) {
	return Exists[M](ctx, x.idb, ops...)
}

func (x *db[M, I]) Tx() (Tx[M, I], error) {
	return NewTxWithDB[M, I](x.idb)
}

func (x *db[M, I]) DB() rdb.IDB {
	return x.idb
}

func (x *db[M, I]) ModelType() *M {
	return x.mm
}

func (x *db[M, I]) TableType() *rdb.TableType {
	return x.tt
}
