package pg

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/sql/scany/dbscan"
	"github.com/stretchr/testify/require"
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
