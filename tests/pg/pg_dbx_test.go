package pg

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/postgres"
	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/scany/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestPg_DBX_Insert_1(t *testing.T) {
	db := getPgDBX()
	r := newRandomRecordForApp("dbx")

	err := db.Model(r).Insert()

	require.NoError(t, err)
}

func TestPg_PGX_Pool_1(t *testing.T) {
	cc, _ := xsql.ToPGXConfig(pgCfg())
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

func TestPg_DBX_Select_Funcs(t *testing.T) {
	db := getPgDBX()

	sqlF := "select %s as payload"
	funcs := getPgFuncsMap()

	for _, fn := range funcs {
		ss := fmt.Sprintf(sqlF, fn)
		q := db.NewQuery(ss)
		var v string
		require.NoError(t, q.Row(&v))
		fmt.Println("SQLite func payload:  ", v)
	}
}

func TestPg_DBX_Select_Version(t *testing.T) {
	db := getPgDBX()

	ver := postgres.QueryVersion(ctx, db.DB().QueryRowContext)
	fmt.Println("Version: ", ver)
}
