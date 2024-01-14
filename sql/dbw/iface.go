package dbw

import (
	"context"
	"database/sql"

	"github.com/ilibs/gosql/v2"
	"github.com/jmoiron/sqlx"
)

type IDB interface {
	DriverName() string
	ShowSql() *gosql.DB
	Begin() (*gosql.DB, error)
	Commit() error
	Rollback() error
	Rebind(query string) string
	Preparex(query string) (*sqlx.Stmt, error)
	Exec(query string, args ...interface{}) (result sql.Result, err error)
	NamedExec(query string, args interface{}) (result sql.Result, err error)
	Queryx(query string, args ...interface{}) (rows *sqlx.Rows, err error)
	QueryRowx(query string, args ...interface{}) (rows *sqlx.Row)
	Get(dest interface{}, query string, args ...interface{}) (err error)
	Select(dest interface{}, query string, args ...interface{}) (err error)
	Txx(ctx context.Context, fn func(ctx context.Context, tx *gosql.DB) error) (err error)
	Tx(fn func(w *gosql.DB) error) (err error)
	Table(t string) *gosql.Mapper
	Model(m interface{}) *gosql.Builder
	WithContext(ctx context.Context) *gosql.Builder
	Relation(name string, fn gosql.BuilderChainFunc) *gosql.DB
}

var _ IDB = (*idb)(nil)
