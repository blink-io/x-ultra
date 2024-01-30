package dbz

import (
	"github.com/stephenafamo/bob"
)

//type DBF interface {
//	BeginTx(ctx context.Context, opts *sql.TxOptions) (bob.Tx, error)
//
//	PrepareContext(ctx context.Context, query string) (bob.Statement, error)
//
//	bob.Executor
//}

type DBF = bob.Executor
