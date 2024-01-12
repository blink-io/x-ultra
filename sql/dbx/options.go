package dbx

import (
	"context"
	"database/sql"
	"time"

	"github.com/pocketbase/dbx"
)

type options struct {
	dbTag string

	// QueryLogFunc is called each time when performing a SQL query that returns data.
	queryLogFunc dbx.QueryLogFunc

	// ExecLogFunc is called each time when a SQL statement is executed.
	// The "t" parameter gives the time that the SQL statement takes to execute,
	// while result and err refer to the result of the execution.
	execLogFunc func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error)
}

type Option func(*options)

func applyOptions(ops ...Option) *options {
	opts := new(options)
	for _, o := range ops {
		o(opts)
	}
	return opts
}

func WithDbTag(dbTag string) Option {
	return func(o *options) {
		o.dbTag = dbTag
	}
}

func WithQueryLogFunc(f dbx.QueryLogFunc) Option {
	return func(o *options) {
		o.queryLogFunc = f
	}
}

func WithExecLogFunc(f dbx.ExecLogFunc) Option {
	return func(o *options) {
		o.execLogFunc = f
	}
}
