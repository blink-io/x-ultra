package sqlite

import (
	"testing"

	"github.com/stephenafamo/scan"
	"github.com/stephenafamo/scan/stdscan"
	"github.com/stretchr/testify/require"
)

func TestScan_1(t *testing.T) {
	db := getSqliteSqlDB()
	a1, err := stdscan.One(ctx, db, scan.StructMapper[*Application](), "SELECT * FROM applications limit 1")
	require.NoError(t, err)
	require.NotNil(t, a1)
}

func TestName(t *testing.T) {

}
