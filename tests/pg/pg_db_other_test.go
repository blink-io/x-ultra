package pg

import (
	"fmt"
	"os"
	"testing"

	"github.com/blink-io/x/sql/scany/dbscan"
	"github.com/qustavo/dotsql"
	"github.com/stretchr/testify/require"
)

func TestCompatible(t *testing.T) {
	//var _ xsql.IDB = (*pgxpool.Pool)(nil)
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
