package sqlite

import (
	"fmt"
	"testing"

	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/scany/dbscan"
	"github.com/huandu/go-sqlbuilder"
	"github.com/stretchr/testify/require"
)

func TestSqlite_GoSqlBuilder_1(t *testing.T) {
	var appStruct = sqlbuilder.NewStruct(new(Application))

	sb := appStruct.SelectFrom("applications")
	sb.Where(sb.Equal("id", "GY7hPxdSg"))

	// Execute the query.
	sql, args := sb.Build()
	//rows, _ := db.Query(sql, args...)

	fmt.Println("SQL: ", sql)
	fmt.Println("Args: ", args)

	db := getSqliteSqlDB()
	rows, err := db.Query(sql, args...)
	require.NoError(t, err)

	var rs []*Application
	require.NoError(t, dbscan.ScanAll(&rs, rows))

	fmt.Println("done")
}

func TestSqlite_DBScan_1(t *testing.T) {
	db := getSqliteDB()

	defer db.Close()

	var rs []map[string]any

	rows, err := db.Query("select * from applications limit 5")
	require.NoError(t, err)
	err = dbscan.ScanAll(&rs, rows)
	require.NoError(t, err)
}

func TestSqlite_DBScan_2(t *testing.T) {
	db := getSqliteDB()

	defer db.Close()

	var rs = new(Application)

	rows, err := db.Query("select * from applications limit 1")
	require.NoError(t, err)
	err = dbscan.ScanOne(rs, rows)
	require.NoError(t, err)
}

func getXSqlOpts(d, n string) *xsql.Config {
	opts := &xsql.Config{
		Dialect: d,
		Name:    n,
	}
	return opts
}

type wrapOptions struct {
	cfg *xsql.Config
}

func TestSqlOptions(t *testing.T) {
	opts1 := getXSqlOpts("ddd", "nnnn")
	opts2 := getXSqlOpts("ffff", "aaaa")
	opts3 := getXSqlOpts("qqqq", "tttt")

	as := []*xsql.Config{
		opts1,
		opts2,
	}

	as = append(as, opts3)

	wopt := wrapOptions{
		cfg: opts1,
	}

	opts1.Dialect = "ninin"
	opts1.Name = "kkdf"

	opts3.Dialect = "gggg"
	opts3.Name = "vvvv"

	require.NotNil(t, opts1)
	require.NotNil(t, wopt)
	require.NotNil(t, as)
}
