package tests

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	xsql "github.com/blink-io/x/sql"
	xdb "github.com/blink-io/x/sql/db"
	"github.com/samber/do/v2"
	"github.com/stretchr/testify/require"
)

var yesErr = errors.New("yes error")

func TestDo_1(t *testing.T) {
	// create DI container
	i := do.New()

	// inject both services into DI container
	do.Provide[*xdb.DB](i, NewDB)
	do.Provide[*xsql.Config](i, NewConfig)

	uname := "uni-opts"
	do.ProvideNamedTransient(i, uname, NewConfig)

	opt1 := do.MustInvoke[*xsql.Config](i)
	opt2 := do.MustInvoke[*xsql.Config](i)
	require.NotNil(t, opt1)
	require.NotNil(t, opt2)

	uopt1 := do.MustInvokeNamed[*xsql.Config](i, uname)
	uopt2 := do.MustInvokeNamed[*xsql.Config](i, uname)
	require.NotNil(t, uopt1)
	require.NotNil(t, uopt2)

	db, err := do.Invoke[*xdb.DB](i)
	require.NoError(t, err)

	row := db.QueryRow("select sqlite_version()")
	var str string
	require.NoError(t, row.Scan(&str))

	fmt.Println("Result: ", str)

	i.Shutdown()
}

func TestDo_2(t *testing.T) {
	i := do.New()
	uname := "abc"
	do.ProvideNamed[*xsql.Config](i, uname, NewConfig)

	uopt1 := do.MustInvokeNamed[*xsql.Config](i, uname)
	uopt2 := do.MustInvokeNamed[*xsql.Config](i, uname)
	require.NotNil(t, uopt1)
	require.NotNil(t, uopt2)
	require.Equal(t, uopt1, uopt2)

	dequal := reflect.DeepEqual(uopt1, uopt2)
	fmt.Println("DeepEqual: ", dequal)

	do.OverrideNamed(i, uname, NewConfig)
	uopt3 := do.MustInvokeNamed[*xsql.Config](i, uname)
	require.NotEqual(t, uopt2, uopt3)
}

func TestDo_3(t *testing.T) {
	i := do.New()
	uname := "abc"
	do.ProvideNamed[*xsql.Config](i, uname, NewConfig)

	uopt1 := do.MustInvokeNamed[*xsql.Config](i, uname)

	uopt2 := new(xsql.Config)

	reflect.Copy(reflect.ValueOf(uopt2), reflect.ValueOf(uopt1))

	fmt.Println("done")
}

func NewDB(i do.Injector) (*xdb.DB, error) {
	return xdb.New(do.MustInvoke[*xsql.Config](i))
}

func NewConfig(i do.Injector) (*xsql.Config, error) {
	var opt = &xsql.Config{
		Dialect: xsql.DialectSQLite,
		Host:    sqlitePath,
		Loc:     time.Local,
	}
	return opt, nil
}
