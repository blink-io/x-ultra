package bun

import (
	"fmt"
	"os"
	"testing"

	"github.com/blink-io/x/sql/scany/dbscan"
	"github.com/qustavo/dotsql"
	"github.com/stretchr/testify/require"
)

func TestPG_Connect_1(t *testing.T) {
	db := getDBWithPG()
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

func TestPG_Connect_2(t *testing.T) {
	fp := "./demo.pg.sql"
	f, err := os.Open(fp)
	require.NoError(t, err)

	dot, err := dotsql.Load(f)
	require.NoError(t, err)

	db := getDBWithPG()

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
