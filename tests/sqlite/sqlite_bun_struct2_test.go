package sqlite

import (
	xbunx "github.com/blink-io/x/bun"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSqlite_Bun_Tuple2SQL_1(t *testing.T) {
	db := getSqliteDB()

	vals, err := xbunx.TypeTuple2SQL[int64, string](ctx, db, "select id, guid from users limit ?", 10)
	require.NoError(t, err)
	require.NotNil(t, vals)
}

func TestSqlite_Bun_Tuple3SQL_1(t *testing.T) {
	db := getSqliteDB()

	vals, err := xbunx.TypeTuple3SQL[int64, string, string](ctx, db, "select id, guid, profile from users limit ?", 10)
	require.NoError(t, err)
	require.NotNil(t, vals)
}
