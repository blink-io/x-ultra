package sql

import (
	"database/sql"
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/require"
)

func TestSQL_Builder_1(t *testing.T) {
	select1 := B().
		Select("name", "age", "level").
		From("users").
		Where(sq.And{
			sq.Eq{"name": "Heison"},
			sq.Eq{"level": "L1"},
		})

	sql1, params1, err := select1.ToSql()
	require.NoError(t, err)

	sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	fmt.Print("SQL1:  ", sql1, "\nParams1:  ", params1, "\n")

	var db *sql.DB
	db.Exec(sql1, params1)
}
