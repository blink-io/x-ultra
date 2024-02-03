package pg

import (
	"testing"

	"github.com/blink-io/x/id"
	xsql "github.com/blink-io/x/sql"
	"github.com/stretchr/testify/require"
)

func TestPg_DBX_ConstraintCheck_Insert_1(t *testing.T) {
	db := getPgDBZ()
	sql := "insert into mymy(status, code) values ($1, $2)"
	args := []any{
		"not-ok",
		id.ShortID(),
	}
	_, err := db.ExecContext(ctx, sql, args...)
	nerr := xsql.WrapError(err)

	require.ErrorIs(t, nerr, xsql.ErrConstraintCheck)
}

func TestPg_DBX_ConstraintUnique_Insert_2(t *testing.T) {
	db := getPgDBZ()
	sql := "insert into mymy(status, code) values ($1, $2)"
	args := []any{
		"ok",
		"1",
	}
	_, err := db.ExecContext(ctx, sql, args...)
	nerr := xsql.WrapError(err)

	require.ErrorIs(t, nerr, xsql.ErrConstraintUnique)
}
