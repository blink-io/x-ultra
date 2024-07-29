package pg

import (
	"testing"

	bunx "github.com/blink-io/x/bun"

	"github.com/stretchr/testify/require"
)

func TestPg_DB_Insert_1(t *testing.T) {
	db := getPgDB()
	r1 := newRandomRecordForApp("bun")
	r2 := newRandomRecordForApp("bun")
	txdb, err := bunx.NewGenericDB[Application, string](db).Tx(ctx, nil)
	require.NoError(t, err)

	err1 := txdb.BulkInsert(ctx, []*Application{r1, r2})
	if err1 != nil {
		require.NoError(t, txdb.Rollback())
	} else {
		require.NoError(t, txdb.Commit())
	}
}
