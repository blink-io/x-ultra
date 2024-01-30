package dbk

import (
	"context"

	"github.com/vingarcia/ksql"
)

type DBF interface {
	Delete(ctx context.Context, table ksql.Table, idOrRecord interface{}) error

	Exec(ctx context.Context, query string, params ...interface{}) (ksql.Result, error)

	Insert(ctx context.Context, table ksql.Table, record interface{}) error

	Patch(ctx context.Context, table ksql.Table, record interface{}) error

	Query(ctx context.Context, records interface{}, query string, params ...interface{}) error

	QueryChunks(ctx context.Context, parser ksql.ChunkParser) error

	QueryOne(ctx context.Context, record interface{}, query string, params ...interface{}) error

	Transaction(ctx context.Context, fn func(ksql.Provider) error) error
}
