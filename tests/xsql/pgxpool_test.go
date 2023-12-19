package bun

import (
	"context"
	"fmt"
	"testing"

	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/scany/pgxscan"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestPxgPool_1(t *testing.T) {
	ctx := context.Background()
	cc := xsql.ToPGXConfig(pgOpt)
	cfg, err := pgxpool.ParseConfig("")
	require.NoError(t, err)

	cfg.ConnConfig = cc

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	conn, err := pool.Acquire(ctx)
	require.NoError(t, err)

	rows, err := conn.Query(ctx, "select version();")

	var str string
	require.NoError(t, pgxscan.ScanOne(&str, rows))

	fmt.Println("DB Version: ", str)
}
