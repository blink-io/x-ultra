package kpgx5

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/sqldialect"
)

// NewFromSQLDB builds a ksql.DB from a *sql.DB instance
func NewFromSQLDB(db *sql.DB) (ksql.DB, error) {
	return ksql.NewWithAdapter(NewSQLAdapter(db), sqldialect.PostgresDialect{})
}

// NewFromPgxPool builds a ksql.DB from a *pgxpool.Pool instance
func NewFromPgxPool(pool *pgxpool.Pool) (db ksql.DB, err error) {
	return ksql.NewWithAdapter(NewPGXAdapter(pool), sqldialect.PostgresDialect{})
}

// New instantiates a new ksql.Client using pgx as the backend driver
func New(
	ctx context.Context,
	connectionString string,
	config ksql.Config,
) (db ksql.DB, err error) {
	config.SetDefaultValues()

	pgxConf, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return ksql.DB{}, err
	}

	pgxConf.MaxConns = int32(config.MaxOpenConns)

	pool, err := pgxpool.NewWithConfig(ctx, pgxConf)
	if err != nil {
		return ksql.DB{}, err
	}
	if err = pool.Ping(ctx); err != nil {
		return ksql.DB{}, err
	}

	return ksql.NewWithAdapter(NewPGXAdapter(pool), sqldialect.PostgresDialect{})
}
