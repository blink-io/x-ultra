package dbq

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type DBF interface {
	Dialect() string
	Begin() (*goqu.TxDatabase, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*goqu.TxDatabase, error)
	WithTx(fn func(*goqu.TxDatabase) error) error
	From(from ...interface{}) *goqu.SelectDataset
	Select(cols ...interface{}) *goqu.SelectDataset
	Update(table interface{}) *goqu.UpdateDataset
	Insert(table interface{}) *goqu.InsertDataset
	Delete(table interface{}) *goqu.DeleteDataset
	Truncate(table ...interface{}) *goqu.TruncateDataset
	Logger(logger goqu.Logger)
	Trace(op, sqlString string, args ...interface{})
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ScanStructs(i interface{}, query string, args ...interface{}) error
	ScanStructsContext(ctx context.Context, i interface{}, query string, args ...interface{}) error
	ScanStruct(i interface{}, query string, args ...interface{}) (bool, error)
	ScanStructContext(ctx context.Context, i interface{}, query string, args ...interface{}) (bool, error)
	ScanVals(i interface{}, query string, args ...interface{}) error
	ScanValsContext(ctx context.Context, i interface{}, query string, args ...interface{}) error
	ScanVal(i interface{}, query string, args ...interface{}) (bool, error)
	ScanValContext(ctx context.Context, i interface{}, query string, args ...interface{}) (bool, error)
}

var _ DBF = (*idb)(nil)
