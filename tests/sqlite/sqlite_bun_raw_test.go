package sqlite

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/blink-io/x/bun"
)

func TestSqlite_Bun_Raw_Select_1(t *testing.T) {
	db := getSqliteDB()

	users := make([]User, 0)

	err := db.NewRaw(
		"SELECT id, guid, profile FROM ? LIMIT ?",
		bun.Ident("users"), 100,
	).Scan(ctx, &users)
	require.NoError(t, err)
}

func TestSqlite_Bun_Raw_Select_Custom_1(t *testing.T) {
	db := getSqliteDB()
	var rs []*IDAndProfile
	err := db.NewRaw("select * from users where id > 0").
		Scan(ctx, &rs)
	require.NoError(t, err)
}
