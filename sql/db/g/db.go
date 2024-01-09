package g

import (
	"context"
	"reflect"

	xdb "github.com/blink-io/x/sql/db"
	"github.com/uptrace/bun/schema"
)

const (
	IDField = "id"
)

type (
	// ID defines the generic type for ID in repository
	ID interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
			~string
	}
	// Model defines the generic type for Model in repository
	Model = any

	IDType = string

	DBer interface {
		DB() *xdb.DB
	}
	Replacer interface {
		Replace(*xdb.DB)
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
		// Exists has record
		Exists(context.Context, ...SelectOption) (bool, error)
	}

	DB[M Model, I ID] interface {
		xdb.IDB
		base[M, I]
		// DB .
		DB() *xdb.DB
		// ModelType defines
		ModelType() *M
		// TableType defines
		TableType() *schema.Table
		// Tx defines
		Tx() (Tx[M, I], error)
	}

	idb = xdb.DB

	db[M Model, I ID] struct {
		*idb
		mm *M
		tt *schema.Table
	}
)

// Do type check
var _ DB[Model, IDType] = (*db[Model, IDType])(nil)

func NewDB[M Model, I ID](idb *xdb.DB) DB[M, I] {
	mm := (*M)(nil)
	idb.RegisterModel(mm)
	tt := idb.Table(reflect.TypeOf(mm))
	return &db[M, I]{idb: idb, mm: mm, tt: tt}
}

func (g *db[M, I]) Insert(ctx context.Context, m *M, ops ...InsertOption) error {
	return Insert[M](ctx, g.idb, m, ops...)
}

func (g *db[M, I]) BulkInsert(ctx context.Context, ms []*M, ops ...InsertOption) error {
	return BulkInsert[M](ctx, g.idb, ms, ops...)
}

func (g *db[M, I]) Update(ctx context.Context, m *M, ops ...UpdateOption) error {
	return Update[M](ctx, g.idb, m, ops...)
}

func (g *db[M, I]) Delete(ctx context.Context, ID I, ops ...DeleteOption) error {
	return Delete[M](ctx, g.idb, ID, IDField, ops...)
}

func (g *db[M, I]) BulkDelete(ctx context.Context, IDs []I, ops ...DeleteOption) error {
	return BulkDelete[M, I](ctx, g.idb, IDs, IDField, ops...)
}

func (g *db[M, I]) Get(ctx context.Context, ID I, ops ...SelectOption) (*M, error) {
	return Get[M, I](ctx, g.idb, ID, IDField, ops...)
}

func (g *db[M, I]) One(ctx context.Context, ops ...SelectOption) (*M, error) {
	return One[M](ctx, g.idb, ops...)
}

func (g *db[M, I]) All(ctx context.Context, ops ...SelectOption) ([]*M, error) {
	return All[M](ctx, g.idb, ops...)
}

func (g *db[M, I]) Count(ctx context.Context, ops ...SelectOption) (int, error) {
	return Count[M](ctx, g.idb, ops...)
}

func (g *db[M, I]) Exists(ctx context.Context, ops ...SelectOption) (bool, error) {
	return Exists[M](ctx, g.idb, ops...)
}

func (g *db[M, I]) Tx() (Tx[M, I], error) {
	return NewTxWithDB[M, I](g.idb)
}

func (g *db[M, I]) DB() *xdb.DB {
	return g.idb
}

func (g *db[M, I]) ModelType() *M {
	return g.mm
}

func (g *db[M, I]) TableType() *schema.Table {
	return g.tt
}
