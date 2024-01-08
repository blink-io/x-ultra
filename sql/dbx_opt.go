package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/pocketbase/dbx"
)

type dbxOptions struct {
	// QueryLogFunc is called each time when performing a SQL query that returns data.
	queryLogFunc dbx.QueryLogFunc

	// ExecLogFunc is called each time when a SQL statement is executed.
	// The "t" parameter gives the time that the SQL statement takes to execute,
	// while result and err refer to the result of the execution.
	execLogFunc func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error)
}

type DBXOption func(*dbxOptions)

func applyDBXOptions(ops ...DBXOption) *dbxOptions {
	opts := new(dbxOptions)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func DBXQueryLogFunc(f dbx.QueryLogFunc) DBXOption {
	return func(o *dbxOptions) {
		o.queryLogFunc = f
	}
}

func DBXExecLogFunc(f dbx.ExecLogFunc) DBXOption {
	return func(o *dbxOptions) {
		o.execLogFunc = f
	}
}
