package repository

import (
	"context"
	"reflect"

	"github.com/blink-io/x/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type (
	// ID defines the generic type for ID in repository
	ID = any
	// Model defines the generic type for Model in repository
	Model = any

	DB[M Model, I ID] interface {
		Tx[M, I]
		// DB .
		DB() *sql.DB
		// Model defines
		Model() *M
		// Table defines
		Table() *schema.Table
		// Tx defines
		Tx() (Tx[M, I], error)
		// TxWith defines
		TxWith(tx bun.Tx) Tx[M, I]
	}

	db[M Model, I ID] struct {
		mt *M
		tb *schema.Table
		rd *sql.DB
	}
)

// Do type check
var _ DB[Model, ID] = (*db[Model, ID])(nil)

func NewDB[M Model, I ID](rd *sql.DB) DB[M, I] {
	mt := (*M)(nil)
	rd.RegisterModel(mt)
	tb := rd.Table(reflect.TypeOf(mt))
	return &db[M, I]{rd: rd, mt: mt, tb: tb}
}

func (g *db[M, I]) Insert(ctx context.Context, m *M, ops ...InsertOption) error {
	return Insert[M](ctx, g.rd, m, ops...)
}

func (g *db[M, I]) BulkInsert(ctx context.Context, ms []*M, ops ...InsertOption) error {
	return BulkInsert[M](ctx, g.rd, ms, ops...)
}

func (g *db[M, I]) Update(ctx context.Context, m *M, ops ...UpdateOption) error {
	return Update[M](ctx, g.rd, m, ops...)
}

func (g *db[M, I]) Delete(ctx context.Context, ID I, ops ...DeleteOption) error {
	return Delete[M](ctx, g.rd, ID, ops...)
}

func (g *db[M, I]) BulkDelete(ctx context.Context, IDs []I, ops ...DeleteOption) error {
	return BulkDelete[M, I](ctx, g.rd, IDs, ops...)
}

func (g *db[M, I]) Get(ctx context.Context, ID I, ops ...SelectOption) (*M, error) {
	return Get[M, I](ctx, g.rd, ID, ops...)
}

func (g *db[M, I]) One(ctx context.Context, ops ...SelectOption) (*M, error) {
	return One[M](ctx, g.rd, ops...)
}

func (g *db[M, I]) All(ctx context.Context, ops ...SelectOption) ([]*M, error) {
	return All[M](ctx, g.rd, ops...)
}

func (g *db[M, I]) Count(ctx context.Context, ops ...SelectOption) (int, error) {
	return Count[M](ctx, g.rd, ops...)
}

func (g *db[M, I]) Exists(ctx context.Context, ops ...SelectOption) (bool, error) {
	return Exists[M](ctx, g.rd, ops...)
}

func (g *db[M, I]) Tx() (Tx[M, I], error) {
	tx, err := g.rd.Begin()
	if err != nil {
		return nil, err
	}
	g.rd.Begin()
	txg := NewTx[M, I](tx)
	return txg, nil
}

func (g *db[M, I]) TxWith(tx bun.Tx) Tx[M, I] {
	txg := NewTx[M, I](tx)
	return txg
}

func (g *db[M, I]) DB() *sql.DB {
	return g.rd
}

func (g *db[M, I]) Model() *M {
	return g.mt
}

func (g *db[M, I]) Table() *schema.Table {
	return g.tb
}
