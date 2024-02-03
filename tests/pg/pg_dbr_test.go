package pg

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/postgres"
	"github.com/stretchr/testify/require"
)

func TestPg_DBR_Select_Funcs(t *testing.T) {
	db := getPgDBX()

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

func TestPg_DBR_Select_Version(t *testing.T) {
	db := getPgDBX()

	ver := postgres.QueryVersion(ctx, db.QueryRowContext)
	fmt.Println("Version: ", ver)
}
