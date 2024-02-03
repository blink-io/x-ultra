package pg

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/postgres"
	"github.com/stretchr/testify/require"
)

func TestPg_DBZ_Select_Funcs(t *testing.T) {
	db := getPgDBR()

	sess := db.NewSession(nil)

	sqlF := "select %s as payload"
	funcs := getPgFuncsMap()

	for k, fstr := range funcs {
		ss := fmt.Sprintf(sqlF, fstr)
		var v string
		r := sess.QueryRow(ss)
		require.NoError(t, r.Scan(&v))
		fmt.Println(k, " => ", v)
	}
}

func TestPg_DBZ_Select_Version(t *testing.T) {
	db := getPgDBZ()

	ver := postgres.QueryVersion(ctx, db.SqlDB().QueryRowContext)
	fmt.Println("Version: ", ver)
}
