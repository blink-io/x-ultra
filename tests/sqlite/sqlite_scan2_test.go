package sqlite

import (
	"fmt"
	"testing"

	"github.com/blockloop/scan/v2"
	"github.com/stretchr/testify/require"
)

func TestSqlite_Scan2_Select_1(t *testing.T) {
	db := getSqliteDB()
	rows, err := db.Query("SELECT * FROM applications")
	require.NoError(t, err)

	var apps []Application
	err = scan.Rows(&apps, rows)
	require.NoError(t, err)

	fmt.Println(apps)
}
