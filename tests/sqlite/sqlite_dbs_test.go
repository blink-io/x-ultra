package sqlite

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/require"
)

func TestSqlite_DBS_1(t *testing.T) {
	db := getSqliteDBS()
	pxy := squirrel.NewStmtCacheProxy(db.SqlDB())
	require.NotNil(t, pxy)

	//var as []*Application
	//var ar *Application
	//
	//row := db.B().Select(appColumns...).
	//	From("applications").
	//	RunWith(db.SqlDB()).Query()
	////require.NoError(t, err)

	fmt.Println("done")
}
