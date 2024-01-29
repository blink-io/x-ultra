package dbx

import (
	"context"
	"database/sql"

	"github.com/pocketbase/dbx"
)

type DBF interface {
	Clone() *dbx.DB
	WithContext(ctx context.Context) *dbx.DB
	Context() context.Context
	DB() *sql.DB
	Close() error
	Begin() (*dbx.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*dbx.Tx, error)
	Wrap(sqlTx *sql.Tx) *dbx.Tx
	Transactional(f func(*dbx.Tx) error) (err error)
	TransactionalContext(ctx context.Context, opts *sql.TxOptions, f func(*dbx.Tx) error) (err error)
	DriverName() string
	QuoteTableName(s string) string
	QuoteColumnName(s string) string
}

var _ DBF = (*idb)(nil)
