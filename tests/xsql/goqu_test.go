package bun

import (
	"fmt"
	"testing"
	"time"

	"github.com/blink-io/x/id"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestGoqu_1(t *testing.T) {
	//Debug = true
	db := getGoquDBWithSQLite()

	sqlF := "select %s as payload"
	funcs := map[string]string{
		"hex":              "hex(randomblob(32))",
		"random":           "random()",
		"version":          "sqlite_version()",
		"changes":          "changes()",
		"total_changes":    "total_changes()",
		"lower":            `lower("HELLO")`,
		"upper":            `upper("hello")`,
		"length":           `length("hello")`,
		"sqlite_source_id": `sqlite_source_id()`,
		//`concat("Hello", ",", "World")`,
	}

	for k, v := range funcs {
		ss := fmt.Sprintf(sqlF, v)
		row := db.QueryRow(ss)
		var v string
		require.NoError(t, row.Scan(&v))
		fmt.Println(k, ":  ", v)
	}
}

func TestGoqu_Insert_1(t *testing.T) {
	db := getGoquDBWithSQLite()

	r1 := new(Application)
	r1.ID = id.ShortID()
	r1.Name = gofakeit.Name()
	r1.Code = gofakeit.Animal()
	r1.Type = gofakeit.Bird()
	r1.Status = gofakeit.Cat()
	r1.CreatedAt = time.Now()
	r1.UpdatedAt = time.Now()

	_, err := db.Insert(r1.TableName()).Rows(r1).Executor().Exec()
	require.NoError(t, err)
}
