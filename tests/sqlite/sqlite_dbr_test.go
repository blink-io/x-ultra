package sqlite

import (
	"fmt"
	"testing"
	"time"

	"github.com/blink-io/x/id"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
)

func TestSqlite_DBR_Select_Funcs(t *testing.T) {
	db := getSqliteDBR()

	sqlF := "select %s as payload"
	funcs := getSqliteFuncMap()

	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		var rt string
		r := db.QueryRow(ss)
		require.NoError(t, r.Scan(&rt))
		fmt.Println(k, "=>", rt)
	}
}

func TestSqlite_DBR_Insert_1(t *testing.T) {
	db := getSqliteDBR()

	r1 := newRandomRecordForApp("dbr")
	r2 := newRandomRecordForApp("dbr")
	_, err := db.InsertInto(r1.Table()).
		Columns(r1.Columns()...).
		Record(r1).Record(r2).
		Exec()
	require.NoError(t, err)
}

func TestSqlite_DBR_Insert_Raw_1(t *testing.T) {
	db := getSqliteDBR()
	model := appModel()
	n := time.Now()

	_, err := db.InsertInto(model.Table()).
		//Columns(model.Columns()...).
		Pair("guid", id.UUID()).
		Pair("status", "hen-good").
		Pair("code", id.ShortUUID()).
		Pair("type", "type-005-"+db.Accessor()).
		Pair("name", gofakeit.Name()).
		Pair("description", gofakeit.ProductName()).
		Pair("created_at", n).
		Pair("updated_at", n).
		Exec()
	require.NoError(t, err)
}

func TestSqlite_DBR_Update_Raw_1(t *testing.T) {
	db := getSqliteDBR()
	mm := appModel()
	n := time.Now()

	var ra = new(Application)

	err := db.Update(mm.TableName()).SetMap(map[string]any{
		"updated_at": n,
		"status":     "111",
	}).Where("id = ?", 8).Returning(mm.Columns()...).Load(ra)
	require.NoError(t, err)
}

func TestSqlite_DBR_Update_Raw_2(t *testing.T) {
	db := getSqliteDBR()
	mm := appModel()
	n := time.Now()

	var ra = new(Application)

	err := db.Update(mm.TableName()).
		Set("status", "fuckll").
		Set("description", "Xilitron is ON").
		Set("updated_at", n).
		Where("id > ?", 8).Returning(mm.Columns()...).Load(ra)
	require.NoError(t, err)
}

func TestSqlite_DBR_Select_1(t *testing.T) {
	db := getSqliteDBR()

	rt := new(Application)

	err := db.Select("*").From(rt.Table()).
		Where("type like ?", "%type-003%").
		Where("? = ?", 1, 1).
		Limit(1).LoadOne(rt)
	require.NoError(t, err)
	fmt.Println("Result ==================>")
	litter.Dump(rt)
}
