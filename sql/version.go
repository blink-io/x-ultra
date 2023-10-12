package sql

import (
	"context"
	"database/sql"
)

func MySQLVersion(ctx context.Context, db *sql.DB) string {
	return doQueryOne(ctx, db, "SELECT version() as version;")
}

func PostgresVersion(ctx context.Context, db *sql.DB) string {
	return doQueryOne(ctx, db, "SELECT version() as version;")
}

func SQLiteVersion(ctx context.Context, db *sql.DB) string {
	return doQueryOne(ctx, db, "SELECT SQLITE_VERSION() as version;")
}

func doQueryOne(ctx context.Context, db *sql.DB, query string) string {
	row := db.QueryRowContext(ctx, query)
	var s string
	_ = row.Scan(&s)
	return s
}
