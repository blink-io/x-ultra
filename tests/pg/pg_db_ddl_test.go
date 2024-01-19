package pg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	modelApp = (*Application)(nil)
)

func TestPg_DB_CreateTable_1(t *testing.T) {
	db := getPgDB()
	m := (*Application)(nil)
	_, err := db.NewCreateTable().IfNotExists().Model(m).Exec(ctx)
	require.NoError(t, err)
}

func TestPg_DB_DropTable_1(t *testing.T) {
	db := getPgDB()
	m := (*Application)(nil)
	_, err := db.NewDropTable().IfExists().Model(m).Exec(ctx)
	require.NoError(t, err)
}

func TestPg_RebuildTable_1(t *testing.T) {
	TestPg_DB_DropTable_1(t)
	TestPg_DB_CreateTable_1(t)
}
