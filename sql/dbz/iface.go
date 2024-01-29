package dbz

import (
	"context"
	"database/sql"

	"github.com/stephenafamo/bob"
)

type DBF interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (bob.Tx, error)

	PrepareContext(ctx context.Context, query string) (bob.Statement, error)

	bob.Executor
}
