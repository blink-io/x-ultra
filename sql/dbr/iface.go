package dbr

import (
	"context"
	"database/sql"
	"time"

	"github.com/gocraft/dbr/v2"
)

type ISession interface {
	DeleteFrom(table string) *dbr.DeleteStmt
	DeleteBySql(query string, value ...interface{}) *dbr.DeleteStmt
	InsertInto(table string) *dbr.InsertStmt
	InsertBySql(query string, value ...interface{}) *dbr.InsertStmt
	Select(column ...string) *dbr.SelectStmt
	SelectBySql(query string, value ...interface{}) *dbr.SelectStmt
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*dbr.Tx, error)
	Begin() (*dbr.Tx, error)
	GetTimeout() time.Duration
	Update(table string) *dbr.UpdateStmt
	UpdateBySql(query string, value ...interface{}) *dbr.UpdateStmt
}

var _ ISession = (*idb)(nil)
