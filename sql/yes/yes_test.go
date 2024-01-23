package yes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFile_1(t *testing.T) {
	path := ""

	qry, err := ParseFile(path)
	require.NoError(t, err)
	require.NotNil(t, qry)
}
