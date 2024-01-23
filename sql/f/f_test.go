package f

import (
	"testing"

	"github.com/leporo/sqlf"
	"github.com/stretchr/testify/require"
)

func TestSqlf(t *testing.T) {
	st := sqlf.New("")
	require.NotNil(t, st)
}
