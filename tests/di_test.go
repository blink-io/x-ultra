package tests

import (
	"errors"
	"fmt"
	"testing"
	"time"

	xsql "github.com/blink-io/x/sql"
	"github.com/samber/do/v2"
	"github.com/stretchr/testify/require"
)

var yesErr = errors.New("yes error")

func TestDo_1(t *testing.T) {
	// create DI container
	i := do.New()

	// inject both services into DI container
	do.Provide[*xsql.DB](i, NewDB)
	do.Provide[*xsql.DBP](i, NewDBPWithErr)
	do.Provide[*xsql.Options](i, NewOptions)

	uname := "uni-opts"
	do.ProvideNamedTransient(i, uname, NewOptions)

	opt1 := do.MustInvoke[*xsql.Options](i)
	opt2 := do.MustInvoke[*xsql.Options](i)
	require.NotNil(t, opt1)
	require.NotNil(t, opt2)

	uopt1 := do.MustInvokeNamed[*xsql.Options](i, uname)
	uopt2 := do.MustInvokeNamed[*xsql.Options](i, uname)
	require.NotNil(t, uopt1)
	require.NotNil(t, uopt2)

	db, err := do.Invoke[*xsql.DB](i)
	require.NoError(t, err)

	dbx, err2 := do.Invoke[*xsql.DBX](i)
	require.Nil(t, dbx)
	require.Error(t, err2)

	dbp, err3 := do.Invoke[*xsql.DBP](i)
	require.Nil(t, dbp)
	require.Equal(t, yesErr, err3)

	row := db.QueryRow("select sqlite_version()")
	var str string
	require.NoError(t, row.Scan(&str))

	fmt.Println("Result: ", str)

	i.Shutdown()
}

func NewDBPWithErr(i do.Injector) (*xsql.DBP, error) {
	return nil, yesErr
}

func NewDB(i do.Injector) (*xsql.DB, error) {
	return xsql.NewDB(do.MustInvoke[*xsql.Options](i))
}

func NewOptions(i do.Injector) (*xsql.Options, error) {
	var opt = &xsql.Options{
		Dialect: xsql.DialectSQLite,
		Host:    sqlitePath,
		DOptions: []xsql.DOption{
			xsql.WithLocation(time.Local),
		},
		Loc: time.Local,
	}
	return opt, nil
}
