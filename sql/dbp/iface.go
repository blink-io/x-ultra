package dbp

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/go-gorp/gorp/v3"
)

type DBF interface {
	TraceOn(prefix string, logger gorp.Logger)
	TraceOff()
	WithContext(ctx context.Context) gorp.SqlExecutor
	CreateIndex() error
	AddTable(i interface{}) *gorp.TableMap
	AddTableWithName(i interface{}, name string) *gorp.TableMap
	AddTableWithNameAndSchema(i interface{}, schema string, name string) *gorp.TableMap
	AddTableDynamic(inp gorp.DynamicTable, schema string) *gorp.TableMap
	CreateTables() error
	CreateTablesIfNotExists() error
	DropTable(table interface{}) error
	DropTableIfExists(table interface{}) error
	DropTables() error
	DropTablesIfExists() error
	TruncateTables() error
	Insert(list ...interface{}) error
	Update(list ...interface{}) (int64, error)
	UpdateColumns(filter gorp.ColumnFilter, list ...interface{}) (int64, error)
	Delete(list ...interface{}) (int64, error)
	Get(i interface{}, keys ...interface{}) (interface{}, error)
	Select(i interface{}, query string, args ...interface{}) ([]interface{}, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	SelectInt(query string, args ...interface{}) (int64, error)
	SelectNullInt(query string, args ...interface{}) (sql.NullInt64, error)
	SelectFloat(query string, args ...interface{}) (float64, error)
	SelectNullFloat(query string, args ...interface{}) (sql.NullFloat64, error)
	SelectStr(query string, args ...interface{}) (string, error)
	SelectNullStr(query string, args ...interface{}) (sql.NullString, error)
	SelectOne(holder interface{}, query string, args ...interface{}) error
	Begin() (*gorp.Transaction, error)
	TableFor(t reflect.Type, checkPK bool) (*gorp.TableMap, error)
	DynamicTableFor(tableName string, checkPK bool) (*gorp.TableMap, error)
	Prepare(query string) (*sql.Stmt, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(q string, args ...interface{}) (*sql.Rows, error)
}

var _ DBF = (*idb)(nil)
