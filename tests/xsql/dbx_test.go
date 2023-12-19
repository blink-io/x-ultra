package bun

import (
	"fmt"
	"testing"
	"time"

	"github.com/blink-io/x/id"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestSQLite3_Select_Funcs_DBX(t *testing.T) {
	db := getDBXWithSQLite()

	sqlF := "select %s as payload"
	funcs := []string{
		"hex(randomblob(32))",
		"random()",
		"sqlite_version()",
		"total_changes()",
		`lower("HELLO")`,
		`upper("hello")`,
		`length("hello")`,
		`length("我是世界")`,
		//`concat("Hello", ",", "World")`,
	}

	for _, fstr := range funcs {
		ss := fmt.Sprintf(sqlF, fstr)
		q := db.NewQuery(ss)
		var v string
		require.NoError(t, q.Row(&v))
		fmt.Println("SQLite func payload:  ", v)
	}
}

func TestDBX_Insert_1(t *testing.T) {
	db := getDBXWithSQLite()

	r1 := new(Application)
	r1.ID = id.ShortID()
	r1.Name = gofakeit.Name()
	r1.Code = gofakeit.Animal()
	r1.Type = gofakeit.Bird()
	r1.Status = gofakeit.Cat()
	r1.CreatedAt = time.Now()
	r1.UpdatedAt = time.Now()

	err := db.Model(r1).Insert()
	require.NoError(t, err)
}
