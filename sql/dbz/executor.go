package dbz

import (
	"context"
	"database/sql"

	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/scan"
)

type ExecWrapper func(bob.Executor) bob.Executor

type ErrorFunc func(e error) error

var _ bob.Executor = (*errorExecutor)(nil)

type errorExecutor struct {
	exec bob.Executor
	fn   ErrorFunc
}

func (e *errorExecutor) QueryContext(ctx context.Context, query string, args ...any) (scan.Rows, error) {
	rows, err := e.exec.QueryContext(ctx, query, args...)
	if err != nil {
		return rows, e.fn(err)
	}
	return rows, err
}

func (e *errorExecutor) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	rows, err := e.exec.ExecContext(ctx, query, args...)
	if err != nil {
		return rows, e.fn(err)
	}
	return rows, err
}

func ExecOnError(exec bob.Executor, fn ErrorFunc) bob.Executor {
	if fn == nil {
		fn = func(e error) error {
			return e
		}
	}
	e := &errorExecutor{
		exec: exec,
		fn:   fn,
	}
	return e
}
