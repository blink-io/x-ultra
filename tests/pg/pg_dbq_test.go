package pg

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/postgres"
	"github.com/doug-martin/goqu/v9"
	"github.com/stretchr/testify/require"
)

func TestPg_DBQ_Insert_1(t *testing.T) {
	db := getPgDBQ()
	r := newRandomRecordForApp("dbq")
	ds := db.From(r.TableName())
	rt, err := ds.Insert().Rows(r).Executor().Exec()

	fmt.Println("SQL Result: ", rt)

	require.NoError(t, err)
}

func TestPg_DBQ_Select_1(t *testing.T) {
	db := getPgDBQ()
	ds := db.From("applications")
	sql, _, err := ds.Select("name", "code", "type").
		Where(goqu.C("code").Like("code-D%")).ToSQL()
	require.NoError(t, err)

	fmt.Println("SQL Generated: ", sql)
}

func TestPg_DBQ_Select_Version(t *testing.T) {
	db := getPgDBQ()

	ver := postgres.QueryVersion(ctx, db.QueryRowContext)
	fmt.Println("Version: ", ver)
}
