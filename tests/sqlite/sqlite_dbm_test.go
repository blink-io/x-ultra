package sqlite

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/sqlite"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
)

func TestSqlite_DBM_Select_Version(t *testing.T) {
	db := getSqliteDBM().SqlDB()
	ver := sqlite.QueryVersion(ctx, db.QueryRowContext)
	fmt.Println("Version: ", ver)
}

func TestSqlite_DBM_Select_Funcs(t *testing.T) {
	db := getSqliteDBM()

	type Result struct {
		Payload string `db:"payload"`
	}

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()
	rt := new(Result)
	for k, v := range funcs {
		ss := rel.SQL(fmt.Sprintf(sqlF, v))
		require.NoError(t, db.Find(ctx, rt, ss))
		fmt.Println(k, "=>", rt.Payload)
	}
}

func TestSqlite_DBM_Select_1(t *testing.T) {
	db := getSqliteDBM()

	type Result struct {
		Verinfo  string `db:"verinfo"`
		Rstr     string `db:"rstr"`
		SourceID string `db:"source_id"`
	}

	var rtmap Result

	sqlstr := `select
random() as rstr,
sqlite_source_id() as source_id, 
sqlite_version() as verinfo;`

	err := db.Find(ctx, &rtmap, rel.SQL(sqlstr))
	require.NoError(t, err)

	fmt.Printf("------------------------------------------------------------\n")

	fmt.Println("Sqlite version: ", rtmap)

	require.NoError(t, db.Ping(ctx))
}

func TestSqlite_DBM_Select_2(t *testing.T) {
	db := getSqliteDBM()
	rt := new(Application)

	sqlF := "select * from applications where type like ? limit 1"
	ss := rel.SQL(sqlF, "%type-01%")
	require.NoError(t, db.Find(ctx, rt, ss))
	fmt.Println("Result ==================>")
	litter.Dump(rt)
}

func TestSqlite_DBM_Select_3(t *testing.T) {
	db := getSqliteDBM()
	rt := new(Application)

	require.NoError(t, db.Find(ctx, rt, where.Like("type", "%type-01%")))
	fmt.Println("Result ==================>")
	litter.Dump(rt)
}

func TestSqlite_DBM_Insert_1(t *testing.T) {
	db := getSqliteDBM()

	r1 := newRandomRecordForApp(db.Accessor())

	err := db.Insert(ctx, r1)
	require.NoError(t, err)
}
