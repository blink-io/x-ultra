package pg

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/postgres"
	"github.com/go-rel/rel"
	"github.com/stretchr/testify/require"
)

func TestPg_DBM_Select_Version(t *testing.T) {
	db := getPgDBM().SqlDB()
	ver := postgres.QueryVersion(ctx, db.QueryRowContext)
	fmt.Println("Version: ", ver)
}

func TestPg_DBM_Select_Funcs(t *testing.T) {
	db := getPgDBM()

	type Result struct {
		Payload string `db:"payload"`
	}

	sqlF := "select %s as payload"
	funcs := getPgFuncsMap()
	rt := new(Result)
	for k, v := range funcs {
		ss := rel.SQL(fmt.Sprintf(sqlF, v))
		require.NoError(t, db.Find(ctx, rt, ss))
		fmt.Println(k, "=>", rt.Payload)
	}
}
