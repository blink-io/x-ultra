package pg

import (
	"testing"

	"github.com/blink-io/x/sql/db/g"
	"github.com/stretchr/testify/require"
)

func TestPg_DB_Insert_1(t *testing.T) {
	db := getPgDB()
	r := newRandomRecordForApp("bun")
	txdb, err := g.NewDB[Application, string](db).Tx()
	require.NoError(t, err)

	err1 := txdb.BulkInsert(ctx, []*Application{r})
	if err1 != nil {
		require.NoError(t, txdb.Rollback())
	} else {
		require.NoError(t, txdb.Commit())
	}
}
