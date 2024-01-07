package pg

import (
	"context"
	"fmt"
	"os"
	"testing"

	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/g"
	"github.com/blink-io/x/sql/scany/dbscan"
	"github.com/blink-io/x/sql/scany/pgxscan"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-rel/rel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qustavo/dotsql"
	"github.com/stretchr/testify/require"
)

var (
	modelApp = (*Application)(nil)
)

func TestPg_DBScan_1(t *testing.T) {
	db := getPgDB()
	row := db.QueryRow("select version();")

	var vinfo string
	err := row.Scan(&vinfo)
	require.NoError(t, err)

	fmt.Println("DB Version: ", vinfo)

	rows, err := db.Query("select now();")
	require.NoError(t, err)

	var nowstr string
	require.NoError(t, dbscan.ScanOne(&nowstr, rows))
	fmt.Println("DB Now: ", nowstr)
}

func TestPg_DotSQL_1(t *testing.T) {
	fp := "./demo.pg.sql"
	f, err := os.Open(fp)
	require.NoError(t, err)

	dot, err := dotsql.Load(f)
	require.NoError(t, err)

	db := getPgDB()

	rows1, err := dot.Query(db, "get-db-version")
	require.NoError(t, err)

	var dbver string
	require.NoError(t, dbscan.ScanOne(&dbver, rows1))

	fmt.Println("DB Version: ", dbver)

	rows2, err := dot.Query(db, "get-db-tsz")
	require.NoError(t, err)

	var dbtsz string
	require.NoError(t, dbscan.ScanOne(&dbtsz, rows2))
	fmt.Println("DB TSZ: ", dbtsz)

	rows3, err := dot.Query(db, "get-db-detail")
	require.NoError(t, err)

	var dbd = make(map[string]any)
	require.NoError(t, dbscan.ScanOne(&dbd, rows3))
	fmt.Println("DB Detail: ", dbd)
}

func TestPg_DB_CreateTable_1(t *testing.T) {
	db := getPgDB()

	_, err := db.NewCreateTable().
		IfNotExists().
		Model(modelApp).Exec(ctx)
	require.NoError(t, err)
}

func TestPg_DB_DropTable_1(t *testing.T) {
	db := getPgDB()

	_, err := db.NewDropTable().IfExists().
		Model(modelApp).Exec(ctx)
	require.NoError(t, err)
}

func TestPg_DB_Insert_1(t *testing.T) {
	db := getPgDB()
	r := newRandomRecordForApp("bun")
	txdb, err := g.NewDB[Application, string](db).Tx()
	require.NoError(t, err)

	err1 := txdb.BulkInsert(ctx, []*Application{r})
	if err1 != nil {
		require.NoError(t, txdb.Rollback())
	} else {
		require.NoError(t, txdb.Commit())
	}
}

func TestPg_DBQ_Insert_1(t *testing.T) {
	db := getPgDBQ()
	r := newRandomRecordForApp("dbq")
	ds := db.From(r.TableName())
	rt, err := ds.Insert().Rows(r).Executor().Exec()

	fmt.Println("SQL Result: ", rt)

	require.NoError(t, err)
}

func TestPg_DBQ_Select_1(t *testing.T) {
	db := getPgDBQ()
	ds := db.From("applications")
	sql, _, err := ds.Select("name", "code", "type").
		Where(goqu.C("code").Like("code-D%")).ToSQL()
	require.NoError(t, err)

	fmt.Println("SQL Generated: ", sql)
}

func TestPg_DBX_Insert_1(t *testing.T) {
	db := getPgDBX()
	r := newRandomRecordForApp("dbx")

	err := db.Model(r).Insert()

	require.NoError(t, err)
}

func TestPg_PGX_Pool_1(t *testing.T) {
	ctx := context.Background()
	cc := xsql.ToPGXConfig(pgCfg())
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

func TestPg_DBR_Select_Funcs(t *testing.T) {
	db := getPgDBR()

	sess := db.NewSession(nil)

	sqlF := "select %s as payload"
	funcs := getPgFuncsMap()

	for k, fstr := range funcs {
		ss := fmt.Sprintf(sqlF, fstr)
		var v string
		r := sess.QueryRow(ss)
		require.NoError(t, r.Scan(&v))
		fmt.Println(k, " => ", v)
	}
}

func TestPg_DBM_Select_Funcs(t *testing.T) {
	db := getPgDBM()

	type Result struct {
		Payload string `db:"payload"`
	}

	sqlF := "select %s as payload"
	funcs := getPgFuncsMap()
	rt := new(Result)
	for k, v := range funcs {
		ss := rel.SQL(fmt.Sprintf(sqlF, v))
		require.NoError(t, db.Find(ctx, rt, ss))
		fmt.Println(k, "=>", rt.Payload)
	}
}
