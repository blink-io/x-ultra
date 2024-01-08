package postgres

import (
	"context"
	"database/sql"
)

func QueryVersion(ctx context.Context, queryRowContext func(ctx context.Context, query string, args ...any) *sql.Row) string {
	row := queryRowContext(ctx, "SELECT version() as version;")
	var str string
	_ = row.Scan(&str)
	return str
}
