package bun

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/blink-io/x/bun/extra/timing"
	xsql "github.com/blink-io/x/sql"
	"github.com/blink-io/x/sql/hooks"
	timinghook "github.com/blink-io/x/sql/hooks/timing"
	"github.com/blink-io/x/sql/scany/dbscan"
	"github.com/qustavo/dotsql"
	"github.com/stretchr/testify/require"
)

var pgOpt = &xsql.Options{
	Context:       context.Background(),
	Dialect:       xsql.DialectPostgres,
	Name:          "blink",
	User:          "blinkbot",
	Port:          15432,
	Host:          "192.168.11.179",
	Loc:           time.Local,
	ValidationSQL: "SELECT 1;",
	ClientName:    "blink-dev",
	DriverHooks: []hooks.Hooks{
		timinghook.New(),
	},
	QueryHooks: []xsql.QueryHook{
		//logging.Func(log.Printf),
		timing.New(),
	},
}

func init() {
	pwd := getPwd()
	pgOpt.Password = pwd
}

func getPwd() string {
	homedir, _ := os.UserHomeDir()

	data, err := os.ReadFile(filepath.Join(homedir, ".passwd.pg"))
	if err != nil {
		panic(err)
	}

	pwd := strings.TrimSuffix(string(data), "\n")
	return pwd
}

func GetPgDB() *xsql.DB {
	db, err1 := xsql.NewDB(pgOpt)
	if err1 != nil {
		panic(err1)
	}

	return db
}

func TestPG_Connect_1(t *testing.T) {
	db := GetPgDB()
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

	db := GetPgDB()

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
}
